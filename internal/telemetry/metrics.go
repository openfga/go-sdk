package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

type Metrics struct {
	meter      metric.Meter
	counters   map[string]metric.Int64Counter
	histograms map[string]metric.Float64Histogram
}

func (m *Metrics) GetCounter(name string, description string) metric.Int64Counter {
	if counter, exists := m.counters[name]; exists {
		return counter
	}
	counter, _ := m.meter.Int64Counter(name, metric.WithDescription(description))
	m.counters[name] = counter
	return counter
}

func (m *Metrics) GetHistogram(name string, description string, unit string) metric.Float64Histogram {
	if histogram, exists := m.histograms[name]; exists {
		return histogram
	}

	histogram, _ := m.meter.Float64Histogram(name, metric.WithDescription(description), metric.WithUnit(unit))
	m.histograms[name] = histogram

	return histogram
}

func (m *Metrics) CredentialsRequest(value int64, attrs map[*Attribute]string) metric.Int64Counter {
	var counter = m.GetCounter(CredentialsRequest.Name, CredentialsRequest.Description)

	counter.Add(context.Background(), value, metric.WithAttributeSet(m.PrepareAttributes(attrs)))

	return counter
}

func (m *Metrics) RequestDuration(value float64, attrs map[*Attribute]string) metric.Float64Histogram {
	var histogram = m.GetHistogram(RequestDuration.Name, RequestDuration.Description, RequestDuration.Unit)

	histogram.Record(context.Background(), value, metric.WithAttributeSet(m.PrepareAttributes(attrs)))

	return histogram
}

func (m *Metrics) QueryDuration(value float64, attrs map[*Attribute]string) metric.Float64Histogram {
	var histogram = m.GetHistogram(QueryDuration.Name, QueryDuration.Description, QueryDuration.Unit)

	histogram.Record(context.Background(), value, metric.WithAttributeSet(m.PrepareAttributes(attrs)))

	return histogram
}
