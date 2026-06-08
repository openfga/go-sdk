---
description: When adding new public methods that call the OpenFGA API
globs: ["client/**", "api_executor.go"]
---

Every new public method that makes API calls must include telemetry. Follow existing methods (Check, BatchCheck, ListObjects) as reference.

- All API methods must use the executor (`api_executor.go`) — no custom HTTP logic. Client methods in `client/` may call API methods, but API methods must all exercise the executor
- Set `OperationName` in `APIExecutorRequest` — this becomes the `fga_client_request_method` attribute
- The executor's `recordTelemetry()` captures `RequestDuration` and `QueryDuration` automatically
- Pass `storeId` through for metric attributes
- If the method introduces a new dimension (like batch size), add an attribute in `telemetry/attributes.go` and `telemetry/configuration.go`
- Add the method to the `SdkClient` interface in `client/client.go`
- Test with `noop.NewMeterProvider()` — verify no panics with nil telemetry config