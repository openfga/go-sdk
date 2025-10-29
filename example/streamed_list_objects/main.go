package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	openfga "github.com/openfga/go-sdk"
	"github.com/openfga/go-sdk/client"
)

func runSync(ctx context.Context, fgaClient *client.OpenFgaClient, authModelId string, relation string, count int, bufferSize *int) {
	fmt.Println("Mode: sync streaming (range over channel)")
	request := client.ClientStreamedListObjectsRequest{Type: "document", Relation: relation, User: "user:anne"}
	options := client.ClientStreamedListObjectsOptions{}
	if authModelId != "" {
		options.AuthorizationModelId = &authModelId
	}
	if bufferSize != nil {
		options.StreamBufferSize = bufferSize
		fmt.Printf("Using custom buffer size: %d\n", *bufferSize)
	}
	response, err := fgaClient.StreamedListObjects(ctx).Body(request).Options(options).Execute()
	if err != nil {
		log.Fatalf("StreamedListObjects failed: %v", err)
	}
	defer response.Close()

	fmt.Println("Streaming objects (sync):")
	received := 0
	for obj := range response.Objects { // synchronous consumption
		received++
		fmt.Printf("  %d. %s\n", received, obj.Object)
	}
	if err := <-response.Errors; err != nil {
		log.Fatalf("Error during streaming: %v", err)
	}
	fmt.Printf("\nTotal objects received (sync): %d (expected up to %d)\n", received, count)
}

// runAsync demonstrates asynchronous consumption using a goroutine; main goroutine can do other work.
func runAsync(ctx context.Context, fgaClient *client.OpenFgaClient, authModelId string, relation string, count int, bufferSize *int) {
	fmt.Println("Mode: async streaming (consume in goroutine)")
	// Use a cancellable context to show cancellation pattern (not cancelling in this example).
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := client.ClientStreamedListObjectsRequest{Type: "document", Relation: relation, User: "user:anne"}
	options := client.ClientStreamedListObjectsOptions{}
	if authModelId != "" {
		options.AuthorizationModelId = &authModelId
	}
	if bufferSize != nil {
		options.StreamBufferSize = bufferSize
		fmt.Printf("Using custom buffer size: %d\n", *bufferSize)
	}

	response, err := fgaClient.StreamedListObjects(ctx).Body(request).Options(options).Execute()
	if err != nil {
		log.Fatalf("StreamedListObjects failed: %v", err)
	}
	defer response.Close()

	done := make(chan struct{})
	received := 0

	go func() {
		defer close(done)
		for obj := range response.Objects {
			received++
			fmt.Printf("  async -> %d. %s\n", received, obj.Object)
		}
		if err := <-response.Errors; err != nil && err != context.Canceled {
			log.Fatalf("Error during async streaming: %v", err)
		}
	}()

	// Simulate doing other work while streaming happens.
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	fmt.Println("Performing other work while streaming...")
	for {
		select {
		case <-done:
			fmt.Printf("\nTotal objects received (async): %d (expected up to %d)\n", received, count)
			return
		case <-ticker.C:
			fmt.Println("  (main goroutine still free to do work)")
		}
	}
}

func createTestData(ctx context.Context, fgaClient *client.OpenFgaClient, authModelId string, relation string, tupleCount int) (string, error) {
	// Ensure relation is either viewer or owner for this simplified example.
	if relation != "viewer" && relation != "owner" {
		return authModelId, fmt.Errorf("unsupported relation '%s' (use viewer or owner)", relation)
	}

	if authModelId == "" {
		fmt.Println("Creating authorization model...")

		// Provide both viewer and owner relations so user can pick.
		relations := map[string]openfga.Userset{
			"viewer": {This: &map[string]interface{}{}},
			"owner":  {This: &map[string]interface{}{}},
		}

		relationMetadata := map[string]openfga.RelationMetadata{
			"viewer": {DirectlyRelatedUserTypes: &[]openfga.RelationReference{{Type: "user"}}},
			"owner":  {DirectlyRelatedUserTypes: &[]openfga.RelationReference{{Type: "user"}}},
		}

		model := openfga.AuthorizationModel{
			SchemaVersion: "1.1",
			TypeDefinitions: []openfga.TypeDefinition{
				{Type: "user"},
				{Type: "document", Relations: &relations, Metadata: &openfga.Metadata{Relations: &relationMetadata}},
			},
		}

		writeModelResp, err := fgaClient.WriteAuthorizationModel(ctx).Body(client.ClientWriteAuthorizationModelRequest{
			SchemaVersion:   model.SchemaVersion,
			TypeDefinitions: model.TypeDefinitions,
		}).Execute()
		if err != nil {
			return authModelId, fmt.Errorf("failed to create authorization model: %w", err)
		}
		authModelId = writeModelResp.AuthorizationModelId
		fmt.Printf("Created authorization model: %s\n\n", authModelId)
	}

	fmt.Printf("Writing %d test tuples for relation '%s'...\n", tupleCount, relation)
	tuples := make([]client.ClientTupleKey, 0, tupleCount)
	for i := 0; i < tupleCount; i++ {
		tuples = append(tuples, client.ClientTupleKey{User: "user:anne", Relation: relation, Object: fmt.Sprintf("document:%d", i)})
	}
	if _, err := fgaClient.WriteTuples(ctx).Body(tuples).Execute(); err != nil {
		return authModelId, fmt.Errorf("failed to write tuples: %w", err)
	}
	fmt.Printf("Wrote %d test tuples\n\n", len(tuples))
	return authModelId, nil
}

