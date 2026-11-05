package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fgaSdkClient "github.com/openfga/go-sdk/client"
)

const (
	clientMethodHeader        = "X-Openfga-Client-Method"
	clientBulkRequestIdHeader = "X-Openfga-Client-Bulk-Request-Id"
)

var uuidV4Pattern = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

// headerRecorder serves responseBody to every request and records each request's headers.
type headerRecorder struct {
	mu       sync.Mutex
	requests []http.Header
}

func (r *headerRecorder) all() []http.Header {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]http.Header(nil), r.requests...)
}

func newRecordingClient(t *testing.T, responseBody string) (*fgaSdkClient.OpenFgaClient, *headerRecorder) {
	t.Helper()

	recorder := &headerRecorder{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		recorder.mu.Lock()
		recorder.requests = append(recorder.requests, req.Header.Clone())
		recorder.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(responseBody))
	}))
	t.Cleanup(server.Close)

	client, err := fgaSdkClient.NewSdkClient(&fgaSdkClient.ClientConfiguration{
		ApiUrl:     server.URL,
		StoreId:    "01H0H015178Y2V4CX10C2KGHF4",
		HTTPClient: &http.Client{},
	})
	require.NoError(t, err)

	return client, recorder
}

// assertSharedBulkRequestId checks that every request carries the same, well-formed bulk request ID and returns it.
func assertSharedBulkRequestId(t *testing.T, requests []http.Header) string {
	t.Helper()

	require.NotEmpty(t, requests)
	bulkId := requests[0].Get(clientBulkRequestIdHeader)
	assert.Regexp(t, uuidV4Pattern, bulkId)
	for _, headers := range requests {
		assert.Equal(t, bulkId, headers.Get(clientBulkRequestIdHeader))
	}
	return bulkId
}

