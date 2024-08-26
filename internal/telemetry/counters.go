package telemetry

const (
	METRIC_COUNTER_CREDENTIALS_REQUEST = "fga-client.credentials.request"
)

var CredentialsRequest = &Counter{
	Name:        METRIC_COUNTER_CREDENTIALS_REQUEST,
	Description: "The total number of times new access tokens have been requested using ClientCredentials.",
}
