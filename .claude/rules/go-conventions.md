---
description: When writing or editing Go source files
globs: ["**/*.go"]
---

Go conventions for this project:

- Format with `gofmt` (not `goimports`) — CI checks this
- Lint with `golangci-lint run` using the repo's `.golangci.yaml`
- Use `openfga.ToPtr[T]()` for optional fields — `PtrString()`, `PtrBool()`, `PtrInt()` are deprecated
- Propagate `context.Context` throughout — never drop or ignore contexts
- Internal packages under `internal/` must not export types intended for external use
- Use `github.com/sourcegraph/conc` for concurrent batch operations (e.g., BatchCheck)