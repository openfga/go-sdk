package telemetry

const (
	METRIC_HISTOGRAM_REQUEST_DURATION = "fga-client.request.duration"
	METRIC_HISTOGRAM_QUERY_DURATION   = "fga-client.query.duration"
)

var (
	RequestDuration = &Histogram{
		Name:        METRIC_HISTOGRAM_REQUEST_DURATION,
		Unit:        "milliseconds",
		Description: "The total time (in milliseconds) it took for the request to complete, including the time it took to send the request and receive the response.",
	}

	QueryDuration = &Histogram{
		Name:        METRIC_HISTOGRAM_QUERY_DURATION,
		Unit:        "milliseconds",
		Description: "The total time it took (in milliseconds) for the FGA server to process and evaluate the request.",
	}
)
