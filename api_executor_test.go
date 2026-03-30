package openfga

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/openfga/go-sdk/internal/constants"
)

// Test helpers

// testRoundTripper to allow stubbing HTTP responses.
type testRoundTripper struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) { return t.fn(req) }

// helper to build a http.Response quickly.
func makeResp(status int, body string, headers map[string]string) *http.Response {
	h := http.Header{}
	for k, v := range headers {
		h.Set(k, v)
	}
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     h,
	}
}

// build a minimal APIClient wired with custom http.Client
func newTestClient(t *testing.T, rt http.RoundTripper, retry *RetryParams) *APIClient {
	t.Helper()
	if retry == nil {
		retry = &RetryParams{MaxRetry: 0, MinWaitInMs: 1}
	}
	cfg := &Configuration{ApiUrl: constants.TestApiUrl, RetryParams: retry, Debug: false, HTTPClient: &http.Client{Transport: rt}}
	return NewAPIClient(cfg)
}

// Tests

func TestValidateRequest(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		request     APIExecutorRequest
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid_request_all_fields",
			request: APIExecutorRequest{
				OperationName: "Check",
				Method:        "POST",
				Path:          "/stores/{store_id}/check",
			},
			expectError: false,
		},
		{
			name: "valid_request_minimal",
			request: APIExecutorRequest{
				OperationName: "Read",
				Method:        "GET",
				Path:          "/read",
			},
			expectError: false,
		},
		{
			name: "missing_operation_name",
			request: APIExecutorRequest{
				Method: "POST",
				Path:   "/check",
			},
			expectError: true,
			errorMsg:    "operationName is required",
		},
		{
			name: "missing_method",
			request: APIExecutorRequest{
				OperationName: "Check",
				Path:          "/check",
			},
			expectError: true,
			errorMsg:    "method is required",
		},
		{
			name: "missing_path",
			request: APIExecutorRequest{
				OperationName: "Check",
				Method:        "POST",
			},
			expectError: true,
			errorMsg:    "path is required",
		},
		{
			name: "empty_operation_name",
			request: APIExecutorRequest{
				OperationName: "",
				Method:        "POST",
				Path:          "/check",
			},
			expectError: true,
			errorMsg:    "operationName is required",
		},
		{
			name: "empty_method",
			request: APIExecutorRequest{
				OperationName: "Check",
				Method:        "",
				Path:          "/check",
			},
			expectError: true,
			errorMsg:    "method is required",
		},
		{
			name: "empty_path",
			request: APIExecutorRequest{
				OperationName: "Check",
				Method:        "POST",
				Path:          "",
			},
			expectError: true,
			errorMsg:    "path is required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := validateRequest(tc.request)

			if tc.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errorMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestBuildPath(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		template     string
		params       map[string]string
		expectedPath string
	}{
		{
			name:         "no_parameters",
			template:     "/stores",
			params:       map[string]string{},
			expectedPath: "/stores",
		},
		{
			name:     "single_parameter",
			template: "/stores/{store_id}",
			params: map[string]string{
				"store_id": "01ARZ3NDEKTSV4RRFFQ69G5FAV",
			},
			expectedPath: "/stores/01ARZ3NDEKTSV4RRFFQ69G5FAV",
		},
		{
			name:     "multiple_parameters",
			template: "/stores/{store_id}/models/{model_id}",
			params: map[string]string{
				"store_id": "store-123",
				"model_id": "model-456",
			},
			expectedPath: "/stores/store-123/models/model-456",
		},
		{
			name:     "parameter_with_special_characters",
			template: "/stores/{store_id}",
			params: map[string]string{
				"store_id": "store id with spaces",
			},
			expectedPath: "/stores/store%20id%20with%20spaces",
		},
		{
			name:     "parameter_with_url_unsafe_characters",
			template: "/items/{id}",
			params: map[string]string{
				"id": "test/with?special&chars",
			},
			expectedPath: "/items/test%2Fwith%3Fspecial&chars",
		},
		{
			name:     "unused_parameters_ignored",
			template: "/stores/{store_id}",
			params: map[string]string{
				"store_id": "123",
				"unused":   "value",
			},
			expectedPath: "/stores/123",
		},
		{
			name:     "parameter_appears_multiple_times",
			template: "/stores/{id}/check/{id}",
			params: map[string]string{
				"id": "abc",
			},
			expectedPath: "/stores/abc/check/abc",
		},
		{
			name:         "nil_params",
			template:     "/stores/{store_id}",
			params:       nil,
			expectedPath: "/stores/{store_id}",
		},
		{
			name:     "empty_parameter_value",
			template: "/stores/{store_id}",
			params: map[string]string{
				"store_id": "",
			},
			expectedPath: "/stores/",
		},
		{
			name:     "parameter_with_unicode",
			template: "/users/{name}",
			params: map[string]string{
				"name": "用户",
			},
			expectedPath: "/users/%E7%94%A8%E6%88%B7",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := buildPath(tc.template, tc.params)
			assert.Equal(t, tc.expectedPath, result)
		})
	}
}

