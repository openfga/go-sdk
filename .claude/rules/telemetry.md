---
description: When editing telemetry, metrics, or OpenTelemetry instrumentation
globs: ["telemetry/**"]
---

The SDK emits OpenTelemetry metrics for all client operations. Metrics are always recorded but only exported if the app provides an OTel exporter.

## Metrics

All metrics use the `fga_client_*` prefix:
- `fga_client_credentials_request` (Counter) — token fetches
- `fga_client_request_duration` (Histogram, ms) — high-level SDK method duration
- `fga_client_query_duration` (Histogram, ms) — query-specific latency
- `fga_client_http_request_duration` (Histogram, ms) — low-level HTTP duration

## Rules

- No PII in metric attributes — no user IDs, tokens, or request bodies
- High-cardinality attributes (e.g., `url_full`, unique IDs per request) must be disabled by default in `DefaultTelemetryConfiguration()` — some providers charge heavily for high cardinality
- Use `noop.NewMeterProvider()` in unit tests — never use live exporters
- All metrics must be user-configurable via `telemetry.Configuration` (enable/disable attributes per metric)
- Any changes or additions to metrics must be called out in the changelog on release

## Adding new metrics

1. Define in `telemetry/configuration.go` as a `MetricConfiguration`
2. Add recording method in `telemetry/metrics.go`
3. Add attribute keys in `telemetry/attributes.go` if new attributes are needed
4. Test with noop provider — verify no panics with nil telemetry config

## Key files

- `telemetry/configuration.go` — metric and attribute definitions, `DefaultTelemetryConfiguration()`
- `telemetry/metrics.go` — metric recording, `BuildTelemetryAttributes()`
- `telemetry/attributes.go` — attribute key constants
- `api_executor.go:recordTelemetry()` — where metrics are recorded per request
