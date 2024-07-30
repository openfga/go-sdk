package telemetry

import (
	"net/http"
	"strconv"

	"go.opentelemetry.io/otel/attribute"
)

var (
	FGAClientRequestClientID  = &Attribute{Name: "fga-client.request.client_id"}
	FGAClientRequestMethod    = &Attribute{Name: "fga-client.request.method"}
	FGAClientRequestModelID   = &Attribute{Name: "fga-client.request.model_id"}
	FGAClientRequestStoreID   = &Attribute{Name: "fga-client.request.store_id"}
	FGAClientResponseModelID  = &Attribute{Name: "fga-client.response.model_id"}
	FGAClientUser             = &Attribute{Name: "fga-client.user"}
	HTTPClientRequestDuration = &Attribute{Name: "http.client.request.duration"}
	HTTPHost                  = &Attribute{Name: "http.host"}
	HTTPRequestMethod         = &Attribute{Name: "http.request.method"}
	HTTPRequestResendCount    = &Attribute{Name: "http.request.resend_count"}
	HTTPResponseStatusCode    = &Attribute{Name: "http.response.status_code"}
	HTTPServerRequestDuration = &Attribute{Name: "http.server.request.duration"}
	URLScheme                 = &Attribute{Name: "url.scheme"}
	URLFull                   = &Attribute{Name: "url.full"}
	UserAgent                 = &Attribute{Name: "user_agent.original"}
)

func (m *Metrics) PrepareAttributes(attrs map[*Attribute]string) attribute.Set {
	var prepared []attribute.KeyValue

	for attr, value := range attrs {
		prepared = append(prepared, attribute.String(attr.Name, value))
	}

	return attribute.NewSet(prepared...)
}

func (m *Metrics) AttributesFromRequest(req *http.Request) map[*Attribute]string {
	return map[*Attribute]string{
		HTTPHost:          req.URL.Host,
		HTTPRequestMethod: req.Method,
		URLFull:           req.URL.String(),
		URLScheme:         req.URL.Scheme,
		UserAgent:         req.UserAgent(),
	}
}

func (m *Metrics) AttributesFromResponse(res *http.Response, attrs map[*Attribute]string) map[*Attribute]string {
	attrs[HTTPResponseStatusCode] = strconv.Itoa(res.StatusCode)

	if res.Header.Get("openfga-authorization-model-id") != "" {
		attrs[FGAClientResponseModelID] = res.Header.Get("openfga-authorization-model-id")
	}

	if res.Header.Get("fga-query-duration-ms") != "" {
		attrs[HTTPServerRequestDuration] = res.Header.Get("fga-query-duration-ms")
	}

	return attrs
}
