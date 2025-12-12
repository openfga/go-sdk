package telemetry

import (
	"context"
	"sync"
	"testing"

	"go.opentelemetry.io/otel/metric"
)

// Mock Int64Counter implementation
type MockInt64Counter struct {
	metric.Int64Counter
	addCalled bool
}

func (m *MockInt64Counter) Add(ctx context.Context, value int64, opts ...metric.AddOption) {
	m.addCalled = true
}

// Mock Float64Histogram implementation
type MockFloat64Histogram struct {
	metric.Float64Histogram
	recordCalled bool
}

func (m *MockFloat64Histogram) Record(ctx context.Context, value float64, opts ...metric.RecordOption) {
	m.recordCalled = true
}

// Mock Meter implementation
type MockMeter struct {
	metric.Meter
	counters   map[string]metric.Int64Counter
	histograms map[string]metric.Float64Histogram
}

func (m *MockMeter) Int64Counter(name string, opts ...metric.Int64CounterOption) (metric.Int64Counter, error) {
	if counter, exists := m.counters[name]; exists {
		return counter, nil
	}
	counter := &MockInt64Counter{}
	m.counters[name] = counter
	return counter, nil
}

func (m *MockMeter) Float64Histogram(name string, opts ...metric.Float64HistogramOption) (metric.Float64Histogram, error) {
	if histogram, exists := m.histograms[name]; exists {
		return histogram, nil
	}
	histogram := &MockFloat64Histogram{}
	m.histograms[name] = histogram
	return histogram, nil
}

func TestGetCounter(t *testing.T) {
	mockMeter := &MockMeter{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	metrics := &Metrics{
		Meter:    mockMeter,
		Counters: make(map[string]metric.Int64Counter),
	}

	counter, err := metrics.GetCounter("test_counter", "A test counter")
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if counter == nil {
		t.Fatalf("Expected a non-nil counter, but got nil")
	}

	counter2, err := metrics.GetCounter("test_counter", "A test counter")
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if counter != counter2 {
		t.Fatalf("Expected the same counter instance to be returned")
	}
}

// Run this test with the "-race" flag.
//
//	go test -race -run ^TestGetCounterRace github.com/openfga/go-sdk/telemetry
func TestGetCounterRace(t *testing.T) {

	mockMeter := &MockMeter{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	metrics := &Metrics{
		Meter:    mockMeter,
		Counters: make(map[string]metric.Int64Counter),
	}

	t.Parallel()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			_, err := metrics.GetCounter("test_counter", "A test counter")
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
				return
			}
			Get(TelemetryFactoryParameters{})
		}()
	}

	wg.Wait()
}

func TestGetHistogram(t *testing.T) {
	mockMeter := &MockMeter{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	metrics := &Metrics{
		Meter:      mockMeter,
		Histograms: make(map[string]metric.Float64Histogram),
	}

	histogram, err := metrics.GetHistogram("test_histogram", "A test histogram", "ms")
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if histogram == nil {
		t.Fatalf("Expected a non-nil histogram, but got nil")
	}

	histogram2, err := metrics.GetHistogram("test_histogram", "A test histogram", "ms")
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if histogram != histogram2 {
		t.Fatalf("Expected the same histogram instance to be returned")
	}
}

// go test -race -run ^TestGetHistogramRace github.com/openfga/go-sdk/telemetry
func TestGetHistogramRace(t *testing.T) {
	mockMeter := &MockMeter{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	metrics := &Metrics{
		Meter:      mockMeter,
		Histograms: make(map[string]metric.Float64Histogram),
	}

	t.Parallel()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			_, err := metrics.GetHistogram("test_histogram", "A test histogram", "ms")
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
				return
			}
		}()
	}

	wg.Wait()
}

func TestCredentialsRequest(t *testing.T) {
	mockMeter := &MockMeter{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	metrics := &Metrics{
		Meter:    mockMeter,
		Counters: make(map[string]metric.Int64Counter),
	}

	attrs := make(map[*Attribute]string)

	counter, err := metrics.CredentialsRequest(1, attrs)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if counter == nil {
		t.Fatalf("Expected a non-nil counter, but got nil")
	}

	mockCounter, ok := counter.(*MockInt64Counter)
	if !ok || !mockCounter.addCalled {
		t.Fatalf("Expected Add method to be called on counter")
	}
}

func TestRequestDuration(t *testing.T) {
	mockMeter := &MockMeter{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	metrics := &Metrics{
		Meter:      mockMeter,
		Histograms: make(map[string]metric.Float64Histogram),
	}

	attrs := make(map[*Attribute]string)

	histogram, err := metrics.RequestDuration(1.0, attrs)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if histogram == nil {
		t.Fatalf("Expected a non-nil histogram, but got nil")
	}

	mockHistogram, ok := histogram.(*MockFloat64Histogram)
	if !ok || !mockHistogram.recordCalled {
		t.Fatalf("Expected Record method to be called on histogram")
	}
}

func TestQueryDuration(t *testing.T) {
	mockMeter := &MockMeter{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	metrics := &Metrics{
		Meter:      mockMeter,
		Histograms: make(map[string]metric.Float64Histogram),
	}

	attrs := make(map[*Attribute]string)

	histogram, err := metrics.QueryDuration(1.0, attrs)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if histogram == nil {
		t.Fatalf("Expected a non-nil histogram, but got nil")
	}

	mockHistogram, ok := histogram.(*MockFloat64Histogram)
	if !ok || !mockHistogram.recordCalled {
		t.Fatalf("Expected Record method to be called on histogram")
	}
}

func TestHTTPRequestDuration(t *testing.T) {
	mockMeter := &MockMeter{
		counters:   make(map[string]metric.Int64Counter),
		histograms: make(map[string]metric.Float64Histogram),
	}
	metrics := &Metrics{
		Meter:      mockMeter,
		Histograms: make(map[string]metric.Float64Histogram),
	}

	attrs := make(map[*Attribute]string)

	histogram, err := metrics.HTTPRequestDuration(1.0, attrs)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if histogram == nil {
		t.Fatalf("Expected a non-nil histogram, but got nil")
	}

	mockHistogram, ok := histogram.(*MockFloat64Histogram)
	if !ok || !mockHistogram.recordCalled {
		t.Fatalf("Expected Record method to be called on histogram")
	}
}
