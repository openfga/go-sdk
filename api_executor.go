package openfga

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/openfga/go-sdk/internal/constants"
	"github.com/openfga/go-sdk/internal/utils/retryutils"
	"github.com/openfga/go-sdk/telemetry"
)

// APIExecutorRequest represents a request to be executed by the API executor.
type APIExecutorRequest struct {
	// OperationName is a descriptive name for the operation (e.g., "Check", "Write", "BatchCheck").
	// Used for logging, telemetry, and error messages.
	OperationName string

	// Method is the HTTP method (GET, POST, PUT, DELETE, etc.).
	Method string

	// Path is the API path (e.g., "/stores/{store_id}/check").
	// Template parameters in curly braces will be replaced using PathParameters.
	Path string

	// PathParameters maps path template variables to their values.
	// For example, for "/stores/{store_id}/check", provide {"store_id": "12345"}.
	PathParameters map[string]string

	// QueryParameters contains URL query parameters.
	QueryParameters url.Values

	// Body is the request payload (will be JSON encoded).
	// Should typically be nil for GET/DELETE requests.
	Body interface{}

	// Headers contains custom HTTP headers.
	// Can override default headers like Content-Type and Accept if needed.
	Headers map[string]string

	// TODO: Add support for request options like per-request timeouts, cancellation, retry options, etc.
}

// APIExecutorRequestBuilder provides a fluent interface for building APIExecutorRequest instances.
type APIExecutorRequestBuilder struct {
	request APIExecutorRequest
}

// NewAPIExecutorRequestBuilder creates a new builder with required fields.
// operationName: descriptive name for the operation (e.g., "Check", "Write")
// method: HTTP method (GET, POST, PUT, DELETE, etc.)
// path: API path with optional template parameters (e.g., "/stores/{store_id}/check")
func NewAPIExecutorRequestBuilder(operationName, method, path string) *APIExecutorRequestBuilder {
	return &APIExecutorRequestBuilder{
		request: APIExecutorRequest{
			OperationName:   operationName,
			Method:          method,
			Path:            path,
			PathParameters:  make(map[string]string),
			QueryParameters: url.Values{},
			Headers:         make(map[string]string),
		},
	}
}

// WithPathParameter adds a single path parameter to the request.
// The parameter will be used to replace template variables in the path.
func (b *APIExecutorRequestBuilder) WithPathParameter(key, value string) *APIExecutorRequestBuilder {
	if b.request.PathParameters == nil {
		b.request.PathParameters = make(map[string]string)
	}

	b.request.PathParameters[key] = value
	return b
}

// WithPathParameters sets all path parameters at once.
// Replaces any previously set path parameters.
func (b *APIExecutorRequestBuilder) WithPathParameters(params map[string]string) *APIExecutorRequestBuilder {
	b.request.PathParameters = params
	return b
}

// WithQueryParameter adds a single query parameter to the request.
func (b *APIExecutorRequestBuilder) WithQueryParameter(key, value string) *APIExecutorRequestBuilder {
	if b.request.QueryParameters == nil {
		b.request.QueryParameters = url.Values{}
	}

	b.request.QueryParameters.Add(key, value)
	return b
}

// WithQueryParameters sets all query parameters at once.
// Replaces any previously set query parameters.
func (b *APIExecutorRequestBuilder) WithQueryParameters(params url.Values) *APIExecutorRequestBuilder {
	b.request.QueryParameters = params
	return b
}

// WithBody sets the request body (will be JSON encoded).
func (b *APIExecutorRequestBuilder) WithBody(body interface{}) *APIExecutorRequestBuilder {
	b.request.Body = body
	return b
}

// WithHeader adds a single custom header to the request.
func (b *APIExecutorRequestBuilder) WithHeader(key, value string) *APIExecutorRequestBuilder {
	if b.request.Headers == nil {
		b.request.Headers = make(map[string]string)
	}

	b.request.Headers[key] = value
	return b
}

