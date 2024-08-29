package telemetry

import (
	"testing"
)

func TestDefaultTelemetryConfiguration(t *testing.T) {
	config := DefaultTelemetryConfiguration()

	if config == nil {
		t.Fatalf("Expected non-nil configuration, but got nil")
	}

	if config.Metrics == nil {
		t.Fatalf("Expected non-nil Metrics configuration, but got nil")
	}

	testMetricConfiguration := func(metricConfig *MetricConfiguration, metricName string) {
		if metricConfig == nil {
			t.Fatalf("Expected non-nil MetricConfiguration for %s, but got nil", metricName)
		}

		if !metricConfig.ATTR_FGA_CLIENT_REQUEST_CLIENT_ID.Enabled {
			t.Errorf("Expected %s.ATTR_FGA_CLIENT_REQUEST_CLIENT_ID to be enabled, but it was not", metricName)
		}
		if !metricConfig.ATTR_HTTP_REQUEST_METHOD.Enabled {
			t.Errorf("Expected %s.ATTR_HTTP_REQUEST_METHOD to be enabled, but it was not", metricName)
		}
		if !metricConfig.ATTR_FGA_CLIENT_REQUEST_MODEL_ID.Enabled {
			t.Errorf("Expected %s.ATTR_FGA_CLIENT_REQUEST_MODEL_ID to be enabled, but it was not", metricName)
		}
		if !metricConfig.ATTR_FGA_CLIENT_REQUEST_STORE_ID.Enabled {
			t.Errorf("Expected %s.ATTR_FGA_CLIENT_REQUEST_STORE_ID to be enabled, but it was not", metricName)
		}
		if !metricConfig.ATTR_FGA_CLIENT_RESPONSE_MODEL_ID.Enabled {
			t.Errorf("Expected %s.ATTR_FGA_CLIENT_RESPONSE_MODEL_ID to be enabled, but it was not", metricName)
		}
		if !metricConfig.ATTR_HTTP_HOST.Enabled {
			t.Errorf("Expected %s.ATTR_HTTP_HOST to be enabled, but it was not", metricName)
		}
		if !metricConfig.ATTR_HTTP_REQUEST_RESEND_COUNT.Enabled {
			t.Errorf("Expected %s.ATTR_HTTP_REQUEST_RESEND_COUNT to be enabled, but it was not", metricName)
		}
		if !metricConfig.ATTR_HTTP_RESPONSE_STATUS_CODE.Enabled {
			t.Errorf("Expected %s.ATTR_HTTP_RESPONSE_STATUS_CODE to be enabled, but it was not", metricName)
		}
		if !metricConfig.ATTR_URL_FULL.Enabled {
			t.Errorf("Expected %s.ATTR_URL_FULL to be enabled, but it was not", metricName)
		}
		if !metricConfig.ATTR_URL_SCHEME.Enabled {
			t.Errorf("Expected %s.ATTR_URL_SCHEME to be enabled, but it was not", metricName)
		}
		if !metricConfig.ATTR_USER_AGENT_ORIGINAL.Enabled {
			t.Errorf("Expected %s.ATTR_USER_AGENT_ORIGINAL to be enabled, but it was not", metricName)
		}
	}

	testMetricConfiguration(config.Metrics.METRIC_COUNTER_CREDENTIALS_REQUEST, "METRIC_COUNTER_CREDENTIALS_REQUEST")
	testMetricConfiguration(config.Metrics.METRIC_HISTOGRAM_REQUEST_DURATION, "METRIC_HISTOGRAM_REQUEST_DURATION")
	testMetricConfiguration(config.Metrics.METRIC_HISTOGRAM_QUERY_DURATION, "METRIC_HISTOGRAM_QUERY_DURATION")
}
