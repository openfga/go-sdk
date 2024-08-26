package telemetry

import (
	"net/http"
	"strconv"

	openfga "github.com/openfga/go-sdk"
	"go.opentelemetry.io/otel/attribute"
)

const (
	ATTR_FGA_CLIENT_REQUEST_CLIENT_ID = "fga-client.request.client_id"
	ATTR_FGA_CLIENT_REQUEST_METHOD    = "fga-client.request.method"
	ATTR_FGA_CLIENT_REQUEST_MODEL_ID  = "fga-client.request.model_id"
	ATTR_FGA_CLIENT_REQUEST_STORE_ID  = "fga-client.request.store_id"
	ATTR_FGA_CLIENT_RESPONSE_MODEL_ID = "fga-client.response.model_id"
	ATTR_FGA_CLIENT_USER              = "fga-client.user"
	ATTR_HTTP_CLIENT_REQUEST_DURATION = "http.client.request.duration"
	ATTR_HTTP_HOST                    = "http.host"
	ATTR_HTTP_REQUEST_METHOD          = "http.request.method"
	ATTR_HTTP_REQUEST_RESEND_COUNT    = "http.request.resend_count"
	ATTR_HTTP_RESPONSE_STATUS_CODE    = "http.response.status_code"
	ATTR_HTTP_SERVER_REQUEST_DURATION = "http.server.request.duration"
	ATTR_URL_SCHEME                   = "url.scheme"
	ATTR_URL_FULL                     = "url.full"
	ATTR_USER_AGENT_ORIGINAL          = "user_agent.original"
)

var (
	FGAClientRequestClientID  = &Attribute{Name: ATTR_FGA_CLIENT_REQUEST_CLIENT_ID}
	FGAClientRequestMethod    = &Attribute{Name: ATTR_FGA_CLIENT_REQUEST_METHOD}
	FGAClientRequestModelID   = &Attribute{Name: ATTR_FGA_CLIENT_REQUEST_MODEL_ID}
	FGAClientRequestStoreID   = &Attribute{Name: ATTR_FGA_CLIENT_REQUEST_STORE_ID}
	FGAClientResponseModelID  = &Attribute{Name: ATTR_FGA_CLIENT_RESPONSE_MODEL_ID}
	FGAClientUser             = &Attribute{Name: ATTR_FGA_CLIENT_USER}
	HTTPClientRequestDuration = &Attribute{Name: ATTR_HTTP_CLIENT_REQUEST_DURATION}
	HTTPHost                  = &Attribute{Name: ATTR_HTTP_HOST}
	HTTPRequestMethod         = &Attribute{Name: ATTR_HTTP_REQUEST_METHOD}
	HTTPRequestResendCount    = &Attribute{Name: ATTR_HTTP_REQUEST_RESEND_COUNT}
	HTTPResponseStatusCode    = &Attribute{Name: ATTR_HTTP_RESPONSE_STATUS_CODE}
	HTTPServerRequestDuration = &Attribute{Name: ATTR_HTTP_SERVER_REQUEST_DURATION}
	URLScheme                 = &Attribute{Name: ATTR_URL_SCHEME}
	URLFull                   = &Attribute{Name: ATTR_URL_FULL}
	UserAgent                 = &Attribute{Name: ATTR_USER_AGENT_ORIGINAL}
)

func (m *Metrics) PrepareAttributes(attrs map[*Attribute]string) (attribute.Set, error) {
	var prepared []attribute.KeyValue

	for attr, value := range attrs {
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

	var attrs = make(map[*Attribute]string)

	if storeId, ok := params["storeId"].(string); ok && storeId != "" {
		attrs[FGAClientRequestStoreID] = storeId
	}

	if authorizationModelId, ok := params["authorizationModelId"].(string); ok && authorizationModelId != "" {
		attrs[FGAClientRequestModelID] = authorizationModelId
	}

	if body, ok := params["body"]; ok {
		switch req := body.(type) {
		case *openfga.CheckRequest:
			if req.TupleKey.User != "" {
				attrs[FGAClientUser] = req.TupleKey.User
			}

			if req.AuthorizationModelId != nil && *req.AuthorizationModelId != "" {
				attrs[FGAClientRequestModelID] = *req.AuthorizationModelId
			}
		case *openfga.ExpandRequest:
		case *openfga.ListObjectsRequest:
		case *openfga.ListUsersRequest:
		case *openfga.WriteRequest:
			if modelId := req.AuthorizationModelId; modelId != nil && *modelId != "" {
				attrs[FGAClientRequestModelID] = *modelId
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
