// This example demonstrates how to use the low-level APIExecutor's ExecuteStreaming method
// to call the StreamedListObjects endpoint with streaming.
//
// This approach is useful when:
// - You want to call a streaming endpoint that is not yet supported by the SDK
// - You are using an earlier version of the SDK that doesn't yet have a typed method
// - You have a custom streaming endpoint deployed that extends the OpenFGA API
// - You need full control over the raw JSON bytes before decoding
//
// For the recommended high-level typed approach, see the streamed_list_objects example.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/openfga/language/pkg/go/transformer"

	openfga "github.com/openfga/go-sdk"
	"github.com/openfga/go-sdk/client"
)

func main() {
	ctx := context.Background()

	// Get API URL from environment or use default
	apiUrl := os.Getenv("FGA_API_URL")
	if apiUrl == "" {
		apiUrl = "http://localhost:8080"
	}

	// Create initial client for store creation
	fgaClient, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl: apiUrl,
	})
	if err != nil {
		handleError(err)
		return
	}

	fmt.Println("Creating temporary store")
	store, err := fgaClient.CreateStore(ctx).Body(client.ClientCreateStoreRequest{
		Name: "api-executor-streaming",
	}).Execute()
	if err != nil {
		handleError(err)
		return
	}

	// Create client with store ID
	clientWithStore, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl:  apiUrl,
		StoreId: store.Id,
	})
	if err != nil {
		handleError(err)
		return
	}

	fmt.Println("Writing authorization model")
	authModel, err := writeAuthorizationModel(ctx, clientWithStore)
	if err != nil {
		handleError(err)
		return
	}

	// Create final client with store ID and authorization model ID
	fga, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl:               apiUrl,
		StoreId:              store.Id,
		AuthorizationModelId: authModel.AuthorizationModelId,
	})
	if err != nil {
		handleError(err)
		return
	}

	fmt.Println("Writing tuples (1000 as owner, 1000 as viewer)")
	if err := writeTuples(ctx, fga); err != nil {
		handleError(err)
		return
	}

	// =========================================================================
	// Using the APIExecutor's ExecuteStreaming for streaming
	// =========================================================================
	fmt.Println("\nStreaming objects via APIExecutor...")
	if err := streamObjectsViaExecutor(ctx, fga); err != nil {
		handleError(err)
		return
	}

	fmt.Println("\nCleaning up...")
	if _, err := fga.DeleteStore(ctx).Execute(); err != nil {
		fmt.Printf("Failed to delete store: %v\n", err)
	}

	fmt.Println("Done")
}

// streamObjectsViaExecutor demonstrates using the low-level APIExecutor to stream objects.
// This gives you raw JSON bytes that you decode yourself.
func streamObjectsViaExecutor(ctx context.Context, fga *client.OpenFgaClient) error {
	consistencyPreference := openfga.CONSISTENCYPREFERENCE_HIGHER_CONSISTENCY

	// Get the API executor from the client
	executor := fga.GetAPIExecutor()

	storeId, err := fga.GetStoreId()
	if err != nil {
		return fmt.Errorf("GetStoreId failed: %w", err)
	}

	// Build the streaming request using the builder pattern
	request := openfga.NewAPIExecutorRequestBuilder("StreamedListObjects", "POST", "/stores/{store_id}/streamed-list-objects").
		WithPathParameter("store_id", storeId).
		WithBody(openfga.ListObjectsRequest{
			User:        "user:anne",
			Relation:    "can_read", // Computed: owner OR viewer
			Type:        "document",
			Consistency: &consistencyPreference,
		}).
		Build()

	// Execute the streaming request
	// The Accept header is automatically set to "application/x-ndjson"
	channel, err := executor.ExecuteStreaming(ctx, request, openfga.DefaultStreamBufferSize)
	if err != nil {
		return fmt.Errorf("ExecuteStreaming failed: %w", err)
	}
	defer channel.Close()

	// Process raw JSON bytes from the stream
	count := 0
	for {
		select {
		case result, ok := <-channel.Results:
			if !ok {
				// Results channel closed — stream completed
				// Check for any final errors
				select {
				case err := <-channel.Errors:
					if err != nil {
						return fmt.Errorf("error during streaming: %w", err)
					}
				default:
				}
				fmt.Printf("✓ Streamed %d objects via APIExecutor\n", count)
				return nil
			}

			// Manually decode the raw JSON bytes into the typed response
			var response openfga.StreamedListObjectsResponse
			if err := json.Unmarshal(result, &response); err != nil {
				return fmt.Errorf("failed to decode stream result: %w", err)
			}

			count++
			if count <= 3 || count%500 == 0 {
				fmt.Printf("  Object: %s\n", response.Object)
			}

		case err := <-channel.Errors:
			if err != nil {
				return fmt.Errorf("error during streaming: %w", err)
			}
		}
	}
}

func writeAuthorizationModel(ctx context.Context, fgaClient *client.OpenFgaClient) (*client.ClientWriteAuthorizationModelResponse, error) {
	dslString := `model
  schema 1.1

type user

type document
  relations
    define owner: [user]
    define viewer: [user]
    define can_read: owner or viewer`

	modelJSON, err := transformer.TransformDSLToJSON(dslString)
	if err != nil {
		return nil, fmt.Errorf("failed to transform DSL to JSON: %w", err)
	}

	var authModel openfga.AuthorizationModel
	if err := json.Unmarshal([]byte(modelJSON), &authModel); err != nil {
		return nil, fmt.Errorf("failed to unmarshal authorization model: %w", err)
	}

	return fgaClient.WriteAuthorizationModel(ctx).Body(openfga.WriteAuthorizationModelRequest{
		SchemaVersion:   authModel.SchemaVersion,
		TypeDefinitions: authModel.TypeDefinitions,
	}).Execute()
}

func writeTuples(ctx context.Context, fga *client.OpenFgaClient) error {
	const batchSize = 100
	totalWritten := 0

	for batch := 0; batch < 10; batch++ {
		tuples := make([]client.ClientTupleKey, 0, batchSize)
		for i := 1; i <= batchSize; i++ {
			tuples = append(tuples, client.ClientTupleKey{
				User:     "user:anne",
				Relation: "owner",
				Object:   fmt.Sprintf("document:%d", batch*batchSize+i),
			})
		}
		if _, err := fga.WriteTuples(ctx).Body(tuples).Execute(); err != nil {
			return fmt.Errorf("failed to write owner tuples: %w", err)
		}
		totalWritten += len(tuples)
	}

	for batch := 0; batch < 10; batch++ {
		tuples := make([]client.ClientTupleKey, 0, batchSize)
		for i := 1; i <= batchSize; i++ {
			tuples = append(tuples, client.ClientTupleKey{
				User:     "user:anne",
				Relation: "viewer",
				Object:   fmt.Sprintf("document:%d", 1000+batch*batchSize+i),
			})
		}
		if _, err := fga.WriteTuples(ctx).Body(tuples).Execute(); err != nil {
			return fmt.Errorf("failed to write viewer tuples: %w", err)
		}
		totalWritten += len(tuples)
	}

	fmt.Printf("Wrote %d tuples\n", totalWritten)
	return nil
}

func handleError(err error) {
	if err.Error() == "connection refused" {
		fmt.Fprintln(os.Stderr, "Is OpenFGA server running? Check FGA_API_URL environment variable or default http://localhost:8080")
	} else {
		fmt.Fprintf(os.Stderr, "An error occurred. [%T]\n", err)
	}
	os.Exit(1)
}