func TestPrepareHeaders(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		customHeaders   map[string]string
		expectedHeaders map[string]string
	}{
		{
			name:          "no_custom_headers",
			customHeaders: map[string]string{},
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
				"Accept":       "application/json",
			},
		},
		{
			name:          "nil_custom_headers",
			customHeaders: nil,
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
				"Accept":       "application/json",
			},
		},
		{
			name: "custom_header_added",
			customHeaders: map[string]string{
				"X-Custom-Header": "custom-value",
			},
			expectedHeaders: map[string]string{
				"Content-Type":    "application/json",
				"Accept":          "application/json",
				"X-Custom-Header": "custom-value",
			},
		},
		{
			name: "override_content_type",
			customHeaders: map[string]string{
				"Content-Type": "application/xml",
			},
			expectedHeaders: map[string]string{
				"Content-Type": "application/xml",
				"Accept":       "application/json",
			},
		},
		{
			name: "override_accept",
			customHeaders: map[string]string{
				"Accept": "application/vnd.api+json",
			},
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
				"Accept":       "application/vnd.api+json",
			},
		},
		{
			name: "override_both_defaults",
			customHeaders: map[string]string{
				"Content-Type": "text/plain",
				"Accept":       "text/html",
			},
			expectedHeaders: map[string]string{
				"Content-Type": "text/plain",
				"Accept":       "text/html",
			},
		},
		{
			name: "multiple_custom_headers",
			customHeaders: map[string]string{
				"Authorization": "Bearer token123",
				"X-Request-ID":  "req-456",
				"X-API-Key":     "key789",
			},
			expectedHeaders: map[string]string{
				"Content-Type":  "application/json",
				"Accept":        "application/json",
				"Authorization": "Bearer token123",
				"X-Request-ID":  "req-456",
				"X-API-Key":     "key789",
			},
		},
		{
			name: "case_sensitive_headers",
			customHeaders: map[string]string{
				"content-type": "should-override",
			},
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
				"Accept":       "application/json",
				"content-type": "should-override",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := prepareHeaders(tc.customHeaders)
			assert.Equal(t, tc.expectedHeaders, result)
		})
	}
}

func TestMakeAPIExecutorResponse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		httpResponse *http.Response
		body         []byte
	}{
		{
			name: "success_response",
			httpResponse: &http.Response{
				StatusCode: 200,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
					"X-Request-ID": []string{"req-123"},
				},
			},
			body: []byte(`{"message":"success"}`),
		},
		{
			name: "error_response",
			httpResponse: &http.Response{
				StatusCode: 404,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
			body: []byte(`{"code":"not_found","message":"Resource not found"}`),
		},
		{
			name: "empty_body",
			httpResponse: &http.Response{
				StatusCode: 204,
				Header:     http.Header{},
			},
			body: []byte{},
		},
		{
			name: "large_body",
			httpResponse: &http.Response{
				StatusCode: 200,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
			body: []byte(strings.Repeat("x", 10000)),
		},
		{
			name: "multiple_header_values",
			httpResponse: &http.Response{
				StatusCode: 200,
				Header: http.Header{
					"Set-Cookie": []string{"session=abc123", "preferences=dark_mode"},
				},
			},
			body: []byte(`{"ok":true}`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := makeAPIExecutorResponse(tc.httpResponse, tc.body)

			require.NotNil(t, result)
			assert.Equal(t, tc.httpResponse, result.HTTPResponse)
			assert.Equal(t, tc.body, result.Body)
			assert.Equal(t, tc.httpResponse.StatusCode, result.StatusCode)
			assert.Equal(t, tc.httpResponse.Header, result.Headers)
		})
	}
}

func TestAPIExecutorRequestBuilder_NilMaps(t *testing.T) {
	t.Parallel()

	t.Run("with_path_parameter_initializes_nil_map", func(t *testing.T) {
		t.Parallel()

		builder := &APIExecutorRequestBuilder{
			request: APIExecutorRequest{
				PathParameters: nil,
			},
		}

		result := builder.WithPathParameter("key", "value")

		assert.NotNil(t, result.request.PathParameters)
		assert.Equal(t, "value", result.request.PathParameters["key"])
	})

	t.Run("with_query_parameter_initializes_nil_map", func(t *testing.T) {
		t.Parallel()

		builder := &APIExecutorRequestBuilder{
			request: APIExecutorRequest{
				QueryParameters: nil,
			},
		}

		result := builder.WithQueryParameter("key", "value")

		assert.NotNil(t, result.request.QueryParameters)
		assert.Equal(t, "value", result.request.QueryParameters.Get("key"))
	})

	t.Run("with_header_initializes_nil_map", func(t *testing.T) {
		t.Parallel()

		builder := &APIExecutorRequestBuilder{
			request: APIExecutorRequest{
				Headers: nil,
			},
		}

		result := builder.WithHeader("key", "value")

		assert.NotNil(t, result.request.Headers)
		assert.Equal(t, "value", result.request.Headers["key"])
	})
}

func TestAPIExecutorRequestBuilder_MultipleQueryValues(t *testing.T) {
	t.Parallel()

	builder := NewAPIExecutorRequestBuilder("Test", "GET", "/test")

	builder.WithQueryParameter("tag", "go").
		WithQueryParameter("tag", "golang").
		WithQueryParameter("tag", "api")

	req := builder.Build()

	tags := req.QueryParameters["tag"]
	assert.Len(t, tags, 3)
	assert.Contains(t, tags, "go")
	assert.Contains(t, tags, "golang")
	assert.Contains(t, tags, "api")
}

func TestAPIExecutorRequestBuilder_PathParameterOverwrite(t *testing.T) {
	t.Parallel()

	builder := NewAPIExecutorRequestBuilder("Test", "GET", "/stores/{store_id}")

	builder.WithPathParameter("store_id", "old-value").
		WithPathParameter("store_id", "new-value")

	req := builder.Build()

	assert.Equal(t, "new-value", req.PathParameters["store_id"])
}

func TestAPIExecutorRequestBuilder_QueryParameterReplace(t *testing.T) {
	t.Parallel()

	builder := NewAPIExecutorRequestBuilder("Test", "GET", "/test")

	builder.WithQueryParameter("page", "1")

	newParams := url.Values{}
	newParams.Add("page", "2")
	newParams.Add("limit", "10")

	builder.WithQueryParameters(newParams)

	req := builder.Build()

	assert.Equal(t, "2", req.QueryParameters.Get("page"))
	assert.Equal(t, "10", req.QueryParameters.Get("limit"))
}