// WithHeaders sets all custom headers at once.
// Replaces any previously set headers.
func (b *APIExecutorRequestBuilder) WithHeaders(headers map[string]string) *APIExecutorRequestBuilder {
	b.request.Headers = headers
	return b
}

// Build returns the constructed APIExecutorRequest.
func (b *APIExecutorRequestBuilder) Build() APIExecutorRequest {
	return b.request
}

// APIExecutorResponse represents the response from an API execution.
type APIExecutorResponse struct {
	// HTTPResponse is the raw HTTP response.
	HTTPResponse *http.Response

	// Body contains the raw response body bytes.
	Body []byte

	// StatusCode is the HTTP status code.
	StatusCode int

	// Headers contains the response headers.
	Headers http.Header
}

// APIExecutor provides a generic interface for executing API requests with retry logic, telemetry, and error handling.
type APIExecutor interface {
	// Execute performs an API request with automatic retry logic, telemetry, and error handling.
	// It returns the raw response that can be decoded manually.
	//
	// Example using struct literal:
	//    openfga.APIExecutorRequest{
	//       OperationName: "Check",
	//       Method:        "POST",
	//       Path:          "/stores/{store_id}/check",
	//       PathParameters: map[string]string{"store_id": storeID},
	//       Body:          checkRequest,
	//   }
	//   response, err := executor.Execute(ctx, request)
	//
	// Example using builder pattern:
	//   request := openfga.NewAPIExecutorRequestBuilder("Check", "POST", "/stores/{store_id}/check").
	//       WithPathParameter("store_id", storeID).
	//       WithBody(checkRequest).
	//       Build()
	//   response, err := executor.Execute(ctx, request)
	Execute(ctx context.Context, request APIExecutorRequest) (*APIExecutorResponse, error)

	// ExecuteWithDecode performs an API request and decodes the response into the provided result pointer.
	// The result parameter must be a pointer to the type you want to decode into.
	//
	// Example using struct literal:
	//   var response openfga.CheckResponse
	//   openfga.APIExecutorRequest{
	//       OperationName: "Check",
	//       Method:        "POST",
	//       Path:          "/stores/{store_id}/check",
	//       PathParameters: map[string]string{"store_id": storeID},
	//       Body:          checkRequest,
	//   }
	//   _, err := executor.ExecuteWithDecode(ctx, request, &response)
	//
	// Example using builder pattern:
	//   var response openfga.CheckResponse
	//   request := openfga.NewAPIExecutorRequestBuilder("Check", "POST", "/stores/{store_id}/check").
	//       WithPathParameter("store_id", storeID).
	//       WithBody(checkRequest).
	//       Build()
	//   _, err := executor.ExecuteWithDecode(ctx, request, &response)
	ExecuteWithDecode(ctx context.Context, request APIExecutorRequest, result interface{}) (*APIExecutorResponse, error)
}

// validateRequest checks that required fields are present in the request.
func validateRequest(request APIExecutorRequest) error {
	if request.OperationName == "" {
		return reportError("operationName is required")
	}
	if request.Method == "" {
		return reportError("method is required")
	}
	if request.Path == "" {
		return reportError("path is required")
	}
	return nil
}

// buildPath replaces template parameters in the path (e.g., {store_id}) with actual values.
func buildPath(template string, params map[string]string) string {
	path := template
	for key, value := range params {
		placeholder := "{" + key + "}"
		path = strings.ReplaceAll(path, placeholder, url.PathEscape(value))
	}
	return path
}

// prepareHeaders creates the header map with defaults and applies custom headers.
func prepareHeaders(customHeaders map[string]string) map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	// Apply custom headers (may override defaults)
	for key, value := range customHeaders {
		headers[key] = value
	}

	return headers
}

