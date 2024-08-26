package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type Telemetry struct {
	Metrics       *Metrics
	Configuration *Configuration
}

var (
	TelemetryInstance *Telemetry
)

func CreateTelemetry(configuration *Configuration) (*Telemetry, error) {
	return &Telemetry{
		Metrics: &Metrics{
			Meter:         otel.Meter("openfga-sdk/0.5.0"),
			Counters:      make(map[string]metric.Int64Counter),
			Histograms:    make(map[string]metric.Float64Histogram),
			Configuration: configuration.Metrics,
		},
		Configuration: configuration,
	}, nil
}

func GetTelemetry(configuration *Configuration) *Telemetry {
	if TelemetryInstance == nil {
		TelemetryInstance, _ = CreateTelemetry(configuration)
	}

	if TelemetryInstance.Configuration != configuration {
		TelemetryInstance.Configuration = configuration
	}

	return TelemetryInstance
}
