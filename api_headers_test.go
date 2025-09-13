package openfga_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	openfga "github.com/openfga/go-sdk"
)

// Test helpers and setup

// Constants to avoid duplication
const (
	apiDefaultHeaderName     = "Default-Header"
	apiDefaultHeaderValue    = "default-value"
	apiOverriddenValue       = "overridden-value"
	apiCustomHeaderName      = "X-Custom-Header"
	apiCustomHeaderValue     = "custom-value"
	apiTestUser              = "user:anne"
	apiTestRelation          = "viewer"
	apiTestObject            = "document:roadmap"
	apiTestStoreId           = "01H0H015178Y2V4CX10C2KGHF4"
	apiRequestFailedMsg      = "API request failed: %v"
	apiExpectedCustomMsg     = "Expected X-Custom-Header to be 'custom-value', got '%s'"
	apiExpectedOverriddenMsg = "Expected Default-Header to be overridden to 'overridden-value', got '%s'"
	apiExpectedDefaultMsg    = "Expected Default-Header to be 'default-value', got '%s'"
)

// createAPITestServer creates a test server that captures headers and returns appropriate responses
func createAPITestServer(t *testing.T, capturedHeaders *map[string]string, responseBody string) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*capturedHeaders = make(map[string]string)
		for name, values := range r.Header {
			if len(values) > 0 {
				(*capturedHeaders)[name] = values[0]
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(responseBody))
	}))
}

// createAPITestClient creates a test API client with the given server URL and default headers
func createAPITestClient(t *testing.T, serverURL string, defaultHeaders map[string]string) *openfga.APIClient {
	t.Helper()

	config, err := openfga.NewConfiguration(openfga.Configuration{
		ApiUrl:         serverURL,
		DefaultHeaders: defaultHeaders,
	})
	if err != nil {
		t.Fatalf("Failed to create API configuration: %v", err)
	}

	return openfga.NewAPIClient(config)
}

// Misc tests

// Test RequestOptions structure directly at API level
func TestAPIRequestOptionsStructure(t *testing.T) {
	t.Run("RequestOptionsWithAllFields", func(t *testing.T) {
		options := openfga.RequestOptions{
			Headers: map[string]string{
				"Test-Header": "test-value",
			},
		}

		if options.Headers["Test-Header"] != "test-value" {
			t.Errorf("Expected Test-Header to be 'test-value', got '%s'", options.Headers["Test-Header"])
		}
	})

	t.Run("RequestOptionsWithNilHeaders", func(t *testing.T) {
		options := openfga.RequestOptions{
			Headers: nil,
		}

		if options.Headers != nil {
			t.Errorf("Expected Headers to be nil, got %v", options.Headers)
		}
	})

	t.Run("RequestOptionsWithEmptyHeaders", func(t *testing.T) {
		options := openfga.RequestOptions{
			Headers: map[string]string{},
		}

		if len(options.Headers) != 0 {
			t.Errorf("Expected Headers to be empty, got %v", options.Headers)
		}
	})
}

// Test header precedence and merging behavior
func TestAPIHeaderPrecedenceHandling(t *testing.T) {
	t.Run("CustomHeadersOverrideDefaults", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"allowed": true}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			"Header-1": "default-1",
			"Header-2": "default-2",
			"Header-3": "default-3",
		})

		_, _, err := client.OpenFgaApi.Check(context.Background(), apiTestStoreId).
			Body(openfga.CheckRequest{
				TupleKey: openfga.CheckRequestTupleKey{
					User:     apiTestUser,
					Relation: apiTestRelation,
					Object:   apiTestObject,
				},
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					"Header-1": "overridden-1", // Override the default
					"Header-4": "custom-4",
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders["Header-1"] != "overridden-1" {
			t.Errorf("Expected Header-1 to be 'overridden-1', got '%s'", capturedHeaders["Header-1"])
		}

		if capturedHeaders["Header-2"] != "default-2" {
			t.Errorf("Expected Header-2 to be 'default-2', got '%s'", capturedHeaders["Header-2"])
		}

		if capturedHeaders["Header-3"] != "default-3" {
			t.Errorf("Expected Header-3 to be 'default-3', got '%s'", capturedHeaders["Header-3"])
		}

		if capturedHeaders["Header-4"] != "custom-4" {
			t.Errorf("Expected Header-4 to be 'custom-4', got '%s'", capturedHeaders["Header-4"])
		}
	})
}

