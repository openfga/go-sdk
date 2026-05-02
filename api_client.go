package openfga

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/openfga/go-sdk/internal/utils/retryutils"
	"github.com/openfga/go-sdk/telemetry"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// ErrorResponse defines the error that will be asserted by FGA API.
// This will only be used for error that is not defined
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// APIClient manages communication with the OpenFGA API v1.x
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	OpenFgaApi OpenFgaApi
}

type service struct {
	client      *APIClient
	RetryParams *RetryParams
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.Telemetry == nil {
		cfg.Telemetry = telemetry.DefaultTelemetryConfiguration()
	}

	// Set default HTTP client if none provided
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	// Process credentials if provided
	if cfg.Credentials != nil {
		cfg.Credentials.Context = context.Background()
		telemetry.Bind(cfg.Credentials.Context, telemetry.Get(telemetry.TelemetryFactoryParameters{Configuration: cfg.Telemetry}))
		var httpClient, headers = cfg.Credentials.GetHttpClientAndHeaderOverrides(retryutils.GetRetryParamsOrDefault(cfg.RetryParams), cfg.Debug, cfg.HTTPClient)

		// Always apply header overrides (works for ApiToken and other header-based auth)
		if len(headers) > 0 {
			for idx := range headers {
				cfg.AddDefaultHeader(headers[idx].Key, headers[idx].Value)
			}
		}

		// Handle OAuth2 client for ClientCredentials
		// GetHttpClientAndHeaderOverrides returns http.DefaultClient for ApiToken/None
		// and returns an OAuth2-enabled client for ClientCredentials
		if httpClient != nil && httpClient != http.DefaultClient {
			// An OAuth2 client was returned (ClientCredentials method)
			// The custom HTTPClient (if provided) is now wrapped by the OAuth2 client
			cfg.HTTPClient = httpClient
		}
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c
	c.common.RetryParams = cfg.RetryParams

	// API Services
	c.OpenFgaApi = (*OpenFgaApiService)(&c.common)

	return c
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	loweredNeedle := strings.ToLower(needle)
	for _, a := range haystack {
		if strings.ToLower(a) == loweredNeedle {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.ReplaceAll(fmt.Sprint(obj), " ", delimiter), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	// Track HTTP request duration
	httpRequestStarted := time.Now()
	resp, err := c.cfg.HTTPClient.Do(request)
	httpRequestDuration := time.Since(httpRequestStarted).Seconds() * 1000

	// Emit HTTP request duration metric for each HTTP call
	{
		metrics := telemetry.GetMetrics(telemetry.TelemetryFactoryParameters{Configuration: c.cfg.Telemetry})
		httpRequestAttrs := map[*telemetry.Attribute]string{
			telemetry.HTTPHost:          request.URL.Host,
			telemetry.HTTPRequestMethod: request.Method,
			telemetry.URLFull:           request.URL.String(),
			telemetry.URLScheme:         request.URL.Scheme,
			telemetry.UserAgent:         request.UserAgent(),
		}
		if resp != nil {
			httpRequestAttrs[telemetry.HTTPResponseStatusCode] = fmt.Sprintf("%d", resp.StatusCode)
		}
		_, _ = metrics.HTTPRequestDuration(httpRequestDuration, httpRequestAttrs)
	}

	if err != nil {
		if resp != nil && resp.Request == nil {
			resp.Request = request
		}

		return resp, err
	}

	if c.cfg.Debug {
		// for debugging, don't dump the body for streaming resp. as it would buffer the entire response
		// only dump headers
		isStreamingResponse := resp.Header.Get("Content-Type") == "application/x-ndjson"
		dump, err := httputil.DumpResponse(resp, !isStreamingResponse)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
		if isStreamingResponse {
			log.Printf("Streaming response body - not dumped to preserve streaming\n")
		}
	}

	if resp.Request == nil {
		resp.Request = request
	}

	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

// GetAPIExecutor returns an APIExecutor that can be used to call any OpenFGA API endpoint
// This is useful for calling endpoints that are not yet supported by the SDK
func (c *APIClient) GetAPIExecutor() APIExecutor {
	return NewAPIExecutor(c)
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// Setup path and query parameters
	uri, err := url.Parse(c.cfg.ApiUrl + path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := uri.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	uri.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, uri.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, uri.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Set("User-Agent", c.cfg.UserAgent)

	for header, value := range c.cfg.DefaultHeaders {
		if localVarRequest.Header.Get(header) == "" {
			localVarRequest.Header.Set(header, value)
		}
	}

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if xmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if jsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

func (c *APIClient) handleAPIError(httpResponse *http.Response, responseBody []byte, requestBody interface{}, operationName string, storeId string) error {
	switch httpResponse.StatusCode {
	case http.StatusBadRequest, http.StatusUnprocessableEntity:
		err := NewFgaApiValidationError(operationName, requestBody, httpResponse, responseBody, storeId)
		var v ValidationErrorMessageResponse
		errBody := c.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
		if errBody != nil {
			err.modelDecodeError = err
			return err
		}
		err.model = v
		err.responseCode = v.GetCode()
		err.error += " with error code " + string(v.GetCode()) + " error message: " + v.GetMessage()
		return err

	case http.StatusUnauthorized, http.StatusForbidden:
		return NewFgaApiAuthenticationError(operationName, requestBody, httpResponse, responseBody, storeId)

	case http.StatusNotFound:
		err := NewFgaApiNotFoundError(operationName, requestBody, httpResponse, responseBody, storeId)
		var v PathUnknownErrorMessageResponse
		errBody := c.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
		if errBody != nil {
			err.modelDecodeError = err
			return err
		}
		err.model = v
		err.responseCode = v.GetCode()
		err.error += " with error code " + string(v.GetCode()) + " error message: " + v.GetMessage()
		return err

	case http.StatusTooManyRequests:
		return NewFgaApiRateLimitExceededError(operationName, requestBody, httpResponse, responseBody, storeId)

	default:
		if httpResponse.StatusCode >= http.StatusInternalServerError {
			err := NewFgaApiInternalError(operationName, requestBody, httpResponse, responseBody, storeId)
			var v InternalErrorMessageResponse
			errBody := c.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if errBody != nil {
				err.modelDecodeError = err
				return err
			}
			err.model = v
			err.responseCode = v.GetCode()
			err.error += " with error code " + string(v.GetCode()) + " error message: " + v.GetMessage()
			return err
		}

		err := NewFgaApiError(operationName, requestBody, httpResponse, responseBody, storeId)
		var v ErrorResponse
		errBody := c.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
		if errBody != nil {
			err.modelDecodeError = err
			return err
		}
		err.model = v
		err.responseCode = v.Code
		err.error += " with error code " + string(v.Code) + " error message: " + v.Message
		return err
	}
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(**os.File); ok {
		_, err = bodyBuf.ReadFrom(*fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}
