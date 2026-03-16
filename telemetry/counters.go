package telemetry

const (
	METRIC_COUNTER_CREDENTIALS_REQUEST string = "fga-client.credentials.request"
	METRIC_COUNTER_REQUEST_COUNT       string = "fga-client.request.count"
)

var CredentialsRequest = &Counter{
	Name:        METRIC_COUNTER_CREDENTIALS_REQUEST,
	Description: "The total number of times new access tokens have been requested using ClientCredentials.",
}

var RequestCount = &Counter{
	Name:        METRIC_COUNTER_REQUEST_COUNT,
	Description: "The total number of HTTP requests made by the SDK.",
}