// makeAPIExecutorResponse creates an APIExecutorResponse from an HTTP response and body.
func makeAPIExecutorResponse(httpResponse *http.Response, body []byte) *APIExecutorResponse {
	return &APIExecutorResponse{
		HTTPResponse: httpResponse,
		Body:         body,
		StatusCode:   httpResponse.StatusCode,
		Headers:      httpResponse.Header,
	}
}

// apiExecutor is the internal implementation of APIExecutor.
type apiExecutor struct {
	client *APIClient
}

// NewAPIExecutor creates a new APIExecutor instance.
// This allows users to call any OpenFGA API endpoint, including those not yet supported by the SDK.
func NewAPIExecutor(client *APIClient) APIExecutor {
	return &apiExecutor{client: client}
}

// Execute performs an API request with automatic retry logic and error handling.
func (e *apiExecutor) Execute(ctx context.Context, request APIExecutorRequest) (*APIExecutorResponse, error) {
	return e.executeInternal(ctx, request, nil)
}

// ExecuteWithDecode performs an API request and decodes the response into the provided result pointer.
func (e *apiExecutor) ExecuteWithDecode(ctx context.Context, request APIExecutorRequest, result interface{}) (*APIExecutorResponse, error) {
	return e.executeInternal(ctx, request, result)
}

// executeInternal is the core execution logic used by both Execute and ExecuteWithDecode.
func (e *apiExecutor) executeInternal(ctx context.Context, request APIExecutorRequest, result interface{}) (*APIExecutorResponse, error) {
	requestStarted := time.Now()

	// Validate required fields
	if err := validateRequest(request); err != nil {
		return nil, err
	}

	// Build request parameters
	path := buildPath(request.Path, request.PathParameters)
	headerParams := prepareHeaders(request.Headers)
	queryParams := request.QueryParameters
	if queryParams == nil {
		queryParams = url.Values{}
	}

	// Get retry configuration
	retryParams := e.getRetryParams()
	storeID := request.PathParameters["store_id"]

	var lastResponse *APIExecutorResponse

	// Execute request with retry logic
	for attemptNum := 0; attemptNum < retryParams.MaxRetry+1; attemptNum++ {
		response, err := e.executeSingleAttempt(ctx, request, path, headerParams, queryParams, attemptNum, requestStarted, storeID)
		if err == nil && response != nil {
			// Decode response if needed
			if result != nil {
				if decodeErr := e.client.decode(result, response.Body, response.Headers.Get("Content-Type")); decodeErr != nil {
					return response, GenericOpenAPIError{
						body:  response.Body,
						error: decodeErr.Error(),
					}
				}
			}
			return response, nil
		}

		lastResponse = response

		// Check if we should retry
		if attemptNum >= retryParams.MaxRetry {
			return lastResponse, err
		}

		// Determine if we should retry and how long to wait
		if shouldRetry, waitDuration := e.determineRetry(err, response, attemptNum, retryParams, request.OperationName); shouldRetry {
			if e.client.cfg.Debug {
				e.logRetry(request, err, response, attemptNum, waitDuration)
			}
			time.Sleep(waitDuration)
			continue
		}

		// Error is not retryable
		return lastResponse, err
	}

	// All retries exhausted
	if lastResponse != nil {
		return lastResponse, reportError("max retries exceeded")
	}
	return nil, reportError("request failed without response")
}

// getRetryParams returns the retry parameters, using defaults if not configured.
func (e *apiExecutor) getRetryParams() RetryParams {
	if e.client.cfg.RetryParams != nil {
		return *e.client.cfg.RetryParams
	}
	return RetryParams{
		MaxRetry:    constants.DefaultMaxRetry,
		MinWaitInMs: constants.DefaultMinWaitInMs,
	}
}

