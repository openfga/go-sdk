module example1

go 1.22.2

require github.com/openfga/go-sdk v0.6.4

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
)

// To reference local build, uncomment below and run `go mod tidy`
replace github.com/openfga/go-sdk v0.6.4 => ../../
