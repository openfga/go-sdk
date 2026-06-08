---
description: When editing the high-level client package
globs: ["client/**"]
---

The `client` package is the primary user-facing SDK interface.

- Maintain the fluent API pattern: `Method(ctx).Body(...).Options(...).Execute()`
- All public methods must accept `context.Context` as first parameter
- Validate StoreId and AuthorizationModelId as ULIDs (see `internal/utils/IsWellFormedUlidString`)
- Wrap errors using typed error types from `errors.go` — never return raw errors
- `OpenFgaClient` must remain safe for concurrent use across goroutines
- Preserve backward compatibility of exported types, methods, and interfaces