---
description: When creating commits or pull requests
---

Follow project conventions for PRs:

- PR titles (not individual commits) must use Conventional Commits format: `type(scope): description`
  - Scope is optional: `feat: add method` or `feat(client): add method` are both valid
- Allowed types: `feat`, `fix`, `refactor`, `perf`, `chore`, `revert`, `release`
  - Use `chore(docs)` instead of `docs`
  - Use `chore(ci)` instead of `ci`
  - Use `chore(test)` instead of `test`
- Every PR must include or update tests that exercise the changed code
- Every PR that changes the public interface or functionality must include corresponding documentation updates: README changes, godoc comments, and a dedicated doc file if the feature warrants it (e.g., `telemetry/` has its own documentation)
- Changes or additions to telemetry metrics must be called out in the changelog on release
- PRs that modify generated files must link a corresponding `openfga/sdk-generator` PR as the source of those changes
- Run `make check` (fmt + lint + test + security) before committing
- The module path is `github.com/openfga/go-sdk` - use this in import paths
- Go version policy: support latest + latest-1 per Go's release policy. Target `(latest-1).0` in `go.mod` but use latest for `toolchain`