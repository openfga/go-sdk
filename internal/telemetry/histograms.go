package telemetry

var (
	RequestDuration = &Histogram{
		Name:        "fga-client.request.duration",
		Unit:        "milliseconds",
		Description: "The total time (in milliseconds) it took for the request to complete, including the time it took to send the request and receive the response.",
	}

	QueryDuration = &Histogram{
		Name:        "fga-client.query.duration",
		Unit:        "milliseconds",
		Description: "The total time it took (in milliseconds) for the FGA server to process and evaluate the request.",
	}
)
