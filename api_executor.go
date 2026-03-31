package openfga

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
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
// Replaces any previously set path parameters. If params is nil, this is a no-op.
func (b *APIExecutorRequestBuilder) WithPathParameters(params map[string]string) *APIExecutorRequestBuilder {
	if params == nil {
		return b
	}
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
// Replaces any previously set query parameters. If params is nil, this is a no-op.
func (b *APIExecutorRequestBuilder) WithQueryParameters(params url.Values) *APIExecutorRequestBuilder {
	if params == nil {
		return b
	}
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
// Replaces any previously set headers. If headers is nil, this is a no-op.
func (b *APIExecutorRequestBuilder) WithHeaders(headers map[string]string) *APIExecutorRequestBuilder {
	if headers == nil {
		return b
	}
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
	// Example:
	//   request := openfga.NewAPIExecutorRequestBuilder("Check", "POST", "/stores/{store_id}/check").
	//       WithPathParameter("store_id", storeID).
	//       WithBody(checkRequest).
	//       Build()
	//   response, err := executor.Execute(ctx, request)
	Execute(ctx context.Context, request APIExecutorRequest) (*APIExecutorResponse, error)

	// ExecuteWithDecode performs an API request and decodes the response into the provided result pointer.
	// The result parameter must be a pointer to the type you want to decode into.
	//
	// Example:
	//   var response openfga.CheckResponse
	//   request := openfga.NewAPIExecutorRequestBuilder("Check", "POST", "/stores/{store_id}/check").
	//       WithPathParameter("store_id", storeID).
	//       WithBody(checkRequest).
	//       Build()
	//   _, err := executor.ExecuteWithDecode(ctx, request, &response)
	ExecuteWithDecode(ctx context.Context, request APIExecutorRequest, result interface{}) (*APIExecutorResponse, error)

	// ExecuteStreaming performs an API request that returns a streaming response.
	// It returns an APIExecutorStreamingChannel that provides results and errors through channels.
	// The caller is responsible for closing the channel when done using defer channel.Close().
	//
	// This method is useful for streaming API endpoints like StreamedListObjects where
	// each line in the response body is a separate JSON object.
	//
	// Parameters:
	//   - ctx: Context for cancellation. When cancelled, the streaming will stop.
	//   - request: The API request configuration. The Accept header is automatically set to
	//     "application/x-ndjson" unless explicitly overridden.
	//   - bufferSize: The buffer size for the results channel. Use DefaultStreamBufferSize (10) for most cases.
	//
	// Example - Calling StreamedListObjects:
	//
	//   executor := openfga.NewAPIExecutor(client)
	//
	//   request := openfga.NewAPIExecutorRequestBuilder("StreamedListObjects", "POST", "/stores/{store_id}/streamed-list-objects").
	//       WithPathParameter("store_id", storeID).
	//       WithBody(openfga.ListObjectsRequest{
	//           AuthorizationModelId: openfga.PtrString(modelID),
	//           Type:                 "document",
	//           Relation:             "viewer",
	//           User:                 "user:alice",
	//       }).
	//       Build()
	//
	//   channel, err := executor.ExecuteStreaming(ctx, request, openfga.DefaultStreamBufferSize)
	//   if err != nil {
	//       return err
	//   }
	//   defer channel.Close()
	//
	//   for {
	//       select {
	//       case result, ok := <-channel.Results:
	//           if !ok {
	//               return nil // Stream completed
	//           }
	//           var response openfga.StreamedListObjectsResponse
	//           if err := json.Unmarshal(result, &response); err != nil {
	//               return err
	//           }
	//           fmt.Printf("Object: %s\n", response.Object)
	//       case err := <-channel.Errors:
	//           if err != nil {
	//               return err
	//           }
	//       }
	//   }
	ExecuteStreaming(ctx context.Context, request APIExecutorRequest, bufferSize int) (*APIExecutorStreamingChannel, error)
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

// Compile-time check that apiExecutor implements APIExecutor.
var _ APIExecutor = (*apiExecutor)(nil)

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

	if strings.Contains(path, "{") || strings.Contains(path, "}") {
		return nil, reportError("not all path parameters were provided for path: %s", path)
	}

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

	_, _ = metrics.RequestCount(1, attrs)

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

	// Context cancellation errors are not retryable - return immediately
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
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

// ============================================================================
// Streaming API Support
// ============================================================================

// DefaultStreamBufferSize is the default buffer size for streaming channels.
const DefaultStreamBufferSize = 10

// StreamResult represents a streaming result wrapper with either a result or an error.
// This is the format used by OpenFGA's streaming responses.
type StreamResult[T any] struct {
	Result *T      `json:"result,omitempty" yaml:"result,omitempty"`
	Error  *Status `json:"error,omitempty" yaml:"error,omitempty"`
}

// StreamingChannel represents a generic channel for streaming responses.
// It provides typed results directly decoded from the stream.
type StreamingChannel[T any] struct {
	Results chan T
	Errors  chan error
	cancel  context.CancelFunc
}

// Close cancels the streaming context and cleans up resources.
func (s *StreamingChannel[T]) Close() {
	if s.cancel != nil {
		s.cancel()
	}
}

// ProcessStreamingResponse processes an HTTP streaming response
// and returns a StreamingChannel with typed results and errors.
//
// This is a convenience wrapper around processStreamingResponseRaw that adds automatic
// JSON unmarshalling of the raw bytes into the target type T.
//
// Parameters:
//   - ctx: The context for cancellation
//   - httpResponse: The HTTP response to process
//   - bufferSize: The buffer size for the channels (default 10 if <= 0)
//
// Returns:
//   - *StreamingChannel[T]: A channel containing streaming results and errors
//   - error: An error if the response is invalid
func ProcessStreamingResponse[T any](ctx context.Context, httpResponse *http.Response, bufferSize int) (*StreamingChannel[T], error) {
	streamCtx, cancel := context.WithCancel(ctx)

	// Use default buffer size of 10 if not specified or invalid
	if bufferSize <= 0 {
		bufferSize = DefaultStreamBufferSize
	}

	channel := &StreamingChannel[T]{
		Results: make(chan T, bufferSize),
		Errors:  make(chan error, 1),
		cancel:  cancel,
	}

	if httpResponse == nil || httpResponse.Body == nil {
		cancel()
		return nil, errors.New("response or response body is nil")
	}

	go func() {
		defer close(channel.Results)
		defer close(channel.Errors)
		defer cancel()
		defer func() { _ = httpResponse.Body.Close() }()

		scanner := bufio.NewScanner(httpResponse.Body)
		// Allow large NDJSON entries (up to 10MB). Tune as needed.
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 10*1024*1024)

		for scanner.Scan() {
			select {
			case <-streamCtx.Done():
				channel.Errors <- streamCtx.Err()
				return
			default:
				line := scanner.Bytes()
				if len(line) == 0 {
					continue
				}

				var streamResult StreamResult[T]
				if err := json.Unmarshal(line, &streamResult); err != nil {
					channel.Errors <- err
					return
				}

				if streamResult.Error != nil {
					msg := "stream error"
					if streamResult.Error.Message != nil {
						msg = *streamResult.Error.Message
					}
					channel.Errors <- errors.New(msg)
					return
				}

				if streamResult.Result != nil {
					select {
					case <-streamCtx.Done():
						channel.Errors <- streamCtx.Err()
						return
					case channel.Results <- *streamResult.Result:
					}
				}
			}
		}

		if err := scanner.Err(); err != nil {
			// Prefer context error if we were canceled to avoid surfacing net/http "use of closed network connection".
			if streamCtx.Err() != nil {
				channel.Errors <- streamCtx.Err()
				return
			}
			channel.Errors <- err
		}
	}()

	return channel, nil
}

// APIExecutorStreamingChannel represents a channel for streaming API responses.
// It provides two channels: Results for successful responses and Errors for any errors encountered.
//
// Usage pattern:
//
//	channel, err := executor.ExecuteStreaming(ctx, request, 10)
//	if err != nil {
//	    return err
//	}
//	defer channel.Close()
//
//	for {
//	    select {
//	    case result, ok := <-channel.Results:
//	        if !ok {
//	            // Channel closed, check for errors
//	            select {
//	            case err := <-channel.Errors:
//	                if err != nil {
//	                    return err
//	                }
//	            default:
//	            }
//	            return nil
//	        }
//	        // Process result (raw JSON bytes)
//	        var response YourResponseType
//	        json.Unmarshal(result, &response)
//	    case err := <-channel.Errors:
//	        if err != nil {
//	            return err
//	        }
//	    }
//	}
type APIExecutorStreamingChannel struct {
	// Results channel receives raw JSON bytes for each streamed result.
	// The channel is closed when the stream ends or an error occurs.
	Results chan []byte

	// Errors channel receives any errors that occur during streaming.
	// Only one error will be sent before the channel is closed.
	Errors chan error

	// cancel is the function to cancel the streaming context
	cancel context.CancelFunc
}

// Close cancels the streaming context and cleans up resources.
// It is safe to call Close multiple times.
// Always defer Close() after successfully creating a streaming channel.
func (s *APIExecutorStreamingChannel) Close() {
	if s.cancel != nil {
		s.cancel()
	}
}

// ExecuteStreaming performs an API request that returns a streaming response.
// It returns an APIExecutorStreamingChannel that provides results and errors through channels.
// The caller is responsible for closing the channel when done using defer channel.Close().
//
// Streaming responses are line-delimited JSON where each line is a JSON object.
// Each line is expected to have either a "result" or "error" field wrapped in a StreamResult structure.
//
// Parameters:
//   - ctx: Context for cancellation. When cancelled, the streaming will stop.
//   - request: The API request configuration. Headers should include "Accept": "application/x-ndjson" for streaming.
//   - bufferSize: The buffer size for the results channel. Use DefaultStreamBufferSize (10) for most cases.
//     A larger buffer can improve throughput but uses more memory.
//
// Example - Calling StreamedListObjects:
//
//	executor := openfga.NewAPIExecutor(client)
//
//	request := openfga.NewAPIExecutorRequestBuilder("StreamedListObjects", "POST", "/stores/{store_id}/streamed-list-objects").
//	    WithPathParameter("store_id", storeID).
//	    WithHeader("Accept", "application/x-ndjson").
//	    WithBody(openfga.ListObjectsRequest{
//	        AuthorizationModelId: openfga.PtrString(modelID),
//	        Type:                 "document",
//	        Relation:             "viewer",
//	        User:                 "user:alice",
//	    }).
//	    Build()
//
//	channel, err := executor.ExecuteStreaming(ctx, request, openfga.DefaultStreamBufferSize)
//	if err != nil {
//	    return err
//	}
//	defer channel.Close()
//
//	for {
//	    select {
//	    case result, ok := <-channel.Results:
//	        if !ok {
//	            // Stream completed
//	            return nil
//	        }
//	        var response openfga.StreamedListObjectsResponse
//	        if err := json.Unmarshal(result, &response); err != nil {
//	            return err
//	        }
//	        fmt.Printf("Object: %s\n", response.Object)
//	    case err := <-channel.Errors:
//	        if err != nil {
//	            return err
//	        }
//	    }
//	}
func (e *apiExecutor) ExecuteStreaming(ctx context.Context, request APIExecutorRequest, bufferSize int) (*APIExecutorStreamingChannel, error) {
	// Validate required fields
	if err := validateRequest(request); err != nil {
		return nil, err
	}

	// Build request parameters
	path := buildPath(request.Path, request.PathParameters)

	if strings.Contains(path, "{") || strings.Contains(path, "}") {
		return nil, reportError("not all path parameters were provided for path: %s", path)
	}

	headerParams := prepareHeaders(request.Headers)
	// Ensure Accept header is set for streaming unless already overridden
	if headerParams["Accept"] == "application/json" {
		headerParams["Accept"] = "application/x-ndjson"
	}

	queryParams := request.QueryParameters
	if queryParams == nil {
		queryParams = url.Values{}
	}

	storeID := request.PathParameters["store_id"]

	// Prepare HTTP request
	req, err := e.client.prepareRequest(ctx, path, request.Method, request.Body, headerParams, queryParams)
	if err != nil {
		return nil, err
	}

	// Execute HTTP request
	httpResponse, err := e.client.callAPI(req)
	if err != nil {
		return nil, err
	}
	if httpResponse == nil {
		return nil, reportError("nil HTTP response from API client")
	}

	// Handle HTTP errors (status >= 300)
	if httpResponse.StatusCode >= http.StatusMultipleChoices {
		responseBody, readErr := io.ReadAll(httpResponse.Body)
		_ = httpResponse.Body.Close()
		if readErr != nil {
			return nil, readErr
		}
		return nil, e.client.handleAPIError(httpResponse, responseBody, request.Body, request.OperationName, storeID)
	}

	// Process streaming response
	return processStreamingResponseRaw(ctx, httpResponse, bufferSize)
}

// processStreamingResponseRaw processes an HTTP streaming response.
// It returns an APIExecutorStreamingChannel with raw JSON bytes for each result.
func processStreamingResponseRaw(ctx context.Context, httpResponse *http.Response, bufferSize int) (*APIExecutorStreamingChannel, error) {
	streamCtx, cancel := context.WithCancel(ctx)

	// Use default buffer size if not specified or invalid
	if bufferSize <= 0 {
		bufferSize = DefaultStreamBufferSize
	}

	channel := &APIExecutorStreamingChannel{
		Results: make(chan []byte, bufferSize),
		Errors:  make(chan error, 1),
		cancel:  cancel,
	}

	if httpResponse == nil || httpResponse.Body == nil {
		cancel()
		return nil, reportError("response or response body is nil")
	}

	go func() {
		defer close(channel.Results)
		defer close(channel.Errors)
		defer cancel()
		defer func() { _ = httpResponse.Body.Close() }()

		scanner := bufio.NewScanner(httpResponse.Body)
		// Allow large NDJSON entries (up to 10MB). Tune as needed.
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 10*1024*1024)

		for scanner.Scan() {
			select {
			case <-streamCtx.Done():
				channel.Errors <- streamCtx.Err()
				return
			default:
				line := scanner.Bytes()
				if len(line) == 0 {
					continue
				}

				// Parse the StreamResult wrapper to check for errors
				var streamResult struct {
					Result json.RawMessage `json:"result,omitempty"`
					Error  *Status         `json:"error,omitempty"`
				}
				if err := json.Unmarshal(line, &streamResult); err != nil {
					channel.Errors <- err
					return
				}

				if streamResult.Error != nil {
					msg := "stream error"
					if streamResult.Error.Message != nil {
						msg = *streamResult.Error.Message
					}
					channel.Errors <- errors.New(msg)
					return
				}

				if streamResult.Result != nil {
					// Make a copy of the raw JSON to send through the channel
					resultCopy := make([]byte, len(streamResult.Result))
					copy(resultCopy, streamResult.Result)

					select {
					case <-streamCtx.Done():
						channel.Errors <- streamCtx.Err()
						return
					case channel.Results <- resultCopy:
					}
				}
			}
		}

		if err := scanner.Err(); err != nil {
			// Prefer context error if we were canceled to avoid surfacing net/http "use of closed network connection".
			if streamCtx.Err() != nil {
				channel.Errors <- streamCtx.Err()
				return
			}
			channel.Errors <- err
		}
	}()

	return channel, nil
}
