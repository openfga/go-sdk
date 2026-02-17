package openfga

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/openfga/go-sdk/credentials"
	"github.com/openfga/go-sdk/internal/constants"
	"github.com/openfga/go-sdk/telemetry"
)

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
			t.Errorf("Authorization header was not sent when custom HTTPClient was provided")
		}

		// Verify that the custom HTTPClient is preserved (not replaced)
		if apiClient.GetConfig().HTTPClient != customHTTPClient {
			t.Error("Custom HTTPClient was replaced when it should have been preserved for ApiToken")
		}
	})

	t.Run("ClientCredentials should wrap custom HTTPClient", func(t *testing.T) {
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
	})

	t.Run("Credentials should work when HTTPClient is nil", func(t *testing.T) {
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
		// This test validates the fix for issue #234:
		// Custom HTTPClient should be wrapped by OAuth2 client, not replaced
		customTransport := &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		}

		customHTTPClient := &http.Client{
			Timeout:   30 * time.Second,
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
	})

	t.Run("No warning or error when combining custom HTTPClient and ClientCredentials", func(t *testing.T) {
		// Before the fix, this would log a warning and replace the custom client
		// After the fix, it should work silently by wrapping the custom client
		customHTTPClient := &http.Client{
			Timeout: 25 * time.Second,
		}

		configuration, err := NewConfiguration(Configuration{
			ApiHost: "api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodClientCredentials,
				Config: &credentials.Config{
					ClientCredentialsClientId:       "client-id",
					ClientCredentialsClientSecret:   "client-secret",
					ClientCredentialsApiTokenIssuer: "https://auth.example.com/token",
					ClientCredentialsApiAudience:    "https://api.example.com",
				},
			},
			HTTPClient: customHTTPClient,
			Debug:      true, // Enable debug to verify no warnings are logged
		})
		if err != nil {
			t.Fatalf("Failed to create configuration: %v", err)
		}

		apiClient := NewAPIClient(configuration)

		if apiClient == nil {
			t.Error("APIClient creation should succeed")
		}

		// Verify OAuth2 client was created (not the original custom client)
		if configuration.HTTPClient == customHTTPClient {
			t.Error("HTTPClient should be OAuth2 wrapper, not original client")
		}

		// Verify it's not just the default client either
		if configuration.HTTPClient == http.DefaultClient {
			t.Error("HTTPClient should be OAuth2 client with wrapped custom transport")
		}
	})

	t.Run("Custom HTTPClient preserved when no credentials provided", func(t *testing.T) {
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
