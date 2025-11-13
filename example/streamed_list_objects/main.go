package main

import (
	"context"
	"fmt"
	"os"

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
	config := client.ClientConfiguration{
		ApiUrl: apiUrl,
	}
	fgaClient, err := client.NewSdkClient(&config)
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

	fmt.Println("Streaming objects via computed 'can_read' relation...")
	if err := streamObjects(ctx, fga); err != nil {
		handleError(err)
		return
	}

	fmt.Println("Cleaning up...")
	if _, err := fga.DeleteStore(ctx).Execute(); err != nil {
		fmt.Printf("Failed to delete store: %v\n", err)
	}

	fmt.Println("Done")
}

func writeAuthorizationModel(ctx context.Context, fgaClient *client.OpenFgaClient) (*client.ClientWriteAuthorizationModelResponse, error) {
	// Define the authorization model with computed relations
	ownerUserset := openfga.Userset{This: &map[string]interface{}{}}
	viewerUserset := openfga.Userset{This: &map[string]interface{}{}}
	canReadUserset := openfga.Userset{
		Union: &openfga.Usersets{
			Child: []openfga.Userset{
				{ComputedUserset: &openfga.ObjectRelation{
					Object:   openfga.PtrString(""),
					Relation: openfga.PtrString("owner"),
				}},
				{ComputedUserset: &openfga.ObjectRelation{
					Object:   openfga.PtrString(""),
					Relation: openfga.PtrString("viewer"),
				}},
			},
		},
	}

	relations := map[string]openfga.Userset{
		"owner":    ownerUserset,
		"viewer":   viewerUserset,
		"can_read": canReadUserset,
	}

	relationMetadata := map[string]openfga.RelationMetadata{
		"owner": {
			DirectlyRelatedUserTypes: &[]openfga.RelationReference{
				{Type: "user"},
			},
		},
		"viewer": {
			DirectlyRelatedUserTypes: &[]openfga.RelationReference{
				{Type: "user"},
			},
		},
		"can_read": {
			DirectlyRelatedUserTypes: &[]openfga.RelationReference{},
		},
	}

	return fgaClient.WriteAuthorizationModel(ctx).Body(openfga.WriteAuthorizationModelRequest{
		SchemaVersion: "1.1",
		TypeDefinitions: []openfga.TypeDefinition{
			{
				Type:      "user",
				Relations: &map[string]openfga.Userset{},
			},
			{
				Type:      "document",
				Relations: &relations,
				Metadata:  &openfga.Metadata{Relations: &relationMetadata},
			},
		},
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

func streamObjects(ctx context.Context, fga *client.OpenFgaClient) error {
	consistencyPreference := openfga.CONSISTENCYPREFERENCE_HIGHER_CONSISTENCY

	response, err := fga.StreamedListObjects(ctx).Body(client.ClientStreamedListObjectsRequest{
		User:     "user:anne",
		Relation: "can_read", // Computed: owner OR viewer
		Type:     "document",
	}).Options(client.ClientStreamedListObjectsOptions{
		Consistency: &consistencyPreference,
	}).Execute()
	if err != nil {
		return fmt.Errorf("StreamedListObjects failed: %w", err)
	}
	defer response.Close()

	count := 0
	for obj := range response.Objects {
		count++
		if count <= 3 || count%500 == 0 {
			fmt.Printf("- %s\n", obj.Object)
		}
	}

	// Check for streaming errors
	if err := <-response.Errors; err != nil {
		return fmt.Errorf("error during streaming: %w", err)
	}

	fmt.Printf("✓ Streamed %d objects\n", count)
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
