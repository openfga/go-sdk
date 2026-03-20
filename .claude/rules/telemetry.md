# Telemetry and OpenTelemetry Patterns

## Overview

The SDK emits OpenTelemetry (OTel) metrics for all client operations. Metrics are always recorded but only exported if the application provides an OTel exporter.

## Metric Names and Types

From `telemetry/configuration.go` and `telemetry/metrics.go`:

| Metric | Type | Unit | Description |
|---|---|---|---|
| `fga_client_credentials_request` | Counter (Int64) | 1 | Incremented each time credentials are requested (e.g., OAuth2 token refresh) |
| `fga_client_request_duration` | Histogram (Float64) | ms | High-level SDK method duration (e.g., Check, Expand) |
| `fga_client_query_duration` | Histogram (Float64) | ms | Query-specific duration (e.g., Check latency) |
| `fga_client_http_request_duration` | Histogram (Float64) | ms | Low-level HTTP request duration (before retries) |

## Attributes

Attributes are attached to each metric. Common attributes:

**Standard attributes** (from `telemetry/attributes.go`):
- `fga_client_request_client_id` — Client identifier (if provided)
- `fga_client_request_method` — SDK method name (e.g., "Check", "Expand")
- `fga_client_request_model_id` — Authorization model ID
- `fga_client_request_store_id` — FGA store ID
- `fga_client_request_batch_check_size` — Number of tuples in batch (if BatchCheck)
- `http_request_method` — HTTP verb (GET, POST, etc.)
- `http_response_status_code` — HTTP status code
- `http_request_resend_count` — Number of retries
- `http_host` — API host
- `url_scheme` — URL scheme (http/https)
- `url_full` — Full request URL
- `user_agent_original` — User-Agent header

## Metric Configuration

Users can enable/disable specific attributes and metrics via `telemetry.Configuration`:

```go
config := &client.ClientConfiguration{
    Telemetry: &telemetry.Configuration{
        Metrics: &telemetry.MetricsConfiguration{
            METRIC_HISTOGRAM_REQUEST_DURATION: &telemetry.MetricConfiguration{
                ATTR_FGA_CLIENT_REQUEST_STORE_ID: &telemetry.AttributeConfiguration{
                    Enabled: false, // Disable store ID in request duration metrics
                },
            },
        },
    },
}
```

Default configuration enables most attributes (see `DefaultTelemetryConfiguration()` in `telemetry/configuration.go`).

## Using OpenTelemetry

### Standard OTel Setup

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/prometheus"
    "go.opentelemetry.io/otel/sdk/metric"
)

// Initialize Prometheus exporter
exporter, _ := prometheus.New()
provider := metric.NewMeterProvider(metric.WithReader(exporter))
otel.SetMeterProvider(provider)

// SDK automatically uses global OTel provider
client, _ := sdk.NewSdkClient(&config)
```

The SDK calls `otel.GetMeterProvider()` to get the global provider and automatically emits metrics.

### Example: Custom Attribute Cardinality Control

High-cardinality attributes like `url_full` can be disabled:

```go
config.Telemetry = &telemetry.Configuration{
    Metrics: &telemetry.MetricsConfiguration{
        METRIC_HISTOGRAM_HTTP_REQUEST_DURATION: &telemetry.MetricConfiguration{
            ATTR_URL_FULL: &telemetry.AttributeConfiguration{
                Enabled: false, // Disable high-cardinality URL attribute
            },
        },
    },
}
```

## Unit Testing with OpenTelemetry

**Critical:** Do not use live OTel exporters in unit tests. Use the `noop` provider:

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/sdk/metric/metricdata"
    "go.opentelemetry.io/otel/metric/noop"
)

func TestCheckWithTelemetry(t *testing.T) {
    // Use noop provider to suppress metric export
    noopProvider := noop.NewMeterProvider()
    otel.SetMeterProvider(noopProvider)

    config := &client.ClientConfiguration{
        ApiUrl: "https://api.fga.example",
        Telemetry: telemetry.DefaultTelemetryConfiguration(),
    }

    c, _ := client.NewSdkClient(config)
    resp, err := c.Check(ctx, storeId, request)
    require.NoError(t, err)
    require.NotNil(t, resp)

    // Metrics are recorded but not exported; no external calls made
}
```