// Test the header handling for the methods

func TestCheckAPIMethodHeaderHandling(t *testing.T) {
	t.Run("CheckAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"allowed": true}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.Check(context.Background(), apiTestStoreId).
			Body(openfga.CheckRequest{
				TupleKey: openfga.CheckRequestTupleKey{
					User:     apiTestUser,
					Relation: apiTestRelation,
					Object:   apiTestObject,
				},
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})

	t.Run("CheckAPIWithoutCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"allowed": true}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.Check(context.Background(), apiTestStoreId).
			Body(openfga.CheckRequest{
				TupleKey: openfga.CheckRequestTupleKey{
					User:     apiTestUser,
					Relation: apiTestRelation,
					Object:   apiTestObject,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiDefaultHeaderName] != apiDefaultHeaderValue {
			t.Errorf(apiExpectedDefaultMsg, capturedHeaders[apiDefaultHeaderName])
		}

		if _, exists := capturedHeaders[apiCustomHeaderName]; exists {
			t.Error("Did not expect X-Custom-Header to be present")
		}
	})

	t.Run("CheckAPIWithEmptyHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"allowed": true}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, nil)

		_, _, err := client.OpenFgaApi.Check(context.Background(), apiTestStoreId).
			Body(openfga.CheckRequest{
				TupleKey: openfga.CheckRequestTupleKey{
					User:     apiTestUser,
					Relation: apiTestRelation,
					Object:   apiTestObject,
				},
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		for header := range capturedHeaders {
			if strings.HasPrefix(header, "X-") || header == apiDefaultHeaderName {
				t.Errorf("Unexpected custom header found: %s", header)
			}
		}
	})
}

func TestBatchCheckAPIMethodHeaderHandling(t *testing.T) {
	t.Run("BatchCheckAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"test-correlation-id": {"allowed": true}}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.BatchCheck(context.Background(), apiTestStoreId).
			Body(openfga.BatchCheckRequest{
				Checks: []openfga.BatchCheckItem{
					{
						TupleKey: openfga.CheckRequestTupleKey{
							User:     apiTestUser,
							Relation: apiTestRelation,
							Object:   apiTestObject,
						},
						CorrelationId: "test-correlation-id",
					},
				},
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})
}

func TestWriteAPIMethodHeaderHandling(t *testing.T) {
	t.Run("WriteAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.Write(context.Background(), apiTestStoreId).
			Body(openfga.WriteRequest{
				Writes: &openfga.WriteRequestWrites{
					TupleKeys: []openfga.TupleKey{
						{
							User:     apiTestUser,
							Relation: apiTestRelation,
							Object:   apiTestObject,
						},
					},
				},
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})
}

func TestReadAPIMethodHeaderHandling(t *testing.T) {
	t.Run("ReadAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"tuples": []}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.Read(context.Background(), apiTestStoreId).
			Body(openfga.ReadRequest{
				TupleKey: &openfga.ReadRequestTupleKey{
					User:     openfga.PtrString(apiTestUser),
					Relation: openfga.PtrString(apiTestRelation),
					Object:   openfga.PtrString(apiTestObject),
				},
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})
}

func TestExpandAPIMethodHeaderHandling(t *testing.T) {
	t.Run("ExpandAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"tree": {"root": {"name": "document:roadmap#viewer"}}}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.Expand(context.Background(), apiTestStoreId).
			Body(openfga.ExpandRequest{
				TupleKey: openfga.ExpandRequestTupleKey{
					Relation: apiTestRelation,
					Object:   apiTestObject,
				},
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})
}

func TestListObjectsAPIMethodHeaderHandling(t *testing.T) {
	t.Run("ListObjectsAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"objects": ["document:roadmap"]}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.ListObjects(context.Background(), apiTestStoreId).
			Body(openfga.ListObjectsRequest{
				User:     apiTestUser,
				Relation: apiTestRelation,
				Type:     "document",
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})
}

func TestListUsersAPIMethodHeaderHandling(t *testing.T) {
	t.Run("ListUsersAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"users": [{"object": {"type": "user", "id": "anne"}}]}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.ListUsers(context.Background(), apiTestStoreId).
			Body(openfga.ListUsersRequest{
				Object: openfga.FgaObject{
					Type: "document",
					Id:   "roadmap",
				},
				Relation: apiTestRelation,
				UserFilters: []openfga.UserTypeFilter{
					{Type: "user"},
				},
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})
}

func TestStoreAPIMethodHeaderHandling(t *testing.T) {
	t.Run("ListStoresAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"stores": [{"id": "01H0H015178Y2V4CX10C2KGHF4", "name": "test"}]}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.ListStores(context.Background()).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})

	t.Run("CreateStoreAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"id": "01H0H015178Y2V4CX10C2KGHF4", "name": "test"}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.CreateStore(context.Background()).
			Body(openfga.CreateStoreRequest{
				Name: "test",
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})

	t.Run("GetStoreAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"id": "01H0H015178Y2V4CX10C2KGHF4", "name": "test"}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.GetStore(context.Background(), apiTestStoreId).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})

	t.Run("DeleteStoreAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, err := client.OpenFgaApi.DeleteStore(context.Background(), apiTestStoreId).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})
}

func TestReadAuthorizationModelAPIMethodHeaderHandling(t *testing.T) {
	t.Run("ReadAuthorizationModelAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"authorization_model": {"id": "01H0H015178Y2V4CX10C2KGHF4", "schema_version": "1.1"}}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.ReadAuthorizationModel(context.Background(), apiTestStoreId, apiTestStoreId).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})
}

func TestWriteAuthorizationModelAPIMethodHeaderHandling(t *testing.T) {
	t.Run("WriteAuthorizationModelAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"authorization_model_id": "01H0H015178Y2V4CX10C2KGHF4"}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.WriteAuthorizationModel(context.Background(), apiTestStoreId).
			Body(openfga.WriteAuthorizationModelRequest{
				SchemaVersion: "1.1",
				TypeDefinitions: []openfga.TypeDefinition{
					{
						Type: "user",
					},
					{
						Type: "document",
						Relations: &map[string]openfga.Userset{
							"viewer": {},
						},
					},
				},
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})
}

func TestReadChangesAPIMethodHeaderHandling(t *testing.T) {
	t.Run("ReadChangesAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"changes": [], "continuation_token": ""}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.ReadChanges(context.Background(), apiTestStoreId).
			Type_("document").
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})
}

func TestAssertionsAPIMethodHeaderHandling(t *testing.T) {
	t.Run("ReadAssertionsAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{"assertions": []}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, _, err := client.OpenFgaApi.ReadAssertions(context.Background(), apiTestStoreId, apiTestStoreId).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})

	t.Run("WriteAssertionsAPIWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createAPITestServer(t, &capturedHeaders, `{}`)
		defer server.Close()

		client := createAPITestClient(t, server.URL, map[string]string{
			apiDefaultHeaderName: apiDefaultHeaderValue,
		})

		_, err := client.OpenFgaApi.WriteAssertions(context.Background(), apiTestStoreId, apiTestStoreId).
			Body(openfga.WriteAssertionsRequest{
				Assertions: []openfga.Assertion{
					{
						TupleKey: openfga.AssertionTupleKey{
							User:     apiTestUser,
							Relation: apiTestRelation,
							Object:   apiTestObject,
						},
						Expectation: true,
					},
				},
			}).
			Options(openfga.RequestOptions{
				Headers: map[string]string{
					apiCustomHeaderName:  apiCustomHeaderValue,
					apiDefaultHeaderName: apiOverriddenValue,
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(apiRequestFailedMsg, err)
		}

		if capturedHeaders[apiCustomHeaderName] != apiCustomHeaderValue {
			t.Errorf(apiExpectedCustomMsg, capturedHeaders[apiCustomHeaderName])
		}

		if capturedHeaders[apiDefaultHeaderName] != apiOverriddenValue {
			t.Errorf(apiExpectedOverriddenMsg, capturedHeaders[apiDefaultHeaderName])
		}
	})
}
