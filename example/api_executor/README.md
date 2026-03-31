# API Executor Example

Demonstrates using the **low-level `APIExecutor`** to call real OpenFGA API
endpoints — both standard request/response and streaming.

This approach is useful when:
- You want to call **any endpoint** (including new or custom ones) not yet
  supported by the SDK's typed client methods
- You are using an **earlier version of the SDK** that doesn't have a typed
  method for a particular endpoint
- You have a **custom endpoint** deployed that extends the OpenFGA API
- You need **full control over the raw request/response** (headers, body bytes, status codes)

> For the **recommended high-level approach**, see the [`example1`](../example1/)
> and [`streamed_list_objects`](../streamed_list_objects/) examples.

## What it covers

The example exercises all three `APIExecutor` methods against a live server:

| # | Operation | HTTP | Method Used |
|---|-----------|------|-------------|
| 1 | **ListStores** | `GET /stores` | `Execute` (raw bytes) |
| 2 | **CreateStore** | `POST /stores` | `ExecuteWithDecode` |
| 3 | **GetStore** | `GET /stores/{store_id}` | `ExecuteWithDecode` |
| 4 | **WriteAuthorizationModel** | `POST /stores/{store_id}/authorization-models` | `ExecuteWithDecode` |
| 5 | **WriteTuples** | `POST /stores/{store_id}/write` | `Execute` |
| 6 | **ReadTuples** | `POST /stores/{store_id}/read` | `ExecuteWithDecode` |
| 7 | **Check** | `POST /stores/{store_id}/check` | `ExecuteWithDecode` (+ custom header) |
| 8 | **ListObjects** | `POST /stores/{store_id}/list-objects` | `ExecuteWithDecode` |
| 9 | **StreamedListObjects** | `POST /stores/{store_id}/streamed-list-objects` | `ExecuteStreaming` |
| 10 | **DeleteStore** | `DELETE /stores/{store_id}` | `Execute` |

## How it works

The `APIExecutor` provides three methods:
1. **`Execute`** — returns raw response bytes (status code, headers, body)
2. **`ExecuteWithDecode`** — returns decoded typed response
3. **`ExecuteStreaming`** — returns a channel of raw JSON bytes (for NDJSON streaming endpoints)

All requests are built using `NewAPIExecutorRequestBuilder` with a fluent API
for setting the operation name, HTTP method, path, path parameters, query
parameters, headers, and body.

## Prerequisites

- OpenFGA server running on `http://localhost:8080` (or set `FGA_API_URL`)

## Running

```bash
cd example/api_executor
go run .
```

## Key Code Patterns

### Standard request with decoded response

```go
executor := fgaClient.GetAPIExecutor()

var checkResp openfga.CheckResponse
_, err := executor.ExecuteWithDecode(ctx,
    openfga.NewAPIExecutorRequestBuilder("Check", "POST", "/stores/{store_id}/check").
        WithPathParameter("store_id", storeID).
        WithHeader("X-Request-ID", "my-request-123").
        WithBody(openfga.CheckRequest{
            TupleKey: openfga.CheckRequestTupleKey{
                User:     "user:alice",
                Relation: "writer",
                Object:   "document:roadmap",
            },
        }).
        Build(),
    &checkResp,
)
fmt.Println(*checkResp.Allowed)
```

### Streaming request

```go
channel, err := executor.ExecuteStreaming(ctx,
    openfga.NewAPIExecutorRequestBuilder("StreamedListObjects", "POST", "/stores/{store_id}/streamed-list-objects").
        WithPathParameter("store_id", storeID).
        WithBody(openfga.ListObjectsRequest{
            User:     "user:alice",
            Relation: "reader",
            Type:     "document",
        }).
        Build(),
    openfga.DefaultStreamBufferSize,
)
if err != nil {
    log.Fatal(err)
}
defer channel.Close()

for {
    select {
    case result, ok := <-channel.Results:
        if !ok {
            return // Stream completed
        }
        var response openfga.StreamedListObjectsResponse
        json.Unmarshal(result, &response)
        fmt.Println(response.Object)
    case err := <-channel.Errors:
        if err != nil {
            log.Fatal(err)
        }
    }
}
```

## Comparison: Client Method vs APIExecutor

| Feature | Client Methods | APIExecutor |
|---|---|---|
| **Typed responses** | Yes, built-in | Manual decode or `ExecuteWithDecode` |
| **Endpoint hardcoded** | Yes, one method per endpoint | No, you specify path, method, params |
| **Custom endpoints** | No, only known endpoints | Yes, any endpoint |
| **Custom headers** | Via `Options()` | Via `WithHeader()` |
| **Retry logic** | Yes | Yes (same retry config) |
| **Streaming** | `StreamedListObjects()` | `ExecuteStreaming()` |
| **Recommended for** | Production use of known endpoints | Custom/new/experimental endpoints |