### Reading Metrics in Tests

If you need to assert on metrics during testing, use `MetricReader`:

```go
import (
    "go.opentelemetry.io/otel/sdk/metric"
    "go.opentelemetry.io/otel/sdk/resource"
    semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func TestCheckMetrics(t *testing.T) {
    reader := metric.NewManualReader()
    provider := metric.NewMeterProvider(
        metric.WithReader(reader),
        metric.WithResource(
            resource.NewWithAttributes(
                context.Background(),
                semconv.ServiceNameKey.String("test"),
            ),
        ),
    )
    otel.SetMeterProvider(provider)

    c, _ := client.NewSdkClient(config)
    resp, err := c.Check(ctx, storeId, request)
    require.NoError(t, err)

    // Collect metrics
    rm := &metricdata.ResourceMetrics{}
    reader.Collect(context.Background(), rm)

    // Assert on recorded metrics
    require.NotNil(t, rm)
}
```

## Recording Custom Metrics

The SDK exposes `telemetry.Metrics` for recording additional metrics:

```go
import "github.com/openfga/go-sdk/telemetry"

metrics := &telemetry.Metrics{
    Meter: otel.GetMeterProvider().Meter("custom"),
}

// Record a custom metric
counter, _ := metrics.GetCounter("custom_operation_count", "Number of operations")
counter.Add(context.Background(), 1)
```

## Metric Naming Convention

All SDK metrics follow the `fga_client_*` prefix:
- `fga_client_credentials_request`
- `fga_client_request_duration`
- `fga_client_query_duration`
- `fga_client_http_request_duration`

Custom metrics should use a different prefix to avoid conflicts:
- `my_app_*` or `my_app_custom_*`

## PII and Security

**Do not include PII in metric attributes:**

- User IDs should NOT be in metrics (use `model_id`, `store_id` instead)
- API tokens should NEVER be logged or exported
- Request/response bodies should NOT be serialized to metrics
- Avoid high-cardinality attributes (unique user sessions, etc.)

Example: Safe vs. unsafe attributes:

```go
// SAFE: Store and model IDs
attrs[telemetry.ATTR_FGA_CLIENT_REQUEST_STORE_ID] = storeId
attrs[telemetry.ATTR_FGA_CLIENT_REQUEST_MODEL_ID] = modelId

// UNSAFE: Would cause high cardinality
// attrs["user_id"] = request.User  // Don't do this
// attrs["full_request"] = json.Marshal(request)  // Don't do this
```

## Telemetry Configuration Files

The SDK's default telemetry setup is in:
- `telemetry/configuration.go` — Metric/attribute definitions
- `telemetry/metrics.go` — Metric recording implementation
- `telemetry/attributes.go` — Attribute key definitions

When adding new metrics:
1. Define the metric in `telemetry/configuration.go` (MetricConfiguration struct)
2. Add the metric type name to the configuration (e.g., `METRIC_HISTOGRAM_CUSTOM`)
3. Implement recording in `telemetry/metrics.go` (e.g., `CustomMetric()` method)
4. Add attribute definitions in `telemetry/attributes.go` if needed
5. Test with `noop.NewMeterProvider()` in unit tests

## Summary

- **Metrics are always recorded** — Configuration controls which attributes are included
- **Use noop provider in tests** — Never use live exporters in unit tests (`noop.NewMeterProvider()`)
- **Configure attributes for cardinality** — Disable high-cardinality attributes if needed
- **Never expose PII** — Keep user IDs, tokens, request bodies out of metrics
- **Follow naming convention** — Use `fga_client_*` prefix for SDK metrics
- **Consult telemetry configuration** — See `telemetry/configuration.go` for defaults
