package telemetry

import (
	"context"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/otel/metric"
)

type Metrics struct {
	Meter          metric.Meter
	countersLock   sync.Mutex
	Counters       map[string]metric.Int64Counter
	histogramsLock sync.Mutex
	Histograms     map[string]metric.Float64Histogram
	Configuration  *MetricsConfiguration
}

type MetricsInterface interface {
	GetCounter(name string, description string) (metric.Int64Counter, error)
	GetHistogram(name string, description string, unit string) (metric.Float64Histogram, error)
	CredentialsRequest(value int64, attrs map[*Attribute]string) (metric.Int64Counter, error)
	RequestDuration(value float64, attrs map[*Attribute]string) (metric.Float64Histogram, error)
	QueryDuration(value float64, attrs map[*Attribute]string) (metric.Float64Histogram, error)
	HTTPRequestDuration(value float64, attrs map[*Attribute]string) (metric.Float64Histogram, error)
	BuildTelemetryAttributes(requestMethod string, methodParameters map[string]interface{}, req *http.Request, res *http.Response, requestStarted time.Time, resendCount int) (map[*Attribute]string, float64, float64, error)
	AttributesFromRequest(req *http.Request, params map[string]interface{}) (map[*Attribute]string, error)
	AttributesFromResponse(res *http.Response, attrs map[*Attribute]string) (map[*Attribute]string, error)
	AttributesFromResendCount(resendCount int, attrs map[*Attribute]string) (map[*Attribute]string, error)
}

func (m *Metrics) GetCounter(name string, description string) (metric.Int64Counter, error) {
	m.countersLock.Lock()
	defer m.countersLock.Unlock()

	if counter, exists := m.Counters[name]; exists {
		return counter, nil
	}
	counter, _ := m.Meter.Int64Counter(name, metric.WithDescription(description))
	m.Counters[name] = counter
	return counter, nil
}

func (m *Metrics) GetHistogram(name string, description string, unit string) (metric.Float64Histogram, error) {
	m.histogramsLock.Lock()
	defer m.histogramsLock.Unlock()

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

func (m *Metrics) HTTPRequestDuration(value float64, attrs map[*Attribute]string) (metric.Float64Histogram, error) {
	var histogram, err = m.GetHistogram(HTTPRequestDuration.Name, HTTPRequestDuration.Description, HTTPRequestDuration.Unit)

	if err == nil {
		attrs, err := m.PrepareAttributes(HTTPRequestDuration, attrs, m.Configuration)

		if err == nil {
			histogram.Record(context.Background(), value, metric.WithAttributeSet(attrs))
		}
	}

	return histogram, err
}