// recordTelemetry records request telemetry metrics.
func (e *apiExecutor) recordTelemetry(operationName string, storeID string, body interface{}, req *http.Request, httpResponse *http.Response, requestStarted time.Time, attemptNum int) {
	metrics := telemetry.GetMetrics(telemetry.TelemetryFactoryParameters{Configuration: e.client.cfg.Telemetry})
	attrs, queryDuration, requestDuration, _ := metrics.BuildTelemetryAttributes(
		operationName,
		map[string]interface{}{
			"storeId": storeID,
			"body":    body,
		},
		req,
		httpResponse,
		requestStarted,
		attemptNum,
	)

	if requestDuration > 0 {
		_, _ = metrics.RequestDuration(requestDuration, attrs)
	}

	if queryDuration > 0 {
		_, _ = metrics.QueryDuration(queryDuration, attrs)
	}
}

// executeSingleAttempt performs a single HTTP request attempt and handles the response.
func (e *apiExecutor) executeSingleAttempt(
	ctx context.Context,
	request APIExecutorRequest,
	path string,
	headerParams map[string]string,
	queryParams url.Values,
	attemptNum int,
	requestStarted time.Time,
	storeID string,
) (*APIExecutorResponse, error) {
	// Prepare HTTP request
	req, err := e.client.prepareRequest(ctx, path, request.Method, request.Body, headerParams, queryParams)
	if err != nil {
		return nil, err
	}

	// Execute HTTP request
	httpResponse, err := e.client.callAPI(req)
	if err != nil || httpResponse == nil {
		return nil, err
	}

	// Read response body
	responseBody, err := io.ReadAll(httpResponse.Body)
	_ = httpResponse.Body.Close()
	httpResponse.Body = io.NopCloser(bytes.NewBuffer(responseBody))
	if err != nil {
		return makeAPIExecutorResponse(httpResponse, responseBody), err
	}

	response := makeAPIExecutorResponse(httpResponse, responseBody)

	// Handle HTTP errors (status >= 300)
	if httpResponse.StatusCode >= http.StatusMultipleChoices {
		apiErr := e.client.handleAPIError(httpResponse, responseBody, request.Body, request.OperationName, storeID)
		return response, apiErr
	}

	// Record telemetry for successful requests
	e.recordTelemetry(request.OperationName, storeID, request.Body, req, httpResponse, requestStarted, attemptNum)

	return response, nil
}

// determineRetry decides whether to retry a failed request and returns the wait duration.
func (e *apiExecutor) determineRetry(
	err error,
	response *APIExecutorResponse,
	attemptNum int,
	retryParams RetryParams,
	operationName string,
) (bool, time.Duration) {
	if err == nil {
		return false, 0
	}

	// Check for rate limit or internal server errors that support retry
	var rateLimitErr FgaApiRateLimitExceededError
	var internalErr FgaApiInternalError

	switch {
	case errors.As(err, &rateLimitErr):
		timeToWait := rateLimitErr.GetTimeToWait(attemptNum, retryParams)
		return timeToWait > 0, timeToWait
	case errors.As(err, &internalErr):
		timeToWait := internalErr.GetTimeToWait(attemptNum, retryParams)
		return timeToWait > 0, timeToWait
	default:
		// Network errors or body read errors
		headers := http.Header{}
		if response != nil {
			headers = response.Headers
		}
		timeToWait := retryutils.GetTimeToWait(attemptNum, retryParams.MaxRetry, retryParams.MinWaitInMs, headers, operationName)
		return timeToWait > 0, timeToWait
	}
}

// logRetry logs retry information for debugging.
func (e *apiExecutor) logRetry(request APIExecutorRequest, err error, response *APIExecutorResponse, attemptNum int, waitDuration time.Duration) {
	if response != nil {
		log.Printf("\nWaiting %v to retry %v (attempt %d, status=%d, error=%v). Request body: %v\n",
			waitDuration, request.OperationName, attemptNum, response.StatusCode, err, request.Body)
	} else {
		log.Printf("\nWaiting %v to retry %v (attempt %d, error=%v). Request body: %v\n",
			waitDuration, request.OperationName, attemptNum, err, request.Body)
	}
}
