package telemetry

import (
	"testing"
)

func TestTelemetryConfiguration(t *testing.T) {
	// Create a new configuration with some attributes enabled
	metricsConfig := &MetricsConfiguration{
		METRIC_COUNTER_CREDENTIALS_REQUEST: &MetricConfiguration{
			ATTR_FGA_CLIENT_REQUEST_CLIENT_ID: &AttributeConfiguration{Enabled: true},
			ATTR_HTTP_REQUEST_METHOD:          &AttributeConfiguration{Enabled: true},
		},
		METRIC_HISTOGRAM_REQUEST_DURATION: &MetricConfiguration{
			ATTR_HTTP_CLIENT_REQUEST_DURATION: &AttributeConfiguration{Enabled: true},
			ATTR_HTTP_RESPONSE_STATUS_CODE:    &AttributeConfiguration{Enabled: false},
		},
	}

	config := &Configuration{
		Metrics: metricsConfig,
	}

	// Verify the configuration is set up correctly
	if config.Metrics == nil {
		t.Fatalf("Expected Metrics configuration to be initialized")
	}

	// Test METRIC_COUNTER_CREDENTIALS_REQUEST
	if !config.Metrics.METRIC_COUNTER_CREDENTIALS_REQUEST.ATTR_FGA_CLIENT_REQUEST_CLIENT_ID.Enabled {
		t.Errorf("Expected ATTR_FGA_CLIENT_REQUEST_CLIENT_ID to be enabled")
	}
	if !config.Metrics.METRIC_COUNTER_CREDENTIALS_REQUEST.ATTR_HTTP_REQUEST_METHOD.Enabled {
		t.Errorf("Expected ATTR_HTTP_REQUEST_METHOD to be enabled")
	}

	// Test METRIC_HISTOGRAM_REQUEST_DURATION
	if !config.Metrics.METRIC_HISTOGRAM_REQUEST_DURATION.ATTR_HTTP_CLIENT_REQUEST_DURATION.Enabled {
		t.Errorf("Expected ATTR_HTTP_CLIENT_REQUEST_DURATION to be enabled")
	}
	if config.Metrics.METRIC_HISTOGRAM_REQUEST_DURATION.ATTR_HTTP_RESPONSE_STATUS_CODE.Enabled {
		t.Errorf("Expected ATTR_HTTP_RESPONSE_STATUS_CODE to be disabled")
	}
}
