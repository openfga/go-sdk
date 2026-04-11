---
description: When working with streaming responses or StreamingChannel
globs: ["streaming.go", "client/*streaming*"]
---

`StreamingChannel[T]` (in `streaming.go`) handles NDJSON streaming responses. Both `Results` and `Errors` channels MUST be drained to prevent goroutine leaks.

## Required pattern

Always use `select` on both channels and always defer `Close()`:

```go
channel, err := client.StreamedListObjects(ctx, storeId, request)
if err != nil { return err }
defer channel.Close()

for {
    select {
    case result, ok := <-channel.Results:
        if !ok { return nil }
        process(result)
    case err := <-channel.Errors:
        return err
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

## Rules

- Never use `for range channel.Results` alone — the Errors channel won't be drained and the sending goroutine leaks
- Always `defer channel.Close()` — otherwise the streaming goroutine and HTTP body stay alive
- Include `<-ctx.Done()` in the select for proper cancellation handling
- On context cancel, the goroutine stops reading, closes the HTTP body, sends `ctx.Err()` to Errors, and closes both channels
- Test streaming with `httptest.NewServer` returning NDJSON responses — not real network calls
