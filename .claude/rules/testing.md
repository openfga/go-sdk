---
description: When writing or modifying test files
globs: ["**/*_test.go"]
---

Test conventions for this repo:

- `t.Parallel()` at both top-level test function and inside `t.Run()` subtests
- Table-driven tests with `t.Run()` for multiple cases
- `testify` for assertions: `require` for fatal checks, `assert` for non-fatal
- Mock HTTP with `github.com/jarcoal/httpmock` — never make real network calls
- Use `httptest.NewServer` for streaming/NDJSON response tests
- Never use hardcoded `time.Sleep` — use channels, contexts, or test helpers
- All test code must be race-safe (CI runs with `-race`)
- Run tests: `go test -race -v ./...` or `make test`