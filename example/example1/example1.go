package main

import (
	"context"
	"fmt"
	"os"

	openfga "github.com/openfga/go-sdk"
	"github.com/openfga/go-sdk/client"
	"github.com/openfga/go-sdk/credentials"
)

func mainInner() error {

	ctx := context.Background()
	creds := credentials.Credentials{}
	if os.Getenv("FGA_CLIENT_ID") != "" {
		creds = credentials.Credentials{
			Method: credentials.CredentialsMethodClientCredentials,
			Config: &credentials.Config{
				ClientCredentialsClientId:       os.Getenv("FGA_CLIENT_ID"),
				ClientCredentialsClientSecret:   os.Getenv("FGA_CLIENT_SECRET"),
				ClientCredentialsApiAudience:    os.Getenv("FGA_API_AUDIENCE"),
				ClientCredentialsApiTokenIssuer: os.Getenv("FGA_API_TOKEN_ISSUER"),
			},
		}
	}

	apiUrl := os.Getenv("FGA_API_URL")
	if apiUrl == "" {
		apiUrl = "http://localhost:8080"
	}
	fgaClient, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl:               apiUrl,
		StoreId:              os.Getenv("FGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
		AuthorizationModelId: os.Getenv("FGA_MODEL_ID"), // optional, recommended to be set for production
		Credentials:          &creds,
	})

	if err != nil {
		return err
	}

	// ListStores
	fmt.Println("Listing Stores")
	stores1, err := fgaClient.ListStores(ctx).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Stores Count: %d\n", len(stores1.GetStores()))

	// CreateStore
	fmt.Println("Creating Test Store")
	store, err := fgaClient.CreateStore(ctx).Body(client.ClientCreateStoreRequest{Name: "Test Store"}).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Test Store ID: %v\n", store.Id)

	// Set the store id
	fgaClient.SetStoreId(store.Id)

	// ListStores after Create
	fmt.Println("Listing Stores")
	stores, err := fgaClient.ListStores(ctx).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Stores Count: %d\n", len(stores.Stores))

	// ListStores with name filter
	fmt.Println("Listing Stores with name filter")
	storesWithFilter, err := fgaClient.ListStores(ctx).Options(client.ClientListStoresOptions{
		Name: openfga.PtrString("Test Store"),
	}).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Stores with name filter Count: %d\n", len(storesWithFilter.Stores))

	// GetStore
	fmt.Println("Getting Current Store")
	currentStore, err := fgaClient.GetStore(ctx).Execute()
	if err != nil {
		return err
	}
	fmt.Println("Current Store Name: %v\n" + currentStore.Name)

	// ReadAuthorizationModels
	fmt.Println("Reading Authorization Models")
	models, err := fgaClient.ReadAuthorizationModels(ctx).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Models Count: %d\n", len(models.AuthorizationModels))

	// ReadLatestAuthorizationModel
	latestAuthorizationModel, err := fgaClient.ReadLatestAuthorizationModel(ctx).Execute()
	if err != nil {
		return err
	}
	if latestAuthorizationModel.AuthorizationModel != nil {
		fmt.Printf("Latest Authorization Model ID: %v\n", (*latestAuthorizationModel.AuthorizationModel).Id)
	} else {
		fmt.Println("Latest Authorization Model not found")
	}

	// WriteAuthorizationModel
	fmt.Println("Writing an Authorization Model")
	body := client.ClientWriteAuthorizationModelRequest{
		SchemaVersion: "1.1",
		TypeDefinitions: []openfga.TypeDefinition{
			{
				Type:      "user",
				Relations: &map[string]openfga.Userset{},
			},
			{
				Type: "document",
				Relations: &map[string]openfga.Userset{
					"writer": {This: &map[string]interface{}{}},
					"viewer": {Union: &openfga.Usersets{
						Child: []openfga.Userset{
							{This: &map[string]interface{}{}},
							{ComputedUserset: &openfga.ObjectRelation{
								Object:   openfga.PtrString(""),
								Relation: openfga.PtrString("writer"),
							}},
						},
					}},
				},
				Metadata: &openfga.Metadata{
					Relations: &map[string]openfga.RelationMetadata{
						"writer": {
							DirectlyRelatedUserTypes: &[]openfga.RelationReference{
								{Type: "user"},
								{Type: "user", Condition: openfga.PtrString("ViewCountLessThan200")},
							},
						},
						"viewer": {
							DirectlyRelatedUserTypes: &[]openfga.RelationReference{
								{Type: "user"},
							},
						},
					},
				},
			},
		},
		Conditions: &map[string]openfga.Condition{
			"ViewCountLessThan200": {
				Name:       "ViewCountLessThan200",
				Expression: "ViewCount < 200",
				Parameters: &map[string]openfga.ConditionParamTypeRef{
					"ViewCount": {
						TypeName: openfga.TYPENAME_INT,
					},
					"Type": {
						TypeName: openfga.TYPENAME_STRING,
					},
					"Name": {
						TypeName: openfga.TYPENAME_STRING,
					},
				},
			},
		},
	}
	authorizationModel, err := fgaClient.WriteAuthorizationModel(ctx).Body(body).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Authorization Model ID: %v\n", authorizationModel.AuthorizationModelId)

	// ReadAuthorizationModels - after Write
	fmt.Println("Reading Authorization Models")
	models, err = fgaClient.ReadAuthorizationModels(ctx).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Models Count: %d\n", len(models.AuthorizationModels))

	// ReadLatestAuthorizationModel - after Write
	latestAuthorizationModel, err = fgaClient.ReadLatestAuthorizationModel(ctx).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Latest Authorization Model ID: %v\n", (*latestAuthorizationModel.AuthorizationModel).Id)

	// Write
	fmt.Println("Writing Tuples")
	_, err = fgaClient.Write(ctx).Body(client.ClientWriteRequest{
		Writes: []client.ClientTupleKey{
			{
				User:     "user:anne",
				Relation: "writer",
				Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				Condition: &openfga.RelationshipCondition{
					Name:    "ViewCountLessThan200",
					Context: &map[string]interface{}{"Name": "Roadmap", "Type": "document"},
				},
			},
		},
	}).Options(client.ClientWriteOptions{
		AuthorizationModelId: &authorizationModel.AuthorizationModelId,
	}).Execute()
	if err != nil {
		return err
	}
	fmt.Println("Done Writing Tuples")

	// Set the model ID
	err = fgaClient.SetAuthorizationModelId(latestAuthorizationModel.AuthorizationModel.Id)
	if err != nil {
		return err
	}

	// Read
	fmt.Println("Reading Tuples")
	readTuples, err := fgaClient.Read(ctx).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Read Tuples: %v\n", readTuples)

	// ReadChanges
	fmt.Println("Reading Tuple Changes")
	readChangesTuples, err := fgaClient.ReadChanges(ctx).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Read Changes Tuples: %v\n", readChangesTuples)

	// Check
	fmt.Println("Checking for access")
	failingCheckResponse, err := fgaClient.Check(ctx).Body(client.ClientCheckRequest{
		User:     "user:anne",
		Relation: "viewer",
		Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
	}).Execute()
	if err != nil {
		fmt.Printf("Failed due to: %w\n", err.Error())
	} else {
		fmt.Printf("Allowed: %v\n", failingCheckResponse.Allowed)
	}

	// Checking for access with context
	fmt.Println("Checking for access with context")
	checkResponse, err := fgaClient.Check(ctx).Body(client.ClientCheckRequest{
		User:     "user:anne",
		Relation: "viewer",
		Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
		Context:  &map[string]interface{}{"ViewCount": 100},
	}).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Allowed: %v\n", checkResponse.Allowed)

	fmt.Println("Checking for access with custom headers")
	checkWithHeadersResponse, err := fgaClient.Check(ctx).Body(client.ClientCheckRequest{
		User:     "user:anne",
		Relation: "viewer",
		Object:   "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
		Context:  &map[string]interface{}{"ViewCount": 100},
	}).Options(client.ClientCheckOptions{
		RequestOptions: client.RequestOptions{
			Headers: map[string]string{
				"X-Request-ID": "example-request-123",
			},
		},
	}).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Allowed (with custom headers): %v\n", checkWithHeadersResponse.Allowed)

	// BatchCheck
	fmt.Println("Batch checking for access")
	batchCheckResponse, err := fgaClient.BatchCheck(ctx).Body(client.ClientBatchCheckRequest{
		Checks: []client.ClientBatchCheckItem{
			{
				CorrelationId: "f278708f-298c-4f43-a893-11a02bbf251c",
				User:          "user:anne",
				Relation:      "viewer",
				Object:        "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				Context:       &map[string]interface{}{"ViewCount": 100},
			},
			{
				CorrelationId: "9f7563d6-2573-4292-9ba2-62d59b97c4d",
				User:          "user:bob",
				Relation:      "viewer",
				Object:        "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
				Context:       &map[string]interface{}{"ViewCount": 100},
			},
		},
	}).Execute()
	if err != nil {
		return err
	}
	fmt.Println("BatchCheck results:")
	for correlationID, result := range batchCheckResponse.GetResult() {
		fmt.Printf("Correlation %s - Allowed: %v\n", correlationID, result.GetAllowed())
	}

	// ListObjects
	fmt.Println("Listing objects user has access to")
	listObjectsResponse, err := fgaClient.ListObjects(ctx).Body(client.ClientListObjectsRequest{
		User:     "user:anne",
		Relation: "viewer",
		Type:     "document",
		Context:  &map[string]interface{}{"ViewCount": 100},
	}).Execute()
	fmt.Printf("Response: Objects = %v\n", listObjectsResponse.Objects)

	// ListRelations
	fmt.Println("Listing relations user has with object")
	listRelationsResponse, err := fgaClient.ListRelations(ctx).Body(client.ClientListRelationsRequest{
		User:      "user:anne",
		Object:    "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
		Relations: []string{"viewer"},
	}).Execute()
	fmt.Printf("Response: Relations = %v\n", listRelationsResponse.Relations)

	// ListUsers
	fmt.Println("Listing user who have access to object")
	listUsersResponse, err := fgaClient.ListUsers(ctx).Body(client.ClientListUsersRequest{
		Relation: "viewer",
		Object: openfga.FgaObject{
			Type: "document",
			Id:   "roadmap",
		},
		UserFilters: []openfga.UserTypeFilter{{
			Type: "user",
		}},
	}).Execute()
	fmt.Printf("Response: Users = %v\n", listUsersResponse.Users)

	// WriteAssertions
	_, err = fgaClient.WriteAssertions(ctx).Body([]client.ClientAssertion{
		{
			User:        "user:carl",
			Relation:    "writer",
			Object:      "document:budget",
			Expectation: true,
			Context:     &map[string]interface{}{"Name": "Roadmap", "Type": "document"},
			ContextualTuples: []client.ClientContextualTupleKey{
				{
					User:     "user:carl",
					Relation: "writer",
					Object:   "document:budget",
				},
			},
		},
		{
			User:        "user:anne",
			Relation:    "viewer",
			Object:      "document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a",
			Expectation: false,
		},
	}).Execute()
	if err != nil {
		return err
	}
	fmt.Println("Assertions updated")

	// ReadAssertions
	fmt.Println("Reading Assertions")
	assertions, err := fgaClient.ReadAssertions(ctx).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Assertions: %v\n", assertions.GetAssertions())

	// DeleteStore
	fmt.Println("Deleting Current Store")
	_, err = fgaClient.DeleteStore(ctx).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("Deleted Store: %v\n", currentStore.Name)

	return nil
}

func main() {
	if err := mainInner(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
