//go:build integration

package integration

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	openfga "github.com/openfga/go-sdk"
	"github.com/openfga/go-sdk/client"
	"github.com/openfga/go-sdk/credentials"
)

const presharedKey = "integration-test-key"

// sharedContainer holds a single OpenFGA container that is reused across all
// tests in this package. It is started lazily on first use and terminated after
// all tests complete via TestMain.
var (
	sharedContainerOnce sync.Once
	sharedAPIURL        string
	sharedContainer     testcontainers.Container
)

// getAPIURL returns the base URL of the shared OpenFGA container, starting it
// if necessary. The container runs with preshared-key authentication enabled.
func getAPIURL(t *testing.T) string {
	t.Helper()

	sharedContainerOnce.Do(func() {
		ctx := context.Background()

		req := testcontainers.ContainerRequest{
			Image:        "openfga/openfga:latest",
			ExposedPorts: []string{"8080/tcp"},
			Cmd: []string{
				"run",
				"--authn-method=preshared",
				"--authn-preshared-keys=" + presharedKey,
			},
			WaitingFor: wait.ForHTTP("/healthz").
				WithPort("8080/tcp").
				WithStartupTimeout(30 * time.Second),
		}

		container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
		if err != nil {
			t.Fatalf("failed to start OpenFGA container: %v", err)
		}
		sharedContainer = container

		host, err := container.Host(ctx)
		if err != nil {
			t.Fatalf("failed to get container host: %v", err)
		}
		port, err := container.MappedPort(ctx, "8080")
		if err != nil {
			t.Fatalf("failed to get mapped port: %v", err)
		}

		sharedAPIURL = fmt.Sprintf("http://%s:%s", host, port.Port())
	})

	if sharedAPIURL == "" {
		t.Fatal("shared OpenFGA container failed to start")
	}

	return sharedAPIURL
}

// terminateSharedContainer stops the shared container. Called from TestMain.
func terminateSharedContainer() {
	if sharedContainer != nil {
		_ = sharedContainer.Terminate(context.Background())
	}
}

// newAuthenticatedCreds returns credentials configured with the preshared key.
func newAuthenticatedCreds() *credentials.Credentials {
	return &credentials.Credentials{
		Method: credentials.CredentialsMethodApiToken,
		Config: &credentials.Config{ApiToken: presharedKey},
	}
}

// setupStoreWithTuples creates a fresh store, writes an authorization model and
// tuples, and returns a fully-configured client along with the store and model IDs.
// The store is deleted when the test finishes.
func setupStoreWithTuples(t *testing.T) (*client.OpenFgaClient, string, string) {
	t.Helper()
	apiURL := getAPIURL(t)
	ctx := context.Background()
	creds := newAuthenticatedCreds()

	// Create a client without a store to bootstrap one.
	bootstrap, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl: apiURL, Credentials: creds,
	})
	if err != nil {
		t.Fatalf("NewSdkClient: %v", err)
	}

	store, err := bootstrap.CreateStore(ctx).Body(client.ClientCreateStoreRequest{
		Name: "integration-" + t.Name(),
	}).Execute()
	if err != nil {
		t.Fatalf("CreateStore: %v", err)
	}
	t.Cleanup(func() {
		c, _ := client.NewSdkClient(&client.ClientConfiguration{
			ApiUrl: apiURL, StoreId: store.Id, Credentials: creds,
		})
		_, _ = c.DeleteStore(context.Background()).Execute()
	})

	// Client with store.
	withStore, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl: apiURL, StoreId: store.Id, Credentials: creds,
	})
	if err != nil {
		t.Fatalf("NewSdkClient (with store): %v", err)
	}

	// Write authorization model.
	modelResp, err := withStore.WriteAuthorizationModel(ctx).Body(openfga.WriteAuthorizationModelRequest{
		SchemaVersion: "1.1",
		TypeDefinitions: []openfga.TypeDefinition{
			{Type: "user", Relations: &map[string]openfga.Userset{}},
			{
				Type: "document",
				Relations: &map[string]openfga.Userset{
					"viewer": {This: &map[string]interface{}{}},
				},
				Metadata: &openfga.Metadata{
					Relations: &map[string]openfga.RelationMetadata{
						"viewer": {
							DirectlyRelatedUserTypes: &[]openfga.RelationReference{{Type: "user"}},
						},
					},
				},
			},
		},
	}).Execute()
	if err != nil {
		t.Fatalf("WriteAuthorizationModel: %v", err)
	}
	modelID := modelResp.AuthorizationModelId

	// Final client with store + model.
	fgaClient, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl: apiURL, StoreId: store.Id, AuthorizationModelId: modelID, Credentials: creds,
	})
	if err != nil {
		t.Fatalf("NewSdkClient (full): %v", err)
	}

	// Write tuples.
	_, err = fgaClient.WriteTuples(ctx).Body([]client.ClientTupleKey{
		{User: "user:anne", Relation: "viewer", Object: "document:1"},
		{User: "user:anne", Relation: "viewer", Object: "document:2"},
		{User: "user:anne", Relation: "viewer", Object: "document:3"},
	}).Execute()
	if err != nil {
		t.Fatalf("WriteTuples: %v", err)
	}

	return fgaClient, store.Id, modelID
}

// requireAuthEnforced is a sanity check that unauthenticated requests are
// rejected by the shared container.
func requireAuthEnforced(t *testing.T) {
	t.Helper()
	resp, err := http.Get(getAPIURL(t) + "/stores")
	if err != nil {
		t.Fatalf("cannot reach OpenFGA: %v", err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected 401 for unauthenticated request, got %d — is authn enabled?", resp.StatusCode)
	}
}
