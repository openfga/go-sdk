// This example demonstrates how to use the low-level APIExecutor to call real
// OpenFGA API endpoints — both standard request/response and streaming.
//
// It exercises Execute, ExecuteWithDecode, and ExecuteStreaming against a live
// OpenFGA server, covering the most common operations:
//
//  1. ListStores       — GET    /stores
//  2. CreateStore      — POST   /stores
//  3. GetStore         — GET    /stores/{store_id}
//  4. WriteAuthModel   — POST   /stores/{store_id}/authorization-models
//  5. WriteTuples      — POST   /stores/{store_id}/write
//  6. ReadTuples       — POST   /stores/{store_id}/read
//  7. Check            — POST   /stores/{store_id}/check
//  8. ListObjects      — POST   /stores/{store_id}/list-objects
//  9. StreamedListObj  — POST   /stores/{store_id}/streamed-list-objects  (streaming)
//  10. DeleteStore     — DELETE /stores/{store_id}
//
// For the recommended high-level typed approach, see the example1 and
// streamed_list_objects examples.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	openfga "github.com/openfga/go-sdk"
	"github.com/openfga/go-sdk/client"
)

func main() {
	ctx := context.Background()

	apiUrl := os.Getenv("FGA_API_URL")
	if apiUrl == "" {
		apiUrl = "http://localhost:8080"
	}

	// We create a thin SDK client only to obtain an APIExecutor.
	// All actual API calls below go through the executor.
	fgaClient, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl: apiUrl,
	})
	if err != nil {
		handleError("NewSdkClient", err)
	}

	executor := fgaClient.GetAPIExecutor()

	fmt.Println("=== OpenFGA APIExecutor Example ===\n")

	// -----------------------------------------------------------------
	// 1. ListStores (GET /stores) — raw Execute
	// -----------------------------------------------------------------
	fmt.Println("1. ListStores (raw Execute)")
	listStoresResp, err := executor.Execute(ctx,
		openfga.NewAPIExecutorRequestBuilder("ListStores", http.MethodGet, "/stores").Build(),
	)
	if err != nil {
		handleError("ListStores", err)
	}
	var listStores openfga.ListStoresResponse
	if err := json.Unmarshal(listStoresResp.Body, &listStores); err != nil {
		handleError("ListStores decode", err)
	}
	fmt.Printf("   Status: %d | Stores count: %d\n\n", listStoresResp.StatusCode, len(listStores.Stores))

	// -----------------------------------------------------------------
	// 2. CreateStore (POST /stores) — ExecuteWithDecode
	// -----------------------------------------------------------------
	fmt.Println("2. CreateStore (ExecuteWithDecode)")
	var createStore openfga.CreateStoreResponse
	createStoreResp, err := executor.ExecuteWithDecode(ctx,
		openfga.NewAPIExecutorRequestBuilder("CreateStore", http.MethodPost, "/stores").
			WithBody(openfga.CreateStoreRequest{Name: "api-executor-example"}).
			Build(),
		&createStore,
	)
	if err != nil {
		handleError("CreateStore", err)
	}
	storeID := createStore.Id
	fmt.Printf("   Status: %d | Store ID: %s | Name: %s\n\n", createStoreResp.StatusCode, storeID, createStore.Name)

	// -----------------------------------------------------------------
	// 3. GetStore (GET /stores/{store_id}) — path parameters
	// -----------------------------------------------------------------
	fmt.Println("3. GetStore (path parameters)")
	var getStore openfga.GetStoreResponse
	getStoreResp, err := executor.ExecuteWithDecode(ctx,
		openfga.NewAPIExecutorRequestBuilder("GetStore", http.MethodGet, "/stores/{store_id}").
			WithPathParameter("store_id", storeID).
			Build(),
		&getStore,
	)
	if err != nil {
		handleError("GetStore", err)
	}
	fmt.Printf("   Status: %d | Name: %s | Created: %s\n\n", getStoreResp.StatusCode, getStore.Name, getStore.CreatedAt.Format("2006-01-02T15:04:05Z"))

	// -----------------------------------------------------------------
	// 4. WriteAuthorizationModel (POST /stores/{store_id}/authorization-models)
	// -----------------------------------------------------------------
	fmt.Println("4. WriteAuthorizationModel")
	var writeModelResp openfga.WriteAuthorizationModelResponse
	_, err = executor.ExecuteWithDecode(ctx,
		openfga.NewAPIExecutorRequestBuilder("WriteAuthorizationModel", http.MethodPost, "/stores/{store_id}/authorization-models").
			WithPathParameter("store_id", storeID).
			WithBody(openfga.WriteAuthorizationModelRequest{
				SchemaVersion: "1.1",
				TypeDefinitions: []openfga.TypeDefinition{
					{
						Type:      "user",
						Relations: &map[string]openfga.Userset{},
					},
					{
						Type: "document",
						Relations: &map[string]openfga.Userset{
							"reader": {This: &map[string]interface{}{}},
							"writer": {This: &map[string]interface{}{}},
						},
						Metadata: &openfga.Metadata{
							Relations: &map[string]openfga.RelationMetadata{
								"reader": {
									DirectlyRelatedUserTypes: &[]openfga.RelationReference{{Type: "user"}},
								},
								"writer": {
									DirectlyRelatedUserTypes: &[]openfga.RelationReference{{Type: "user"}},
								},
							},
						},
					},
				},
			}).
			Build(),
		&writeModelResp,
	)
	if err != nil {
		handleError("WriteAuthorizationModel", err)
	}
	modelID := writeModelResp.AuthorizationModelId
	fmt.Printf("   Model ID: %s\n\n", modelID)

	// -----------------------------------------------------------------
	// 5. Write tuples (POST /stores/{store_id}/write)
	// -----------------------------------------------------------------
	fmt.Println("5. WriteTuples")
	_, err = executor.Execute(ctx,
		openfga.NewAPIExecutorRequestBuilder("Write", http.MethodPost, "/stores/{store_id}/write").
			WithPathParameter("store_id", storeID).
			WithBody(openfga.WriteRequest{
				Writes: &openfga.WriteRequestWrites{
					TupleKeys: []openfga.TupleKey{
						{User: "user:alice", Relation: "writer", Object: "document:roadmap"},
						{User: "user:bob", Relation: "reader", Object: "document:roadmap"},
					},
				},
				AuthorizationModelId: &modelID,
			}).
			Build(),
	)
	if err != nil {
		handleError("Write", err)
	}
	fmt.Println("   Tuples written: user:alice→writer, user:bob→reader on document:roadmap\n")

	// -----------------------------------------------------------------
	// 6. Read tuples (POST /stores/{store_id}/read)
	// -----------------------------------------------------------------
	fmt.Println("6. ReadTuples")
	var readResp openfga.ReadResponse
	_, err = executor.ExecuteWithDecode(ctx,
		openfga.NewAPIExecutorRequestBuilder("Read", http.MethodPost, "/stores/{store_id}/read").
			WithPathParameter("store_id", storeID).
			WithBody(openfga.ReadRequest{
				TupleKey: &openfga.ReadRequestTupleKey{
					Object: openfga.PtrString("document:roadmap"),
				},
			}).
			Build(),
		&readResp,
	)
	if err != nil {
		handleError("Read", err)
	}
	fmt.Printf("   Found %d tuple(s):\n", len(readResp.Tuples))
	for _, t := range readResp.Tuples {
		fmt.Printf("     - %s is %s of %s\n", t.Key.User, t.Key.Relation, t.Key.Object)
	}
	fmt.Println()

	// -----------------------------------------------------------------
	// 7. Check (POST /stores/{store_id}/check) — with custom header
	// -----------------------------------------------------------------
	fmt.Println("7. Check (with custom header)")
	var checkResp openfga.CheckResponse
	_, err = executor.ExecuteWithDecode(ctx,
		openfga.NewAPIExecutorRequestBuilder("Check", http.MethodPost, "/stores/{store_id}/check").
			WithPathParameter("store_id", storeID).
			WithHeader("X-Request-ID", "example-check-123").
			WithBody(openfga.CheckRequest{
				TupleKey: openfga.CheckRequestTupleKey{
					User:     "user:alice",
					Relation: "writer",
					Object:   "document:roadmap",
				},
				AuthorizationModelId: &modelID,
			}).
			Build(),
		&checkResp,
	)
	if err != nil {
		handleError("Check", err)
	}
	fmt.Printf("   user:alice writer document:roadmap → Allowed: %v\n", *checkResp.Allowed)

	// Also check a user who should NOT have access
	var checkResp2 openfga.CheckResponse
	_, err = executor.ExecuteWithDecode(ctx,
		openfga.NewAPIExecutorRequestBuilder("Check", http.MethodPost, "/stores/{store_id}/check").
			WithPathParameter("store_id", storeID).
			WithBody(openfga.CheckRequest{
				TupleKey: openfga.CheckRequestTupleKey{
					User:     "user:bob",
					Relation: "writer",
					Object:   "document:roadmap",
				},
				AuthorizationModelId: &modelID,
			}).
			Build(),
		&checkResp2,
	)
	if err != nil {
		handleError("Check (bob)", err)
	}
	fmt.Printf("   user:bob   writer document:roadmap → Allowed: %v\n\n", *checkResp2.Allowed)

	// -----------------------------------------------------------------
	// 8. ListObjects (POST /stores/{store_id}/list-objects)
	// -----------------------------------------------------------------
	fmt.Println("8. ListObjects")
	var listObjectsResp openfga.ListObjectsResponse
	_, err = executor.ExecuteWithDecode(ctx,
		openfga.NewAPIExecutorRequestBuilder("ListObjects", http.MethodPost, "/stores/{store_id}/list-objects").
			WithPathParameter("store_id", storeID).
			WithBody(openfga.ListObjectsRequest{
				AuthorizationModelId: &modelID,
				User:                 "user:alice",
				Relation:             "writer",
				Type:                 "document",
			}).
			Build(),
		&listObjectsResp,
	)
	if err != nil {
		handleError("ListObjects", err)
	}
	fmt.Printf("   Objects user:alice can write: %v\n\n", listObjectsResp.Objects)

	// -----------------------------------------------------------------
	// 9. StreamedListObjects (POST /stores/{store_id}/streamed-list-objects)
	//    Write more tuples first so we have something meaningful to stream.
	// -----------------------------------------------------------------
	fmt.Println("9. StreamedListObjects (ExecuteStreaming)")
	fmt.Println("   Writing 200 additional tuples for streaming demo...")
	for batch := 0; batch < 2; batch++ {
		tuples := make([]openfga.TupleKey, 0, 100)
		for i := 1; i <= 100; i++ {
			tuples = append(tuples, openfga.TupleKey{
				User:     "user:alice",
				Relation: "reader",
				Object:   fmt.Sprintf("document:doc-%d", batch*100+i),
			})
		}
		_, err = executor.Execute(ctx,
			openfga.NewAPIExecutorRequestBuilder("Write", http.MethodPost, "/stores/{store_id}/write").
				WithPathParameter("store_id", storeID).
				WithBody(openfga.WriteRequest{
					Writes:               &openfga.WriteRequestWrites{TupleKeys: tuples},
					AuthorizationModelId: &modelID,
				}).
				Build(),
		)
		if err != nil {
			handleError("Write (batch)", err)
		}
	}

	channel, err := executor.ExecuteStreaming(ctx,
		openfga.NewAPIExecutorRequestBuilder("StreamedListObjects", http.MethodPost, "/stores/{store_id}/streamed-list-objects").
			WithPathParameter("store_id", storeID).
			WithBody(openfga.ListObjectsRequest{
				AuthorizationModelId: &modelID,
				User:                 "user:alice",
				Relation:             "reader",
				Type:                 "document",
			}).
			Build(),
		openfga.DefaultStreamBufferSize,
	)
	if err != nil {
		handleError("ExecuteStreaming", err)
	}
	defer channel.Close()

	count := 0
	for {
		select {
		case result, ok := <-channel.Results:
			if !ok {
				select {
				case err := <-channel.Errors:
					if err != nil {
						handleError("StreamedListObjects stream", err)
					}
				default:
				}
				fmt.Printf("   ✓ Streamed %d objects\n\n", count)
				goto streamDone
			}
			var obj openfga.StreamedListObjectsResponse
			if err := json.Unmarshal(result, &obj); err != nil {
				handleError("decode stream result", err)
			}
			count++
			if count <= 3 || count%50 == 0 {
				fmt.Printf("     Object: %s\n", obj.Object)
			}
		case err := <-channel.Errors:
			if err != nil {
				handleError("StreamedListObjects error", err)
			}
		}
	}
streamDone:

	// -----------------------------------------------------------------
	// 10. DeleteStore (DELETE /stores/{store_id})
	// -----------------------------------------------------------------
	fmt.Println("10. DeleteStore (cleanup)")
	deleteResp, err := executor.Execute(ctx,
		openfga.NewAPIExecutorRequestBuilder("DeleteStore", http.MethodDelete, "/stores/{store_id}").
			WithPathParameter("store_id", storeID).
			Build(),
	)
	if err != nil {
		handleError("DeleteStore", err)
	}
	fmt.Printf("    Status: %d | Store deleted\n\n", deleteResp.StatusCode)

	fmt.Println("=== All examples completed successfully! ===")
}

func handleError(context string, err error) {
	fmt.Fprintf(os.Stderr, "\nError in %s: %v\n", context, err)
	fmt.Fprintln(os.Stderr, "\nMake sure OpenFGA is running on localhost:8080 (or set FGA_API_URL)")
	fmt.Fprintln(os.Stderr, "Run: docker run -p 8080:8080 openfga/openfga:latest run")
	os.Exit(1)
}
