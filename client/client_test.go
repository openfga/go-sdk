package client_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/openfga/go-sdk"
	. "github.com/openfga/go-sdk/client"
	"net/http"
	"testing"
)

type TestDefinition struct {
	Name           string
	JsonResponse   string
	ResponseStatus int
	Method         string
	RequestPath    string
}

func TestOpenFgaClient(t *testing.T) {
	fgaClient, err := NewSdkClient(&ClientConfiguration{
		ApiHost: "api.fga.example",
		StoreId: "6c181474-aaa1-4df7-8929-6e7b3a992754",
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	/* Stores */
	t.Run("ListStores", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ListStores",
			JsonResponse:   `{"stores":[{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw","name":"Test Store","created_at":"2023-01-01T23:23:23.000000000Z","updated_at":"2023-01-01T23:23:23.000000000Z"}]}`,
			ResponseStatus: http.StatusOK,
			Method:         "GET",
		}

		var expectedResponse openfga.ListStoresResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		got, response, err := fgaClient.ListStores(context.Background()).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		if len(*got.Stores) != 1 {
			t.Fatalf("%v", err)
		}

		if *((*got.Stores)[0].Id) != *((*expectedResponse.Stores)[0].Id) {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, *((*got.Stores)[0].Id), *((*expectedResponse.Stores)[0].Id))
		}
	})

	t.Run("CreateStore", func(t *testing.T) {
		test := TestDefinition{
			Name:           "CreateStore",
			JsonResponse:   `{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw","name":"Test Store","created_at":"2023-01-01T23:23:23.000000000Z","updated_at":"2023-01-01T23:23:23.000000000Z"}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
		}
		requestBody := ClientCreateStoreRequest{
			Name: "Test Store",
		}

		var expectedResponse openfga.CreateStoreResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := fgaClient.CreateStore(context.Background()).Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		_, err = got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}
		if *got.Name != *expectedResponse.Name {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, *got.Name, *expectedResponse.Name)
		}
	})

	t.Run("GetStore", func(t *testing.T) {
		test := TestDefinition{
			Name:           "GetStore",
			JsonResponse:   `{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw","name":"Test Store","created_at":"2023-01-01T23:23:23.000000000Z","updated_at":"2023-01-01T23:23:23.000000000Z"}`,
			ResponseStatus: http.StatusOK,
			Method:         "GET",
		}

		var expectedResponse openfga.GetStoreResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		got, response, err := fgaClient.GetStore(context.Background()).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		if *got.Id != *expectedResponse.Id {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, *got.Id, *expectedResponse.Id)
		}
	})

	t.Run("DeleteStore", func(t *testing.T) {
		test := TestDefinition{
			Name:           "DeleteStore",
			JsonResponse:   ``,
			ResponseStatus: http.StatusNoContent,
			Method:         "DELETE",
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, "{}")
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, response, err := fgaClient.DeleteStore(context.Background()).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}
	})

	/* Authorization Models */
	t.Run("ReadAuthorizationModels", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadAuthorizationModels",
			JsonResponse:   `{"authorization_models":[{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw","schema_version":"1.1","type_definitions":[]}]}`,
			ResponseStatus: http.StatusOK,
			Method:         "GET",
			RequestPath:    "authorization-models",
		}

		var expectedResponse openfga.ReadAuthorizationModelsResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		got, response, err := fgaClient.ReadAuthorizationModels(context.Background()).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		if len(*got.AuthorizationModels) != 1 {
			t.Fatalf("%v", err)
		}

		if *((*got.AuthorizationModels)[0].Id) != *((*expectedResponse.AuthorizationModels)[0].Id) {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, *((*got.AuthorizationModels)[0].Id), *((*expectedResponse.AuthorizationModels)[0].Id))
		}
	})

	t.Run("WriteAuthorizationModel", func(t *testing.T) {
		test := TestDefinition{
			Name:           "WriteAuthorizationModel",
			JsonResponse:   `{"authorization_model_id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw"}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "authorization-models",
		}
		requestBody := ClientWriteAuthorizationModelRequest{
			TypeDefinitions: []openfga.TypeDefinition{{
				Type: "github-repo",
				Relations: &map[string]openfga.Userset{
					"repo_writer": {
						This: &map[string]interface{}{},
					},
					"viewer": {Union: &openfga.Usersets{
						Child: &[]openfga.Userset{
							{This: &map[string]interface{}{}},
							{ComputedUserset: &openfga.ObjectRelation{
								Object:   openfga.PtrString(""),
								Relation: openfga.PtrString("repo_writer"),
							}},
						},
					}},
				},
			}},
		}

		var expectedResponse openfga.WriteAuthorizationModelResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := fgaClient.WriteAuthorizationModel(context.Background()).Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		_, err = got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}
	})

	t.Run("ReadAuthorizationModel", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadAuthorizationModel",
			JsonResponse:   `{"authorization_model":{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw","schema_version":"1.1","type_definitions":[{"type":"github-repo", "relations":{"viewer":{"this":{}}}}]}}`,
			ResponseStatus: http.StatusOK,
			Method:         "GET",
			RequestPath:    "authorization-models",
		}

		var expectedResponse openfga.ReadAuthorizationModelResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}
		modelId := *(*expectedResponse.AuthorizationModel).Id

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath, modelId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		options := ClientReadAuthorizationModelOptions{
			AuthorizationModelId: openfga.PtrString(modelId),
		}
		got, response, err := fgaClient.ReadAuthorizationModel(context.Background()).Options(options).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if *(*got.AuthorizationModel).Id != modelId {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("ReadLatestAuthorizationModel", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadAuthorizationModels",
			JsonResponse:   `{"authorization_models":[{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw","schema_version":"1.1","type_definitions":[{"type":"github-repo", "relations":{"viewer":{"this":{}}}}]}]}`,
			ResponseStatus: http.StatusOK,
			Method:         "GET",
			RequestPath:    "authorization-models",
		}

		var expectedResponse openfga.ReadAuthorizationModelsResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}
		modelId := *((*expectedResponse.AuthorizationModels)[0].Id)

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := fgaClient.ReadLatestAuthorizationModel(context.Background()).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if (*(*got.AuthorizationModel).Id) != modelId {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	/* Relationship Tuples */
	t.Run("ReadChanges", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ReadChanges",
			JsonResponse:   `{"changes":[{"tuple_key":{"user":"user:81684243-9356-4421-8fbf-a4f8d36aa31b","relation":"viewer","object":"document:roadmap"},"operation":"TUPLE_OPERATION_WRITE","timestamp": "2000-01-01T00:00:00Z"}],"continuation_token":"eyJwayI6IkxBVEVTVF9OU0NPTkZJR19hdXRoMHN0b3JlIiwic2siOiIxem1qbXF3MWZLZExTcUoyN01MdTdqTjh0cWgifQ=="}`,
			ResponseStatus: http.StatusOK,
			Method:         "GET",
			RequestPath:    "changes",
		}

		var expectedResponse openfga.ReadChangesResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		body := ClientReadChangesRequest{
			Type: "repo",
		}
		options := ClientReadChangesOptions{ContinuationToken: openfga.PtrString("eyJwayI6IkxBVEVTVF9OU0NPTkZJR19hdXRoMHN0b3JlIiwic2siOiIxem1qbXF3MWZLZExTcUoyN01MdTdqTjh0cWgifQ=="), PageSize: openfga.PtrInt32(25)}
		got, response, err := fgaClient.ReadChanges(context.Background()).Body(body).Options(options).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if len(*got.Changes) != len(*expectedResponse.Changes) {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("Read", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Read",
			JsonResponse:   `{"tuples":[{"key":{"user":"user:81684243-9356-4421-8fbf-a4f8d36aa31b","relation":"viewer","object":"document:roadmap"},"timestamp": "2000-01-01T00:00:00Z"}]}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "read",
		}

		requestBody := ClientReadRequest{
			User:     openfga.PtrString("user:81684243-9356-4421-8fbf-a4f8d36aa31b"),
			Relation: openfga.PtrString("viewer"),
			Object:   openfga.PtrString("document:roadmap"),
		}

		var expectedResponse openfga.ReadResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := fgaClient.Read(context.Background()).Body(requestBody).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if len(*got.Tuples) != len(*expectedResponse.Tuples) {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("Write", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "write",
		}
		requestBody := ClientWriteRequest{
			Writes: &[]ClientTupleKey{{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "viewer",
				Object:   "document:roadmap",
			}},
		}
		options := ClientWriteOptions{
			AuthorizationModelId: openfga.PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, err := fgaClient.Write(context.Background()).Body(requestBody).Options(options).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		//if response.StatusCode != test.ResponseStatus {
		//	t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		//}
	})

	t.Run("WriteNonTransaction", func(t *testing.T) {
		t.Skip()
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "write",
		}
		requestBody := ClientWriteRequest{
			Writes: &[]ClientTupleKey{{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "viewer",
				Object:   "document:roadmap",
			}},
		}
		options := ClientWriteOptions{
			AuthorizationModelId: openfga.PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
			Transaction:          &TransactionOptions{Disable: true},
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, err := fgaClient.Write(context.Background()).Body(requestBody).Options(options).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		//if response.StatusCode != test.ResponseStatus {
		//	t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		//}
	})

	t.Run("WriteTuples", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "write",
		}
		requestBody := ClientWriteRequest{
			Writes: &[]ClientTupleKey{{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "viewer",
				Object:   "document:roadmap",
			}},
		}
		options := ClientWriteOptions{
			AuthorizationModelId: openfga.PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, err := fgaClient.Write(context.Background()).Body(requestBody).Options(options).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		//if response.StatusCode != test.ResponseStatus {
		//	t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		//}
	})

	t.Run("DeleteTuples", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Write",
			JsonResponse:   `{}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "write",
		}

		requestBody := []ClientTupleKey{{
			User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
			Relation: "viewer",
			Object:   "document:roadmap",
		}}
		options := ClientWriteOptions{
			AuthorizationModelId: openfga.PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse map[string]interface{}
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		_, err := fgaClient.DeleteTuples(context.Background()).Body(requestBody).Options(options).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		//if response.StatusCode != test.ResponseStatus {
		//	t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		//}
	})

	/* Relationship Queries */
	t.Run("Check", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := ClientCheckRequest{
			User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
			Relation: "viewer",
			Object:   "document:roadmap",
		}

		options := ClientCheckOptions{
			AuthorizationModelId: openfga.PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse openfga.CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := fgaClient.Check(context.Background()).Body(requestBody).Options(options).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if *got.Allowed != *expectedResponse.Allowed {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("BatchCheck", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Check",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "check",
		}
		requestBody := ClientBatchCheckBody{{
			User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
			Relation: "viewer",
			Object:   "document:roadmap",
			ContextualTuples: &[]ClientTupleKey{{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "editor",
				Object:   "document:roadmap",
			}},
		}, {
			User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
			Relation: "admin",
			Object:   "document:roadmap",
			ContextualTuples: &[]ClientTupleKey{{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "editor",
				Object:   "document:roadmap",
			}},
		}, {
			User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
			Relation: "creator",
			Object:   "document:roadmap",
		}, {
			User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
			Relation: "deleter",
			Object:   "document:roadmap",
		}}

		options := ClientBatchCheckOptions{
			AuthorizationModelId: openfga.PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse openfga.CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, err := fgaClient.BatchCheck(context.Background()).Body(requestBody).Options(options).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if httpmock.GetTotalCallCount() != 4 {
			t.Fatalf("OpenFgaClient.%v() - wanted %v calls to /check, got %v", test.Name, 4, httpmock.GetTotalCallCount())
		}

		if len(got) != len(requestBody) {
			t.Fatalf("OpenFgaClient.%v() - Response Length = %v, want %v", test.Name, len(got), len(requestBody))
		}

		for index := 0; index < len(got); index++ {
			response := got[index]
			if response.Error != nil {
				t.Fatalf("OpenFgaClient.%v()|%d/ %v", test.Name, index, response.Error)
			}
			if response.HttpResponse.StatusCode != test.ResponseStatus {
				t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.HttpResponse.StatusCode, test.ResponseStatus)
			}

			responseJson, err := response.MarshalJSON()
			if err != nil {
				t.Fatalf("%v", err)
			}

			if *response.Allowed != *expectedResponse.Allowed {
				t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
			}
		}
	})

	t.Run("Expand", func(t *testing.T) {
		test := TestDefinition{
			Name:           "Expand",
			JsonResponse:   `{"tree":{"root":{"name":"document:roadmap#viewer","union":{"nodes":[{"name": "document:roadmap#viewer","leaf":{"users":{"users":["user:81684243-9356-4421-8fbf-a4f8d36aa31b"]}}}]}}}}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "expand",
		}

		requestBody := ClientExpandRequest{
			Relation: "viewer",
			Object:   "document:roadmap",
		}
		options := ClientExpandOptions{
			AuthorizationModelId: openfga.PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse openfga.ExpandResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := fgaClient.Expand(context.Background()).Body(requestBody).Options(options).Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		_, err = got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}
	})

	t.Run("ListObjects", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ListObjects",
			JsonResponse:   `{"objects":["document:roadmap"]}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "list-objects",
		}

		requestBody := ClientListObjectsRequest{
			User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
			Relation: "can_read",
			Type:     "document",
			ContextualTuples: &[]ClientTupleKey{{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "editor",
				Object:   "folder:product",
			}, {
				User:     "folder:product",
				Relation: "parent",
				Object:   "document:roadmap",
			}},
		}
		options := ClientListObjectsOptions{
			AuthorizationModelId: openfga.PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse openfga.ListObjectsResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := fgaClient.ListObjects(context.Background()).
			Body(requestBody).
			Options(options).
			Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if len(*got.Objects) != len(*expectedResponse.Objects) || (*got.Objects)[0] != (*expectedResponse.Objects)[0] {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("ListRelations", func(t *testing.T) {
		test := TestDefinition{
			Name:           "ListRelations",
			JsonResponse:   `{"allowed":true, "resolution":""}`,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "check",
		}

		requestBody := ClientListRelationsRequest{
			User:      "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
			Object:    "document:roadmap",
			Relations: []string{"can_view", "can_edit", "can_delete", "can_rename"},
			ContextualTuples: &[]ClientTupleKey{{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "editor",
				Object:   "document:roadmap",
			}},
		}
		options := ClientListRelationsOptions{
			AuthorizationModelId: openfga.PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse openfga.CheckResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterMatcherResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			httpmock.BodyContainsString(`"relation":"can_delete"`),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, openfga.CheckResponse{Allowed: openfga.PtrBool(false)})
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)

		got, err := fgaClient.ListRelations(context.Background()).
			Body(requestBody).
			Options(options).
			Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if httpmock.GetTotalCallCount() != 4 {
			t.Fatalf("OpenFgaClient.%v() - wanted %v calls to /check, got %v", test.Name, 4, httpmock.GetTotalCallCount())
		}

		// TODO
		//responseJson, err := got.MarshalJSON()
		//if err != nil {
		//	t.Fatalf("%v", err)
		//}
		//
		if len(got.Relations) != 3 {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, len(got.Relations), 3)
		}
	})

	t.Run("ListRelationsNoRelationsProvided", func(t *testing.T) {
		t.Skip()
		test := TestDefinition{
			Name:           "ListRelations",
			JsonResponse:   ``,
			ResponseStatus: http.StatusOK,
			Method:         "POST",
			RequestPath:    "check",
		}

		requestBody := ClientListRelationsRequest{
			User:      "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
			Object:    "document:roadmap",
			Relations: []string{},
			ContextualTuples: &[]ClientTupleKey{{
				User:     "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation: "editor",
				Object:   "document:roadmap",
			}},
		}
		options := ClientListRelationsOptions{
			AuthorizationModelId: openfga.PtrString("01GAHCE4YVKPQEKZQHT2R89MQV"),
		}

		var expectedResponse ClientListRelationsResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		_, err := fgaClient.ListRelations(context.Background()).
			Body(requestBody).
			Options(options).
			Execute()

		if err == nil {
			t.Fatalf("OpenFgaClient.%v() - expected an error but received none", test.Name)
		}
	})

	/* Assertions */
	t.Run("ReadAssertions", func(t *testing.T) {
		modelId := "01GAHCE4YVKPQEKZQHT2R89MQV"
		test := TestDefinition{
			Name:           "ReadAssertions",
			JsonResponse:   fmt.Sprintf(`{"assertions":[{"tuple_key":{"user":"user:anna","relation":"can_view","object":"document:roadmap"},"expectation":true}],"authorization_model_id":"%s"}`, modelId),
			ResponseStatus: http.StatusOK,
			Method:         http.MethodGet,
			RequestPath:    "assertions",
		}

		options := ClientReadAssertionsOptions{
			AuthorizationModelId: openfga.PtrString(modelId),
		}

		var expectedResponse openfga.ReadAssertionsResponse
		if err := json.Unmarshal([]byte(test.JsonResponse), &expectedResponse); err != nil {
			t.Fatalf("%v", err)
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath, modelId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, expectedResponse)
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		got, response, err := fgaClient.ReadAssertions(context.Background()).
			Options(options).
			Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}

		responseJson, err := got.MarshalJSON()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if len(*got.Assertions) != len(*expectedResponse.Assertions) || (*got.Assertions)[0].Expectation != (*expectedResponse.Assertions)[0].Expectation {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, string(responseJson), test.JsonResponse)
		}
	})

	t.Run("WriteAssertions", func(t *testing.T) {
		modelId := "01GAHCE4YVKPQEKZQHT2R89MQV"
		test := TestDefinition{
			Name:           "WriteAssertions",
			JsonResponse:   "",
			ResponseStatus: http.StatusNoContent,
			Method:         http.MethodPut,
			RequestPath:    "assertions",
		}

		requestBody := ClientWriteAssertionsRequest{
			ClientAssertion{
				User:        "user:81684243-9356-4421-8fbf-a4f8d36aa31b",
				Relation:    "can_view",
				Object:      "document:roadmap",
				Expectation: true,
			},
		}
		options := ClientWriteAssertionsOptions{
			AuthorizationModelId: openfga.PtrString(modelId),
		}

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(test.Method, fmt.Sprintf("%s://%s/stores/%s/%s/%s", fgaClient.Config.ApiScheme, fgaClient.Config.ApiHost, fgaClient.Config.StoreId, test.RequestPath, modelId),
			func(req *http.Request) (*http.Response, error) {
				resp, err := httpmock.NewJsonResponse(test.ResponseStatus, "")
				if err != nil {
					return httpmock.NewStringResponse(500, ""), nil
				}
				return resp, nil
			},
		)
		response, err := fgaClient.WriteAssertions(context.Background()).
			Body(requestBody).
			Options(options).
			Execute()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if response.StatusCode != test.ResponseStatus {
			t.Fatalf("OpenFgaClient.%v() = %v, want %v", test.Name, response.StatusCode, test.ResponseStatus)
		}
	})
}
