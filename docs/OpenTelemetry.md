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

## Example

You can find a basic example integration in the [examples/opentelemetry](../../examples/opentelemetry) directory, which demonstrates how to configure the OpenFGA SDK with OpenTelemetry.

Please see [the OpenTelemetry documentation](https://opentelemetry.io/docs/languages/go/) for additional details on how to further configure their SDK for your applications.
