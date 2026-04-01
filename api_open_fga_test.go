package openfga

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"

	"github.com/openfga/go-sdk/credentials"
	"github.com/openfga/go-sdk/internal/constants"
	"github.com/openfga/go-sdk/internal/utils/retryutils"
)

type TestDefinition struct {
	Name           string
	JsonResponse   string
	ResponseStatus int
	Method         string
	RequestPath    string
}

func TestOpenFgaApiConfiguration(t *testing.T) {
	t.Run("Providing no store id should not error", func(t *testing.T) {
		_, err := NewConfiguration(Configuration{
			ApiHost: "api." + constants.SampleBaseDomain,
		})

		if err != nil {
			t.Fatalf("%v", err)
		}
	})

	t.Run("Providing no ApiHost should error", func(t *testing.T) {
		_, err := NewConfiguration(Configuration{})

		if err == nil {
			t.Fatalf("Expected an error when storeId is required but not provided")
		}
	})

	t.Run("ApiHost should be valid", func(t *testing.T) {
		_, err := NewConfiguration(Configuration{
			ApiHost: "https://api." + constants.SampleBaseDomain,
		})

		if err == nil {
			t.Fatalf("Expected an error when ApiHost is invalid (scheme is part of the host)")
		}
	})

	t.Run("RetryParams should be valid", func(t *testing.T) {
		tests := []struct {
			retryParams   *RetryParams
			expectedError bool
		}{
			{
				retryParams: &RetryParams{
					MaxRetry:    1,
					MinWaitInMs: 0,
				},
				expectedError: true,
			},
			{
				retryParams: &RetryParams{
					MaxRetry:    100,
					MinWaitInMs: 1,
				},
				expectedError: true,
			},
			{
				retryParams: &RetryParams{
					MaxRetry:    -1,
					MinWaitInMs: 1,
				},
				expectedError: true,
			},
			{
				retryParams: &RetryParams{
					MaxRetry:    1,
					MinWaitInMs: -1,
				},
				expectedError: true,
			},
			{
				retryParams: &RetryParams{
					MaxRetry:    1,
					MinWaitInMs: -1,
				},
				expectedError: true,
			},
			{
				retryParams: &RetryParams{
					MaxRetry:    1,
					MinWaitInMs: 1,
				},
				expectedError: false,
			},
			{
				retryParams: &RetryParams{
					MaxRetry:    0,
					MinWaitInMs: 1,
				},
				expectedError: false,
			},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("RetryParams: %v", *test.retryParams), func(t *testing.T) {
				config, err := NewConfiguration(Configuration{
					ApiUrl:      "https://api." + constants.SampleBaseDomain,
					RetryParams: test.retryParams,
				})

				if test.expectedError && err == nil {
					t.Fatalf("Expected an error when RetryParams are invalid, got none")
				}

				if !test.expectedError && err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}

				if !test.expectedError {
					if config.RetryParams == nil {
						t.Fatalf("Expected RetryParams on the config to be non-nil")
					}

					appliedRetryParams := *config.RetryParams
					if appliedRetryParams.MaxRetry != test.retryParams.MaxRetry {
						t.Fatalf("Expected MaxRetry to be %v, got %v", test.retryParams.MaxRetry, appliedRetryParams.MaxRetry)
					}

					if appliedRetryParams.MinWaitInMs != test.retryParams.MinWaitInMs {
						t.Fatalf("Expected MinWaitInMs to be %v, got %v", test.retryParams.MinWaitInMs, appliedRetryParams.MinWaitInMs)
					}

					appliedRetryParams = config.GetRetryParams()
					if appliedRetryParams.MaxRetry != test.retryParams.MaxRetry {
						t.Fatalf("Expected MaxRetry to be %v, got %v", test.retryParams.MaxRetry, appliedRetryParams.MaxRetry)
					}

					if appliedRetryParams.MinWaitInMs != test.retryParams.MinWaitInMs {
						t.Fatalf("Expected MinWaitInMs to be %v, got %v", test.retryParams.MinWaitInMs, appliedRetryParams.MinWaitInMs)
					}
				}
			})
		}
	})

	t.Run("RetryParams is default if not set", func(t *testing.T) {
		config, err := NewConfiguration(Configuration{
			ApiUrl: "https://api." + constants.SampleBaseDomain,
		})
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if config.RetryParams == nil {
			t.Fatalf("Expected RetryParams on the config to be non-nil")
		}

		appliedRetryParams := config.GetRetryParams()
		defaultParams := retryutils.GetRetryParamsOrDefault(nil)
		if appliedRetryParams.MaxRetry != defaultParams.MaxRetry {
			t.Fatalf("Expected MaxRetry to be %v, got %v", defaultParams.MaxRetry, appliedRetryParams.MaxRetry)
		}

		if appliedRetryParams.MinWaitInMs != defaultParams.MinWaitInMs {
			t.Fatalf("Expected MinWaitInMs to be %v, got %v", defaultParams.MinWaitInMs, appliedRetryParams.MinWaitInMs)
		}
	})

	t.Run("In ApiToken credential method, apiToken is required in the Credentials Config", func(t *testing.T) {
		_, err := NewConfiguration(Configuration{
			ApiHost: "https://api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodApiToken,
			},
		})

		if err == nil {
			t.Fatalf("Expected an error when apiToken is missing but the credential method is ApiToken")
		}
	})

	t.Run("should issue a successful network call when using ApiToken credential method", func(t *testing.T) {
		configuration, err := NewConfiguration(Configuration{
			ApiHost: "api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodApiToken,
				Config: &credentials.Config{
					ApiToken: "some-token",
				},
			},
		})
		if err != nil {
			t.Fatalf("%v", err)
		}

		apiClient := NewAPIClient(configuration)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s/stores/%s/authorization-models", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2"),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, ReadAuthorizationModelsResponse{AuthorizationModels: []AuthorizationModel{
					{
						Id:              "01GXSA8YR785C4FYS3C0RTG7B1",
						TypeDefinitions: []TypeDefinition{},
					},
					{
						Id:              "01GXSBM5PVYHCJNRNKXMB4QZTW",
						TypeDefinitions: []TypeDefinition{},
					},
				}})
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		if _, _, err = apiClient.OpenFgaApi.ReadAuthorizationModels(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Execute(); err != nil {
			t.Fatalf("%v", err)
		}
	})

	t.Run("In ClientCredentials method, providing no client id, secret or issuer should error", func(t *testing.T) {
		_, err := NewConfiguration(Configuration{
			ApiHost: "https://api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodApiToken,
				Config: &credentials.Config{
					ClientCredentialsClientSecret: "some-secret",
				},
			},
		})

		if err == nil {
			t.Fatalf("Expected an error: client id is required")
		}

		_, err = NewConfiguration(Configuration{
			ApiHost: "https://api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodApiToken,
				Config: &credentials.Config{
					ClientCredentialsClientId:       "some-id",
					ClientCredentialsApiTokenIssuer: "some-issuer",
					ClientCredentialsApiAudience:    "some-audience",
				},
			},
		})

		if err == nil {
			t.Fatalf("Expected an error: client secret is required")
		}

		_, err = NewConfiguration(Configuration{
			ApiHost: "https://api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodApiToken,
				Config: &credentials.Config{
					ClientCredentialsClientId:     "some-id",
					ClientCredentialsClientSecret: "some-secret",
					ClientCredentialsApiAudience:  "some-audience",
				},
			},
		})

		if err == nil {
			t.Fatalf("Expected an error: api token issuer is required")
		}

		_, err = NewConfiguration(Configuration{
			ApiHost: "api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodClientCredentials,
				Config: &credentials.Config{
					ClientCredentialsClientId:       "some-id",
					ClientCredentialsClientSecret:   "some-secret",
					ClientCredentialsApiAudience:    "some-audience",
					ClientCredentialsApiTokenIssuer: "some-issuer." + constants.SampleBaseDomain,
				},
			},
		})

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("NewCredentials should validate properly", func(t *testing.T) {
		// Passing valid credentials to NewCredentials should not error
		creds, err := credentials.NewCredentials(credentials.Credentials{
			Method: credentials.CredentialsMethodApiToken,
			Config: &credentials.Config{
				ApiToken: "some-token",
			},
		})

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if creds == nil {
			t.Fatalf("Expected creds to be non-nil")
		}

		if creds.Method != credentials.CredentialsMethodApiToken {
			t.Fatalf("Expected method to be %v, got %v", credentials.CredentialsMethodApiToken, creds.Method)
		}

		if creds.Config.ApiToken != "some-token" {
			t.Fatalf("Expected ApiToken to be %v, got %v", "some-token", creds.Config.ApiToken)
		}

		// Passing invalid credentials to NewCredentials should error
		_, err = credentials.NewCredentials(credentials.Credentials{
			Method: credentials.CredentialsMethodApiToken,
			Config: &credentials.Config{
				ClientCredentialsClientSecret: "some-secret",
			},
		})

		if err == nil {
			t.Fatalf("Expected validation error")
		}
	})

	clientCredentialsFirstRequestTest := func(t *testing.T, config Configuration, expectedTokenEndpoint string) {
		configuration, err := NewConfiguration(config)
		if err != nil {
			t.Fatalf("%v", err)
		}

		apiClient := NewAPIClient(configuration)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s/stores/%s/authorization-models", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2"),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, ReadAuthorizationModelsResponse{AuthorizationModels: []AuthorizationModel{
					{
						Id:              "01GXSA8YR785C4FYS3C0RTG7B1",
						TypeDefinitions: []TypeDefinition{},
					},
					{
						Id:              "01GXSBM5PVYHCJNRNKXMB4QZTW",
						TypeDefinitions: []TypeDefinition{},
					},
				}})
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		httpmock.RegisterResponder("POST", expectedTokenEndpoint,
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, struct {
					AccessToken string `json:"access_token"`
				}{AccessToken: "abcde"})
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		if _, _, err = apiClient.OpenFgaApi.ReadAuthorizationModels(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Execute(); err != nil {
			t.Fatalf("%v", err)
		}

		info := httpmock.GetCallCountInfo()
		numCalls := info[fmt.Sprintf("POST %s", expectedTokenEndpoint)]
		if numCalls != 1 {
			t.Fatalf("Expected call to get access token to be made exactly once, saw: %d", numCalls)
		}
		numCalls = info[fmt.Sprintf("GET %s/stores/%s/authorization-models", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2")]
		if numCalls != 1 {
			t.Fatalf("Expected call to get authorization models to be made exactly once, saw: %d", numCalls)
		}
	}

	tokenIssuers := map[string]string{
		"issuer." + constants.SampleBaseDomain:                                 "https://issuer." + constants.SampleBaseDomain + "/oauth/token",
		"https://issuer." + constants.SampleBaseDomain:                         "https://issuer." + constants.SampleBaseDomain + "/oauth/token",
		"https://issuer." + constants.SampleBaseDomain + "/":                   "https://issuer." + constants.SampleBaseDomain + "/oauth/token",
		"https://issuer." + constants.SampleBaseDomain + ":8080":               "https://issuer." + constants.SampleBaseDomain + ":8080/oauth/token",
		"https://issuer." + constants.SampleBaseDomain + ":8080/":              "https://issuer." + constants.SampleBaseDomain + ":8080/oauth/token",
		"issuer." + constants.SampleBaseDomain + "/some_endpoint":              "https://issuer." + constants.SampleBaseDomain + "/some_endpoint",
		"https://issuer." + constants.SampleBaseDomain + "/some_endpoint":      "https://issuer." + constants.SampleBaseDomain + "/some_endpoint",
		"https://issuer." + constants.SampleBaseDomain + ":8080/some_endpoint": "https://issuer." + constants.SampleBaseDomain + ":8080/some_endpoint",
	}

	for tokenIssuer, expectedTokenURL := range tokenIssuers {
		t.Run("should issue a network call to get the token at the first request if client id is provided", func(t *testing.T) {
			t.Run("with Auth0 configuration", func(t *testing.T) {
				clientCredentialsFirstRequestTest(t, Configuration{
					ApiUrl: "http://api." + constants.SampleBaseDomain,
					Credentials: &credentials.Credentials{
						Method: credentials.CredentialsMethodClientCredentials,
						Config: &credentials.Config{
							ClientCredentialsClientId:       "some-id",
							ClientCredentialsClientSecret:   "some-secret",
							ClientCredentialsApiAudience:    "some-audience",
							ClientCredentialsApiTokenIssuer: tokenIssuer,
						},
					},
				}, expectedTokenURL)
			})
			t.Run("with OAuth2 configuration", func(t *testing.T) {
				clientCredentialsFirstRequestTest(t, Configuration{
					ApiUrl: "http://api." + constants.SampleBaseDomain,
					Credentials: &credentials.Credentials{
						Method: credentials.CredentialsMethodClientCredentials,
						Config: &credentials.Config{
							ClientCredentialsClientId:       "some-id",
							ClientCredentialsClientSecret:   "some-secret",
							ClientCredentialsScopes:         "scope1 scope2",
							ClientCredentialsApiTokenIssuer: tokenIssuer,
						},
					},
				}, expectedTokenURL)
			})
		})
	}
	t.Run("should not issue a network call to get the token at the first request if the clientId is not provided", func(t *testing.T) {
		configuration, err := NewConfiguration(Configuration{
			ApiHost: "api." + constants.SampleBaseDomain,
			Credentials: &credentials.Credentials{
				Method: credentials.CredentialsMethodNone,
				Config: &credentials.Config{ClientCredentialsApiTokenIssuer: "tokenissuer.api.example"},
			},
		})
		if err != nil {
			t.Fatalf("%v", err)
		}
		configuration.ApiHost = "api." + constants.SampleBaseDomain

		apiClient := NewAPIClient(configuration)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s/stores/%s/authorization-models", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2"),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(200, ReadAuthorizationModelsResponse{AuthorizationModels: []AuthorizationModel{
					{
						Id:              "01GXSA8YR785C4FYS3C0RTG7B1",
						TypeDefinitions: []TypeDefinition{},
					},
					{
						Id:              "01GXSBM5PVYHCJNRNKXMB4QZTW",
						TypeDefinitions: []TypeDefinition{},
					},
				}})
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		if _, _, err = apiClient.OpenFgaApi.ReadAuthorizationModels(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Execute(); err != nil {
			t.Fatalf("%v", err)
		}

		info := httpmock.GetCallCountInfo()
		numCalls := info[fmt.Sprintf("POST https://%s/oauth/token", configuration.Credentials.Config.ClientCredentialsApiTokenIssuer)]
		if numCalls != 0 {
			t.Fatalf("Unexpected call to get access token made. Expected 0, saw: %d", numCalls)
		}
		numCalls = info[fmt.Sprintf("GET %s/stores/%s/authorization-models", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2")]
		if numCalls != 1 {
			t.Fatalf("Expected call to get authorization models to be made exactly once, saw: %d", numCalls)
		}
	})
}

func TestOpenFgaApi(t *testing.T) {
	configuration, err := NewConfiguration(Configuration{
		ApiHost: "api." + constants.SampleBaseDomain,
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	apiClient := NewAPIClient(configuration)

	t.Run("ReadAuthorizationModels", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadAuthorizationModels",
			JsonResponse:   `{"authorization_models":[{"id":"01GXSA8YR785C4FYS3C0RTG7B1","schema_version":"1.1","type_definitions":[]}]}`,
			ResponseStatus: 200,
			Method:         "GET",
			RequestPath:    "authorization-models",
		}

		var expectedResponse ReadAuthorizationModelsResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		got, response, err := apiClient.OpenFgaApi.ReadAuthorizationModels(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		if len(got.AuthorizationModels) != 1 {
			t.Fatalf("%v", err)
		}

		if got.AuthorizationModels[0].Id != expectedResponse.AuthorizationModels[0].Id {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, got.AuthorizationModels[0].Id, expectedResponse.AuthorizationModels[0].Id)
		}
	})

	t.Run("WriteAuthorizationModel", func(t *testing.T) {
		test := TestDefinition{
			Name:           "WriteAuthorizationModel",
			JsonResponse:   `{"authorization_model_id":"01GXSA8YR785C4FYS3C0RTG7B1"}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "authorization-models",
		}
		requestBody := WriteAuthorizationModelRequest{
			TypeDefinitions: []TypeDefinition{{
				Type: "github-repo",
				Relations: &map[string]Userset{
					"repo_writer": {
						This: &map[string]interface{}{},
					},
					"viewer": {Union: &Usersets{
						Child: []Userset{
							{This: &map[string]interface{}{}},
							{ComputedUserset: &ObjectRelation{
								Object:   PtrString(""),
								Relation: PtrString("repo_writer"),
							}},
						},
					}},
				},
			}},
		}

		var expectedResponse WriteAuthorizationModelResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.OpenFgaApi.WriteAuthorizationModel(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		_, err = got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}
	})

	t.Run("ReadAuthorizationModel", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadAuthorizationModel",
			JsonResponse:   `{"authorization_model":{"id":"01GXSA8YR785C4FYS3C0RTG7B1", "schema_version":"1.1", "type_definitions":[{"type":"github-repo", "relations":{"viewer":{"this":{}}}}]}}`,
			ResponseStatus: 200,
			Method:         "GET",
			RequestPath:    "authorization-models",
		}

		var expectedResponse ReadAuthorizationModelResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}
		modelId := expectedResponse.AuthorizationModel.Id

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath, modelId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.OpenFgaApi.ReadAuthorizationModel(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2", modelId).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if got.AuthorizationModel.Id != modelId {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("Check", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: CheckRequestTupleKey{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "viewer",
				Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
			},
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.OpenFgaApi.Check(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if *got.Allowed != *expectedResponse.Allowed {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("Write (Write Tuple)", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "write",
		}
		requestBody := WriteRequest{
			Writes: &WriteRequestWrites{
				TupleKeys: []TupleKey{{
					User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
					Relation: "viewer",
					Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				}},
			},
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, response, err := apiClient.OpenFgaApi.Write(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}
	})

	t.Run("Write (Delete Tuple)", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "write",
		}

		requestBody := WriteRequest{
			Deletes: &WriteRequestDeletes{
				TupleKeys: []TupleKeyWithoutCondition{{
					User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
					Relation: "viewer",
					Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				}},
			},
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, response, err := apiClient.OpenFgaApi.Write(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}
	})

	t.Run("Write (Write Tuple with OnDuplicate ignore)", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "write",
		}
		onDuplicateIgnore := "ignore"
		requestBody := WriteRequest{
			Writes: &WriteRequestWrites{
				TupleKeys: []TupleKey{{
					User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
					Relation: "viewer",
					Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				}},
				OnDuplicate: &onDuplicateIgnore,
			},
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				// Verify the request body contains the OnDuplicate field
				var body WriteRequest
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Errorf("Failed to decode request body: %v", err)
				}
				if body.Writes.OnDuplicate == nil || *body.Writes.OnDuplicate != "ignore" {
					t.Errorf("Expected OnDuplicate to be 'ignore', got %v", body.Writes.OnDuplicate)
				}

				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, response, err := apiClient.OpenFgaApi.Write(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}
	})

	t.Run("Write (Write Tuple with OnDuplicate error)", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "write",
		}
		onDuplicateError := "error"
		requestBody := WriteRequest{
			Writes: &WriteRequestWrites{
				TupleKeys: []TupleKey{{
					User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
					Relation: "viewer",
					Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				}},
				OnDuplicate: &onDuplicateError,
			},
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				// Verify the request body contains the OnDuplicate field
				var body WriteRequest
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Errorf("Failed to decode request body: %v", err)
				}
				if body.Writes.OnDuplicate == nil || *body.Writes.OnDuplicate != "error" {
					t.Errorf("Expected OnDuplicate to be 'error', got %v", body.Writes.OnDuplicate)
				}

				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, response, err := apiClient.OpenFgaApi.Write(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}
	})

	t.Run("Write (Delete Tuple with OnMissing ignore)", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "write",
		}
		onMissingIgnore := "ignore"
		requestBody := WriteRequest{
			Deletes: &WriteRequestDeletes{
				TupleKeys: []TupleKeyWithoutCondition{{
					User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
					Relation: "viewer",
					Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				}},
				OnMissing: &onMissingIgnore,
			},
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				// Verify the request body contains the OnMissing field
				var body WriteRequest
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Errorf("Failed to decode request body: %v", err)
				}
				if body.Deletes.OnMissing == nil || *body.Deletes.OnMissing != "ignore" {
					t.Errorf("Expected OnMissing to be 'ignore', got %v", body.Deletes.OnMissing)
				}

				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, response, err := apiClient.OpenFgaApi.Write(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}
	})

	t.Run("Write (Delete Tuple with OnMissing error)", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "write",
		}
		onMissingError := "error"
		requestBody := WriteRequest{
			Deletes: &WriteRequestDeletes{
				TupleKeys: []TupleKeyWithoutCondition{{
					User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
					Relation: "viewer",
					Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				}},
				OnMissing: &onMissingError,
			},
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				// Verify the request body contains the OnMissing field
				var body WriteRequest
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Errorf("Failed to decode request body: %v", err)
				}
				if body.Deletes.OnMissing == nil || *body.Deletes.OnMissing != "error" {
					t.Errorf("Expected OnMissing to be 'error', got %v", body.Deletes.OnMissing)
				}

				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, response, err := apiClient.OpenFgaApi.Write(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}
	})

	t.Run("Write (Mixed writes and deletes with conflict options)", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "write",
		}
		onDuplicateIgnore := "ignore"
		onMissingIgnore := "ignore"
		requestBody := WriteRequest{
			Writes: &WriteRequestWrites{
				TupleKeys: []TupleKey{{
					User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
					Relation: "viewer",
					Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				}},
				OnDuplicate: &onDuplicateIgnore,
			},
			Deletes: &WriteRequestDeletes{
				TupleKeys: []TupleKeyWithoutCondition{{
					User:     "user:another-user",
					Relation: "viewer",
					Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				}},
				OnMissing: &onMissingIgnore,
			},
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				// Verify the request body contains both OnDuplicate and OnMissing fields
				var body WriteRequest
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Errorf("Failed to decode request body: %v", err)
				}
				if body.Writes.OnDuplicate == nil || *body.Writes.OnDuplicate != "ignore" {
					t.Errorf("Expected OnDuplicate to be 'ignore', got %v", body.Writes.OnDuplicate)
				}
				if body.Deletes.OnMissing == nil || *body.Deletes.OnMissing != "ignore" {
					t.Errorf("Expected OnMissing to be 'ignore', got %v", body.Deletes.OnMissing)
				}

				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, response, err := apiClient.OpenFgaApi.Write(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}
	})

	t.Run("Expand", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Expand",
			JsonResponse:   `{"tree":{"root":{"name":"document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a#viewer","union":{"nodes":[{"name": "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a#viewer","leaf":{"users":{"users":["user:81684243-9356-4421-8fbf-a4f8d36aa31b"]}}}]}}}}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "expand",
		}

		requestBody := ExpandRequest{
			TupleKey: ExpandRequestTupleKey{
				Relation: "viewer",
				Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
			},
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse ExpandResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.OpenFgaApi.Expand(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		_, err = got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}
	})

	t.Run("Read", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Read",
			JsonResponse:   `{"tuples":[{"key":{"user":"user:81684243-9356-4421-8fbf-a4f8d36aa31b","relation":"viewer","object":"document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a"},"timestamp": "2000-01-01T00:00:00Z"}]}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "read",
		}

		requestBody := ReadRequest{
			TupleKey: &ReadRequestTupleKey{
				User:     PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: PtrString("viewer"),
				Object:   PtrString("document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a"),
			},
		}

		var expectedResponse ReadResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.OpenFgaApi.Read(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if len(got.Tuples) != len(expectedResponse.Tuples) {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("ReadChanges", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadChanges",
			JsonResponse:   `{"changes":[{"tuple_key":{"user":"user:81684243-9356-4421-8fbf-a4f8d36aa31b","relation":"viewer","object":"document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a"},"operation":"TUPLE_OPERATION_WRITE","timestamp": "2000-01-01T00:00:00Z"}],"continuation_token":"eyJwayI6IkxBVEVTVF9OU0NPTkZJR19hdXRoMHN0b3JlIiwic2siOiIxem1qbXF3MWZLZExTcUoyN01MdTdqTjh0cWgifQ=="}`,
			ResponseStatus: 200,
			Method:         "GET",
			RequestPath:    "changes",
		}

		var expectedResponse ReadChangesResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		startTime, _ := time.Parse(time.RFC3339, "2022-01-01T00:00:00Z")
		got, response, err := apiClient.OpenFgaApi.ReadChanges(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").
			Type_("repo").
			PageSize(25).
			StartTime(startTime).
			ContinuationToken("eyJwayI6IkxBVEVTVF9OU0NPTkZJR19hdXRoMHN0b3JlIiwic2siOiIxem1qbXF3MWZLZExTcUoyN01MdTdqTjh0cWgifQ==").
			Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if len(got.Changes) != len(expectedResponse.Changes) {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("ListObjects", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ListObjects",
			JsonResponse:   `{"objects":["document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a"]}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "list-objects",
		}

		requestBody := ListObjectsRequest{
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
			User:                 "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
			Relation:             "can_read",
			Type:                 "document",
			ContextualTuples: &ContextualTupleKeys{
				TupleKeys: []TupleKey{{
					User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
					Relation: "editor",
					Object:   "folder:product",
				}, {
					User:     "folder:product",
					Relation: "parent",
					Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				}},
			},
		}

		var expectedResponse ListObjectsResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.OpenFgaApi.ListObjects(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").
			Body(requestBody).
			Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if len(got.Objects) != len(expectedResponse.Objects) || (got.Objects)[0] != (expectedResponse.Objects)[0] {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("ListUsers", func(t *testing.T) {
		test := TestDefinition{
			Name: "ListUsers",
			// A real API would not return all these for the filter provided, these are just for test purposes
			JsonResponse:   `{"users":[{"object":{"id":"81684243-9356-4421-8fbf-a4f8d36aa31b","type":"user"}},{"userset":{"id":"fga","relation":"member","type":"team"}},{"wildcard":{"type":"user"}}]}`,
			ResponseStatus: http.StatusOK,
			Method:         http.MethodPost,
			RequestPath:    "list-users",
		}

		requestBody := ListUsersRequest{
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
			Object: FgaObject{
				Type: "document",
				Id:   "roadmap",
			},
			Relation: "can_read",
			// API does not allow sending multiple filters, done here for testing purposes
			UserFilters: []UserTypeFilter{{
				Type: "user",
			}, {
				Type:     "team",
				Relation: PtrString("member"),
			}},
			ContextualTuples: &[]TupleKey{{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "editor",
				Object:   "folder:product",
			}, {
				User:     "folder:product",
				Relation: "parent",
				Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
			}},
			Context: &map[string]interface{}{"ViewCount": 100},
		}

		var expectedResponse ListUsersResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := apiClient.OpenFgaApi.ListUsers(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").
			Body(requestBody).
			Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		_, err = got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if len(got.Users) != len(expectedResponse.Users) {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, got.GetUsers(), expectedResponse.GetUsers())
		}

		if got.Users[0].GetObject().Type != expectedResponse.Users[0].GetObject().Type || got.Users[0].GetObject().Id != expectedResponse.Users[0].GetObject().Id {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v (%v)", test.Name, got.Users[0], expectedResponse.Users[0], "object: { type: \"user\", id: \"81684243-9356-4421-8fbf-a4f8d36aa31b\" }")
		}

		if got.Users[1].GetUserset().Type != expectedResponse.Users[1].GetUserset().Type || got.Users[1].GetUserset().Id != expectedResponse.Users[1].GetUserset().Id || got.Users[1].GetUserset().Relation != expectedResponse.Users[1].GetUserset().Relation {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v (%v)", test.Name, got.Users[1], expectedResponse.Users[1], "wildcard: { type: \"team\", id: \"fga\", relation: \"member\" }")
		}

		if got.Users[2].GetWildcard().Type != expectedResponse.Users[2].GetWildcard().Type {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v (%v)", test.Name, got.Users[2], expectedResponse.Users[2], "wildcard: { type: \"user\" }")
		}
	})

	t.Run("Check with 400 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 400,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: CheckRequestTupleKey{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "viewer",
				Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		storeId := "01GXSB9YR785C4FYS3C0RTG7B2"
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, storeId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				errObj := ErrorResponse{
					Code:    "validation_error",
					Message: "Foo",
				}
				return httpmock.NewJsonResponse(400, errObj)
			},
		)
		_, _, err := apiClient.OpenFgaApi.Check(context.Background(), storeId).Body(requestBody).Execute()
		if err == nil {
			t.Fatalf("Expected error with 400 request but there is none")
		}
		validationError, ok := err.(FgaApiValidationError)
		if !ok {
			t.Fatalf("Expected validation Error but type is incorrect %v", err)
		}
		// Do some basic validation of the error itself

		if validationError.StoreId() != storeId {
			t.Fatalf("Expected store id to be %s but actual %s", storeId, validationError.StoreId())
		}

		if validationError.EndpointCategory() != "Check" {
			t.Fatalf("Expected category to be Check but actual %s", validationError.EndpointCategory())
		}

		if validationError.RequestMethod() != "POST" {
			t.Fatalf("Expected category to be POST but actual %s", validationError.RequestMethod())
		}

		if validationError.ResponseStatusCode() != 400 {
			t.Fatalf("Expected status code to be 400 but actual %d", validationError.ResponseStatusCode())
		}

		if validationError.ResponseCode() != ERRORCODE_VALIDATION_ERROR {
			t.Fatalf("Expected response code to be ERRORCODE_VALIDATION_ERROR but actual %s", validationError.ResponseCode())
		}
	})

	t.Run("Check with 401 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 401,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: CheckRequestTupleKey{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "viewer",
				Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		storeId := "01GXSB9YR785C4FYS3C0RTG7B2"
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, storeId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				errObj := ErrorResponse{
					Code:    "auth_failure",
					Message: "Foo",
				}
				return httpmock.NewJsonResponse(401, errObj)
			},
		)
		_, _, err := apiClient.OpenFgaApi.Check(context.Background(), storeId).Body(requestBody).Execute()
		if err == nil {
			t.Fatalf("Expected error with 401 request but there is none")
		}
		authenticationError, ok := err.(FgaApiAuthenticationError)
		if !ok {
			t.Fatalf("Expected authentication Error but type is incorrect %v", err)
		}
		// Do some basic validation of the error itself

		if authenticationError.StoreId() != storeId {
			t.Fatalf("Expected store id to be %s but actual %s", storeId, authenticationError.StoreId())
		}

		if authenticationError.EndpointCategory() != "Check" {
			t.Fatalf("Expected category to be Check but actual %s", authenticationError.EndpointCategory())
		}

		if authenticationError.ResponseStatusCode() != 401 {
			t.Fatalf("Expected status code to be 401 but actual %d", authenticationError.ResponseStatusCode())
		}

	})

	t.Run("Check with 404 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 404,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: CheckRequestTupleKey{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "viewer",
				Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		storeId := "01GXSB9YR785C4FYS3C0RTG7B2"
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, storeId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				errObj := ErrorResponse{
					Code:    "undefined_endpoint",
					Message: "Foo",
				}
				return httpmock.NewJsonResponse(404, errObj)
			},
		)
		_, _, err := apiClient.OpenFgaApi.Check(context.Background(), storeId).Body(requestBody).Execute()
		if err == nil {
			t.Fatalf("Expected error with 404 request but there is none")
		}
		notFoundError, ok := err.(FgaApiNotFoundError)
		if !ok {
			t.Fatalf("Expected not found Error but type is incorrect %v", err)
		}
		// Do some basic validation of the error itself

		if notFoundError.StoreId() != storeId {
			t.Fatalf("Expected store id to be %s but actual %s", storeId, notFoundError.StoreId())
		}

		if notFoundError.EndpointCategory() != "Check" {
			t.Fatalf("Expected category to be Check but actual %s", notFoundError.EndpointCategory())
		}

		if notFoundError.RequestMethod() != "POST" {
			t.Fatalf("Expected category to be POST but actual %s", notFoundError.RequestMethod())
		}

		if notFoundError.ResponseStatusCode() != 404 {
			t.Fatalf("Expected status code to be 404 but actual %d", notFoundError.ResponseStatusCode())
		}

		if notFoundError.ResponseCode() != NOTFOUNDERRORCODE_UNDEFINED_ENDPOINT {
			t.Fatalf("Expected response code to be NOTFOUNDERRORCODE_UNDEFINED_ENDPOINT but actual %s", notFoundError.ResponseCode())
		}
	})

	t.Run("Check with 429 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 429,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: CheckRequestTupleKey{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "viewer",
				Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		storeId := "01GXSB9YR785C4FYS3C0RTG7B2"
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, storeId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				errObj := ErrorResponse{
					Code:    "rate_limit_exceeded",
					Message: "Foo",
				}
				return httpmock.NewJsonResponse(429, errObj)
			},
		)

		updatedConfiguration, err := NewConfiguration(Configuration{
			ApiHost: "api." + constants.SampleBaseDomain,
			RetryParams: &RetryParams{
				MaxRetry:    3,
				MinWaitInMs: 5,
			},
		})
		if err != nil {
			t.Fatalf("%v", err)
		}

		updatedApiClient := NewAPIClient(updatedConfiguration)

		_, _, err = updatedApiClient.OpenFgaApi.Check(context.Background(), storeId).Body(requestBody).Execute()
		if err == nil {
			t.Fatalf("Expected error with 429 request but there is none")
		}
		rateLimitError, ok := err.(FgaApiRateLimitExceededError)
		if !ok {
			t.Fatalf("Expected rate limit exceeded Error but type is incorrect %v", err)
		}
		// Do some basic validation of the error itself

		if rateLimitError.StoreId() != storeId {
			t.Fatalf("Expected store id to be %s but actual %s", storeId, rateLimitError.StoreId())
		}

		if rateLimitError.EndpointCategory() != "Check" {
			t.Fatalf("Expected category to be Check but actual %s", rateLimitError.EndpointCategory())
		}

		if rateLimitError.ResponseStatusCode() != 429 {
			t.Fatalf("Expected status code to be 429 but actual %d", rateLimitError.ResponseStatusCode())
		}

	})

	t.Run("Check with initial 429 but eventually resolved", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: CheckRequestTupleKey{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "viewer",
				Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		firstMock := httpmock.NewStringResponder(429, "")
		secondMock, _ := httpmock.NewJsonResponder(200, expectedResponse)
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			firstMock.Then(firstMock).Then(firstMock).Then(secondMock),
		)
		updatedConfiguration, err := NewConfiguration(Configuration{
			ApiHost: "api." + constants.SampleBaseDomain,
			RetryParams: &RetryParams{
				MaxRetry:    3,
				MinWaitInMs: 5,
			},
		})
		if err != nil {
			t.Fatalf("%v", err)
		}

		updatedApiClient := NewAPIClient(updatedConfiguration)

		got, response, err := updatedApiClient.OpenFgaApi.Check(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()

		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if *got.Allowed != *expectedResponse.Allowed {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("Check with initial 429 but eventually resolved with default config", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 200,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: CheckRequestTupleKey{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "viewer",
				Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		firstMock := httpmock.NewStringResponder(429, "")
		secondMock, _ := httpmock.NewJsonResponder(200, expectedResponse)
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, "01GXSB9YR785C4FYS3C0RTG7B2", test.RequestPath),
			firstMock.Then(secondMock),
		)

		got, response, err := apiClient.OpenFgaApi.Check(context.Background(), "01GXSB9YR785C4FYS3C0RTG7B2").Body(requestBody).Execute()

		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if *got.Allowed != *expectedResponse.Allowed {
			t.Fatalf("OpenFga%v().Execute() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("Check with 500 error", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: 500,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := CheckRequest{
			TupleKey: CheckRequestTupleKey{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "viewer",
				Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
			},
		}

		var expectedResponse CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		storeId := "01GXSB9YR785C4FYS3C0RTG7B2"
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s/stores/%s/%s", configuration.ApiUrl, storeId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				errObj := ErrorResponse{
					Code:    "internal_error",
					Message: "Foo",
				}
				return httpmock.NewJsonResponse(500, errObj)
			},
		)
		_, _, err := apiClient.OpenFgaApi.Check(context.Background(), storeId).Body(requestBody).Execute()
		if err == nil {
			t.Fatalf("Expected error with 500 request but there is none")
		}
		internalError, ok := err.(FgaApiInternalError)
		if !ok {
			t.Fatalf("Expected internal Error but type is incorrect %v", err)
		}
		// Do some basic validation of the error itself

		if internalError.StoreId() != storeId {
			t.Fatalf("Expected store id to be %s but actual %s", storeId, internalError.StoreId())
		}

		if internalError.EndpointCategory() != "Check" {
			t.Fatalf("Expected category to be Check but actual %s", internalError.EndpointCategory())
		}

		if internalError.RequestMethod() != "POST" {
			t.Fatalf("Expected category to be POST but actual %s", internalError.RequestMethod())
		}

		if internalError.ResponseStatusCode() != 500 {
			t.Fatalf("Expected status code to be 500 but actual %d", internalError.ResponseStatusCode())
		}

		if internalError.ResponseCode() != INTERNALERRORCODE_INTERNAL_ERROR {
			t.Fatalf("Expected response code to be INTERNALERRORCODE_INTERNAL_ERROR but actual %s", internalError.ResponseCode())
		}
	})

	t.Run("Retry on 500 error", func(t *testing.T) {
		storeID := "01H0H015178Y2V4CX10C2KGHF4"

		var attempts int32

		// First two attempts return 500, third succeeds.
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cur := atomic.AddInt32(&attempts, 1)
			w.Header().Set("Content-Type", "application/json")
			if cur < 3 {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"code":"internal_error","message":"transient"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"allowed":true}`))
		}))
		defer server.Close()

		cfg, err := NewConfiguration(Configuration{
			ApiUrl: server.URL,
			RetryParams: &RetryParams{ // allow enough retries (default max is 3, which is fine too)
				MaxRetry:    4,
				MinWaitInMs: 10, // keep test fast
			},
			HTTPClient: &http.Client{},
		})
		if err != nil {
			t.Fatalf("failed to create configuration: %v", err)
		}

		apiClient := NewAPIClient(cfg)

		resp, httpResp, reqErr := apiClient.OpenFgaApi.Check(context.Background(), storeID).
			Body(CheckRequest{TupleKey: CheckRequestTupleKey{User: "user:anne", Relation: "viewer", Object: "document:doc"}}).
			Execute()

		if reqErr != nil {
			t.Fatalf("expected eventual success after internal error retries, got: %v", reqErr)
		}
		if httpResp == nil || httpResp.StatusCode != http.StatusOK {
			t.Fatalf("expected final HTTP 200, got %+v", httpResp)
		}
		if resp.Allowed == nil || !*resp.Allowed {
			t.Fatalf("expected Allowed true in final response")
		}

		gotAttempts := int(atomic.LoadInt32(&attempts))
		if gotAttempts != 3 { // 1 initial + 2 retries
			t.Fatalf("expected 3 attempts (500 retried), got %d", gotAttempts)
		}
	})

	t.Run("Do not retry on 501 error", func(t *testing.T) {
		var attempts int32
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt32(&attempts, 1)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotImplemented) // 501
			_, _ = w.Write([]byte(`{"code":"not_implemented","message":"no retry"}`))
		}))
		defer server.Close()

		cfg, _ := NewConfiguration(Configuration{ApiUrl: server.URL, RetryParams: &RetryParams{MaxRetry: 4, MinWaitInMs: 10}, HTTPClient: &http.Client{}})
		apiClient := NewAPIClient(cfg)

		_, _, err := apiClient.OpenFgaApi.Check(context.Background(), "store").Body(CheckRequest{TupleKey: CheckRequestTupleKey{User: "u", Relation: "r", Object: "o"}}).Execute()
		if err == nil {
			t.Fatalf("expected error")
		}
		if got := int(atomic.LoadInt32(&attempts)); got != 1 {
			t.Fatalf("expected 1 attempt, got %d", got)
		}
	})

	t.Run("Retry on 429 error with Retry-After", func(t *testing.T) {
		storeID := "01H0H015178Y2V4CX10C2KGHF4"

		var attempts int32

		// We simulate two 429 responses providing Retry-After, then a success.
		retryAfterSeconds := 1.0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			current := atomic.AddInt32(&attempts, 1)
			w.Header().Set("Content-Type", "application/json")
			if current < 3 { // first two attempts fail with rate limit and Retry-After header
				w.Header().Set("Retry-After", "1") // 1 second
				w.WriteHeader(http.StatusTooManyRequests)
				_, _ = w.Write([]byte(`{"code":"rate_limit_exceeded","message":"Rate limit exceeded, retry after some time"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"allowed":true}`))
		}))
		defer server.Close()

		cfg, err := NewConfiguration(Configuration{
			ApiUrl: server.URL,
			RetryParams: &RetryParams{
				MaxRetry:    3,
				MinWaitInMs: 5, // this won't be used due to Retry-After headers
			},
			HTTPClient: &http.Client{},
		})
		if err != nil {
			t.Fatalf("failed to build configuration: %v", err)
		}

		apiClient := NewAPIClient(cfg)

		start := time.Now()
		resp, _, reqErr := apiClient.OpenFgaApi.Check(context.Background(), storeID).
			Body(CheckRequest{TupleKey: CheckRequestTupleKey{User: "user:anne", Relation: "viewer", Object: "document:doc"}}).
			Execute()
		elapsed := time.Since(start)

		if reqErr != nil {
			t.Fatalf("expected success after retries, got error: %v", reqErr)
		}
		if resp.Allowed == nil || !*resp.Allowed {
			t.Fatalf("expected Allowed true in final response")
		}

		gotAttempts := int(atomic.LoadInt32(&attempts))
		if gotAttempts != 3 { // 1 initial + 2 retries
			t.Fatalf("expected 3 attempts total, got %d", gotAttempts)
		}

		// 2xe Retry-After header (1s each) -> expect >= 2s total
		minExpected := time.Duration(2*retryAfterSeconds) * time.Second
		if elapsed < minExpected {
			t.Fatalf("expected elapsed >= %v due to Retry-After headers, got %v", minExpected, elapsed)
		}
	})
}

func TestStreamedListObjectsExecute(t *testing.T) {
	t.Parallel()

	configuration, err := NewConfiguration(Configuration{
		ApiUrl: constants.TestApiUrl,
	})
	if err != nil {
		t.Fatalf("failed to create configuration: %v", err)
	}
	apiClient := NewAPIClient(configuration)
	storeID := "01GXSB9YR785C4FYS3C0RTG7B2"

	t.Run("successful streaming with multiple objects", func(t *testing.T) {
		// Streaming response with multiple streamed objects
		responseBody := `{"result":{"object":"document:doc1"}}` + "\n" +
			`{"result":{"object":"document:doc2"}}` + "\n" +
			`{"result":{"object":"document:doc3"}}` + "\n"

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("POST", fmt.Sprintf("%s/stores/%s/streamed-list-objects", configuration.ApiUrl, storeID),
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, responseBody)
				resp.Header.Set("Content-Type", "application/x-ndjson")
				return resp, nil
			},
		)

		requestBody := ListObjectsRequest{
			AuthorizationModelId: PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
			User:                 "user:anne",
			Relation:             "viewer",
			Type:                 "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Body(requestBody).
			Execute()
		if err != nil {
			t.Fatalf("StreamedListObjects failed: %v", err)
		}
		defer channel.Close()

		var objects []string
		for obj := range channel.Objects {
			objects = append(objects, obj.Object)
		}

		// Check for errors
		select {
		case err := <-channel.Errors:
			if err != nil {
				t.Fatalf("unexpected stream error: %v", err)
			}
		default:
		}

		if len(objects) != 3 {
			t.Fatalf("expected 3 objects, got %d", len(objects))
		}
		if objects[0] != "document:doc1" || objects[1] != "document:doc2" || objects[2] != "document:doc3" {
			t.Fatalf("unexpected objects: %v", objects)
		}
	})

	t.Run("streaming with single object", func(t *testing.T) {
		responseBody := `{"result":{"object":"document:single"}}` + "\n"

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("POST", fmt.Sprintf("%s/stores/%s/streamed-list-objects", configuration.ApiUrl, storeID),
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, responseBody)
				resp.Header.Set("Content-Type", "application/x-ndjson")
				return resp, nil
			},
		)

		requestBody := ListObjectsRequest{
			User:     "user:bob",
			Relation: "editor",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Body(requestBody).
			Execute()
		if err != nil {
			t.Fatalf("StreamedListObjects failed: %v", err)
		}
		defer channel.Close()

		var objects []string
		for obj := range channel.Objects {
			objects = append(objects, obj.Object)
		}

		if len(objects) != 1 || objects[0] != "document:single" {
			t.Fatalf("expected [document:single], got %v", objects)
		}
	})

	t.Run("streaming with empty result", func(t *testing.T) {
		// Empty streaming response (no objects match)
		responseBody := ""

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("POST", fmt.Sprintf("%s/stores/%s/streamed-list-objects", configuration.ApiUrl, storeID),
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, responseBody)
				resp.Header.Set("Content-Type", "application/x-ndjson")
				return resp, nil
			},
		)

		requestBody := ListObjectsRequest{
			User:     "user:nobody",
			Relation: "viewer",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Body(requestBody).
			Execute()
		if err != nil {
			t.Fatalf("StreamedListObjects failed: %v", err)
		}
		defer channel.Close()

		var objects []string
		for obj := range channel.Objects {
			objects = append(objects, obj.Object)
		}

		if len(objects) != 0 {
			t.Fatalf("expected 0 objects, got %d", len(objects))
		}
	})

	t.Run("streaming with stream error", func(t *testing.T) {
		// Streaming response with an error in the stream
		responseBody := `{"result":{"object":"document:doc1"}}` + "\n" +
			`{"error":{"message":"Internal stream error occurred"}}` + "\n"

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("POST", fmt.Sprintf("%s/stores/%s/streamed-list-objects", configuration.ApiUrl, storeID),
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(200, responseBody)
				resp.Header.Set("Content-Type", "application/x-ndjson")
				return resp, nil
			},
		)

		requestBody := ListObjectsRequest{
			User:     "user:anne",
			Relation: "viewer",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Body(requestBody).
			Execute()
		if err != nil {
			t.Fatalf("StreamedListObjects failed: %v", err)
		}
		defer channel.Close()

		// First object should come through
		obj := <-channel.Objects
		if obj.Object != "document:doc1" {
			t.Fatalf("expected first object document:doc1, got %s", obj.Object)
		}

		// Then we should get an error
		streamErr := <-channel.Errors
		if streamErr == nil {
			t.Fatal("expected stream error, got nil")
		}
		if streamErr.Error() != "Internal stream error occurred" {
			t.Fatalf("expected 'Internal stream error occurred', got %v", streamErr)
		}
	})

	t.Run("HTTP error response", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("POST", fmt.Sprintf("%s/stores/%s/streamed-list-objects", configuration.ApiUrl, storeID),
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewJsonResponse(400, map[string]interface{}{
					"code":    "validation_error",
					"message": "Invalid request body",
				})
			},
		)

		requestBody := ListObjectsRequest{
			User:     "invalid",
			Relation: "viewer",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Body(requestBody).
			Execute()

		if err == nil {
			if channel != nil {
				channel.Close()
			}
			t.Fatal("expected error for 400 response, got nil")
		}
		if channel != nil {
			t.Fatal("expected nil channel on error")
		}
	})

	t.Run("missing store ID validation", func(t *testing.T) {
		requestBody := ListObjectsRequest{
			User:     "user:anne",
			Relation: "viewer",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), "").
			Body(requestBody).
			Execute()

		if err == nil {
			if channel != nil {
				channel.Close()
			}
			t.Fatal("expected error for missing store ID, got nil")
		}
	})

	t.Run("missing body validation", func(t *testing.T) {
		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Execute()

		if err == nil {
			if channel != nil {
				channel.Close()
			}
			t.Fatal("expected error for missing body, got nil")
		}
	})

	t.Run("context cancellation", func(t *testing.T) {
		// Create a slow response that will be cancelled
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("POST", fmt.Sprintf("%s/stores/%s/streamed-list-objects", configuration.ApiUrl, storeID),
			func(req *http.Request) (*http.Response, error) {
				// Return first result, then the response will be closed when context is cancelled
				responseBody := `{"result":{"object":"document:doc1"}}` + "\n"
				resp := httpmock.NewStringResponse(200, responseBody)
				resp.Header.Set("Content-Type", "application/x-ndjson")
				return resp, nil
			},
		)

		ctx, cancel := context.WithCancel(context.Background())

		requestBody := ListObjectsRequest{
			User:     "user:anne",
			Relation: "viewer",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(ctx, storeID).
			Body(requestBody).
			Execute()
		if err != nil {
			t.Fatalf("StreamedListObjects failed: %v", err)
		}

		// Read first object
		obj := <-channel.Objects
		if obj.Object != "document:doc1" {
			t.Fatalf("expected document:doc1, got %s", obj.Object)
		}

		// Cancel the context
		cancel()

		// Channel should close
		channel.Close()
	})

	t.Run("with custom headers", func(t *testing.T) {
		var capturedHeaders http.Header

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("POST", fmt.Sprintf("%s/stores/%s/streamed-list-objects", configuration.ApiUrl, storeID),
			func(req *http.Request) (*http.Response, error) {
				capturedHeaders = req.Header.Clone()
				responseBody := `{"result":{"object":"document:doc1"}}` + "\n"
				resp := httpmock.NewStringResponse(200, responseBody)
				resp.Header.Set("Content-Type", "application/x-ndjson")
				return resp, nil
			},
		)

		requestBody := ListObjectsRequest{
			User:     "user:anne",
			Relation: "viewer",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Body(requestBody).
			Options(StreamingRequestOptions{
				RequestOptions: RequestOptions{
					Headers: map[string]string{
						"X-Custom-Header": "custom-value",
						"X-Request-ID":    "req-123",
					},
				},
			}).
			Execute()
		if err != nil {
			t.Fatalf("StreamedListObjects failed: %v", err)
		}
		defer channel.Close()

		// Drain the channel
		for range channel.Objects {
		}

		if capturedHeaders.Get("X-Custom-Header") != "custom-value" {
			t.Fatalf("expected X-Custom-Header to be 'custom-value', got '%s'", capturedHeaders.Get("X-Custom-Header"))
		}
		if capturedHeaders.Get("X-Request-ID") != "req-123" {
			t.Fatalf("expected X-Request-ID to be 'req-123', got '%s'", capturedHeaders.Get("X-Request-ID"))
		}
	})

	t.Run("retry on 500 error", func(t *testing.T) {
		var attempts int32

		// First two attempts return 500, third succeeds with streaming response.
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cur := atomic.AddInt32(&attempts, 1)
			if cur < 3 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"code":"internal_error","message":"transient"}`))
				return
			}
			w.Header().Set("Content-Type", "application/x-ndjson")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"result":{"object":"document:retried"}}` + "\n"))
		}))
		defer server.Close()

		cfg, err := NewConfiguration(Configuration{
			ApiUrl:      server.URL,
			RetryParams: &RetryParams{MaxRetry: 4, MinWaitInMs: 1},
			HTTPClient:  &http.Client{},
		})
		if err != nil {
			t.Fatalf("failed to create configuration: %v", err)
		}
		apiClient := NewAPIClient(cfg)

		requestBody := ListObjectsRequest{
			User:     "user:anne",
			Relation: "viewer",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Body(requestBody).
			Execute()
		if err != nil {
			t.Fatalf("expected eventual success after retries, got: %v", err)
		}
		defer channel.Close()

		var objects []string
		for obj := range channel.Objects {
			objects = append(objects, obj.Object)
		}

		select {
		case err := <-channel.Errors:
			if err != nil {
				t.Fatalf("unexpected stream error: %v", err)
			}
		default:
		}

		gotAttempts := int(atomic.LoadInt32(&attempts))
		if gotAttempts != 3 {
			t.Fatalf("expected 3 attempts (2 x 500 + 1 success), got %d", gotAttempts)
		}
		if len(objects) != 1 || objects[0] != "document:retried" {
			t.Fatalf("expected [document:retried], got %v", objects)
		}
	})

	t.Run("retry on 429 rate limit error", func(t *testing.T) {
		var attempts int32

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cur := atomic.AddInt32(&attempts, 1)
			if cur < 2 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				_, _ = w.Write([]byte(`{"code":"rate_limit_exceeded","message":"Rate limit exceeded"}`))
				return
			}
			w.Header().Set("Content-Type", "application/x-ndjson")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"result":{"object":"document:after-rate-limit"}}` + "\n"))
		}))
		defer server.Close()

		cfg, err := NewConfiguration(Configuration{
			ApiUrl:      server.URL,
			RetryParams: &RetryParams{MaxRetry: 3, MinWaitInMs: 1},
			HTTPClient:  &http.Client{},
		})
		if err != nil {
			t.Fatalf("failed to create configuration: %v", err)
		}
		apiClient := NewAPIClient(cfg)

		requestBody := ListObjectsRequest{
			User:     "user:anne",
			Relation: "viewer",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Body(requestBody).
			Execute()
		if err != nil {
			t.Fatalf("expected eventual success after rate limit retry, got: %v", err)
		}
		defer channel.Close()

		var objects []string
		for obj := range channel.Objects {
			objects = append(objects, obj.Object)
		}

		gotAttempts := int(atomic.LoadInt32(&attempts))
		if gotAttempts != 2 {
			t.Fatalf("expected 2 attempts (1 x 429 + 1 success), got %d", gotAttempts)
		}
		if len(objects) != 1 || objects[0] != "document:after-rate-limit" {
			t.Fatalf("expected [document:after-rate-limit], got %v", objects)
		}
	})

	t.Run("retry on 400 validation error via default case", func(t *testing.T) {
		// Note: 400 errors fall through to determineRetry's default case, which retries
		// them just like the non-streaming Check endpoint does. This is consistent behavior.
		var attempts int32

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt32(&attempts, 1)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"code":"validation_error","message":"Invalid request"}`))
		}))
		defer server.Close()

		cfg, err := NewConfiguration(Configuration{
			ApiUrl:      server.URL,
			RetryParams: &RetryParams{MaxRetry: 2, MinWaitInMs: 1},
			HTTPClient:  &http.Client{},
		})
		if err != nil {
			t.Fatalf("failed to create configuration: %v", err)
		}
		apiClient := NewAPIClient(cfg)

		requestBody := ListObjectsRequest{
			User:     "user:anne",
			Relation: "viewer",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Body(requestBody).
			Execute()

		if err == nil {
			if channel != nil {
				channel.Close()
			}
			t.Fatal("expected error for persistent 400 responses, got nil")
		}

		gotAttempts := int(atomic.LoadInt32(&attempts))
		// 400 errors are retried via the default case in determineRetry (matching non-streaming behavior)
		if gotAttempts != 3 {
			t.Fatalf("expected 3 attempts (1 initial + 2 retries, matching non-streaming behavior), got %d", gotAttempts)
		}
	})

	t.Run("retry on transport error then succeed", func(t *testing.T) {
		var attempts int32

		// Start a real server for the successful attempt
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/x-ndjson")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"result":{"object":"document:recovered"}}` + "\n"))
		}))
		defer server.Close()

		// Use a custom round tripper that fails once then delegates to real transport
		transport := &http.Transport{}
		cfg, err := NewConfiguration(Configuration{
			ApiUrl:      server.URL,
			RetryParams: &RetryParams{MaxRetry: 3, MinWaitInMs: 1},
			HTTPClient: &http.Client{
				Transport: &testStreamRetryTransport{
					attempts:      &attempts,
					failUntil:     2,
					realTransport: transport,
				},
			},
		})
		if err != nil {
			t.Fatalf("failed to create configuration: %v", err)
		}
		apiClient := NewAPIClient(cfg)

		requestBody := ListObjectsRequest{
			User:     "user:anne",
			Relation: "viewer",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Body(requestBody).
			Execute()
		if err != nil {
			t.Fatalf("expected eventual success after transport retry, got: %v", err)
		}
		defer channel.Close()

		var objects []string
		for obj := range channel.Objects {
			objects = append(objects, obj.Object)
		}

		gotAttempts := int(atomic.LoadInt32(&attempts))
		if gotAttempts != 2 {
			t.Fatalf("expected 2 attempts (1 transport error + 1 success), got %d", gotAttempts)
		}
		if len(objects) != 1 || objects[0] != "document:recovered" {
			t.Fatalf("expected [document:recovered], got %v", objects)
		}
	})

	t.Run("exhausts retries on persistent 500", func(t *testing.T) {
		var attempts int32

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt32(&attempts, 1)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"code":"internal_error","message":"persistent failure"}`))
		}))
		defer server.Close()

		cfg, err := NewConfiguration(Configuration{
			ApiUrl:      server.URL,
			RetryParams: &RetryParams{MaxRetry: 2, MinWaitInMs: 1},
			HTTPClient:  &http.Client{},
		})
		if err != nil {
			t.Fatalf("failed to create configuration: %v", err)
		}
		apiClient := NewAPIClient(cfg)

		requestBody := ListObjectsRequest{
			User:     "user:anne",
			Relation: "viewer",
			Type:     "document",
		}

		channel, err := apiClient.OpenFgaApi.StreamedListObjects(context.Background(), storeID).
			Body(requestBody).
			Execute()

		if err == nil {
			if channel != nil {
				channel.Close()
			}
			t.Fatal("expected error after exhausting retries, got nil")
		}

		gotAttempts := int(atomic.LoadInt32(&attempts))
		if gotAttempts != 3 {
			t.Fatalf("expected 3 attempts (1 initial + 2 retries), got %d", gotAttempts)
		}
	})
}

// testStreamRetryTransport is a test helper that fails the first N requests with a transport error
// and then delegates to a real transport for subsequent requests.
type testStreamRetryTransport struct {
	attempts      *int32
	failUntil     int32
	realTransport http.RoundTripper
}

func (t *testStreamRetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	cur := atomic.AddInt32(t.attempts, 1)
	if cur < t.failUntil {
		return nil, fmt.Errorf("connection refused")
	}
	return t.realTransport.RoundTrip(req)
}
