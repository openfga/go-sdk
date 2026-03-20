# OAuth2 Credentials and Token Lifecycle

## Credential Types

The SDK supports three credential methods (from `credentials/credentials.go`):

| Type | Usage | Token Refresh | Best For |
|---|---|---|---|
| `CredentialsMethodNone` | No authentication | N/A | Development, public APIs |
| `CredentialsMethodApiToken` | Static bearer token | Never | Service-to-service with long-lived tokens |
| `CredentialsMethodClientCredentials` | OAuth2 client credentials | Automatic | Secure production deployments |

## Configuration

```go
import "github.com/openfga/go-sdk/credentials"

// No credentials
creds := &credentials.Credentials{
    Method: credentials.CredentialsMethodNone,
}

// Static API token
creds := &credentials.Credentials{
    Method: credentials.CredentialsMethodApiToken,
    Config: &credentials.Config{
        ApiToken: "your-token-here",
    },
}

// OAuth2 client credentials
creds := &credentials.Credentials{
    Method: credentials.CredentialsMethodClientCredentials,
    Config: &credentials.Config{
        ClientCredentialsApiTokenIssuer: "https://issuer.fga.example",
        ClientCredentialsApiAudience:    "https://api.fga.example",
        ClientCredentialsClientId:       "client-id",
        ClientCredentialsClientSecret:   "client-secret",
        ClientCredentialsScopes:         "read write",  // Optional
    },
}

config := &client.ClientConfiguration{
    ApiUrl:      "https://api.fga.example",
    Credentials: creds,
}

sdkClient, err := client.NewSdkClient(config)
```

## Token Lifecycle

### OAuth2 Token Refresh

From `internal/constants/constants.go`:

```go
TokenExpiryThresholdBufferInSec = 300   // Refresh if expiry is within 5 min
TokenExpiryJitterInSec = 300            // Add 5 min random jitter
```

**Combined effect:** Token may refresh **5–10 minutes early**:
- Threshold check: If token expires in <= 300 seconds, start refresh
- Jitter: Add 0–300 seconds random jitter, so actual refresh occurs 5–10 minutes before expiry

### Why Jitter?

Jitter prevents **thundering herd**: If many SDK instances refresh at exactly expiry time, they'd hammer the token issuer simultaneously. Random jitter spreads refresh requests.

### Token Refresh Flow

1. **Check expiry** on every request to an OAuth2-protected endpoint
2. **If expired or within threshold + jitter**, request a new token
3. **Bearer token injected** in `Authorization: Bearer <token>` header
4. **Metrics recorded:** `fga_client_credentials_request` (counter)

From `oauth2/clientcredentials/clientcredentials.go`:

```go
func (r *tokenRefresher) GetToken(ctx context.Context) (*Token, error) {
    // Check if token is stale (expires within threshold + jitter)
    // If stale, request new token from issuer
    // Cache token until near expiry
    // Return cached or new token
}
```

### Token Caching

Tokens are cached until near expiry, so multiple calls don't trigger redundant refreshes.

## RoundTripper Pattern

The SDK wraps `http.Client` with a custom `RoundTripper` to inject tokens:

From `oauth2/transport.go`:

```go
type oAuth2Transport struct {
    base http.RoundTripper
    // token refresher instance
}

func (t *oAuth2Transport) RoundTrip(req *http.Request) (*http.Response, error) {
    // Get current token (refresh if needed)
    token, err := t.GetToken(req.Context())
    if err != nil {
        return nil, err
    }

    // Inject Bearer token
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

    // Delegate to underlying RoundTripper
    return t.base.RoundTrip(req)
}
```

This pattern ensures:
- **Tokens are refreshed transparently** before each request
- **No manual token management** needed by the application
- **Metrics recorded** for each credential request

## Critical: Do Not Write Time-Based Tests

Because tokens refresh 5–10 minutes early (due to threshold + jitter), DO NOT write tests that assert on exact token expiry times:

```go
// WRONG: Assumes token valid until exact expiry time
now := time.Now()
expiryTime := now.Add(100 * time.Second)  // Expires in 100 seconds

token := &oauth2.Token{
    AccessToken: "test-token",
    Expiry:      expiryTime,
}

// SDK may refresh immediately (threshold 300s < 100s is false, but jitter might push it)
// Test is flaky and unpredictable
```

