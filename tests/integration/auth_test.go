// auth_test.go — Verify that the Authorization header is sent for every request
// path when using preshared-key authentication.

//go:build integration

package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	openfga "github.com/openfga/go-sdk"
	"github.com/openfga/go-sdk/client"
)

func TestAuthOnCheck(t *testing.T) {
	requireAuthEnforced(t)
	fgaClient, _, _ := setupStoreWithTuples(t)

	resp, err := fgaClient.Check(context.Background()).Body(client.ClientCheckRequest{
		User: "user:anne", Relation: "viewer", Object: "document:1",
	}).Execute()

	require.NoError(t, err, "Check should succeed with auth")
	assert.True(t, resp.GetAllowed(), "expected allowed=true")
}

func TestAuthOnListObjects(t *testing.T) {
	requireAuthEnforced(t)
	fgaClient, _, _ := setupStoreWithTuples(t)

	resp, err := fgaClient.ListObjects(context.Background()).Body(client.ClientListObjectsRequest{
		User: "user:anne", Relation: "viewer", Type: "document",
	}).Execute()

	require.NoError(t, err, "ListObjects should succeed with auth")
	assert.Len(t, resp.GetObjects(), 3)
}

func TestAuthOnStreamedListObjects(t *testing.T) {
	requireAuthEnforced(t)
	fgaClient, _, _ := setupStoreWithTuples(t)

	consistency := openfga.CONSISTENCYPREFERENCE_HIGHER_CONSISTENCY
	response, err := fgaClient.StreamedListObjects(context.Background()).Body(client.ClientStreamedListObjectsRequest{
		User: "user:anne", Relation: "viewer", Type: "document",
	}).Options(client.ClientStreamedListObjectsOptions{
		Consistency: &consistency,
	}).Execute()
	require.NoError(t, err, "StreamedListObjects should succeed with auth")
	defer response.Close()

	var objects []string
	for obj := range response.Objects {
		objects = append(objects, obj.Object)
	}
	require.NoError(t, <-response.Errors, "stream should complete without error")
	assert.Len(t, objects, 3)
}

func TestAuthOnExecutorExecute(t *testing.T) {
	requireAuthEnforced(t)
	fgaClient, storeID, modelID := setupStoreWithTuples(t)

	executor := fgaClient.GetAPIExecutor()
	var checkResp openfga.CheckResponse
	_, err := executor.ExecuteWithDecode(context.Background(),
		openfga.NewAPIExecutorRequestBuilder("Check", http.MethodPost, "/stores/{store_id}/check").
			WithPathParameter("store_id", storeID).
			WithBody(openfga.CheckRequest{
				TupleKey: openfga.CheckRequestTupleKey{
					User: "user:anne", Relation: "viewer", Object: "document:1",
				},
				AuthorizationModelId: &modelID,
			}).
			Build(),
		&checkResp,
	)

	require.NoError(t, err, "Executor.Execute should succeed with auth")
	assert.True(t, *checkResp.Allowed, "expected allowed=true")
}

func TestAuthOnExecutorExecuteStreaming(t *testing.T) {
	requireAuthEnforced(t)
	fgaClient, storeID, modelID := setupStoreWithTuples(t)

	executor := fgaClient.GetAPIExecutor()
	channel, err := executor.ExecuteStreaming(context.Background(),
		openfga.NewAPIExecutorRequestBuilder("StreamedListObjects", http.MethodPost, "/stores/{store_id}/streamed-list-objects").
			WithPathParameter("store_id", storeID).
			WithBody(openfga.ListObjectsRequest{
				AuthorizationModelId: &modelID,
				User:                 "user:anne",
				Relation:             "viewer",
				Type:                 "document",
				Consistency:          openfga.CONSISTENCYPREFERENCE_HIGHER_CONSISTENCY.Ptr(),
			}).
			Build(),
		openfga.DefaultStreamBufferSize,
	)
	require.NoError(t, err, "Executor.ExecuteStreaming should succeed with auth")
	defer channel.Close()

	var objects []string
	for {
		select {
		case result, ok := <-channel.Results:
			if !ok {
				select {
				case err := <-channel.Errors:
					require.NoError(t, err, "stream should complete without error")
				default:
				}
				goto done
			}
			var obj openfga.StreamedListObjectsResponse
			require.NoError(t, json.Unmarshal(result, &obj))
			objects = append(objects, obj.Object)
		case err := <-channel.Errors:
			require.NoError(t, err, "unexpected stream error")
		}
	}
done:
	assert.Len(t, objects, 3)
}

func TestAuthRejectedWithoutCredentials(t *testing.T) {
	requireAuthEnforced(t)

	fgaClient, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl: getAPIURL(t),
	})
	require.NoError(t, err)

	_, err = fgaClient.ListStores(context.Background()).Execute()
	require.Error(t, err, "unauthenticated request should be rejected")
}
