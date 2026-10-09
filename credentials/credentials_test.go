package credentials

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/openfga/go-sdk/internal/utils/retryutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientCredentialsApiTokenIssuerParams(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		params url.Values
		aud    string
		scopes string
		want   url.Values
	}{
		{
			name:   "extra params only",
			params: url.Values{"store_id": {"store-123"}, "resource": {"one", "two"}},
			want:   url.Values{"grant_type": {"client_credentials"}, "store_id": {"store-123"}, "resource": {"one", "two"}},
		},
		{
			name:   "extra params plus audience",
			params: url.Values{"store_id": {"store-123"}},
			aud:    "https://api.example.com",
			want:   url.Values{"grant_type": {"client_credentials"}, "store_id": {"store-123"}, "audience": {"https://api.example.com"}},
		},
		{
			name: "audience only",
			aud:  "https://api.example.com",
			want: url.Values{"grant_type": {"client_credentials"}, "audience": {"https://api.example.com"}},
		},
		{
			name:   "extra params plus scopes",
			params: url.Values{"store_id": {"store-123"}},
			scopes: "read write",
			want:   url.Values{"grant_type": {"client_credentials"}, "store_id": {"store-123"}, "scope": {"read write"}},
		},
		{
			name: "nil params",
			want: url.Values{"grant_type": {"client_credentials"}},
		},
		{
			name:   "empty params",
			params: url.Values{},
			want:   url.Values{"grant_type": {"client_credentials"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got url.Values
			issuer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/token" {
					if err := r.ParseForm(); err != nil {
						t.Errorf("failed parsing token request form: %v", err)
					} else {
						got = r.PostForm
					}
					w.Header().Set("Content-Type", "application/json")
					_, _ = w.Write([]byte(`{"access_token":"test-token","token_type":"Bearer","expires_in":3600}`))
					return
				}
				w.WriteHeader(http.StatusNoContent)
			}))
			defer issuer.Close()

			creds, err := NewCredentials(Credentials{
				Method: CredentialsMethodClientCredentials,
				Config: &Config{
					ClientCredentialsClientId:             "client-id",
					ClientCredentialsClientSecret:         "client-secret",
					ClientCredentialsApiTokenIssuer:       issuer.URL + "/token",
					ClientCredentialsApiAudience:          tt.aud,
					ClientCredentialsScopes:               tt.scopes,
					ClientCredentialsApiTokenIssuerParams: tt.params,
				},
			})
			require.NoError(t, err)

			client, _ := creds.GetHttpClientAndHeaderOverrides(retryutils.RetryParams{}, false)
			resp, err := client.Get(issuer.URL + "/resource")
			require.NoError(t, err)
			_ = resp.Body.Close()

			assert.True(t, reflect.DeepEqual(got, tt.want), "token request form = %#v; want %#v", got, tt.want)
		})
	}
}

func TestClientCredentialsApiTokenIssuerParamsRejectReservedKeys(t *testing.T) {
	t.Parallel()

	for _, key := range []string{"", "grant_type", "client_id", "client_secret", "scope", "audience"} {
		t.Run(key, func(t *testing.T) {
			t.Parallel()

			const secretValue = "sensitive-value"
			_, err := NewCredentials(Credentials{
				Method: CredentialsMethodClientCredentials,
				Config: &Config{
					ClientCredentialsClientId:             "client-id",
					ClientCredentialsClientSecret:         "client-secret",
					ClientCredentialsApiTokenIssuer:       "https://issuer.example.com/token",
					ClientCredentialsApiTokenIssuerParams: url.Values{key: {secretValue}},
				},
			})
			require.Error(t, err)
			assert.Contains(t, err.Error(), fmt.Sprintf("%q", key), "error should identify reserved key")
			assert.NotContains(t, err.Error(), secretValue, "error should not include parameter value")
		})
	}
}