**Solution: Mock the token provider instead:**

```go
// CORRECT: Mock the tokenRefresher
type mockTokenRefresher struct {
    token *oauth2.Token
}

func (m *mockTokenRefresher) GetToken(ctx context.Context) (*oauth2.Token, error) {
    return m.token, nil
}

// Use mock in tests; no time-dependent logic
```

## Handling Expired Credentials

If credentials are invalid or tokens cannot be refreshed:

```go
resp, err := client.Check(ctx, storeId, request)

if err != nil {
    if authErr, ok := err.(*openfga.FgaApiAuthenticationError); ok {
        // Credentials invalid or token refresh failed
        log.Fatalf("Auth failed for store %s: %s", authErr.StoreId(), authErr.Error())
    }
}
```

Common causes:
- **Wrong client ID/secret**: Check `ClientCredentialsClientId` and `ClientCredentialsClientSecret`
- **Expired secret**: Rotate credentials and update configuration
- **Issuer unreachable**: Check `ClientCredentialsApiTokenIssuer` URL and network connectivity
- **Insufficient scope**: Add required scopes to `ClientCredentialsScopes`

## Token Headers

The API token is sent in the `Authorization` header:

```
Authorization: Bearer <access_token>
```

For OAuth2, the token is obtained from the configured issuer and used the same way.

## Testing OAuth2 Flows

Use mocks from `oauth2/` package or `httpmock` to simulate token responses:

```go
import (
    "github.com/jarcoal/httpmock"
    "github.com/openfga/go-sdk/oauth2/clientcredentials"
)

func TestOAuth2TokenRefresh(t *testing.T) {
    mock := httpmock.NewMockHTTP()
    defer mock.Close()

    // Mock token issuer
    mock.WithStatusCode(200).
        WithHeader("Content-Type", "application/json").
        WithResponse(`{
            "access_token": "new-token",
            "token_type": "Bearer",
            "expires_in": 3600
        }`)

    config := &client.ClientConfiguration{
        Credentials: &credentials.Credentials{
            Method: credentials.CredentialsMethodClientCredentials,
            Config: &credentials.Config{
                ClientCredentialsApiTokenIssuer: mock.URL + "/token",
                ClientCredentialsClientId:       "test-client",
                ClientCredentialsClientSecret:   "test-secret",
            },
        },
    }

    c, _ := client.NewSdkClient(config)
    resp, err := c.Check(ctx, storeId, request)
    require.NoError(t, err)
}
```

## Secret Security

**Never commit secrets to the repository:**

- Do not hardcode credentials in code
- Use environment variables or secret management services
- Configure credentials from `~/.fga/config` or environment at runtime

Example with environment variables:

```go
config := &client.ClientConfiguration{
    ApiUrl: os.Getenv("FGA_API_URL"),
    Credentials: &credentials.Credentials{
        Method: credentials.CredentialsMethodClientCredentials,
        Config: &credentials.Config{
            ClientCredentialsApiTokenIssuer: os.Getenv("FGA_ISSUER"),
            ClientCredentialsClientId:       os.Getenv("FGA_CLIENT_ID"),
            ClientCredentialsClientSecret:   os.Getenv("FGA_CLIENT_SECRET"),
        },
    },
}
```

## Implementation Files

- **`credentials/credentials.go`** — Credential types and factory functions
- **`oauth2/clientcredentials/clientcredentials.go`** — OAuth2 token refresh logic
- **`oauth2/transport.go`** — RoundTripper for Bearer token injection
- **`oauth2/token.go`** — Token struct and lifecycle
- **`internal/constants/constants.go`** — Token expiry threshold and jitter constants

## Summary

- **Three credential types:** None, API token, OAuth2 client credentials
- **OAuth2 tokens refresh automatically** 5–10 minutes before expiry
- **Do not write time-based tests** — Mock the token provider instead
- **Use RoundTripper pattern** — Tokens injected transparently
- **Never commit secrets** — Use environment variables or secret management
- **Check authentication errors** for credential issues
- **Metrics recorded** for every credential request (`fga_client_credentials_request`)