func TestAPIExecutorRequestBuilder_HeaderReplace(t *testing.T) {
	t.Parallel()

	builder := NewAPIExecutorRequestBuilder("Test", "GET", "/test")

	builder.WithHeader("X-Old", "old-value")

	newHeaders := map[string]string{
		"X-New": "new-value",
		"X-API": "api-key",
	}

	builder.WithHeaders(newHeaders)

	req := builder.Build()

	assert.Equal(t, "new-value", req.Headers["X-New"])
	assert.Equal(t, "api-key", req.Headers["X-API"])
	_, exists := req.Headers["X-Old"]
	assert.False(t, exists, "Old header should be replaced")
}

func TestAPIExecutorRequestBuilder_BodyTypes(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		body interface{}
	}{
		{
			name: "string_body",
			body: "test string",
		},
		{
			name: "struct_body",
			body: struct {
				Name  string
				Value int
			}{Name: "test", Value: 123},
		},
		{
			name: "map_body",
			body: map[string]interface{}{"key": "value", "number": 42},
		},
		{
			name: "slice_body",
			body: []string{"a", "b", "c"},
		},
		{
			name: "nil_body",
			body: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			builder := NewAPIExecutorRequestBuilder("Test", "POST", "/test")
			builder.WithBody(tc.body)
			req := builder.Build()

			assert.Equal(t, tc.body, req.Body)
		})
	}
}

