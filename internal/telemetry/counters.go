package telemetry

var CredentialsRequest = &Counter{
	Name:        "fga-client.credentials.request",
	Description: "The total number of times new access tokens have been requested using ClientCredentials.",
}
