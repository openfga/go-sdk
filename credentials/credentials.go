package credentials

import (
    "context"
    "fmt"
    "net/http"
	"net/url"

    "github.com/openfga/go-sdk/oauth2/clientcredentials"
)

const ApiTokenHeaderKey = "Authorization"
const ApiTokenHeaderValuePrefix = "Bearer"

// Avaialable credential methods
type CredentialsMethod string

const (
    // No credentials (default)
    CredentialsMethodNone              CredentialsMethod = "none"
    // API Token credentials (will be sent in "Authorization: Bearer $TOKEN" header)
    CredentialsMethodApiToken          CredentialsMethod = "api_token"
    // Client Credentials flow will be performed, resulting token will be sent in "Authorization: Bearer $TOKEN" header
    CredentialsMethodClientCredentials CredentialsMethod = "client_credentials"
)

type Config struct {
    ApiToken                        string `json:"apiToken,omitempty"`
    ClientCredentialsApiTokenIssuer string `json:"apiTokenIssuer,omitempty"`
    ClientCredentialsApiAudience    string `json:"apiAudience,omitempty"`
    ClientCredentialsClientId       string `json:"clientId,omitempty"`
    ClientCredentialsClientSecret   string `json:"clientSecret,omitempty"`
}

type Credentials struct {
    Method CredentialsMethod `json:"method,omitempty"`
    Config *Config           `json:"config,omitempty"`
}

func isWellFormedUri(uriString string) bool {
    uri, err := url.Parse(uriString)

    if (err != nil) || (uri.Scheme != "http" && uri.Scheme != "https") || ((uri.Scheme + "://" + uri.Host) != uriString) {
        return false
    }

    return true
}

func NewCredentials(config Credentials) (*Credentials, error) {
    creds := &Credentials{
        Method: config.Method,
        Config: config.Config,
    }

    if creds.Method == "" {
        creds.Method = CredentialsMethodNone
    }

    err := creds.ValidateCredentialsConfig()

    if err != nil {
        return creds, nil
    }

    return nil, err
}

func (c *Credentials) ValidateCredentialsConfig() error {
    conf := c.Config
    if c.Method == CredentialsMethodApiToken && (conf == nil || conf.ApiToken == "") {
        return fmt.Errorf("CredentialsConfig.ApiToken is required when CredentialsMethod is CredentialsMethodApiToken (%s)", c.Method)
    } else if c.Method == CredentialsMethodClientCredentials {
        if conf == nil ||
            conf.ClientCredentialsClientId == "" ||
            conf.ClientCredentialsClientSecret == "" ||
            conf.ClientCredentialsApiTokenIssuer == "" ||
            conf.ClientCredentialsApiAudience == "" {
            return fmt.Errorf("all of CredentialsConfig.ClientId, CredentialsConfig.ClientSecret, CredentialsConfig.ApiAudience and CredentialsConfig.ApiTokenIssuer are required when CredentialsMethod is CredentialsMethodClientCredentials (%s)", c.Method)
        }
        if !isWellFormedUri("https://" + conf.ClientCredentialsApiTokenIssuer) {
            return fmt.Errorf("CredentialsConfig.ApiTokenIssuer (%s) is in an invalid format", "https://"+conf.ClientCredentialsApiTokenIssuer)
        }
    }

    return nil
}

type HeaderParams struct {
    Key   string
    Value string
}

func (c *Credentials) GetApiTokenHeader() *HeaderParams {
    if c.Method != CredentialsMethodApiToken {
        return nil
    }

    return &HeaderParams{
        Key:   ApiTokenHeaderKey,
        Value: ApiTokenHeaderValuePrefix + " " + c.Config.ApiToken,
    }
}

// GetHttpClientAndHeaderOverrides
// The main export the client uses to get a configuration with the necessary
// httpClient and header overrides based on the chosen credential method
func (c *Credentials) GetHttpClientAndHeaderOverrides() (*http.Client, []*HeaderParams) {
	var headers []*HeaderParams
	var client = http.DefaultClient
	switch c.Method {
	case CredentialsMethodClientCredentials:
		ccConfig := clientcredentials.Config{
			ClientID:     c.Config.ClientCredentialsClientId,
			ClientSecret: c.Config.ClientCredentialsClientSecret,
			TokenURL:     fmt.Sprintf("https://%s/oauth/token", c.Config.ClientCredentialsApiTokenIssuer),
			EndpointParams: map[string][]string{
				"audience": {c.Config.ClientCredentialsApiAudience},
			},
		}
		client = ccConfig.Client(context.Background())
	case CredentialsMethodApiToken:
		var header = c.GetApiTokenHeader()
		if header != nil {
			headers = append(headers, header)
		}
	case CredentialsMethodNone:
	default:
	}

	return client, headers
}
