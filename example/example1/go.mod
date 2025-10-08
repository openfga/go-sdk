module example1

go 1.24.0

toolchain go1.25.1

// To reference published build, comment below and run `go mod tidy`
replace github.com/openfga/go-sdk v0.7.3 => ../../

require github.com/openfga/go-sdk v0.7.3

require (
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.38.0 // indirect
	go.opentelemetry.io/otel/metric v1.38.0 // indirect
	go.opentelemetry.io/otel/trace v1.38.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
)
