# Error Handling and Retry Patterns

## Error Type Hierarchy

The SDK defines five error types (from `errors.go`). All are rich with context:

| Type | When Raised | Retryable? | Fields |
|---|---|---|---|
| `GenericOpenAPIError` | Fallback generic error | No | body, error, model |
| `FgaApiAuthenticationError` | HTTP 401/403 | No | body, storeId, requestHost, responseStatusCode, requestId, etc. |
| `FgaApiValidationError` | HTTP 400 | No | body, storeId, requestBody, responseStatusCode, responseCode, etc. |
| `FgaApiNotFoundError` | HTTP 404 | No | body, storeId, requestBody, responseStatusCode, etc. |
| `FgaApiInternalError` | HTTP 5xx or 5xx-like | Conditional | body, storeId, retryAfterDurationInMs, retryAfterEpoc, etc. |
| `FgaApiRateLimitExceededError` | HTTP 429 | Yes | body, storeId, rateLimit, rateLimitResetEpoch, retryAfterDurationInMs, etc. |

All errors implement:
- `Error() string` — Error message
- `Body() []byte` — Raw response body
- `Model() interface{}` — Unmarshalled error response (if available)
- `StoreId() string` — Store ID from request
- `RequestId() string` — FGA request ID (from response header)
- `ResponseStatusCode() int` — HTTP status code
- Additional fields per type (e.g., `RateLimit()`, `RetryAfterDurationInMs()`)

## Retry Eligibility

### Always Retried

- **HTTP 429 (Rate Limit Exceeded)**: `FgaApiRateLimitExceededError.ShouldRetry()` returns `true`
- **HTTP 5xx (Server Error)**: `FgaApiInternalError.ShouldRetry()` returns `true` except for 501 (Not Implemented)

### Never Retried

- **HTTP 4xx (Client Error)**: `FgaApiValidationError`, `FgaApiNotFoundError`, `FgaApiAuthenticationError` — client must fix request
- **HTTP 501 (Not Implemented)**: `FgaApiInternalError.ShouldRetry()` returns `false` — endpoint does not exist

### Example: Check ShouldRetry

```go
resp, err := client.Check(ctx, storeId, request)
if err != nil {
    // Check if retryable
    if apiErr, ok := err.(FgaApiInternalError); ok {
        if apiErr.ShouldRetry() {
            // Retry after backoff
        }
    }
    if apiErr, ok := err.(FgaApiRateLimitExceededError); ok {
        // Always retried by SDK, but custom retry possible
    }
}
```

## Retry-After Header Handling

When the server responds with a `Retry-After` header, the SDK respects it:

- **Format**: Seconds as integer (e.g., `Retry-After: 30`) or HTTP-date (e.g., `Retry-After: Fri, 31 Dec 2021 23:59:59 GMT`)
- **Parsed by**: `retryutils.ParseRetryAfterHeaderValue()` in `internal/utils/retryutils/retryutils.go`
- **Maximum respected**: 1800 seconds (30 minutes); headers exceeding this are ignored
- **Returned in error**: `FgaApiInternalError.RetryAfterDurationInMs()` and `RetryAfterEpoc()`

Do not override retry headers with custom backoff logic; the SDK already does this optimally.

## Exponential Backoff

From `internal/utils/retryutils/retryutils.go`:

```
Backoff = exponential, with jitter
Formula: randomTime(loopCount, minWaitInMs)
    minTimeToWait = 2^loopCount * minWaitInMs
    maxTimeToWait = 2^(loopCount+1) * minWaitInMs
    return random duration between min and max
```

**Configuration:**
- `DefaultMaxRetry = 3` — max 3 retry attempts
- `DefaultMinWaitInMs = 100` — minimum backoff 100ms
- `MaxBackoffTimeInSec = 120` — max backoff capped at 120 seconds

**Custom retry params:**

```go
config.RetryParams = &openfga.RetryParams{
    MaxRetry:    5,
    MinWaitInMs: 50,
}
```

## Retry-After Priority

The SDK checks headers in this order and uses the first found:

1. `Retry-After` (standard HTTP header)
2. `X-RateLimit-Reset` (rate limit reset timestamp)
3. `X-Rate-Limit-Reset` (alternative rate limit reset)

If any is found and within 1800 seconds, it overrides exponential backoff.

## Error Context Example

```go
resp, err := client.Check(ctx, storeId, request)

if err != nil {
    switch e := err.(type) {
    case *FgaApiRateLimitExceededError:
        log.Printf("Rate limited. Retry after %dms. Limit: %d/%s",
            e.RetryAfterDurationInMs(),
            e.RateLimit(),
            e.RateUnit())

    case *FgaApiInternalError:
        if e.ShouldRetry() {
            log.Printf("Server error %d. Retry after %dms",
                e.ResponseStatusCode(),
                e.RetryAfterDurationInMs())
        } else {
            log.Printf("Non-retryable server error: %v", e)
        }

    case *FgaApiValidationError:
        log.Printf("Validation error: %s. Code: %s",
            e.Error(),
            e.ResponseCode())

    case *FgaApiAuthenticationError:
        log.Printf("Auth failed for %s: %s",
            e.StoreId(),
            e.Error())

    default:
        log.Printf("Unknown error: %v", err)
    }
}
```

## Do Not Implement Custom Retry

The SDK **already retries automatically** with intelligent backoff. Custom retry logic is usually unnecessary:

```go
// WRONG: Manual retry loop
for i := 0; i < 5; i++ {
    resp, err := client.Check(ctx, storeId, request)
    if err == nil {
        return resp
    }
    time.Sleep(time.Duration(100*(1<<i)) * time.Millisecond)
}

// RIGHT: Let SDK retry; only handle persistent errors
resp, err := client.Check(ctx, storeId, request)
if err != nil {
    // Non-retryable or max retries exceeded
    return err
}
```

## Testing Error Scenarios

Use `httpmock` (in `go.mod`) to simulate error responses:

```go
func TestRateLimitRetry(t *testing.T) {
    mock := httpmock.NewMockHTTP()
    defer mock.Close()

    mock.WithStatusCode(429).
        WithHeader("Retry-After", "5").
        WithResponse(`{"error": "rate_limited"}`)

    _, err := client.Check(ctx, storeId, request)
    require.Error(t, err)

    rateLimitErr, ok := err.(*FgaApiRateLimitExceededError)
    require.True(t, ok)
    require.True(t, rateLimitErr.ShouldRetry())
    require.Equal(t, 5000, rateLimitErr.RetryAfterDurationInMs())
}
```

## Status Code Reference

| Code | Type | Retryable | Action |
|---|---|---|---|
| 400 | `FgaApiValidationError` | No | Fix request |
| 401 | `FgaApiAuthenticationError` | No | Fix credentials |
| 403 | `FgaApiAuthenticationError` | No | Check permissions |
| 404 | `FgaApiNotFoundError` | No | Resource not found |
| 429 | `FgaApiRateLimitExceededError` | Yes | Wait and retry |
| 500 | `FgaApiInternalError` | Yes | Retry with backoff |
| 501 | `FgaApiInternalError` | No | Endpoint not implemented |
| 502, 503, 504 | `FgaApiInternalError` | Yes | Retry with backoff |

## Summary

- **Check error type** to determine if retryable
- **Use `ShouldRetry()` method** on internal/rate limit errors
- **Respect Retry-After header** — SDK does this automatically
- **Do not implement custom retry** — SDK already does exponential backoff
- **Extract context** from error fields (storeId, requestId, responseCode)
- **Test with httpmock** to simulate error responses
