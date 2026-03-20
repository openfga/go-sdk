# Streaming Channel Patterns

## Overview

The `StreamingChannel[T]` type handles responses from streaming endpoints (e.g., `StreamedListObjects`). The SDK uses NDJSON (newline-delimited JSON) for streaming.

**Critical:** Both `Results` and `Errors` channels must be drained to prevent goroutine leaks.

## StreamingChannel Type

From `streaming.go`:

```go
type StreamingChannel[T any] struct {
    Results chan T      // Successful results
    Errors  chan error   // Errors (only one sent before channel closes)
    cancel  context.CancelFunc
}

// Close cancels the streaming context and cleans up resources
func (s *StreamingChannel[T]) Close() {
    if s.cancel != nil {
        s.cancel()
    }
}
```

## Correct Usage Pattern

Always use a `select` statement to drain both channels:

```go
channel, err := client.StreamedListObjects(ctx, storeId, request)
if err != nil {
    return err
}
defer channel.Close()

for {
    select {
    case result, ok := <-channel.Results:
        if !ok {
            // Results channel closed (no more results)
            break
        }
        // Process result
        processObject(result)
    case err := <-channel.Errors:
        // Error received; loop must terminate
        return fmt.Errorf("stream error: %w", err)
    }
}
```

## Why Both Channels Matter

### Results Channel
- Receives parsed objects from the NDJSON stream
- Closed when stream ends normally
- Blocks if consumer is slow

### Errors Channel
- Receives parse errors or stream errors
- Capacity 1 (buffered); only one error sent
- If not drained, the goroutine sending the error will block indefinitely

## Common Mistakes

### Mistake 1: Only Reading Results

```go
// WRONG: Ignores errors; goroutine may leak if error occurs
for result := range channel.Results {
    process(result)
}
```

**Problem:** If an error occurs, it's sent to `channel.Errors`, but nothing is reading it. The error-sending goroutine blocks forever.

**Fix:** Always select from both channels.

### Mistake 2: Not Closing the Channel

```go
// INCOMPLETE: Context is not cancelled
channel, _ := client.StreamedListObjects(ctx, storeId, req)
for ... { }
// Context keeps running in background
```

**Problem:** The streaming goroutine continues running even after the consumer stops reading.

**Fix:** Always defer `channel.Close()`.

### Mistake 3: Breaking Without Fully Draining

```go
// RISKY: Early exit without draining channels
select {
case result := <-channel.Results:
    if someCondition {
        return result  // Doesn't drain Errors channel
    }
case err := <-channel.Errors:
    return err
}
```

**Problem:** If you exit early, the goroutine may still be sending to the other channel.

**Fix:** Defer `channel.Close()` and let the select loop exit naturally, or ensure both channels are drained before returning.

## Context Cancellation

The streaming goroutine respects context cancellation:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

channel, err := client.StreamedListObjects(ctx, storeId, request)
if err != nil {
    return err
}

for {
    select {
    case result := <-channel.Results:
        process(result)
    case err := <-channel.Errors:
        return err
    case <-ctx.Done():
        // Context expired or cancelled
        return ctx.Err()
    }
}
```

When context is cancelled, the goroutine will:
1. Stop reading from the HTTP response
2. Close the HTTP body
3. Send `ctx.Err()` to the `Errors` channel
4. Close both channels

## Implementation Details

The `ProcessStreamingResponse` function (from `streaming.go`) implements the streaming logic:

1. Creates a buffered `Results` channel (default buffer 10)
2. Creates a buffered `Errors` channel (buffer 1)
3. Launches a goroutine that:
   - Reads NDJSON lines from the HTTP response body
   - Unmarshals each line into a `StreamResult[T]`
   - Sends results to `Results` channel
   - Sends errors to `Errors` channel
   - Closes both channels when done
   - Closes the HTTP body
   - Cancels the context

The goroutine respects context cancellation at each select point, so cancelling the context will cleanly shut down the stream.

## Testing Streaming Endpoints

Mock the HTTP response or use `httpmock` (already in `go.mod`):

```go
func TestStreamedListObjects(t *testing.T) {
    mock := httpmock.NewMockHTTP()
    defer mock.Close()

    // Mock returns NDJSON
    mock.WithHeader("Content-Type", "application/x-ndjson").
        WithResponse(`{"result":{"object":"user:alice"}}
{"result":{"object":"user:bob"}}
`)

    channel, _ := client.StreamedListObjects(ctx, storeId, req)
    defer channel.Close()

    count := 0
    for {
        select {
        case result := <-channel.Results:
            count++
            require.NotNil(t, result)
        case err := <-channel.Errors:
            require.NoError(t, err)
        }
        if count == 2 {
            break
        }
    }
}
```

## Summary

- **Always drain both `Results` and `Errors` channels** using a `select` statement
- **Always defer `channel.Close()`** to cancel the context
- **Respect context cancellation** by including a `<-ctx.Done()` case in the select (optional but recommended)
- **Test with mocked HTTP responses** to avoid external dependencies