func parseArgs() (mode string, count int, relation string, bufferSize *int) {
	mode = "sync" // default
	relation = "viewer"
	count = 3
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}
	if len(os.Args) > 2 {
		if c, err := strconv.Atoi(os.Args[2]); err == nil && c > 0 {
			count = c
		}
	} else if envCount := os.Getenv("FGA_TUPLE_COUNT"); envCount != "" {
		if c, err := strconv.Atoi(envCount); err == nil && c > 0 {
			count = c
		}
	}
	if len(os.Args) > 3 {
		relation = os.Args[3]
	} else if envRel := os.Getenv("FGA_RELATION"); envRel != "" {
		relation = envRel
	}
	if len(os.Args) > 4 {
		if b, err := strconv.Atoi(os.Args[4]); err == nil && b > 0 {
			bufferSize = &b
		}
	} else if envBuffer := os.Getenv("FGA_BUFFER_SIZE"); envBuffer != "" {
		if b, err := strconv.Atoi(envBuffer); err == nil && b > 0 {
			bufferSize = &b
		}
	}
	return
}

func main() {
	ctx := context.Background()
	apiUrl := os.Getenv("FGA_API_URL")
	if apiUrl == "" {
		apiUrl = "http://localhost:8080"
	}
	config := client.ClientConfiguration{ApiUrl: apiUrl}
	fgaClient, err := client.NewSdkClient(&config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create store unless provided via env var.
	storeId := os.Getenv("FGA_STORE_ID")
	createdTempStore := false
	if storeId == "" {
		fmt.Println("Creating Test Store for streamed list objects")
		store, err := fgaClient.CreateStore(ctx).Body(client.ClientCreateStoreRequest{Name: "Test Store"}).Execute()
		if err != nil {
			log.Fatalf("failed to create store: %v", err)
		}
		storeId = store.Id
		createdTempStore = true
	}

	// Re-init client with storeId
	config = client.ClientConfiguration{ApiUrl: apiUrl, StoreId: storeId}
	fgaClient, err = client.NewSdkClient(&config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Println("OpenFGA StreamedListObjects Example")
	fmt.Println("____________________________________")
	fmt.Printf("API URL: %s\n", apiUrl)
	fmt.Printf("Store ID: %s\n", storeId)
	fmt.Println()

	authModelId := os.Getenv("FGA_MODEL_ID")
	mode, tupleCount, relation, bufferSize := parseArgs()
	if authModelId != "" {
		fmt.Printf("Authorization Model ID (provided): %s\n\n", authModelId)
	}
	authModelId, err = createTestData(ctx, fgaClient, authModelId, relation, tupleCount)
	if err != nil {
		log.Printf("Warning: Failed to create test data: %v", err)
		log.Println("Continuing with example...")
	}

	fmt.Printf("Selected mode: %s | relation: %s | tuple count: %d", mode, relation, tupleCount)
	if bufferSize != nil {
		fmt.Printf(" | buffer size: %d", *bufferSize)
	}
	fmt.Printf(" (pass 'sync|async [count] [relation] [bufferSize]')\n\n")

	switch mode {
	case "async":
		runAsync(ctx, fgaClient, authModelId, relation, tupleCount, bufferSize)
	case "sync":
		runSync(ctx, fgaClient, authModelId, relation, tupleCount, bufferSize)
	default:
		fmt.Printf("Unknown mode '%s'. Use 'sync' or 'async'.\n", mode)
		os.Exit(1)
	}

	if createdTempStore {
		fmt.Println("\nDeleting temporary store...")
		if _, err := fgaClient.DeleteStore(ctx).Execute(); err != nil {
			fmt.Printf("Failed to delete store: %v\n", err)
		} else {
			fmt.Printf("Deleted temporary store (%s)\n", storeId)
		}
	}
	fmt.Println("\nDone.")
}
