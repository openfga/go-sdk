package telemetry

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBuildTelemetryAttributes(t *testing.T) {
	metrics := &Metrics{}

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	req.Header.Set("User-Agent", "test-agent")

	res := &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
	}
	res.Header.Set("openfga-authorization-model-id", "test-model-id")
	res.Header.Set("fga-query-duration-ms", "123")

	methodParameters := map[string]interface{}{
		"storeId":              "test-store-id",
		"authorizationModelId": "test-model-id",
	}

	requestStarted := time.Now().Add(-500 * time.Millisecond)

	resendCount := 2

	attrs, queryDuration, requestDuration, err := metrics.BuildTelemetryAttributes("TestMethod", methodParameters, req, res, requestStarted, resendCount)

	if err != nil {
		t.Errorf("Expected no error from BuildTelemetryAttributes, but got %v", err)
	}

	if attrs[FGAClientRequestMethod] != "TestMethod" {
		t.Errorf("Expected method to be 'TestMethod', but got %v", attrs[FGAClientRequestMethod])
	}

	if attrs[FGAClientRequestStoreID] != "test-store-id" {
		t.Errorf("Expected store ID to be 'test-store-id', but got %v", attrs[FGAClientRequestStoreID])
	}

	if attrs[FGAClientRequestModelID] != "test-model-id" {
		t.Errorf("Expected model ID to be 'test-model-id', but got %v", attrs[FGAClientRequestModelID])
	}

	if attrs[FGAClientResponseModelID] != "test-model-id" {
		t.Errorf("Expected model ID in response to be 'test-model-id', but got %v", attrs[FGAClientResponseModelID])
	}

	if attrs[HTTPServerRequestDuration] != "123" {
		t.Errorf("Expected query duration to be '123', but got %v", attrs[HTTPServerRequestDuration])
	}

	if attrs[HTTPRequestResendCount] != "2" {
		t.Errorf("Expected resend count to be '2', but got %v", attrs[HTTPRequestResendCount])
	}

	if requestDuration <= 0 {
		t.Errorf("Expected positive request duration, but got %v", requestDuration)
	}

	if queryDuration != 123.0 {
		t.Errorf("Expected query duration to be 123.0, but got %v", queryDuration)
	}
}
