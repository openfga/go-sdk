package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

type Metrics struct {
	Meter         metric.Meter
	Counters      map[string]metric.Int64Counter
	Histograms    map[string]metric.Float64Histogram
	Configuration *MetricsConfiguration
}

func (m *Metrics) GetCounter(name string, description string) (metric.Int64Counter, error) {
	if counter, exists := m.Counters[name]; exists {
		return counter, nil
	}
	counter, _ := m.Meter.Int64Counter(name, metric.WithDescription(description))
	m.Counters[name] = counter
	return counter, nil
}

func (m *Metrics) GetHistogram(name string, description string, unit string) (metric.Float64Histogram, error) {
	if histogram, exists := m.Histograms[name]; exists {
		return histogram, nil
	}

	histogram, _ := m.Meter.Float64Histogram(name, metric.WithDescription(description), metric.WithUnit(unit))
	m.Histograms[name] = histogram

	return histogram, nil
}

func (m *Metrics) CredentialsRequest(value int64, attrs map[*Attribute]string) (metric.Int64Counter, error) {
	var counter, err = m.GetCounter(CredentialsRequest.Name, CredentialsRequest.Description)

	if err == nil {
		attrs, err := m.PrepareAttributes(CredentialsRequest, attrs, m.Configuration)

		if err == nil {
			counter.Add(context.Background(), value, metric.WithAttributeSet(attrs))
		}
	}

	return counter, err
}

func (m *Metrics) RequestDuration(value float64, attrs map[*Attribute]string) (metric.Float64Histogram, error) {
	var histogram, err = m.GetHistogram(RequestDuration.Name, RequestDuration.Description, RequestDuration.Unit)

	if err == nil {
		attrs, err := m.PrepareAttributes(RequestDuration, attrs, m.Configuration)

		if err == nil {
			histogram.Record(context.Background(), value, metric.WithAttributeSet(attrs))
		}
	}

	return histogram, err
}

func (m *Metrics) QueryDuration(value float64, attrs map[*Attribute]string) (metric.Float64Histogram, error) {
	var histogram, err = m.GetHistogram(QueryDuration.Name, QueryDuration.Description, QueryDuration.Unit)

	if err == nil {
		attrs, err := m.PrepareAttributes(QueryDuration, attrs, m.Configuration)

		if err == nil {
			histogram.Record(context.Background(), value, metric.WithAttributeSet(attrs))
		}
	}

	return histogram, err
}
