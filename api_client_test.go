package openfga

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/openfga/go-sdk/credentials"
	"github.com/openfga/go-sdk/internal/constants"
	"github.com/openfga/go-sdk/oauth2"
	"github.com/openfga/go-sdk/telemetry"
)

type countingRoundTripper struct {
	base  http.RoundTripper
	count *int32
}

func (c *countingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	atomic.AddInt32(c.count, 1)
	return c.base.RoundTrip(req)
}

// Not parallel: uses httptest servers whose Close() reads the global
// http.DefaultTransport, which the httpmock-based parallel tests mutate.
func TestClientCredentialsEndToEndWithCustomClient(t *testing.T) {
	tokenRequests := int32(0)
	issuer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&tokenRequests, 1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"e2e-token","token_type":"Bearer","expires_in":3600}`))
	}))
	defer issuer.Close()

	var gotAuthHeader atomic.Value
	gotAuthHeader.Store("")
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuthHeader.Store(r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"authorization_models":[]}`))
	}))
	defer api.Close()

	var transportHits int32
	customHTTPClient := &http.Client{
		Timeout:   30 * time.Second,
		Transport: &countingRoundTripper{base: http.DefaultTransport, count: &transportHits},
	}

	configuration, err := NewConfiguration(Configuration{
		ApiUrl: api.URL,
		Credentials: &credentials.Credentials{
			Method: credentials.CredentialsMethodClientCredentials,
			Config: &credentials.Config{
				ClientCredentialsClientId:       "client-id",
				ClientCredentialsClientSecret:   "client-secret",
				ClientCredentialsApiTokenIssuer: issuer.URL,
				ClientCredentialsApiAudience:    "https://api.example.com",
			},
		},
		HTTPClient: customHTTPClient,
	})
	if err != nil {
		t.Fatalf("Failed to create configuration: %v", err)
	}

	apiClient := NewAPIClient(configuration)

	_, _, err = apiClient.OpenFgaApi.ReadAuthorizationModels(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Execute()
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	if got := gotAuthHeader.Load().(string); got != "Bearer e2e-token" {
		t.Errorf("API request did not carry the bearer token. Got: %q", got)
	}
	if atomic.LoadInt32(&tokenRequests) == 0 {
		t.Error("Token issuer was never called")
	}
	// The custom transport must be traversed by both the API call and the
	// token fetch, since baseClient is injected via oauth2.HTTPClient.
	if atomic.LoadInt32(&transportHits) == 0 {
		t.Error("Custom transport was not used for the API request")
	}
}

func TestApiClientCreatedWithDefaultTelemetry(t *testing.T) {
	cfg := Configuration{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		ApiUrl:     "http://localhost:8080/",
	}
	_ = NewAPIClient(&cfg)

	telemetry1 := telemetry.Get(telemetry.TelemetryFactoryParameters{Configuration: cfg.Telemetry})
	telemetry2 := telemetry.Get(telemetry.TelemetryFactoryParameters{Configuration: cfg.Telemetry})

	if telemetry1 != telemetry2 {
		t.Fatalf("Telemetry instance should be the same")
	}
}

