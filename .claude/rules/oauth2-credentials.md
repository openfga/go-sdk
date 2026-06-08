---
description: When editing OAuth2 token handling or credential configuration
globs: ["oauth2/**", "credentials/**"]
---

Three credential types: `CredentialsMethodNone`, `CredentialsMethodApiToken`, `CredentialsMethodClientCredentials`. All configured via `credentials.Credentials` in `credentials/credentials.go`.

## Token lifecycle

- OAuth2 tokens refresh automatically 5-10 min before expiry (300s threshold + 0-300s random jitter)
- Do not remove jitter — it prevents thundering herd against the token issuer
- Tokens are cached and reused until near expiry
- Bearer token injected via custom RoundTripper in `oauth2/transport.go`
- `fga_client_credentials_request` metric recorded on each token fetch

## Rules

- Never log or include tokens/secrets in error messages or debug output
- All three credential types must be handled in switch statements — add new types to all switches
- Token refresh must be thread-safe (OpenFgaClient is shared across goroutines)
- Context must be passed through to credential operations for cancellation support
- Do not write time-based test assertions on token expiry — mock the token provider instead
- Test OAuth2 flows with `httpmock` to simulate token issuer responses
- Never hardcode credentials — use environment variables or secret management

## Key files

- `credentials/credentials.go` — Credential types and validation
- `oauth2/clientcredentials/clientcredentials.go` — Token refresh logic
- `oauth2/transport.go` — RoundTripper for Bearer injection
- `internal/constants/constants.go` — Expiry threshold and jitter constants
