# OpenTelemetry

This SDK produces [metrics](https://opentelemetry.io/docs/concepts/signals/metrics/) using [OpenTelemetry](https://opentelemetry.io/) that allow you to view data such as request timings. These metrics also include attributes for the model and store ID, as well as the API called to allow you to build reporting.

When an OpenTelemetry SDK instance is configured, the metrics will be exported and sent to the collector configured as part of your applications configuration. If you are not using OpenTelemetry, the metric functionality is a no-op and the events are never sent.

In cases when metrics events are sent, they will not be viewable outside of infrastructure configured in your application, and are never available to the OpenFGA team or contributors.

## Metrics

### Supported Metrics

| Metric Name                      | Type      | Enabled by Default | Description                                                                       |
| -------------------------------- | --------- | ------------------ | --------------------------------------------------------------------------------- |
| `fga-client.request.duration`    | Histogram | Yes                | Total request time for FGA requests, in milliseconds                              |
| `fga-client.query.duration`      | Histogram | Yes                | Time taken by the FGA server to process and evaluate the request, in milliseconds |
| `fga-client.credentials.request` | Counter   | Yes                | Total number of new token requests initiated using the Client Credentials flow    |

### Supported Attributes

| Attribute Name                 | Type   | Enabled by Default | Description                                                                       |
| ------------------------------ | ------ | ------------------ | --------------------------------------------------------------------------------- |
| `fga-client.request.client_id` | string | Yes                | Client ID associated with the request, if any                                     |
| `fga-client.request.method`    | string | Yes                | FGA method/action that was performed (e.g., Check, ListObjects) in TitleCase      |
| `fga-client.request.model_id`  | string | Yes                | Authorization model ID that was sent as part of the request, if any               |
| `fga-client.request.store_id`  | string | Yes                | Store ID that was sent as part of the request                                     |
| `fga-client.response.model_id` | string | Yes                | Authorization model ID that the FGA server used                                   |
| `fga-client.user`              | string | No                 | User associated with the action of the request for check and list users           |
| `http.client.request.duration` | int    | No                 | Duration for the SDK to complete the request, in milliseconds                     |
| `http.host`                    | string | Yes                | Host identifier of the origin the request was sent to                             |
| `http.request.method`          | string | Yes                | HTTP method for the request                                                       |
| `http.request.resend_count`    | int    | Yes                | Number of retries attempted, if any                                               |
| `http.response.status_code`    | int    | Yes                | Status code of the response (e.g., `200` for success)                             |
| `http.server.request.duration` | int    | No                 | Time taken by the FGA server to process and evaluate the request, in milliseconds |
| `url.scheme`                   | string | Yes                | HTTP scheme of the request (`http`/`https`)                                       |
| `url.full`                     | string | Yes                | Full URL of the request                                                           |
| `user_agent.original`          | string | Yes                | User Agent used in the query                                                      |

## Customizing Reporting

To control which metrics and attributes are reported by the SDK, you can provide your own `TelemetryConfiguration` instance during initialization, as shown in the example above. The `TelemetryConfiguration` class allows you to configure the metrics and attributes that are reported by the SDK, as outlined in [the tables above](#metrics).

## Usage

### 1. Install Dependencies

Install the OpenFGA SDK and OpenTelemetry SDK in your application using `pip`:

```sh
go get "github.com/openfg/go-sdk" \
  "go.opentelemetry.io/otel" \
  "go.opentelemetry.io/otel/sdk/metric"
  ```

You must also install an OpenTelemetry exporter; for example, the OTLP gRPC exporter:

```sh
go get "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
```

### 2. Configure OpenTelemetry

Configure your application to use OpenTelemetry, and set up the metrics provider to export metrics using an exporter:

```go
package main

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
)

// Configure OpenTelemetry
metricExporter, _ := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))

meterProvider := sdkmetric.NewMeterProvider(
  sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)),
  sdkmetric.WithResource(res),
)
otel.SetMeterProvider(meterProvider)
defer meterProvider.Shutdown()
```

### 3. Configure OpenFGA

Configure the OpenFGA client, and (optionally) customize what metrics and attributes are reported:

```go
import (
  "github.com/openfga/go-sdk/client"
  "os"
)

fgaClient, err := NewSdkClient(&ClientConfiguration{
  ApiUrl:  os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
  StoreId: os.Getenv("FGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
  AuthorizationModelId: os.Getenv("FGA_MODEL_ID"), // optional, recommended to be set for production
  Telemetry: &telemetry.Configuration{
    Metrics: &telemetry.MetricsConfiguration{
      METRIC_HISTOGRAM_REQUEST_DURATION: &telemetry.MetricConfiguration{
        ATTR_FGA_CLIENT_REQUEST_CLIENT_ID: &telemetry.AttributeConfiguration{Enabled: true},
        ATTR_HTTP_RESPONSE_STATUS_CODE:    &telemetry.AttributeConfiguration{Enabled: true},
      }
    }
  }
})

```

## Example Integration

An [example integration](../example/opentelemetry) is provided that also demonstrates how to configure an application with OpenFGA and OpenTelemetry. Please refer to [the README](../example/opentelemetry/README.md) for more information.
