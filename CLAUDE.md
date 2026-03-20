@AGENTS.md

## Claude-Specific Notes

Claude Code also loads `.claude/rules/` files that contain detailed coding instructions for specific topics:

- **`code-generation.md`** — Guardrails for generated vs. hand-written files; which files get wiped on regeneration; where to submit changes
- **`streaming.md`** — StreamingChannel patterns, goroutine leak prevention, context cancellation
- **`error-handling.md`** — Error types, retry eligibility, status code behavior, retry-after header handling
- **`telemetry.md`** — OpenTelemetry metric patterns, attributes, test setup with noop providers
- **`oauth2-credentials.md`** — Token lifecycle, expiry jitter, credential patterns, RoundTripper wrapping

When working on this SDK, consult both AGENTS.md (project context) and the applicable rules files (implementation guardrails).
