# Streamed List Objects Example

Demonstrates using `StreamedListObjects` to retrieve objects via the streaming API in the Go SDK.

## What is StreamedListObjects?

The Streamed ListObjects API is very similar to the ListObjects API, with two key differences:

1. **Streaming Results**: Instead of collecting all objects before returning a response, it streams them to the client as they are collected.
2. **No Pagination Limit**: Returns all results without the 1000-object limit of the standard ListObjects API.

This makes it ideal for scenarios where you need to retrieve large numbers of objects, especially when querying computed relations.

## Prerequisites

- OpenFGA server running on `http://localhost:8080` (or set `FGA_API_URL`)

## Running

```bash
# From the example directory
cd example/streamed_list_objects
go run .
```

## What it does

- Creates a temporary store
- Writes an authorization model with **computed relations**
- Adds 2000 tuples (1000 owners + 1000 viewers)
- Queries the **computed `can_read` relation** via `StreamedListObjects`
- Shows all 2000 results (demonstrating computed relations)
- Shows progress (first 3 objects and every 500th)
- Cleans up the store

## Authorization Model

The example demonstrates OpenFGA's **computed relations**:

```
type user

type document
  relations
    define owner: [user]
    define viewer: [user]
    define can_read: owner or viewer
```

**Why this matters:**
- We write tuples to `owner` and `viewer` (base permissions)
- We query `can_read` (computed from owner OR viewer)

**Example flow:**
1. Write: `user:anne owner document:1-1000`
2. Write: `user:anne viewer document:1001-2000`
3. Query: `StreamedListObjects(user:anne, relation:can_read, type:document)`
4. Result: All 2000 documents (because `can_read = owner OR viewer`)

## Key Features Demonstrated

### Channel-based Streaming Pattern

The `StreamedListObjects` method returns a response with channels, which is the idiomatic Go way to handle streaming data:

```go
response, err := fgaClient.StreamedListObjects(ctx).Body(request).Execute()
if err != nil {
    log.Fatal(err)
}
defer response.Close()

for obj := range response.Objects {
    fmt.Printf("Received: %s\n", obj.Object)
}

// Check for errors
if err := <-response.Errors; err != nil {
    log.Fatal(err)
}
```

### Early Break and Cleanup

The streaming implementation properly handles early termination:

```go
for obj := range response.Objects {
    fmt.Println(obj.Object)
    if someCondition {
        break // Stream is automatically cleaned up via defer response.Close()
    }
}
```

### Context Cancellation Support

Full support for `context.Context`:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

response, err := fgaClient.StreamedListObjects(ctx).Body(request).Execute()
if err != nil {
    log.Fatal(err)
}
defer response.Close()

for obj := range response.Objects {
    fmt.Println(obj.Object)
}

if err := <-response.Errors; err != nil && err != context.Canceled {
    log.Fatal(err)
}
```

## Benefits Over ListObjects

- **No Pagination**: Retrieve all objects in a single streaming request
- **Lower Memory**: Objects are processed as they arrive, not held in memory
- **Early Termination**: Can stop streaming at any point without wasting resources
- **Better for Large Results**: Ideal when expecting hundreds or thousands of objects

## Performance Considerations

- Streaming starts immediately - no need to wait for all results
- HTTP connection remains open during streaming
- Properly handles cleanup if consumer stops early
- Supports all the same options as `ListObjects` (consistency, contextual tuples, etc.)

## Error Handling

The example includes robust error handling that:
- Catches configuration errors
- Detects connection issues
- Avoids logging sensitive data
- Provides helpful messages for common issues

