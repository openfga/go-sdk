@AGENTS.md

## Claude-Specific Notes

Claude Code also loads `.claude/rules/` files that contain detailed coding instructions for specific topics:

- **`code-generation.md`** — Guardrails for generated vs. hand-written files; which files get wiped on regeneration; where to submit changes
- **`streaming.md`** — StreamingChannel patterns, goroutine leak prevention, context cancellation
- **`error-handling.md`** — Error types, retry eligibility, status code behavior, retry-after header handling
- **`telemetry.md`** — OpenTelemetry metric patterns, attributes, test setup with noop providers
- **`oauth2-credentials.md`** — Token lifecycle, expiry jitter, credential patterns, RoundTripper wrapping
- **`security.md`** — Security rules for credentials, oauth2, error handling, and HTTP client code
- **`client-api.md`** — High-level client package conventions (fluent API, ULID validation, thread safety)
- **`testing.md`** — Test conventions (t.Parallel, table-driven, testify, httpmock, race-safe)
- **`go-conventions.md`** — Go coding standards (gofmt, golangci-lint, context propagation)
- **`new-api-methods.md`** — Checklist for new public API methods (telemetry, operation name, SdkClient interface)
- **`pr-conventions.md`** — Commit and PR conventions (conventional commits, make check)

When working on this SDK, consult both AGENTS.md (project context) and the applicable rules files (implementation guardrails).
