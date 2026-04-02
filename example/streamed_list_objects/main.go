// This example demonstrates how to use the concrete StreamedListObjects client method
// to stream objects via the typed high-level API.
//
// This is the recommended approach for calling StreamedListObjects - it provides
// typed responses (StreamedListObjectsResponse) directly, without requiring
// manual JSON unmarshalling.
//
// For an example using the low-level APIExecutor with raw NDJSON streaming,
// see the api_executor example.

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
		Name: "streamed-list-objects",
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
	// Using the concrete StreamedListObjects client method (recommended)
	// =========================================================================
	fmt.Println("\nStreaming objects via computed 'can_read' relation...")
	if err := streamObjectsViaClient(ctx, fga); err != nil {
		handleError(err)
		return
	}

	fmt.Println("\nCleaning up...")
	if _, err := fga.DeleteStore(ctx).Execute(); err != nil {
		fmt.Printf("Failed to delete store: %v\n", err)
	}

	fmt.Println("Done")
}

// streamObjectsViaClient demonstrates using the typed StreamedListObjects client method.
// This is the recommended approach - it provides typed responses directly.
func streamObjectsViaClient(ctx context.Context, fga *client.OpenFgaClient) error {
	consistency := openfga.CONSISTENCYPREFERENCE_HIGHER_CONSISTENCY

	// Call StreamedListObjects using the fluent client API
	response, err := fga.StreamedListObjects(ctx).Body(client.ClientStreamedListObjectsRequest{
		User:     "user:anne",
		Relation: "can_read", // Computed: owner OR viewer
		Type:     "document",
	}).Options(client.ClientStreamedListObjectsOptions{
		Consistency: &consistency,
	}).Execute()
	if err != nil {
		return fmt.Errorf("StreamedListObjects failed: %w", err)
	}
	defer response.Close()

	// Process typed responses directly - no manual JSON unmarshalling needed!
	count := 0
	for obj := range response.Objects {
		count++
		if count <= 3 || count%500 == 0 {
			fmt.Printf("  Object: %s\n", obj.Object)
		}
	}

	// Check for errors after the stream completes
	if err := <-response.Errors; err != nil {
		return fmt.Errorf("error during streaming: %w", err)
	}

	fmt.Printf("✓ Streamed %d objects via client method\n", count)
	return nil
}

func writeAuthorizationModel(ctx context.Context, fgaClient *client.OpenFgaClient) (*client.ClientWriteAuthorizationModelResponse, error) {
	// Define the authorization model using OpenFGA DSL
	dslString := `model
  schema 1.1

type user

type document
  relations
    define owner: [user]
    define viewer: [user]
    define can_read: owner or viewer`

	// Transform DSL to JSON string
	modelJSON, err := transformer.TransformDSLToJSON(dslString)
	if err != nil {
		return nil, fmt.Errorf("failed to transform DSL to JSON: %w", err)
	}

	// Parse the JSON into the authorization model request
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

	// Write 1000 documents where anne is the owner
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

	// Write 1000 documents where anne is a viewer
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
	// Avoid logging sensitive data; only display generic info
	if err.Error() == "connection refused" {
		fmt.Fprintln(os.Stderr, "Is OpenFGA server running? Check FGA_API_URL environment variable or default http://localhost:8080")
	} else {
		fmt.Fprintf(os.Stderr, "An error occurred. [%T]\n", err)
	}
	os.Exit(1)
}