func TestNewAPIExecutor(t *testing.T) {
	t.Parallel()

	t.Run("creates_executor_with_valid_client", func(t *testing.T) {
		t.Parallel()

		client := newTestClient(t, &testRoundTripper{fn: func(r *http.Request) (*http.Response, error) {
			return makeResp(200, `{"ok":true}`, nil), nil
		}}, nil)

		executor := NewAPIExecutor(client)

		assert.NotNil(t, executor)
	})

	t.Run("executor_can_execute_request", func(t *testing.T) {
		t.Parallel()

		client := newTestClient(t, &testRoundTripper{fn: func(r *http.Request) (*http.Response, error) {
			return makeResp(200, `{"message":"ok"}`, map[string]string{"Content-Type": "application/json"}), nil
		}}, nil)

		executor := NewAPIExecutor(client)
		resp, err := executor.Execute(context.Background(), APIExecutorRequest{
			OperationName: "Test",
			Method:        "GET",
			Path:          "/test",
		})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("executor_error_on_missing_path_params", func(t *testing.T) {
		t.Parallel()

		client := newTestClient(t, &testRoundTripper{fn: func(r *http.Request) (*http.Response, error) {
			return makeResp(200, `{"message":"ok"}`, map[string]string{"Content-Type": "application/json"}), nil
		}}, nil)

		executor := NewAPIExecutor(client)
		resp, err := executor.Execute(context.Background(), APIExecutorRequest{
			OperationName: "Test",
			Method:        "GET",
			Path:          "/stores/{store_id}/test",
		})

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestAPIExecutor_GetRetryParams(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		configRetryParams   *RetryParams
		expectedMaxRetry    int
		expectedMinWaitInMs int
	}{
		{
			name: "uses_configured_retry_params",
			configRetryParams: &RetryParams{
				MaxRetry:    5,
				MinWaitInMs: 100,
			},
			expectedMaxRetry:    5,
			expectedMinWaitInMs: 100,
		},
		{
			name: "uses_custom_values",
			configRetryParams: &RetryParams{
				MaxRetry:    10,
				MinWaitInMs: 500,
			},
			expectedMaxRetry:    10,
			expectedMinWaitInMs: 500,
		},
		{
			name: "uses_different_retry_values",
			configRetryParams: &RetryParams{
				MaxRetry:    2,
				MinWaitInMs: 200,
			},
			expectedMaxRetry:    2,
			expectedMinWaitInMs: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := newTestClient(t, &testRoundTripper{fn: func(r *http.Request) (*http.Response, error) {
				return makeResp(200, `{"ok":true}`, nil), nil
			}}, tc.configRetryParams)

			executor := NewAPIExecutor(client).(*apiExecutor)
			retryParams := executor.getRetryParams()

			assert.Equal(t, tc.expectedMaxRetry, retryParams.MaxRetry)
			assert.Equal(t, tc.expectedMinWaitInMs, retryParams.MinWaitInMs)
		})
	}
}

func TestAPIExecutor_DetermineRetry(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		err                 error
		response            *APIExecutorResponse
		attemptNum          int
		retryParams         RetryParams
		operationName       string
		expectShouldRetry   bool
		expectWaitDuration  bool // Whether we expect a non-zero wait duration
		minExpectedDuration int  // Minimum expected duration in ms
	}{
		{
			name:               "no_error_no_retry",
			err:                nil,
			response:           nil,
			attemptNum:         0,
			retryParams:        RetryParams{MaxRetry: 3, MinWaitInMs: 50},
			operationName:      "Test",
			expectShouldRetry:  false,
			expectWaitDuration: false,
		},
		{
			name:                "generic_error_retries",
			err:                 errors.New("network error"),
			response:            nil,
			attemptNum:          0,
			retryParams:         RetryParams{MaxRetry: 3, MinWaitInMs: 50},
			operationName:       "Test",
			expectShouldRetry:   true,
			expectWaitDuration:  true,
			minExpectedDuration: 50,
		},
		{
			name:                "connection_error_retries",
			err:                 errors.New("connection refused"),
			response:            nil,
			attemptNum:          0,
			retryParams:         RetryParams{MaxRetry: 3, MinWaitInMs: 100},
			operationName:       "Test",
			expectShouldRetry:   true,
			expectWaitDuration:  true,
			minExpectedDuration: 100,
		},
		{
			name:                "below_max_attempts",
			err:                 errors.New("network error"),
			response:            nil,
			attemptNum:          2,
			retryParams:         RetryParams{MaxRetry: 5, MinWaitInMs: 50},
			operationName:       "Test",
			expectShouldRetry:   true,
			expectWaitDuration:  true,
			minExpectedDuration: 50,
		},
		{
			name:                "high_attempt_number",
			err:                 errors.New("timeout"),
			response:            nil,
			attemptNum:          10,
			retryParams:         RetryParams{MaxRetry: 15, MinWaitInMs: 50},
			operationName:       "Test",
			expectShouldRetry:   true,
			expectWaitDuration:  true,
			minExpectedDuration: 50,
		},
		{
			name:               "context_canceled_no_retry",
			err:                context.Canceled,
			response:           nil,
			attemptNum:         0,
			retryParams:        RetryParams{MaxRetry: 3, MinWaitInMs: 50},
			operationName:      "Test",
			expectShouldRetry:  false,
			expectWaitDuration: false,
		},
		{
			name:               "context_deadline_exceeded_no_retry",
			err:                context.DeadlineExceeded,
			response:           nil,
			attemptNum:         0,
			retryParams:        RetryParams{MaxRetry: 3, MinWaitInMs: 50},
			operationName:      "Test",
			expectShouldRetry:  false,
			expectWaitDuration: false,
		},
		{
			name:               "wrapped_context_canceled_no_retry",
			err:                fmt.Errorf("operation failed: %w", context.Canceled),
			response:           nil,
			attemptNum:         0,
			retryParams:        RetryParams{MaxRetry: 3, MinWaitInMs: 50},
			operationName:      "Test",
			expectShouldRetry:  false,
			expectWaitDuration: false,
		},
		{
			name:               "wrapped_context_deadline_exceeded_no_retry",
			err:                fmt.Errorf("operation failed: %w", context.DeadlineExceeded),
			response:           nil,
			attemptNum:         0,
			retryParams:        RetryParams{MaxRetry: 3, MinWaitInMs: 50},
			operationName:      "Test",
			expectShouldRetry:  false,
			expectWaitDuration: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := newTestClient(t, &testRoundTripper{fn: func(r *http.Request) (*http.Response, error) {
				return makeResp(200, `{"ok":true}`, nil), nil
			}}, &tc.retryParams)

			executor := NewAPIExecutor(client).(*apiExecutor)
			shouldRetry, waitDuration := executor.determineRetry(
				tc.err,
				tc.response,
				tc.attemptNum,
				tc.retryParams,
				tc.operationName,
			)

			assert.Equal(t, tc.expectShouldRetry, shouldRetry)
			if tc.expectWaitDuration {
				assert.Greater(t, waitDuration.Milliseconds(), int64(tc.minExpectedDuration-1))
			} else {
				assert.Equal(t, int64(0), waitDuration.Milliseconds())
			}
		})
	}
}

func TestBuildPath_EdgeCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		template     string
		params       map[string]string
		expectedPath string
	}{
		{
			name:         "path_with_no_placeholders",
			template:     "/api/v1/stores",
			params:       map[string]string{"store_id": "123"},
			expectedPath: "/api/v1/stores",
		},
		{
			name:         "placeholder_not_in_params",
			template:     "/stores/{store_id}",
			params:       map[string]string{"other_id": "123"},
			expectedPath: "/stores/{store_id}",
		},
		{
			name:     "multiple_slashes_preserved",
			template: "/stores//{store_id}//check",
			params: map[string]string{
				"store_id": "123",
			},
			expectedPath: "/stores//123//check",
		},
		{
			name:     "placeholder_at_start",
			template: "{store_id}/check",
			params: map[string]string{
				"store_id": "123",
			},
			expectedPath: "123/check",
		},
		{
			name:     "placeholder_at_end",
			template: "/stores/{store_id}",
			params: map[string]string{
				"store_id": "123",
			},
			expectedPath: "/stores/123",
		},
		{
			name:     "adjacent_placeholders",
			template: "/api/{version}{store_id}",
			params: map[string]string{
				"version":  "v1",
				"store_id": "123",
			},
			expectedPath: "/api/v1123",
		},
		{
			name:     "placeholder_with_underscores_and_numbers",
			template: "/stores/{store_id_1}/models/{model_id_2}",
			params: map[string]string{
				"store_id_1": "abc",
				"model_id_2": "xyz",
			},
			expectedPath: "/stores/abc/models/xyz",
		},
		{
			name:     "url_encoded_value_with_percent",
			template: "/items/{id}",
			params: map[string]string{
				"id": "100%",
			},
			expectedPath: "/items/100%25",
		},
		{
			name:     "value_with_curly_braces",
			template: "/items/{id}",
			params: map[string]string{
				"id": "{test}",
			},
			expectedPath: "/items/%7Btest%7D",
		},
		{
			name:     "value_with_plus_sign",
			template: "/search/{query}",
			params: map[string]string{
				"query": "hello+world",
			},
			expectedPath: "/search/hello+world",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := buildPath(tc.template, tc.params)
			assert.Equal(t, tc.expectedPath, result)
		})
	}
}

func TestPrepareHeaders_EdgeCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		customHeaders map[string]string
		checkHeader   string
		expectedValue string
	}{
		{
			name: "empty_string_header_value",
			customHeaders: map[string]string{
				"X-Empty": "",
			},
			checkHeader:   "X-Empty",
			expectedValue: "",
		},
		{
			name: "header_with_special_characters",
			customHeaders: map[string]string{
				"X-Special": "value with spaces and !@#$%",
			},
			checkHeader:   "X-Special",
			expectedValue: "value with spaces and !@#$%",
		},
		{
			name: "very_long_header_value",
			customHeaders: map[string]string{
				"X-Long": strings.Repeat("a", 1000),
			},
			checkHeader:   "X-Long",
			expectedValue: strings.Repeat("a", 1000),
		},
		{
			name: "header_with_unicode",
			customHeaders: map[string]string{
				"X-Unicode": "Hello 世界 🌍",
			},
			checkHeader:   "X-Unicode",
			expectedValue: "Hello 世界 🌍",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := prepareHeaders(tc.customHeaders)
			assert.Equal(t, tc.expectedValue, result[tc.checkHeader])
		})
	}
}

