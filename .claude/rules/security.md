---
description: When editing credentials, oauth2, error handling, or HTTP client code
globs: ["credentials/**", "oauth2/**", "errors.go", "api_client.go", "api_executor.go"]
---

This is a public SDK that handles user credentials.

- Never log or include tokens, client secrets, or credentials in error messages
- Never expose sensitive headers in string representations or debug output
- OAuth2 token refresh must be thread-safe (OpenFgaClient is shared across goroutines)
- Always close HTTP response bodies in all code paths, including error paths
- Retry only on 429 and 5xx — never retry 4xx client errors (except 429)
- Respect `Retry-After` and `X-RateLimit-Reset` headers for backoff timing
