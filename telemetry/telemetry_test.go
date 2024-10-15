package telemetry

import (
	"context"
	"net/http"
	"sync"
	"testing"
	"time"

	"go.opentelemetry.io/otel/metric"
)

// Mock implementation of MetricsInterface
type MockMetrics struct {
	counters   map[string]metric.Int64Counter
	histograms map[string]metric.Float64Histogram
}

func (m *MockMetrics) GetCounter(name string, description string) (metric.Int64Counter, error) {
	if counter, exists := m.counters[name]; exists {
		return counter, nil
	}
	counter := &MockInt64Counter{}
	m.counters[name] = counter
	return counter, nil
}

func (m *MockMetrics) GetHistogram(name string, description string, unit string) (metric.Float64Histogram, error) {
	if histogram, exists := m.histograms[name]; exists {
		return histogram, nil
	}
	histogram := &MockFloat64Histogram{}
	m.histograms[name] = histogram
	return histogram, nil
}

func (m *MockMetrics) CredentialsRequest(value int64, attrs map[*Attribute]string) (metric.Int64Counter, error) {
	counter, _ := m.GetCounter("credentials_request", "A credentials request")
	return counter, nil
}

func (m *MockMetrics) RequestDuration(value float64, attrs map[*Attribute]string) (metric.Float64Histogram, error) {
	histogram, _ := m.GetHistogram("request_duration", "A request duration", "ms")
	return histogram, nil
}

func (m *MockMetrics) QueryDuration(value float64, attrs map[*Attribute]string) (metric.Float64Histogram, error) {
	histogram, _ := m.GetHistogram("query_duration", "A query duration", "ms")
	return histogram, nil
}

func (m *MockMetrics) BuildTelemetryAttributes(requestMethod string, methodParameters map[string]interface{}, req *http.Request, res *http.Response, requestStarted time.Time, resendCount int) (map[*Attribute]string, float64, float64, error) {
	attrs := map[*Attribute]string{
		HTTPRequestMethod: requestMethod,
	}

	requestDuration := float64(100)
	queryDuration := float64(50)

	return attrs, queryDuration, requestDuration, nil
}

func TestConfigure(t *testing.T) {
	config := &Configuration{}
	telemetry, err := Configure(config)

	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if telemetry == nil {
		t.Fatalf("Expected telemetry to be non-nil")
	}

	if telemetry.Configuration != config {
		t.Fatalf("Expected configuration to be set correctly")
	}
}

func TestGet(t *testing.T) {
	config := &Configuration{}
	factoryParams := TelemetryFactoryParameters{Configuration: config}

	telemetry := Get(factoryParams)

	if telemetry == nil {
		t.Fatalf("Expected telemetry to be non-nil")
	}

	if telemetry.Configuration != config {
		t.Fatalf("Expected configuration to be set correctly")
	}
}

func TestBindAndExtract(t *testing.T) {
	config := &Configuration{}
	telemetry := &Telemetry{Configuration: config}

	ctx := Bind(context.Background(), telemetry)
	extractedTelemetry := Extract(ctx)

	if extractedTelemetry == nil {
		t.Fatalf("Expected extracted telemetry to be non-nil")
	}

	if extractedTelemetry.Configuration != config {
		t.Fatalf("Expected extracted telemetry configuration to be set correctly")
	}
}

func TestUnbind(t *testing.T) {
	config := &Configuration{}
	telemetry := &Telemetry{Configuration: config}

	ctx := Bind(context.Background(), telemetry)
	ctx = Unbind(ctx)
	extractedTelemetry := Extract(ctx)

	if extractedTelemetry != nil {
		t.Fatalf("Expected extracted telemetry to be nil after unbinding")
	}
}

func TestCredentialsRequestMetric(t *testing.T) {
	config := &Configuration{}
	metrics := &MockMetrics{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	telemetry := &Telemetry{Metrics: metrics, Configuration: config}
	Bind(context.Background(), telemetry)

	factoryParams := CredentialsRequestMetricParameters{
		Value:                      1,
		Attrs:                      make(map[*Attribute]string),
		TelemetryFactoryParameters: TelemetryFactoryParameters{Configuration: config},
	}

	counter, err := CredentialsRequestMetric(factoryParams)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if counter == nil {
		t.Fatalf("Expected counter to be non-nil")
	}
}

func TestRequestDurationMetric(t *testing.T) {
	config := &Configuration{}
	metrics := &MockMetrics{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	telemetry := &Telemetry{Metrics: metrics, Configuration: config}
	Bind(context.Background(), telemetry)

	factoryParams := RequestDurationMetricParameters{
		Value:                      1.0,
		Attrs:                      make(map[*Attribute]string),
		TelemetryFactoryParameters: TelemetryFactoryParameters{Configuration: config},
	}

	histogram, err := RequestDurationMetric(factoryParams)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if histogram == nil {
		t.Fatalf("Expected histogram to be non-nil")
	}
}

func TestQueryDurationMetric(t *testing.T) {
	config := &Configuration{}
	metrics := &MockMetrics{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	telemetry := &Telemetry{Metrics: metrics, Configuration: config}
	Bind(context.Background(), telemetry)

	factoryParams := QueryDurationMetricParameters{
		Value:                      1.0,
		Attrs:                      make(map[*Attribute]string),
		TelemetryFactoryParameters: TelemetryFactoryParameters{Configuration: config},
	}

	histogram, err := QueryDurationMetric(factoryParams)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if histogram == nil {
		t.Fatalf("Expected histogram to be non-nil")
	}
}

func TestBuildTelemetryAttributesMethod(t *testing.T) {
	config := &Configuration{}
	metrics := &MockMetrics{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	telemetry := &Telemetry{Metrics: metrics, Configuration: config}
	Bind(context.Background(), telemetry)

	req := &http.Request{}
	res := &http.Response{}
	requestStarted := time.Now()

	attrs, queryDuration, requestDuration, err := metrics.BuildTelemetryAttributes("GET", make(map[string]interface{}), req, res, requestStarted, 0)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if len(attrs) == 0 {
		t.Fatalf("Expected non-empty attributes")
	}

	if queryDuration != 50 {
		t.Fatalf("Expected queryDuration to be 50, but got %v", queryDuration)
	}

	if requestDuration != 100 {
		t.Fatalf("Expected requestDuration to be 100, but got %v", requestDuration)
	}
}

// Run this test with the "-race" flag.
//
//	go test -race -run ^TestGetRace$ github.com/openfga/go-sdk/telemetry
func TestGetRace(t *testing.T) {
	t.Parallel()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			Get(TelemetryFactoryParameters{})
		}()
	}

	wg.Wait()
}
