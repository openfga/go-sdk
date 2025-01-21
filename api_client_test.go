package openfga

import (
	"github.com/openfga/go-sdk/telemetry"
	"net/http"
	"testing"
	"time"
)

func TestApiClientCreatedWithDefaultTelemetry(t *testing.T) {
	cfg := Configuration{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		ApiUrl:     "http://localhost:8080/",
	}
	_ = NewAPIClient(&cfg)

	telemetry1 := telemetry.Get(telemetry.TelemetryFactoryParameters{Configuration: cfg.Telemetry})
	telemetry2 := telemetry.Get(telemetry.TelemetryFactoryParameters{Configuration: cfg.Telemetry})

	if telemetry1 != telemetry2 {
		t.Fatalf("Telemetry instance should be the same")
	}
}
