# StreamedListObjects Example

This example demonstrates how to use the `StreamedListObjects` API in the OpenFGA Go SDK in both:
- Synchronous mode (range over the channel)
- Asynchronous mode (consume in a goroutine)

It creates (if not provided) a temporary store and authorization model, writes mock tuples, streams objects, and optionally cleans up.

## Prerequisites

1. An OpenFGA server running (default: `http://localhost:8080`)
2. (Optional) Existing store and authorization model IDs

## Environment Variables

- `FGA_API_URL` (default: `http://localhost:8080`)
- `FGA_STORE_ID` (optional; if absent a new store is created and later deleted)
- `FGA_MODEL_ID` (optional; if absent a simple model is created)
- `FGA_TUPLE_COUNT` (optional; number of tuples to write; overridden by CLI arg if passed)
- `FGA_RELATION` (optional; relation used when writing and listing; must be `viewer` or `owner`; overridden by CLI arg if passed)
- `FGA_BUFFER_SIZE` (optional; buffer size for streaming channel; defaults to 10 if not set; overridden by CLI arg if passed)

## CLI Arguments

```
go run . [mode] [tupleCount] [relation] [bufferSize]
```

- `mode`: `sync` (default) or `async`
- `tupleCount`: positive integer (default: 3 if omitted and env var not set)
- `relation`: `viewer` or `owner` (default: `viewer`)
- `bufferSize`: positive integer (optional; sets the streaming channel buffer size; defaults to 10)

Examples:

```bash
# Basic sync (defaults to 3 tuples, relation viewer)
go run .

# Explicit sync, 10 tuples, relation owner
go run . sync 10 owner

# Async mode with 50 tuples and viewer relation
go run . async 50 viewer

# Sync mode with custom buffer size of 100
go run . sync 10 viewer 100

# Using environment variables (relation owner, 25 tuples, buffer size 50)
export FGA_TUPLE_COUNT=25
export FGA_RELATION=owner
export FGA_BUFFER_SIZE=50
go run . async
```

## Buffer Size Configuration

The `bufferSize` parameter (4th CLI argument or `FGA_BUFFER_SIZE` env var) controls the size of the internal channel buffer used for streaming responses:

- **Larger buffers** (e.g., 100+) improve throughput for high-volume streams but use more memory
- **Smaller buffers** (e.g., 1-10) reduce memory usage but may decrease throughput
- **Default value** is 10, providing a balanced approach for most use cases

Example with large buffer for high-volume streaming:
```bash
go run . async 1000 viewer 200
```

## What Happens Internally

1. Store creation (if `FGA_STORE_ID` not provided)
2. Authorization model creation (if `FGA_MODEL_ID` not provided) with `viewer` and `owner` relations
3. Tuple writes: `user:anne` assigned chosen relation for `document:0 .. document:N-1`
4. Streaming request (`StreamedListObjects`) for the chosen relation
5. Consumption pattern:
    - Sync: range directly over `response.Objects`
    - Async: consume in a goroutine while main goroutine reports progress
6. Final error check via `response.Errors`
7. Temporary store deletion (only if the example created it)

## Expected Output (Sync Mode Example)

```
OpenFGA StreamedListObjects Example
====================================
API URL: http://localhost:8080
Store ID: 01ARZ3NDEKTSV4RRFFQ69G5FAV

Creating authorization model...
Created authorization model: 01ARZ3NDEKTSV4RRFFQ69G5FAX

Writing 3 test tuples for relation 'viewer'...
Wrote 3 test tuples

Selected mode: sync | relation: viewer | tuple count: 3 (pass 'sync|async [count] [relation] [bufferSize]')
Mode: sync streaming (range over channel)
Streaming objects (sync):
  1. document:0
  2. document:1
  3. document:2

Total objects received (sync): 3 (expected up to 3)

Deleting temporary store...
Deleted temporary store (01ARZ3NDEKTSV4RRFFQ69G5FAV)

Done.
```

## Expected Output (Async Mode with Custom Buffer Size)

```
OpenFGA StreamedListObjects Example
====================================
API URL: http://localhost:8080
Store ID: 01ARZ3NDEKTSV4RRFFQ69G5FAV

Creating authorization model...
Created authorization model: 01ARZ3NDEKTSV4RRFFQ69G5FAX

Writing 5 test tuples for relation 'owner'...
Wrote 5 test tuples

Selected mode: async | relation: owner | tuple count: 5 | buffer size: 50 (pass 'sync|async [count] [relation] [bufferSize]')
Mode: async streaming (consume in goroutine)
Using custom buffer size: 50
Performing other work while streaming...
  (main goroutine still free to do work)
  async -> 1. document:0
  (main goroutine still free to do work)
  async -> 2. document:1
  async -> 3. document:2
  (main goroutine still free to do work)
  async -> 4. document:3
  async -> 5. document:4

Total objects received (async): 5 (expected up to 5)

Deleting temporary store...
Deleted temporary store (01ARZ3NDEKTSV4RRFFQ69G5FAV)

Done.
```
