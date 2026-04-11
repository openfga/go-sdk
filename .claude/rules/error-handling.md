---
description: When editing error types, retry logic, or status code handling
globs: ["errors.go", "api_executor.go", "internal/utils/retryutils/**"]
---

Error types are defined in `errors.go`. The SDK retries automatically with exponential backoff — do not implement custom retry loops.

## Error types and retry eligibility

| Code | Error Type | Retryable | `ShouldRetry()` |
|------|-----------|-----------|-----------------|
| 400 | `FgaApiValidationError` | No | N/A |
| 401/403 | `FgaApiAuthenticationError` | No | N/A |
| 404 | `FgaApiNotFoundError` | No | N/A |
| 429 | `FgaApiRateLimitExceededError` | Yes | always `true` |
| 500/502/503/504 | `FgaApiInternalError` | Yes | `true` |
| 501 | `FgaApiInternalError` | No | `false` |

All error types expose: `Error()`, `Body()`, `StoreId()`, `RequestId()`, `ResponseStatusCode()`.
Rate limit and internal errors also expose `RetryAfterDurationInMs()`.

## Retry behavior

- Defaults: 3 retries, 100ms min backoff, 120s max backoff, with jitter
- Retry-After headers checked in order: `Retry-After` > `X-RateLimit-Reset` > `X-Rate-Limit-Reset` (capped at 1800s)
- Retry-After overrides exponential backoff when present
- Custom params via `config.RetryParams = &openfga.RetryParams{MaxRetry: 5, MinWaitInMs: 50}`

## Rules

- Never retry 4xx errors (except 429) — the client must fix the request
- Do not implement manual retry loops — the SDK handles retries automatically
- Do not weaken Retry-After header respect or remove jitter (prevents thundering herd)
- Use `errors.As()` or type switches to inspect error types — do not match on error strings
- Test error scenarios with `httpmock` to simulate status codes and Retry-After headers
