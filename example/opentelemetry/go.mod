module openfga-opentelemetry-example

go 1.24.0

toolchain go1.25.1

// To reference published build, comment below and run `go mod tidy`
replace github.com/openfga/go-sdk v0.7.1 => ../../

require (
	github.com/joho/godotenv v1.5.1
	github.com/openfga/go-sdk v0.7.1
	go.opentelemetry.io/otel v1.38.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.38.0
	go.opentelemetry.io/otel/sdk v1.38.0
	go.opentelemetry.io/otel/sdk/metric v1.38.0
	google.golang.org/grpc v1.75.0
)
