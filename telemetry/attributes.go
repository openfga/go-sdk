package telemetry

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/otel/attribute"
)

const (
	ATTR_FGA_CLIENT_REQUEST_CLIENT_ID        = "fga-client.request.client_id"
	ATTR_FGA_CLIENT_REQUEST_METHOD           = "fga-client.request.method"
	ATTR_FGA_CLIENT_REQUEST_MODEL_ID         = "fga-client.request.model_id"
	ATTR_FGA_CLIENT_REQUEST_STORE_ID         = "fga-client.request.store_id"
	ATTR_FGA_CLIENT_REQUEST_BATCH_CHECK_SIZE = "fga-client.request.batch_check_size"
	ATTR_FGA_CLIENT_RESPONSE_MODEL_ID        = "fga-client.response.model_id"
	ATTR_FGA_CLIENT_USER                     = "fga-client.user"
	ATTR_HTTP_CLIENT_REQUEST_DURATION        = "http.client.request.duration"
	ATTR_HTTP_HOST                           = "http.host"
	ATTR_HTTP_REQUEST_METHOD                 = "http.request.method"
	ATTR_HTTP_REQUEST_RESEND_COUNT           = "http.request.resend_count"
	ATTR_HTTP_RESPONSE_STATUS_CODE           = "http.response.status_code"
	ATTR_HTTP_SERVER_REQUEST_DURATION        = "http.server.request.duration"
	ATTR_URL_SCHEME                          = "url.scheme"
	ATTR_URL_FULL                            = "url.full"
	ATTR_USER_AGENT_ORIGINAL                 = "user_agent.original"
)

var (
	FGAClientRequestClientID       = &Attribute{Name: ATTR_FGA_CLIENT_REQUEST_CLIENT_ID}
	FGAClientRequestMethod         = &Attribute{Name: ATTR_FGA_CLIENT_REQUEST_METHOD}
	FGAClientRequestModelID        = &Attribute{Name: ATTR_FGA_CLIENT_REQUEST_MODEL_ID}
	FGAClientRequestStoreID        = &Attribute{Name: ATTR_FGA_CLIENT_REQUEST_STORE_ID}
	FGAClientRequestBatchCheckSize = &Attribute{Name: ATTR_FGA_CLIENT_REQUEST_BATCH_CHECK_SIZE}
	FGAClientResponseModelID       = &Attribute{Name: ATTR_FGA_CLIENT_RESPONSE_MODEL_ID}
	FGAClientUser                  = &Attribute{Name: ATTR_FGA_CLIENT_USER}
	HTTPClientRequestDuration      = &Attribute{Name: ATTR_HTTP_CLIENT_REQUEST_DURATION}
	HTTPHost                       = &Attribute{Name: ATTR_HTTP_HOST}
	HTTPRequestMethod              = &Attribute{Name: ATTR_HTTP_REQUEST_METHOD}
	HTTPRequestResendCount         = &Attribute{Name: ATTR_HTTP_REQUEST_RESEND_COUNT}
	HTTPResponseStatusCode         = &Attribute{Name: ATTR_HTTP_RESPONSE_STATUS_CODE}
	HTTPServerRequestDuration      = &Attribute{Name: ATTR_HTTP_SERVER_REQUEST_DURATION}
	URLScheme                      = &Attribute{Name: ATTR_URL_SCHEME}
	URLFull                        = &Attribute{Name: ATTR_URL_FULL}
	UserAgent                      = &Attribute{Name: ATTR_USER_AGENT_ORIGINAL}
)

