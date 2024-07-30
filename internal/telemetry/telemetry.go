package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type Telemetry struct {
	metrics *Metrics
}

func (t *Telemetry) Metrics() *Metrics {
	if t.metrics == nil {
		t.metrics = &Metrics{
			meter:      otel.Meter("openfga-sdk/0.5.0"),
			counters:   make(map[string]metric.Int64Counter),
			histograms: make(map[string]metric.Float64Histogram),
		}
	}

	return t.metrics
}
