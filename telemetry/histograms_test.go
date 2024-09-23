package telemetry

import (
	"testing"
)

func TestRequestDurationHistogram(t *testing.T) {
	expectedName := METRIC_HISTOGRAM_REQUEST_DURATION
	expectedUnit := "milliseconds"
	expectedDescription := "The total time (in milliseconds) it took for the request to complete, including the time it took to send the request and receive the response."

	if RequestDuration == nil {
		t.Fatalf("Expected RequestDuration to be initialized, but got nil")
	}

	if RequestDuration.GetName() != expectedName {
		t.Errorf("Expected RequestDuration Name to be '%s', but got '%s'", expectedName, RequestDuration.GetName())
	}

	if RequestDuration.GetUnit() != expectedUnit {
		t.Errorf("Expected RequestDuration Unit to be '%s', but got '%s'", expectedUnit, RequestDuration.GetUnit())
	}

	if RequestDuration.GetDescription() != expectedDescription {
		t.Errorf("Expected RequestDuration Description to be '%s', but got '%s'", expectedDescription, RequestDuration.GetDescription())
	}
}

func TestQueryDurationHistogram(t *testing.T) {
	expectedName := METRIC_HISTOGRAM_QUERY_DURATION
	expectedUnit := "milliseconds"
	expectedDescription := "The total time it took (in milliseconds) for the FGA server to process and evaluate the request."

	if QueryDuration == nil {
		t.Fatalf("Expected QueryDuration to be initialized, but got nil")
	}

	if QueryDuration.GetName() != expectedName {
		t.Errorf("Expected QueryDuration Name to be '%s', but got '%s'", expectedName, QueryDuration.GetName())
	}

	if QueryDuration.GetUnit() != expectedUnit {
		t.Errorf("Expected QueryDuration Unit to be '%s', but got '%s'", expectedUnit, QueryDuration.GetUnit())
	}

	if QueryDuration.GetDescription() != expectedDescription {
		t.Errorf("Expected QueryDuration Description to be '%s', but got '%s'", expectedDescription, QueryDuration.GetDescription())
	}
}
