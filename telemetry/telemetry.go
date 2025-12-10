package telemetry

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type TelemetryInterface interface {
	Configure(configuration *Configuration) (*Telemetry, error)
	Get(configuration *Configuration) *Telemetry
}

type Telemetry struct {
	Metrics       MetricsInterface
	Configuration *Configuration
}

type TelemetryFactoryParameters struct {
	Configuration *Configuration
}

type CredentialsRequestMetricParameters struct {
	Value int64
	Attrs map[*Attribute]string
	TelemetryFactoryParameters
}

type RequestDurationMetricParameters struct {
	Value float64
	Attrs map[*Attribute]string
	TelemetryFactoryParameters
}

type QueryDurationMetricParameters struct {
	Value float64
	Attrs map[*Attribute]string
	TelemetryFactoryParameters
}

type TelemetryContextKey struct{}

var (
	telemetryInstancesLock sync.Mutex
	// Warning: do not use directly, it may cause data race.
	// Deprecated: this map will be renamed to telemetryInstances.
	TelemetryInstances map[*Configuration]*Telemetry
	TelemetryContext   TelemetryContextKey
)

func Configure(configuration *Configuration) (*Telemetry, error) {
	return &Telemetry{
		Metrics: &Metrics{
			Meter:         otel.Meter("openfga-sdk"),
			Counters:      make(map[string]metric.Int64Counter),
			Histograms:    make(map[string]metric.Float64Histogram),
			Configuration: configuration.Metrics,
		},
		Configuration: configuration,
	}, nil
}

func Get(factory TelemetryFactoryParameters) *Telemetry {
	configuration := factory.Configuration

	if configuration == nil {
		configuration = DefaultTelemetryConfiguration()
	}

	telemetryInstancesLock.Lock()
	defer telemetryInstancesLock.Unlock()

	if TelemetryInstances == nil {
		TelemetryInstances = make(map[*Configuration]*Telemetry)
	}

	if _, exists := TelemetryInstances[configuration]; !exists {
		telemetry, _ := Configure(configuration)
		TelemetryInstances[configuration] = telemetry
	}

	return TelemetryInstances[configuration]
}

func Bind(ctx context.Context, instance *Telemetry) context.Context {
	return context.WithValue(ctx, TelemetryContext, instance)
}

func Unbind(ctx context.Context) context.Context {
	return context.WithValue(ctx, TelemetryContext, nil)
}

func Extract(ctx context.Context) *Telemetry {
	if ctx == nil {
		return nil
	}

	if instance, ok := ctx.Value(TelemetryContext).(*Telemetry); ok {
		return instance
	}

	return nil
}

func GetMetrics(factory TelemetryFactoryParameters) MetricsInterface {
	return Get(factory).Metrics
}

func CredentialsRequestMetric(factory CredentialsRequestMetricParameters) (metric.Int64Counter, error) {
	return GetMetrics(TelemetryFactoryParameters{Configuration: factory.Configuration}).CredentialsRequest(factory.Value, factory.Attrs)
}

func RequestDurationMetric(factory RequestDurationMetricParameters) (metric.Float64Histogram, error) {
	return GetMetrics(TelemetryFactoryParameters{Configuration: factory.Configuration}).RequestDuration(factory.Value, factory.Attrs)
}

func QueryDurationMetric(factory QueryDurationMetricParameters) (metric.Float64Histogram, error) {
	return GetMetrics(TelemetryFactoryParameters{Configuration: factory.Configuration}).QueryDuration(factory.Value, factory.Attrs)
}

type HttpRequestDurationMetricParameters struct {
	Value float64
	Attrs map[*Attribute]string
	TelemetryFactoryParameters
}

func HttpRequestDurationMetric(factory HttpRequestDurationMetricParameters) (metric.Float64Histogram, error) {
	return GetMetrics(TelemetryFactoryParameters{Configuration: factory.Configuration}).HttpRequestDuration(factory.Value, factory.Attrs)
}