func TestApiClientWithCredentials(t *testing.T) {
	t.Run("ApiToken credentials should be applied with custom HTTPClient", func(t *testing.T) {
		// Not parallel: httpmock mutates the custom client's transport.
		customHTTPClient := &http.Client{
			Timeout: 30 * time.Second,
		}

		configuration, err := NewConfiguration(Configuration{
			ApiHost: "api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodApiToken,
				Config: &credentials.Config{
					ApiToken: "test-api-token",
				},
			},
			HTTPClient: customHTTPClient,
		})
		if err != nil {
			t.Fatalf("Failed to create configuration: %v", err)
		}

		apiClient := NewAPIClient(configuration)

		httpmock.ActivateNonDefault(customHTTPClient)
		defer httpmock.DeactivateAndReset()

		authHeaderReceived := false
		expectedAuthHeader := "Bearer test-api-token"

		httpmock.RegisterResponder("GET", fmt.Sprintf("%s/stores/%s/authorization-models", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2"),
			func(req *http.Request) (*http.Response, error) {
				authHeader := req.Header.Get("Authorization")
				if authHeader == expectedAuthHeader {
					authHeaderReceived = true
				}

				resp, err := httpmock.NewJsonResponse(200, ReadAuthorizationModelsResponse{
					AuthorizationModels: []AuthorizationModel{
						{
							Id:              "01GXSA8YR785C4FYS3C0RTG7B1",
							TypeDefinitions: []TypeDefinition{},
						},
					},
				})
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		_, _, err = apiClient.OpenFgaApi.ReadAuthorizationModels(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Execute()
		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		if !authHeaderReceived {
			t.Errorf("Authorization header was not sent when custom HTTPClient was provided")
		}

		// Verify that the custom HTTPClient is preserved (not replaced)
		if apiClient.GetConfig().HTTPClient != customHTTPClient {
			t.Error("Custom HTTPClient was replaced when it should have been preserved for ApiToken")
		}
	})

	t.Run("ClientCredentials should wrap custom HTTPClient", func(t *testing.T) {
		t.Parallel()
		customHTTPClient := &http.Client{
			Timeout: 30 * time.Second,
		}

		configuration, err := NewConfiguration(Configuration{
			ApiHost: "api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodClientCredentials,
				Config: &credentials.Config{
					ClientCredentialsClientId:       "test-client-id",
					ClientCredentialsClientSecret:   "test-client-secret",
					ClientCredentialsApiTokenIssuer: "https://issuer.example.com",
					ClientCredentialsApiAudience:    "https://api.example.com",
				},
			},
			HTTPClient: customHTTPClient,
		})
		if err != nil {
			t.Fatalf("Failed to create configuration: %v", err)
		}

		apiClient := NewAPIClient(configuration)

		// The HTTPClient should be replaced with an OAuth2-enabled client
		if configuration.HTTPClient == customHTTPClient {
			t.Error("HTTPClient was not replaced with OAuth2-enabled client when ClientCredentials were provided")
		}

		// The OAuth2 client should not be nil or default
		if configuration.HTTPClient == nil || configuration.HTTPClient == http.DefaultClient {
			t.Error("OAuth2 client should not be nil or default client")
		}

		// Verify the client was created successfully
		if apiClient == nil {
			t.Error("APIClient should not be nil")
		}

		// Verify that the OAuth2 transport actually wraps the custom client's transport
		if configuration.HTTPClient.Transport != nil {
			oauthTransport, ok := configuration.HTTPClient.Transport.(*oauth2.Transport)
			if !ok {
				t.Error("HTTPClient.Transport should be an *oauth2.Transport")
			} else {
				// customHTTPClient.Transport is nil, so Base should be nil or http.DefaultTransport
				if customHTTPClient.Transport == nil {
					// When custom transport is nil, OAuth2 uses http.DefaultTransport
					if oauthTransport.Base != nil && oauthTransport.Base != http.DefaultTransport {
						t.Errorf("OAuth2 transport Base should be nil or http.DefaultTransport when custom transport is nil, got: %T", oauthTransport.Base)
					}
				} else {
					// When custom transport is not nil, OAuth2 should wrap it
					if oauthTransport.Base != customHTTPClient.Transport {
						t.Errorf("OAuth2 transport Base does not match custom transport. Expected: %v, Got: %v", customHTTPClient.Transport, oauthTransport.Base)
					}
				}
			}
		}
	})

	t.Run("Credentials should work when HTTPClient is nil", func(t *testing.T) {
		// Not parallel: httpmock.Activate() swaps the global http.DefaultTransport.
		configuration, err := NewConfiguration(Configuration{
			ApiHost: "api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodApiToken,
				Config: &credentials.Config{
					ApiToken: "test-api-token",
				},
			},
		})
		if err != nil {
			t.Fatalf("Failed to create configuration: %v", err)
		}

		apiClient := NewAPIClient(configuration)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		authHeaderReceived := false
		expectedAuthHeader := "Bearer test-api-token"

		httpmock.RegisterResponder("GET", fmt.Sprintf("%s/stores/%s/authorization-models", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2"),
			func(req *http.Request) (*http.Response, error) {
				authHeader := req.Header.Get("Authorization")
				if authHeader == expectedAuthHeader {
					authHeaderReceived = true
				}

				resp, err := httpmock.NewJsonResponse(200, ReadAuthorizationModelsResponse{
					AuthorizationModels: []AuthorizationModel{
						{
							Id:              "01GXSA8YR785C4FYS3C0RTG7B1",
							TypeDefinitions: []TypeDefinition{},
						},
					},
				})
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		_, _, err = apiClient.OpenFgaApi.ReadAuthorizationModels(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Execute()
		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		if !authHeaderReceived {
			t.Errorf("Authorization header was not sent")
		}
	})

	t.Run("ApiToken with custom transport and authentication", func(t *testing.T) {
		t.Parallel()
		customClient := &http.Client{
			Timeout: 45 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		}

		configuration, err := NewConfiguration(Configuration{
			ApiUrl: "https://api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodApiToken,
				Config: &credentials.Config{
					ApiToken: "production-token",
				},
			},
			HTTPClient: customClient,
		})
		if err != nil {
			t.Fatalf("Failed to create configuration: %v", err)
		}

		apiClient := NewAPIClient(configuration)

		authHeader, exists := apiClient.GetConfig().DefaultHeaders["Authorization"]
		if !exists || authHeader != "Bearer production-token" {
			t.Error("Authorization header not properly set with custom HTTPClient")
		}

		// Verify that the custom HTTPClient with custom Transport is preserved
		if apiClient.GetConfig().HTTPClient != customClient {
			t.Error("Custom HTTPClient was not preserved")
		}
	})

	t.Run("ClientCredentials with custom transport settings wraps correctly", func(t *testing.T) {
		t.Parallel()
		// This test validates the fix for issue #234:
		// Custom HTTPClient should be wrapped by OAuth2 client, not replaced
		customTransport := &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		}

		customTimeout := 30 * time.Second
		customHTTPClient := &http.Client{
			Timeout:   customTimeout,
			Transport: customTransport,
		}

		configuration, err := NewConfiguration(Configuration{
			ApiHost: "api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodClientCredentials,
				Config: &credentials.Config{
					ClientCredentialsClientId:       "test-client-id",
					ClientCredentialsClientSecret:   "test-client-secret",
					ClientCredentialsApiTokenIssuer: "https://issuer.example.com/token",
					ClientCredentialsApiAudience:    "https://api.example.com",
				},
			},
			HTTPClient: customHTTPClient,
		})
		if err != nil {
			t.Fatalf("Failed to create configuration: %v", err)
		}

		apiClient := NewAPIClient(configuration)

		// The HTTPClient should be replaced with OAuth2-enabled client
		if configuration.HTTPClient == customHTTPClient {
			t.Error("HTTPClient was not replaced with OAuth2-enabled client when ClientCredentials were provided")
		}

		// The OAuth2 client should not be nil or default
		if configuration.HTTPClient == nil || configuration.HTTPClient == http.DefaultClient {
			t.Error("OAuth2 client should not be nil or default client")
		}

		// Verify the client was created successfully without warnings
		if apiClient == nil {
			t.Error("APIClient should not be nil")
		}

		// The key validation: OAuth2 client wraps the custom client
		// This means both credentials AND custom transport settings work together
		if apiClient.GetConfig().HTTPClient == customHTTPClient {
			t.Error("HTTPClient should be the OAuth2 wrapper, not the original custom client")
		}

		// Verify that the OAuth2 transport actually wraps the customTransport
		finalClient := apiClient.GetConfig().HTTPClient
		if finalClient.Transport != nil {
			oauthTransport, ok := finalClient.Transport.(*oauth2.Transport)
			if !ok {
				t.Errorf("HTTPClient.Transport should be an *oauth2.Transport, got: %T", finalClient.Transport)
			} else {
				// Verify that the Base transport matches the customTransport
				if oauthTransport.Base != customTransport {
					t.Errorf("OAuth2 transport Base does not match customTransport. Expected: %p, Got: %p", customTransport, oauthTransport.Base)
				}
			}
		} else {
			t.Error("HTTPClient.Transport should not be nil")
		}

		if finalClient.Timeout != customTimeout {
			t.Errorf("Custom HTTPClient Timeout was not preserved. Expected: %v, Got: %v", customTimeout, finalClient.Timeout)
		}

		if customHTTPClient.Transport != customTransport {
			t.Error("Original custom HTTPClient.Transport was mutated; it should be left untouched")
		}
	})

	t.Run("Custom HTTPClient preserved when no credentials provided", func(t *testing.T) {
		t.Parallel()
		customHTTPClient := &http.Client{
			Timeout: 20 * time.Second,
		}

		configuration, err := NewConfiguration(Configuration{
			ApiHost:    "api." + constants.SampleBaseDomain,
			HTTPClient: customHTTPClient,
		})
		if err != nil {
			t.Fatalf("Failed to create configuration: %v", err)
		}

		_ = NewAPIClient(configuration)

		// With no credentials, custom client should be preserved as-is
		if configuration.HTTPClient != customHTTPClient {
			t.Error("Custom HTTPClient should be preserved when no credentials provided")
		}
	})
}
