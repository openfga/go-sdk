package telemetry

import (
	"testing"
)

func TestCredentialsRequestCounter(t *testing.T) {
	expectedName := METRIC_COUNTER_CREDENTIALS_REQUEST
	expectedDescription := "The total number of times new access tokens have been requested using ClientCredentials."

	if CredentialsRequest == nil {
		t.Fatalf("Expected CredentialsRequest to be initialized, but got nil")
	}

	if CredentialsRequest.GetName() != expectedName {
		t.Errorf("Expected Name to be '%s', but got '%s'", expectedName, CredentialsRequest.GetName())
	}

	if CredentialsRequest.GetDescription() != expectedDescription {
		t.Errorf("Expected Description to be '%s', but got '%s'", expectedDescription, CredentialsRequest.GetDescription())
	}
}
