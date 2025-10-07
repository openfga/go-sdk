package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	fgaSdk "github.com/openfga/go-sdk"
	fgaSdkClient "github.com/openfga/go-sdk/client"
)

// Test helpers and setup

// Constants to avoid duplication
const (
	defaultHeaderName           = "Default-Header"
	defaultHeaderValue          = "default-value"
	overriddenValue             = "overridden-value"
	customHeaderName            = "X-Custom-Header"
	customHeaderValue           = "custom-value"
	testUser                    = "user:anne"
	testRelation                = "viewer"
	testObject                  = "document:roadmap"
	testHeaderName              = "Test-Header"
	testHeaderValue             = "test-value"
	checkRequestFailedMsg       = "Check request failed: %v"
	expectedCustomHeaderMsg     = "Expected X-Custom-Header to be 'custom-value', got '%s'"
	expectedOverriddenHeaderMsg = "Expected Default-Header to be overridden to 'overridden-value', got '%s'"
	expectedDefaultHeaderMsg    = "Expected Default-Header to be 'default-value', got '%s'"
)

func createTestServer(t *testing.T, capturedHeaders *map[string]string, responseBody string) *httptest.Server {
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

func createTestClient(t *testing.T, serverURL string, defaultHeaders map[string]string) *fgaSdkClient.OpenFgaClient {
	t.Helper()

	config := fgaSdkClient.ClientConfiguration{
		ApiUrl:         serverURL,
		DefaultHeaders: defaultHeaders,
		StoreId:        "01H0H015178Y2V4CX10C2KGHF4",
		HTTPClient:     &http.Client{},
	}

	client, err := fgaSdkClient.NewSdkClient(&config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	return client
}

// Test RequestOptions
func TestRequestOptionsStructure(t *testing.T) {
	t.Run("RequestOptionsEmbedding", func(t *testing.T) {
		options := fgaSdkClient.ClientCheckOptions{
			RequestOptions: fgaSdkClient.RequestOptions{
				Headers: map[string]string{
					testHeaderName: testHeaderValue,
				},
			},
			AuthorizationModelId: fgaSdk.PtrString("01H0H015178Y2V4CX10C2KGHF4"),
			Consistency:          nil,
		}

		if options.Headers[testHeaderName] != testHeaderValue {
			t.Errorf("Expected %s to be '%s', got '%s'", testHeaderName, testHeaderValue, options.Headers[testHeaderName])
		}

		if options.AuthorizationModelId == nil || *options.AuthorizationModelId != "01H0H015178Y2V4CX10C2KGHF4" {
			t.Errorf("Expected AuthorizationModelId to be set correctly")
		}
	})

	t.Run("RequestWithNilHeaders", func(t *testing.T) {
		options := fgaSdkClient.ClientCheckOptions{
			RequestOptions: fgaSdkClient.RequestOptions{
				Headers: nil,
			},
		}

		if options.Headers != nil {
			t.Errorf("Expected Headers to be nil, got %v", options.Headers)
		}
	})
}

// Test the header handling for the methods

func TestCheckMethodHeaderHandling(t *testing.T) {
	t.Run("CheckWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, `{"allowed": true}`)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.Check(context.Background()).
			Body(fgaSdkClient.ClientCheckRequest{
				User:     testUser,
				Relation: testRelation,
				Object:   testObject,
			}).
			Options(fgaSdkClient.ClientCheckOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})

	t.Run("CheckWithoutCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, `{"allowed": true}`)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.Check(context.Background()).
			Body(fgaSdkClient.ClientCheckRequest{
				User:     testUser,
				Relation: testRelation,
				Object:   testObject,
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[defaultHeaderName] != defaultHeaderValue {
			t.Errorf(expectedDefaultHeaderMsg, capturedHeaders[defaultHeaderName])
		}

		if _, exists := capturedHeaders[customHeaderName]; exists {
			t.Error("Did not expect X-Custom-Header to be present")
		}
	})

	t.Run("CheckWithEmptyHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, `{"allowed": true}`)
		defer server.Close()

		client := createTestClient(t, server.URL, nil)

		_, err := client.Check(context.Background()).
			Body(fgaSdkClient.ClientCheckRequest{
				User:     testUser,
				Relation: testRelation,
				Object:   testObject,
			}).
			Options(fgaSdkClient.ClientCheckOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		// Only standard headers should be present
		for header := range capturedHeaders {
			if strings.HasPrefix(header, "X-") || header == defaultHeaderName {
				t.Errorf("Unexpected custom header found: %s", header)
			}
		}
	})

	t.Run("CheckWithNilOptions", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, `{"allowed": true}`)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		request := client.Check(context.Background()).
			Body(fgaSdkClient.ClientCheckRequest{
				User:     testUser,
				Relation: testRelation,
				Object:   testObject,
			})

		// Options not set
		_, err := request.Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[defaultHeaderName] != defaultHeaderValue {
			t.Errorf(expectedDefaultHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}

func TestWriteMethodHeaderHandling(t *testing.T) {
	const writeResponse = `{}`

	t.Run("WriteWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, writeResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.Write(context.Background()).
			Body(fgaSdkClient.ClientWriteRequest{
				Writes: []fgaSdkClient.ClientTupleKey{
					{
						User:     testUser,
						Relation: testRelation,
						Object:   testObject,
					},
				},
			}).
			Options(fgaSdkClient.ClientWriteOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})

	t.Run("WriteWithoutCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, writeResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.Write(context.Background()).
			Body(fgaSdkClient.ClientWriteRequest{
				Writes: []fgaSdkClient.ClientTupleKey{
					{
						User:     testUser,
						Relation: testRelation,
						Object:   testObject,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[defaultHeaderName] != defaultHeaderValue {
			t.Errorf(expectedDefaultHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}

func TestReadMethodHeaderHandling(t *testing.T) {
	const readResponse = `{"tuples": []}`

	t.Run("ReadWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, readResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.Read(context.Background()).
			Body(fgaSdkClient.ClientReadRequest{
				User:     fgaSdk.PtrString(testUser),
				Relation: fgaSdk.PtrString(testRelation),
				Object:   fgaSdk.PtrString(testObject),
			}).
			Options(fgaSdkClient.ClientReadOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}

func TestExpandMethodHeaderHandling(t *testing.T) {
	const expandResponse = `{"tree": {"root": {"name": "document:roadmap#viewer"}}}`

	t.Run("ExpandWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, expandResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.Expand(context.Background()).
			Body(fgaSdkClient.ClientExpandRequest{
				Relation: testRelation,
				Object:   testObject,
			}).
			Options(fgaSdkClient.ClientExpandOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}

func TestListObjectsMethodHeaderHandling(t *testing.T) {
	const listObjectsResponse = `{"objects": ["document:roadmap"]}`

	t.Run("ListObjectsWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, listObjectsResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.ListObjects(context.Background()).
			Body(fgaSdkClient.ClientListObjectsRequest{
				User:     testUser,
				Relation: testRelation,
				Type:     "document",
			}).
			Options(fgaSdkClient.ClientListObjectsOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}

func TestListUsersMethodHeaderHandling(t *testing.T) {
	const listUsersResponse = `{"users": [{"object": {"type": "user", "id": "anne"}}]}`

	t.Run("ListUsersWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, listUsersResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.ListUsers(context.Background()).
			Body(fgaSdkClient.ClientListUsersRequest{
				Object: fgaSdk.FgaObject{
					Type: "document",
					Id:   "roadmap",
				},
				Relation: testRelation,
				UserFilters: []fgaSdk.UserTypeFilter{
					{Type: "user"},
				},
			}).
			Options(fgaSdkClient.ClientListUsersOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}

func TestBatchCheckMethodHeaderHandling(t *testing.T) {
	const batchCheckResponse = `{"result":{"corr-id-123":{"allowed": true}}}`

	t.Run("BatchCheckWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, batchCheckResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		checks := fgaSdkClient.ClientBatchCheckRequest{
			Checks: []fgaSdkClient.ClientBatchCheckItem{{
                CorrelationId: "corr-id-123",
				User:     testUser,
				Relation: testRelation,
				Object:   testObject,
			}},
		}

		_, err := client.BatchCheck(context.Background()).
			Body(checks).
			Options(fgaSdkClient.BatchCheckOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}

func TestReadAuthorizationModelMethodHeaderHandling(t *testing.T) {
	const authModelResponse = `{"authorization_model": {"id": "01H0H015178Y2V4CX10C2KGHF4", "schema_version": "1.1"}}`

	t.Run("ReadAuthorizationModelWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, authModelResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.ReadAuthorizationModel(context.Background()).
			Options(fgaSdkClient.ClientReadAuthorizationModelOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
				AuthorizationModelId: fgaSdk.PtrString("01H0H015178Y2V4CX10C2KGHF4"),
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}

func TestWriteAuthorizationModelMethodHeaderHandling(t *testing.T) {
	const writeAuthModelResponse = `{"authorization_model_id": "01H0H015178Y2V4CX10C2KGHF4"}`

	t.Run("WriteAuthorizationModelWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, writeAuthModelResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.WriteAuthorizationModel(context.Background()).
			Body(fgaSdkClient.ClientWriteAuthorizationModelRequest{
				SchemaVersion: "1.1",
				TypeDefinitions: []fgaSdk.TypeDefinition{
					{
						Type: "user",
					},
					{
						Type: "document",
						Relations: &map[string]fgaSdk.Userset{
							"viewer": {},
						},
					},
				},
			}).
			Options(fgaSdkClient.ClientWriteAuthorizationModelOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}

func TestStoreMethodHeaderHandling(t *testing.T) {
	t.Run("ListStoresWithCustomHeaders", func(t *testing.T) {
		const listStoresResponse = `{"stores": [{"id": "01H0H015178Y2V4CX10C2KGHF4", "name": "test"}]}`
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, listStoresResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.ListStores(context.Background()).
			Options(fgaSdkClient.ClientListStoresOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})

	t.Run("CreateStoreWithCustomHeaders", func(t *testing.T) {
		const createStoreResponse = `{"id": "01H0H015178Y2V4CX10C2KGHF4", "name": "test"}`
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, createStoreResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.CreateStore(context.Background()).
			Body(fgaSdkClient.ClientCreateStoreRequest{
				Name: "test",
			}).
			Options(fgaSdkClient.ClientCreateStoreOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})

	t.Run("GetStoreWithCustomHeaders", func(t *testing.T) {
		const getStoreResponse = `{"id": "01H0H015178Y2V4CX10C2KGHF4", "name": "test"}`
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, getStoreResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.GetStore(context.Background()).
			Options(fgaSdkClient.ClientGetStoreOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})

	t.Run("DeleteStoreWithCustomHeaders", func(t *testing.T) {
		const deleteStoreResponse = `{}`
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, deleteStoreResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.DeleteStore(context.Background()).
			Options(fgaSdkClient.ClientDeleteStoreOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}

func TestReadChangesMethodHeaderHandling(t *testing.T) {
	const readChangesResponse = `{"changes": [], "continuation_token": ""}`

	t.Run("ReadChangesWithCustomHeaders", func(t *testing.T) {
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, readChangesResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.ReadChanges(context.Background()).
			Body(fgaSdkClient.ClientReadChangesRequest{
				Type: "document",
			}).
			Options(fgaSdkClient.ClientReadChangesOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}

func TestAssertionsMethodHeaderHandling(t *testing.T) {
	t.Run("ReadAssertionsWithCustomHeaders", func(t *testing.T) {
		const readAssertionsResponse = `{"assertions": []}`
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, readAssertionsResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		_, err := client.ReadAssertions(context.Background()).
			Options(fgaSdkClient.ClientReadAssertionsOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
				AuthorizationModelId: fgaSdk.PtrString("01H0H015178Y2V4CX10C2KGHF4"),
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})

	t.Run("WriteAssertionsWithCustomHeaders", func(t *testing.T) {
		const writeAssertionsResponse = `{}`
		var capturedHeaders map[string]string
		server := createTestServer(t, &capturedHeaders, writeAssertionsResponse)
		defer server.Close()

		client := createTestClient(t, server.URL, map[string]string{
			defaultHeaderName: defaultHeaderValue,
		})

		assertions := []fgaSdkClient.ClientAssertion{
			{
				User:        testUser,
				Relation:    testRelation,
				Object:      testObject,
				Expectation: true,
			},
		}

		_, err := client.WriteAssertions(context.Background()).
			Body(assertions).
			Options(fgaSdkClient.ClientWriteAssertionsOptions{
				RequestOptions: fgaSdkClient.RequestOptions{
					Headers: map[string]string{
						customHeaderName:  customHeaderValue,
						defaultHeaderName: overriddenValue,
					},
				},
				AuthorizationModelId: fgaSdk.PtrString("01H0H015178Y2V4CX10C2KGHF4"),
			}).
			Execute()

		if err != nil {
			t.Fatalf(checkRequestFailedMsg, err)
		}

		if capturedHeaders[customHeaderName] != customHeaderValue {
			t.Errorf(expectedCustomHeaderMsg, capturedHeaders[customHeaderName])
		}

		if capturedHeaders[defaultHeaderName] != overriddenValue {
			t.Errorf(expectedOverriddenHeaderMsg, capturedHeaders[defaultHeaderName])
		}
	})
}