func TestAPIExecutorResponse_Fields(t *testing.T) {
	t.Parallel()

	t.Run("all_fields_populated", func(t *testing.T) {
		t.Parallel()

		httpResp := &http.Response{
			StatusCode: 201,
			Header: http.Header{
				"Content-Type":   []string{"application/json"},
				"X-Request-Id":   []string{"req-123"},
				"Content-Length": []string{"100"},
			},
		}
		body := []byte(`{"created":true,"id":"abc123"}`)

		resp := makeAPIExecutorResponse(httpResp, body)

		assert.Equal(t, 201, resp.StatusCode)
		assert.Equal(t, body, resp.Body)
		assert.Equal(t, httpResp, resp.HTTPResponse)
		assert.Equal(t, httpResp.Header, resp.Headers)
		// Note: Header comparison validates that all headers are preserved
		assert.Contains(t, resp.Headers["Content-Type"], "application/json")
		assert.Contains(t, resp.Headers["X-Request-Id"], "req-123")
	})

	t.Run("access_body_directly", func(t *testing.T) {
		t.Parallel()

		body := []byte(`{"data":"test"}`)
		resp := makeAPIExecutorResponse(&http.Response{StatusCode: 200, Header: http.Header{}}, body)

		assert.Equal(t, `{"data":"test"}`, string(resp.Body))
	})

	t.Run("response_with_redirect_status", func(t *testing.T) {
		t.Parallel()

		httpResp := &http.Response{
			StatusCode: 302,
			Header: http.Header{
				"Location": []string{"/new-location"},
			},
		}
		body := []byte{}

		resp := makeAPIExecutorResponse(httpResp, body)

		assert.Equal(t, 302, resp.StatusCode)
		assert.Equal(t, "/new-location", resp.Headers.Get("Location"))
		assert.Empty(t, resp.Body)
	})
}

func TestAPIExecutorRequestBuilder_Chaining(t *testing.T) {
	t.Parallel()

	t.Run("complete_chain", func(t *testing.T) {
		t.Parallel()

		req := NewAPIExecutorRequestBuilder("ComplexOp", "POST", "/stores/{store_id}/check").
			WithPathParameter("store_id", "store-123").
			WithPathParameter("model_id", "model-456"). // Extra param
			WithQueryParameter("expand", "true").
			WithQueryParameter("limit", "10").
			WithHeader("Authorization", "Bearer token").
			WithHeader("X-API-Version", "v1").
			WithBody(map[string]string{"user": "user:anne"}).
			Build()

		assert.Equal(t, "ComplexOp", req.OperationName)
		assert.Equal(t, "POST", req.Method)
		assert.Equal(t, "/stores/{store_id}/check", req.Path)
		assert.Equal(t, "store-123", req.PathParameters["store_id"])
		assert.Equal(t, "model-456", req.PathParameters["model_id"])
		assert.Equal(t, "true", req.QueryParameters.Get("expand"))
		assert.Equal(t, "10", req.QueryParameters.Get("limit"))
		assert.Equal(t, "Bearer token", req.Headers["Authorization"])
		assert.Equal(t, "v1", req.Headers["X-API-Version"])
		assert.NotNil(t, req.Body)
	})

	t.Run("empty_chain", func(t *testing.T) {
		t.Parallel()

		req := NewAPIExecutorRequestBuilder("Empty", "GET", "/empty").Build()

		assert.Equal(t, "Empty", req.OperationName)
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/empty", req.Path)
		assert.NotNil(t, req.PathParameters)
		assert.NotNil(t, req.QueryParameters)
		assert.NotNil(t, req.Headers)
		assert.Nil(t, req.Body)
	})

	t.Run("build_multiple_times", func(t *testing.T) {
		t.Parallel()

		builder := NewAPIExecutorRequestBuilder("Multi", "GET", "/test")
		builder.WithPathParameter("key", "value1")

		req1 := builder.Build()
		builder.WithPathParameter("key", "value2")
		req2 := builder.Build()

		// Both builds should reflect the current state
		assert.Equal(t, "value2", req1.PathParameters["key"])
		assert.Equal(t, "value2", req2.PathParameters["key"])
	})
}

func TestAPIExecutorRequestBuilder_Overrides(t *testing.T) {
	t.Parallel()

	t.Run("path_parameters_replacement", func(t *testing.T) {
		t.Parallel()

		builder := NewAPIExecutorRequestBuilder("Test", "GET", "/test")

		builder.WithPathParameter("id", "1")
		builder.WithPathParameter("name", "test")

		newParams := map[string]string{
			"id":        "2",
			"different": "value",
		}
		builder.WithPathParameters(newParams)

		req := builder.Build()

		assert.Equal(t, "2", req.PathParameters["id"])
		assert.Equal(t, "value", req.PathParameters["different"])
		_, hasName := req.PathParameters["name"]
		assert.False(t, hasName, "name parameter should be replaced")
	})

	t.Run("query_parameters_replacement", func(t *testing.T) {
		t.Parallel()

		builder := NewAPIExecutorRequestBuilder("Test", "GET", "/test")

		builder.WithQueryParameter("page", "1")
		builder.WithQueryParameter("limit", "10")

		newParams := url.Values{}
		newParams.Add("page", "2")
		newParams.Add("sort", "asc")

		builder.WithQueryParameters(newParams)

		req := builder.Build()

		assert.Equal(t, "2", req.QueryParameters.Get("page"))
		assert.Equal(t, "asc", req.QueryParameters.Get("sort"))
		assert.Empty(t, req.QueryParameters.Get("limit"), "limit should be replaced")
	})

	t.Run("headers_replacement", func(t *testing.T) {
		t.Parallel()

		builder := NewAPIExecutorRequestBuilder("Test", "GET", "/test")

		builder.WithHeader("X-Old", "old")
		builder.WithHeader("X-Keep", "keep")

		newHeaders := map[string]string{
			"X-New": "new",
		}

		builder.WithHeaders(newHeaders)

		req := builder.Build()

		assert.Equal(t, "new", req.Headers["X-New"])
		_, hasOld := req.Headers["X-Old"]
		assert.False(t, hasOld, "X-Old should be replaced")
		_, hasKeep := req.Headers["X-Keep"]
		assert.False(t, hasKeep, "X-Keep should be replaced")
	})
}

