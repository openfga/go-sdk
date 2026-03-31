# API Executor Streaming Example

Demonstrates using the **low-level `APIExecutor.ExecuteStreaming`** method to call the `StreamedListObjects` endpoint with streaming.

This approach is useful when:
- You want to call a **new streaming endpoint** that is not yet supported by the SDK
- You are using an **earlier version of the SDK** that doesn't have a typed method for a particular endpoint
- You have a **custom streaming endpoint** deployed that extends the OpenFGA API
- You need **full control over the raw JSON bytes** before decoding

> For the **recommended high-level approach** using the concrete `StreamedListObjects` client method, see the [`streamed_list_objects`](../streamed_list_objects/) example.

## How it works

The `APIExecutor` provides three methods:
1. **`Execute`** — returns raw response bytes (for non-streaming endpoints)
2. **`ExecuteWithDecode`** — returns decoded response (for non-streaming endpoints)
3. **`ExecuteStreaming`** — returns a channel of raw JSON bytes (for streaming endpoints)

This example uses `ExecuteStreaming`, which:
- Automatically sets the `Accept: application/x-ndjson` header
- Returns an `APIExecutorStreamingChannel` with `Results` and `Errors` channels
- Handles the `StreamResult` wrapper (extracting the `result` field, surfacing `error` fields)
- Requires manual `json.Unmarshal` of the raw bytes into your target type

## Prerequisites

- OpenFGA server running on `http://localhost:8080` (or set `FGA_API_URL`)

## Running

```bash
cd example/api_executor_streaming
go run .
```

## Key Code Pattern

```go
// Get the API executor from the client
executor := fgaClient.GetAPIExecutor()

storeId, _ := fgaClient.GetStoreId()

// Build the request manually using the builder
request := openfga.NewAPIExecutorRequestBuilder("StreamedListObjects", "POST", "/stores/{store_id}/streamed-list-objects").
    WithPathParameter("store_id", storeId).
    WithBody(openfga.ListObjectsRequest{
        User:     "user:anne",
        Relation: "can_read",
        Type:     "document",
    }).
    Build()

// Execute streaming — returns raw JSON bytes per result
channel, err := executor.ExecuteStreaming(ctx, request, openfga.DefaultStreamBufferSize)
if err != nil {
    log.Fatal(err)
}
defer channel.Close()

for {
    select {
    case result, ok := <-channel.Results:
        if !ok {
            // Stream completed
            return
        }
        // Manually decode raw JSON
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

| Feature | Client Method (`StreamedListObjects`) | APIExecutor (`ExecuteStreaming`) |
|---|---|---|
| **Typed responses** | Yes, `StreamedListObjectsResponse` directly | No, raw `[]byte`, manual unmarshal |
| **Endpoint hardcoded** | Yes, built-in | No, you specify path, method, params |
| **Custom endpoints** | No, only known endpoints | Yes, any endpoint |
| **Recommended for** | Production use of known endpoints | Custom/new/experimental endpoints |