func (m *Metrics) PrepareAttributes(metric MetricInterface, attrs map[*Attribute]string, config *MetricsConfiguration) (attribute.Set, error) {
	var prepared []attribute.KeyValue
	var allowed *MetricConfiguration

	if config == nil || metric == nil || attrs == nil {
		return *attribute.EmptySet(), nil
	}

	switch metric.GetName() {
	case METRIC_COUNTER_CREDENTIALS_REQUEST:
		if config.METRIC_COUNTER_CREDENTIALS_REQUEST == nil {
			return *attribute.EmptySet(), nil
		}

		allowed = config.METRIC_COUNTER_CREDENTIALS_REQUEST
	case METRIC_HISTOGRAM_REQUEST_DURATION:
		if config.METRIC_HISTOGRAM_REQUEST_DURATION == nil {
			return *attribute.EmptySet(), nil
		}

		allowed = config.METRIC_HISTOGRAM_REQUEST_DURATION
	case METRIC_HISTOGRAM_QUERY_DURATION:
		if config.METRIC_HISTOGRAM_QUERY_DURATION == nil {
			return *attribute.EmptySet(), nil
		}

		allowed = config.METRIC_HISTOGRAM_QUERY_DURATION
	case METRIC_HISTOGRAM_HTTP_REQUEST_DURATION:
		if config.METRIC_HISTOGRAM_HTTP_REQUEST_DURATION == nil {
			return *attribute.EmptySet(), nil
		}

		allowed = config.METRIC_HISTOGRAM_HTTP_REQUEST_DURATION
	}

	if allowed == nil {
		return *attribute.EmptySet(), nil
	}

	for attr, value := range attrs {
		if attr == nil {
			continue
		}

		switch attr {
		case FGAClientRequestClientID:
			if allowed.ATTR_FGA_CLIENT_REQUEST_CLIENT_ID == nil || !allowed.ATTR_FGA_CLIENT_REQUEST_CLIENT_ID.Enabled {
				continue
			}
		case FGAClientRequestMethod:
			if allowed.ATTR_HTTP_REQUEST_METHOD == nil || !allowed.ATTR_HTTP_REQUEST_METHOD.Enabled {
				continue
			}
		case FGAClientRequestModelID:
			if allowed.ATTR_FGA_CLIENT_REQUEST_MODEL_ID == nil || !allowed.ATTR_FGA_CLIENT_REQUEST_MODEL_ID.Enabled {
				continue
			}
		case FGAClientRequestStoreID:
			if allowed.ATTR_FGA_CLIENT_REQUEST_STORE_ID == nil || !allowed.ATTR_FGA_CLIENT_REQUEST_STORE_ID.Enabled {
				continue
			}
		case FGAClientRequestBatchCheckSize:
			if allowed.ATTR_FGA_CLIENT_REQUEST_BATCH_CHECK_SIZE == nil || !allowed.ATTR_FGA_CLIENT_REQUEST_BATCH_CHECK_SIZE.Enabled {
				continue
			}
		case FGAClientResponseModelID:
			if allowed.ATTR_FGA_CLIENT_RESPONSE_MODEL_ID == nil || !allowed.ATTR_FGA_CLIENT_RESPONSE_MODEL_ID.Enabled {
				continue
			}
		case FGAClientUser:
			if allowed.ATTR_FGA_CLIENT_USER == nil || !allowed.ATTR_FGA_CLIENT_USER.Enabled {
				continue
			}
		case HTTPClientRequestDuration:
			if allowed.ATTR_HTTP_CLIENT_REQUEST_DURATION == nil || !allowed.ATTR_HTTP_CLIENT_REQUEST_DURATION.Enabled {
				continue
			}
		case HTTPHost:
			if allowed.ATTR_HTTP_HOST == nil || !allowed.ATTR_HTTP_HOST.Enabled {
				continue
			}
		case HTTPRequestMethod:
			if allowed.ATTR_HTTP_REQUEST_METHOD == nil || !allowed.ATTR_HTTP_REQUEST_METHOD.Enabled {
				continue
			}
		case HTTPRequestResendCount:
			if allowed.ATTR_HTTP_REQUEST_RESEND_COUNT == nil || !allowed.ATTR_HTTP_REQUEST_RESEND_COUNT.Enabled {
				continue
			}
		case HTTPResponseStatusCode:
			if allowed.ATTR_HTTP_RESPONSE_STATUS_CODE == nil || !allowed.ATTR_HTTP_RESPONSE_STATUS_CODE.Enabled {
				continue
			}
		case HTTPServerRequestDuration:
			if allowed.ATTR_HTTP_SERVER_REQUEST_DURATION == nil || !allowed.ATTR_HTTP_SERVER_REQUEST_DURATION.Enabled {
				continue
			}
		case URLScheme:
			if allowed.ATTR_URL_SCHEME == nil || !allowed.ATTR_URL_SCHEME.Enabled {
				continue
			}
		case URLFull:
			if allowed.ATTR_URL_FULL == nil || !allowed.ATTR_URL_FULL.Enabled {
				continue
			}
		case UserAgent:
			if allowed.ATTR_USER_AGENT_ORIGINAL == nil || !allowed.ATTR_USER_AGENT_ORIGINAL.Enabled {
				continue
			}
		}

		prepared = append(prepared, attribute.String(attr.Name, value))
	}

	return attribute.NewSet(prepared...), nil
}