func TestValidateRequest_AllFieldCombinations(t *testing.T) {
	t.Parallel()

	t.Run("only_operation_name", func(t *testing.T) {
		t.Parallel()

		err := validateRequest(APIExecutorRequest{
			OperationName: "Test",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "method is required")
	})

	t.Run("only_method", func(t *testing.T) {
		t.Parallel()

		err := validateRequest(APIExecutorRequest{
			Method: "GET",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "operationName is required")
	})

	t.Run("only_path", func(t *testing.T) {
		t.Parallel()

		err := validateRequest(APIExecutorRequest{
			Path: "/test",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "operationName is required")
	})

	t.Run("operation_name_and_method", func(t *testing.T) {
		t.Parallel()

		err := validateRequest(APIExecutorRequest{
			OperationName: "Test",
			Method:        "GET",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "path is required")
	})

	t.Run("operation_name_and_path", func(t *testing.T) {
		t.Parallel()

		err := validateRequest(APIExecutorRequest{
			OperationName: "Test",
			Path:          "/test",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "method is required")
	})

	t.Run("method_and_path", func(t *testing.T) {
		t.Parallel()

		err := validateRequest(APIExecutorRequest{
			Method: "GET",
			Path:   "/test",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "operationName is required")
	})

	t.Run("all_required_fields_with_optional_fields", func(t *testing.T) {
		t.Parallel()

		err := validateRequest(APIExecutorRequest{
			OperationName:   "Test",
			Method:          "POST",
			Path:            "/test",
			PathParameters:  map[string]string{"id": "123"},
			QueryParameters: url.Values{"page": []string{"1"}},
			Body:            map[string]string{"data": "test"},
			Headers:         map[string]string{"X-Test": "value"},
		})
		assert.NoError(t, err)
	})
}

func TestBuildPath_SpecialCases(t *testing.T) {
	t.Parallel()

	t.Run("empty_template", func(t *testing.T) {
		t.Parallel()

		result := buildPath("", map[string]string{"id": "123"})
		assert.Equal(t, "", result)
	})

	t.Run("template_with_only_placeholder", func(t *testing.T) {
		t.Parallel()

		result := buildPath("{id}", map[string]string{"id": "123"})
		assert.Equal(t, "123", result)
	})

	t.Run("nested_braces", func(t *testing.T) {
		t.Parallel()

		result := buildPath("/api/{{id}}", map[string]string{"id": "123"})
		assert.Equal(t, "/api/{123}", result)
	})

	t.Run("placeholder_with_dash", func(t *testing.T) {
		t.Parallel()

		result := buildPath("/stores/{store-id}", map[string]string{"store-id": "123"})
		assert.Equal(t, "/stores/123", result)
	})

	t.Run("empty_params_map", func(t *testing.T) {
		t.Parallel()

		result := buildPath("/stores/{store_id}", map[string]string{})
		assert.Equal(t, "/stores/{store_id}", result)
	})

	t.Run("value_with_equals_sign", func(t *testing.T) {
		t.Parallel()

		result := buildPath("/query/{q}", map[string]string{"q": "key=value"})
		assert.Contains(t, result, "=")
	})

	t.Run("value_with_ampersand", func(t *testing.T) {
		t.Parallel()

		result := buildPath("/query/{q}", map[string]string{"q": "a&b"})
		assert.Contains(t, result, "&")
	})
}

// ============================================================================
// Streaming API Tests
// ============================================================================

func TestAPIExecutorStreamingChannel_Close(t *testing.T) {
	t.Parallel()

	t.Run("close_with_cancel_function", func(t *testing.T) {
		t.Parallel()

		cancelCalled := false
		channel := &APIExecutorStreamingChannel{
			Results: make(chan []byte),
			Errors:  make(chan error),
			cancel: func() {
				cancelCalled = true
			},
		}

		channel.Close()
		assert.True(t, cancelCalled, "cancel function should be called")
	})

	t.Run("close_with_nil_cancel", func(t *testing.T) {
		t.Parallel()

		channel := &APIExecutorStreamingChannel{
			Results: make(chan []byte),
			Errors:  make(chan error),
			cancel:  nil,
		}

		// Should not panic
		assert.NotPanics(t, func() {
			channel.Close()
		})
	})

	t.Run("close_multiple_times", func(t *testing.T) {
		t.Parallel()

		callCount := 0
		channel := &APIExecutorStreamingChannel{
			Results: make(chan []byte),
			Errors:  make(chan error),
			cancel: func() {
				callCount++
			},
		}

		channel.Close()
		channel.Close()
		channel.Close()

		assert.Equal(t, 3, callCount, "cancel function should be called each time")
	})
}

func TestDefaultStreamBufferSize(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 10, DefaultStreamBufferSize, "default buffer size should be 10")
}

func TestProcessStreamingResponseRaw(t *testing.T) {
	t.Parallel()

	t.Run("nil_response_returns_error", func(t *testing.T) {
		t.Parallel()

		channel, err := processStreamingResponseRaw(context.Background(), nil, 10)

		assert.Error(t, err)
		assert.Nil(t, channel)
		assert.Contains(t, err.Error(), "response or response body is nil")
	})

	t.Run("nil_body_returns_error", func(t *testing.T) {
		t.Parallel()

		resp := &http.Response{
			StatusCode: 200,
			Body:       nil,
		}

		channel, err := processStreamingResponseRaw(context.Background(), resp, 10)

		assert.Error(t, err)
		assert.Nil(t, channel)
		assert.Contains(t, err.Error(), "response or response body is nil")
	})

	t.Run("uses_default_buffer_size_when_zero", func(t *testing.T) {
		t.Parallel()

		body := io.NopCloser(strings.NewReader(`{"result":{"object":"doc:1"}}`))
		resp := &http.Response{
			StatusCode: 200,
			Body:       body,
		}

		channel, err := processStreamingResponseRaw(context.Background(), resp, 0)

		require.NoError(t, err)
		require.NotNil(t, channel)
		defer channel.Close()

		// Check buffer size (indirectly through cap)
		assert.Equal(t, DefaultStreamBufferSize, cap(channel.Results))
	})

	t.Run("uses_default_buffer_size_when_negative", func(t *testing.T) {
		t.Parallel()

		body := io.NopCloser(strings.NewReader(`{"result":{"object":"doc:1"}}`))
		resp := &http.Response{
			StatusCode: 200,
			Body:       body,
		}

		channel, err := processStreamingResponseRaw(context.Background(), resp, -5)

		require.NoError(t, err)
		require.NotNil(t, channel)
		defer channel.Close()

		assert.Equal(t, DefaultStreamBufferSize, cap(channel.Results))
	})

	t.Run("uses_custom_buffer_size", func(t *testing.T) {
		t.Parallel()

		body := io.NopCloser(strings.NewReader(`{"result":{"object":"doc:1"}}`))
		resp := &http.Response{
			StatusCode: 200,
			Body:       body,
		}

		channel, err := processStreamingResponseRaw(context.Background(), resp, 25)

		require.NoError(t, err)
		require.NotNil(t, channel)
		defer channel.Close()

		assert.Equal(t, 25, cap(channel.Results))
	})

	t.Run("processes_single_result", func(t *testing.T) {
		t.Parallel()

		ndjson := `{"result":{"object":"document:1"}}` + "\n"
		body := io.NopCloser(strings.NewReader(ndjson))
		resp := &http.Response{
			StatusCode: 200,
			Body:       body,
		}

		channel, err := processStreamingResponseRaw(context.Background(), resp, 10)

		require.NoError(t, err)
		require.NotNil(t, channel)
		defer channel.Close()

		// Collect results
		var results [][]byte
		for result := range channel.Results {
			results = append(results, result)
		}

		// Check for errors
		select {
		case err := <-channel.Errors:
			assert.NoError(t, err)
		default:
		}

		assert.Len(t, results, 1)
		assert.JSONEq(t, `{"object":"document:1"}`, string(results[0]))
	})

	t.Run("processes_multiple_results", func(t *testing.T) {
		t.Parallel()

		ndjson := `{"result":{"object":"document:1"}}` + "\n" +
			`{"result":{"object":"document:2"}}` + "\n" +
			`{"result":{"object":"document:3"}}` + "\n"
		body := io.NopCloser(strings.NewReader(ndjson))
		resp := &http.Response{
			StatusCode: 200,
			Body:       body,
		}

		channel, err := processStreamingResponseRaw(context.Background(), resp, 10)

		require.NoError(t, err)
		require.NotNil(t, channel)
		defer channel.Close()

		var results [][]byte
		for result := range channel.Results {
			results = append(results, result)
		}

		select {
		case err := <-channel.Errors:
			assert.NoError(t, err)
		default:
		}

		assert.Len(t, results, 3)
		assert.JSONEq(t, `{"object":"document:1"}`, string(results[0]))
		assert.JSONEq(t, `{"object":"document:2"}`, string(results[1]))
		assert.JSONEq(t, `{"object":"document:3"}`, string(results[2]))
	})

	t.Run("handles_stream_error_response", func(t *testing.T) {
		t.Parallel()

		ndjson := `{"result":{"object":"document:1"}}` + "\n" +
			`{"error":{"message":"Something went wrong"}}` + "\n"
		body := io.NopCloser(strings.NewReader(ndjson))
		resp := &http.Response{
			StatusCode: 200,
			Body:       body,
		}

		channel, err := processStreamingResponseRaw(context.Background(), resp, 10)

		require.NoError(t, err)
		require.NotNil(t, channel)
		defer channel.Close()

		// First result should come through
		result := <-channel.Results
		assert.JSONEq(t, `{"object":"document:1"}`, string(result))

		// Then we should get an error
		streamErr := <-channel.Errors
		assert.Error(t, streamErr)
		assert.Contains(t, streamErr.Error(), "Something went wrong")
	})

	t.Run("handles_invalid_json", func(t *testing.T) {
		t.Parallel()

		ndjson := `{"result":{"object":"document:1"}}` + "\n" +
			`invalid json` + "\n"
		body := io.NopCloser(strings.NewReader(ndjson))
		resp := &http.Response{
			StatusCode: 200,
			Body:       body,
		}

		channel, err := processStreamingResponseRaw(context.Background(), resp, 10)

		require.NoError(t, err)
		require.NotNil(t, channel)
		defer channel.Close()

		// First result should come through
		result := <-channel.Results
		assert.JSONEq(t, `{"object":"document:1"}`, string(result))

		// Then we should get a JSON parsing error
		streamErr := <-channel.Errors
		assert.Error(t, streamErr)
	})

	t.Run("skips_empty_lines", func(t *testing.T) {
		t.Parallel()

		ndjson := `{"result":{"object":"document:1"}}` + "\n" +
			`` + "\n" +
			`{"result":{"object":"document:2"}}` + "\n" +
			`` + "\n"
		body := io.NopCloser(strings.NewReader(ndjson))
		resp := &http.Response{
			StatusCode: 200,
			Body:       body,
		}

		channel, err := processStreamingResponseRaw(context.Background(), resp, 10)

		require.NoError(t, err)
		require.NotNil(t, channel)
		defer channel.Close()

		var results [][]byte
		for result := range channel.Results {
			results = append(results, result)
		}

		assert.Len(t, results, 2)
	})

	t.Run("context_cancellation_stops_streaming", func(t *testing.T) {
		t.Parallel()

		// Create a pipe where we control when data is written
		pr, pw := io.Pipe()

		resp := &http.Response{
			StatusCode: 200,
			Body:       pr,
		}

		ctx, cancel := context.WithCancel(context.Background())
		channel, err := processStreamingResponseRaw(ctx, resp, 10)

		require.NoError(t, err)
		require.NotNil(t, channel)

		// Write one result first
		go func() {
			_, _ = pw.Write([]byte(`{"result":{"object":"doc:1"}}` + "\n"))
			// Give time for the result to be processed, then close the pipe to unblock the scanner
			time.Sleep(50 * time.Millisecond)
			cancel()
			pw.Close()
		}()

		// Read the first result
		result := <-channel.Results
		assert.JSONEq(t, `{"object":"doc:1"}`, string(result))

		// The channel should close after context is cancelled and pipe is closed
		// Wait for channels to close
		select {
		case _, ok := <-channel.Results:
			if ok {
				// Got another result, that's fine
			}
			// Channel closed, success
		case <-time.After(2 * time.Second):
			t.Fatal("timeout waiting for channel to close")
		}
	})
}

func TestAPIExecutor_ExecuteStreaming(t *testing.T) {
	t.Parallel()

	t.Run("validates_request", func(t *testing.T) {
		t.Parallel()

		client := newTestClient(t, &testRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
			return makeResp(200, "", nil), nil
		}}, nil)

		executor := NewAPIExecutor(client)

		// Missing operation name
		channel, err := executor.ExecuteStreaming(context.Background(), APIExecutorRequest{
			Method: "POST",
			Path:   "/test",
		}, 10)

		assert.Error(t, err)
		assert.Nil(t, channel)
		assert.Contains(t, err.Error(), "operationName is required")
	})

	t.Run("validates_path_parameters", func(t *testing.T) {
		t.Parallel()

		client := newTestClient(t, &testRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
			return makeResp(200, "", nil), nil
		}}, nil)

		executor := NewAPIExecutor(client)

		channel, err := executor.ExecuteStreaming(context.Background(), APIExecutorRequest{
			OperationName: "StreamedListObjects",
			Method:        "POST",
			Path:          "/stores/{store_id}/test",
			// Missing path parameter
		}, 10)

		assert.Error(t, err)
		assert.Nil(t, channel)
		assert.Contains(t, err.Error(), "not all path parameters were provided")
	})

	t.Run("sets_ndjson_accept_header", func(t *testing.T) {
		t.Parallel()

		var capturedReq *http.Request
		client := newTestClient(t, &testRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
			capturedReq = req
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"result":{"object":"doc:1"}}` + "\n")),
				Header:     http.Header{},
			}, nil
		}}, nil)

		executor := NewAPIExecutor(client)

		channel, err := executor.ExecuteStreaming(context.Background(), APIExecutorRequest{
			OperationName:  "StreamedListObjects",
			Method:         "POST",
			Path:           "/stores/{store_id}/test",
			PathParameters: map[string]string{"store_id": "123"},
		}, 10)

		require.NoError(t, err)
		require.NotNil(t, channel)
		defer channel.Close()

		// Drain the channel
		for range channel.Results {
		}

		assert.Equal(t, "application/x-ndjson", capturedReq.Header.Get("Accept"))
	})

	t.Run("handles_http_error", func(t *testing.T) {
		t.Parallel()

		client := newTestClient(t, &testRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 400,
				Body:       io.NopCloser(strings.NewReader(`{"code":"validation_error","message":"Invalid request"}`)),
				Header:     http.Header{"Content-Type": []string{"application/json"}},
			}, nil
		}}, nil)

		executor := NewAPIExecutor(client)

		channel, err := executor.ExecuteStreaming(context.Background(), APIExecutorRequest{
			OperationName:  "StreamedListObjects",
			Method:         "POST",
			Path:           "/stores/{store_id}/test",
			PathParameters: map[string]string{"store_id": "123"},
		}, 10)

		assert.Error(t, err)
		assert.Nil(t, channel)
	})

	t.Run("successful_streaming", func(t *testing.T) {
		t.Parallel()

		ndjson := `{"result":{"object":"doc:1"}}` + "\n" +
			`{"result":{"object":"doc:2"}}` + "\n"

		client := newTestClient(t, &testRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(ndjson)),
				Header:     http.Header{},
			}, nil
		}}, nil)

		executor := NewAPIExecutor(client)

		channel, err := executor.ExecuteStreaming(context.Background(), APIExecutorRequest{
			OperationName:  "StreamedListObjects",
			Method:         "POST",
			Path:           "/stores/{store_id}/streamed-list-objects",
			PathParameters: map[string]string{"store_id": "123"},
			Body:           map[string]string{"type": "document"},
		}, 10)

		require.NoError(t, err)
		require.NotNil(t, channel)
		defer channel.Close()

		var results [][]byte
		for result := range channel.Results {
			results = append(results, result)
		}

		assert.Len(t, results, 2)
		assert.JSONEq(t, `{"object":"doc:1"}`, string(results[0]))
		assert.JSONEq(t, `{"object":"doc:2"}`, string(results[1]))
	})
}