func TestClientMethodHeaders(t *testing.T) {
	t.Parallel()

	t.Run("ClientBatchCheck", func(t *testing.T) {
		t.Parallel()
		client, recorder := newRecordingClient(t, `{"allowed":true}`)

		_, err := client.ClientBatchCheck(context.Background()).Body(fgaSdkClient.ClientBatchCheckClientBody{
			{User: "user:anne", Relation: "viewer", Object: "document:1"},
			{User: "user:anne", Relation: "viewer", Object: "document:2"},
			{User: "user:anne", Relation: "viewer", Object: "document:3"},
		}).Execute()
		require.NoError(t, err)

		requests := recorder.all()
		require.Len(t, requests, 3)
		for _, headers := range requests {
			assert.Equal(t, "ClientBatchCheck", headers.Get(clientMethodHeader))
		}
		assertSharedBulkRequestId(t, requests)
	})

	t.Run("ListRelations", func(t *testing.T) {
		t.Parallel()
		client, recorder := newRecordingClient(t, `{"allowed":true}`)

		_, err := client.ListRelations(context.Background()).Body(fgaSdkClient.ClientListRelationsRequest{
			User:      "user:anne",
			Object:    "document:1",
			Relations: []string{"viewer", "editor"},
		}).Execute()
		require.NoError(t, err)

		requests := recorder.all()
		require.Len(t, requests, 2)
		for _, headers := range requests {
			assert.Equal(t, "ListRelations", headers.Get(clientMethodHeader))
		}
		assertSharedBulkRequestId(t, requests)
	})

	t.Run("BatchCheck", func(t *testing.T) {
		t.Parallel()
		client, recorder := newRecordingClient(t, `{"result":{}}`)

		batchSize := int32(1)
		_, err := client.BatchCheck(context.Background()).Body(fgaSdkClient.ClientBatchCheckRequest{
			Checks: []fgaSdkClient.ClientBatchCheckItem{
				{User: "user:anne", Relation: "viewer", Object: "document:1", CorrelationId: "1"},
				{User: "user:anne", Relation: "viewer", Object: "document:2", CorrelationId: "2"},
			},
		}).Options(fgaSdkClient.BatchCheckOptions{MaxBatchSize: &batchSize}).Execute()
		require.NoError(t, err)

		requests := recorder.all()
		require.Len(t, requests, 2)
		for _, headers := range requests {
			assert.Empty(t, headers.Get(clientMethodHeader))
		}
		assertSharedBulkRequestId(t, requests)
	})

	t.Run("WriteWithTransactionsDisabled", func(t *testing.T) {
		t.Parallel()
		client, recorder := newRecordingClient(t, `{}`)

		_, err := client.Write(context.Background()).Body(fgaSdkClient.ClientWriteRequest{
			Writes: []fgaSdkClient.ClientTupleKey{
				{User: "user:anne", Relation: "viewer", Object: "document:1"},
				{User: "user:anne", Relation: "viewer", Object: "document:2"},
			},
		}).Options(fgaSdkClient.ClientWriteOptions{
			Transaction: &fgaSdkClient.TransactionOptions{Disable: true, MaxPerChunk: 1},
		}).Execute()
		require.NoError(t, err)

		requests := recorder.all()
		require.Len(t, requests, 2)
		for _, headers := range requests {
			assert.Equal(t, "Write", headers.Get(clientMethodHeader))
		}
		assertSharedBulkRequestId(t, requests)
	})

	t.Run("WriteInTransaction", func(t *testing.T) {
		t.Parallel()
		client, recorder := newRecordingClient(t, `{}`)

		_, err := client.Write(context.Background()).Body(fgaSdkClient.ClientWriteRequest{
			Writes: []fgaSdkClient.ClientTupleKey{{User: "user:anne", Relation: "viewer", Object: "document:1"}},
		}).Execute()
		require.NoError(t, err)

		requests := recorder.all()
		require.Len(t, requests, 1)
		assert.Empty(t, requests[0].Get(clientMethodHeader))
		assert.Empty(t, requests[0].Get(clientBulkRequestIdHeader))
	})

	t.Run("WriteTuples", func(t *testing.T) {
		t.Parallel()
		client, recorder := newRecordingClient(t, `{}`)

		_, err := client.WriteTuples(context.Background()).Body(fgaSdkClient.ClientWriteTuplesBody{
			{User: "user:anne", Relation: "viewer", Object: "document:1"},
		}).Execute()
		require.NoError(t, err)

		requests := recorder.all()
		require.Len(t, requests, 1)
		assert.Equal(t, "WriteTuples", requests[0].Get(clientMethodHeader))
		assert.Empty(t, requests[0].Get(clientBulkRequestIdHeader))
	})

	t.Run("DeleteTuples", func(t *testing.T) {
		t.Parallel()
		client, recorder := newRecordingClient(t, `{}`)

		_, err := client.DeleteTuples(context.Background()).Body(fgaSdkClient.ClientDeleteTuplesBody{
			{User: "user:anne", Relation: "viewer", Object: "document:1"},
		}).Execute()
		require.NoError(t, err)

		requests := recorder.all()
		require.Len(t, requests, 1)
		assert.Equal(t, "DeleteTuples", requests[0].Get(clientMethodHeader))
	})

	t.Run("ReadLatestAuthorizationModel", func(t *testing.T) {
		t.Parallel()
		client, recorder := newRecordingClient(t, `{"authorization_models":[]}`)

		_, err := client.ReadLatestAuthorizationModel(context.Background()).Execute()
		require.NoError(t, err)

		requests := recorder.all()
		require.Len(t, requests, 1)
		assert.Equal(t, "ReadLatestAuthorizationModel", requests[0].Get(clientMethodHeader))
		assert.Empty(t, requests[0].Get(clientBulkRequestIdHeader))
	})

	t.Run("CallerHeadersAreKept", func(t *testing.T) {
		t.Parallel()
		client, recorder := newRecordingClient(t, `{"allowed":true}`)

		callerHeaders := map[string]string{
			"X-OpenFGA-Client-Method":          "MyMethod",
			"X-OpenFGA-Client-Bulk-Request-Id": "my-bulk-id",
		}
		_, err := client.ClientBatchCheck(context.Background()).Body(fgaSdkClient.ClientBatchCheckClientBody{
			{User: "user:anne", Relation: "viewer", Object: "document:1"},
		}).Options(fgaSdkClient.ClientBatchCheckClientOptions{
			RequestOptions: fgaSdkClient.RequestOptions{Headers: callerHeaders},
		}).Execute()
		require.NoError(t, err)

		requests := recorder.all()
		require.Len(t, requests, 1)
		assert.Equal(t, "MyMethod", requests[0].Get(clientMethodHeader))
		assert.Equal(t, "my-bulk-id", requests[0].Get(clientBulkRequestIdHeader))
		assert.Len(t, callerHeaders, 2, "the caller's headers map must not be modified")
	})

	t.Run("EachCallGetsItsOwnBulkRequestId", func(t *testing.T) {
		t.Parallel()
		client, recorder := newRecordingClient(t, `{"allowed":true}`)

		body := fgaSdkClient.ClientBatchCheckClientBody{{User: "user:anne", Relation: "viewer", Object: "document:1"}}
		_, err := client.ClientBatchCheck(context.Background()).Body(body).Execute()
		require.NoError(t, err)
		_, err = client.ClientBatchCheck(context.Background()).Body(body).Execute()
		require.NoError(t, err)

		requests := recorder.all()
		require.Len(t, requests, 2)
		assert.NotEqual(t, requests[0].Get(clientBulkRequestIdHeader), requests[1].Get(clientBulkRequestIdHeader))
	})
}