func (m *Metrics) AttributesFromRequest(req *http.Request, params map[string]interface{}) (map[*Attribute]string, error) {
	var request = map[*Attribute]string{
		HTTPHost:          req.URL.Host,
		HTTPRequestMethod: req.Method,
		URLFull:           req.URL.String(),
		URLScheme:         req.URL.Scheme,
		UserAgent:         req.UserAgent(),
	}

	if storeId, ok := params["storeId"].(string); ok && storeId != "" {
		request[FGAClientRequestStoreID] = storeId
	}

	if authorizationModelId, ok := params["authorizationModelId"].(string); ok && authorizationModelId != "" {
		request[FGAClientRequestModelID] = authorizationModelId
	}

	if body, ok := params["body"]; ok {
		requestType := fmt.Sprintf("%T", body)

		switch requestType {
		case "*openfga.BatchCheckRequest":
			if req, ok := body.(interface{ GetChecks() []interface{} }); ok {
				checks := req.GetChecks()
				if len(checks) > 0 {
					request[FGAClientRequestBatchCheckSize] = fmt.Sprintf("%d", len(checks))
				}
			}
		case "*openfga.CheckRequest":
			if req, ok := body.(CheckRequestInterface); ok {
				if tupleKey := req.GetTupleKey(); tupleKey != nil {
					if user := tupleKey.GetUser(); user != nil {
						request[FGAClientUser] = *user
					}
				}

				if modelId := req.GetAuthorizationModelId(); modelId != nil {
					request[FGAClientRequestModelID] = *modelId
				}
			}
		case "*openfga.ExpandRequest":
		case "*openfga.ListObjectsRequest":
		case "*openfga.ListUsersRequest":
		case "*openfga.WriteRequest":
			if req, ok := body.(RequestAuthorizationModelIdInterface); ok {
				if modelId := req.GetAuthorizationModelId(); modelId != nil {
					request[FGAClientRequestModelID] = *modelId
				}
			}
		}
	}

	return request, nil
}

func (m *Metrics) AttributesFromResponse(res *http.Response, attrs map[*Attribute]string) (map[*Attribute]string, error) {
	attrs[HTTPResponseStatusCode] = strconv.Itoa(res.StatusCode)

	if res.Header.Get("openfga-authorization-model-id") != "" {
		attrs[FGAClientResponseModelID] = res.Header.Get("openfga-authorization-model-id")
	}

	if res.Header.Get("fga-query-duration-ms") != "" {
		attrs[HTTPServerRequestDuration] = res.Header.Get("fga-query-duration-ms")
	}

	return attrs, nil
}

func (m *Metrics) AttributesFromRequestDuration(requestStarted time.Time, attrs map[*Attribute]string) (float64, map[*Attribute]string, error) {
	requestDurationFloat := time.Since(requestStarted).Seconds() * 1000
	attrs[HTTPClientRequestDuration] = strconv.FormatFloat(requestDurationFloat, 'f', -1, 64)

	return requestDurationFloat, attrs, nil
}

func (m *Metrics) AttributesFromQueryDuration(attrs map[*Attribute]string) (float64, map[*Attribute]string, error) {
	if attrs[HTTPServerRequestDuration] == "" {
		return 0, attrs, nil
	}

	queryDurationFloat, queryDurationFloatErr := strconv.ParseFloat(attrs[HTTPServerRequestDuration], 64)

	if queryDurationFloatErr != nil {
		return 0, attrs, queryDurationFloatErr
	}

	return queryDurationFloat, attrs, nil
}

func (m *Metrics) AttributesFromResendCount(resendCount int, attrs map[*Attribute]string) (map[*Attribute]string, error) {
	if resendCount > 0 {
		attrs[HTTPRequestResendCount] = strconv.Itoa(resendCount)
	}

	return attrs, nil
}

func (m *Metrics) BuildTelemetryAttributes(requestMethod string, methodParameters map[string]interface{}, req *http.Request, res *http.Response, requestStarted time.Time, resendCount int) (map[*Attribute]string, float64, float64, error) {
	var attrs map[*Attribute]string

	attrs, _ = m.AttributesFromRequest(req, methodParameters)
	attrs, _ = m.AttributesFromResponse(res, attrs)
	attrs, _ = m.AttributesFromResendCount(resendCount, attrs)

	var requestDuration, queryDuration float64
	queryDuration, attrs, _ = m.AttributesFromQueryDuration(attrs)
	requestDuration, attrs, _ = m.AttributesFromRequestDuration(requestStarted, attrs)

	attrs[FGAClientRequestMethod] = requestMethod

	return attrs, queryDuration, requestDuration, nil
}
