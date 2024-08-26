package telemetry

type AttributeConfiguration struct {
	Enabled bool `json:"enabled,omitempty"`
}

type MetricConfiguration struct {
	ATTR_FGA_CLIENT_REQUEST_CLIENT_ID *AttributeConfiguration `json:"fga_client_request_client_id,omitempty"`
	ATTR_FGA_CLIENT_REQUEST_METHOD    *AttributeConfiguration `json:"fga_client_request_method,omitempty"`
	ATTR_FGA_CLIENT_REQUEST_MODEL_ID  *AttributeConfiguration `json:"fga_client_request_model_id,omitempty"`
	ATTR_FGA_CLIENT_REQUEST_STORE_ID  *AttributeConfiguration `json:"fga_client_request_store_id,omitempty"`
	ATTR_FGA_CLIENT_RESPONSE_MODEL_ID *AttributeConfiguration `json:"fga_client_response_model_id,omitempty"`
	ATTR_FGA_CLIENT_USER              *AttributeConfiguration `json:"fga_client_user,omitempty"`
	ATTR_HTTP_CLIENT_REQUEST_DURATION *AttributeConfiguration `json:"http_client_request_duration,omitempty"`
	ATTR_HTTP_HOST                    *AttributeConfiguration `json:"http_host,omitempty"`
	ATTR_HTTP_REQUEST_METHOD          *AttributeConfiguration `json:"http_request_method,omitempty"`
	ATTR_HTTP_REQUEST_RESEND_COUNT    *AttributeConfiguration `json:"http_request_resend_count,omitempty"`
	ATTR_HTTP_RESPONSE_STATUS_CODE    *AttributeConfiguration `json:"http_response_status_code,omitempty"`
	ATTR_HTTP_SERVER_REQUEST_DURATION *AttributeConfiguration `json:"http_server_request_duration,omitempty"`
	ATTR_URL_SCHEME                   *AttributeConfiguration `json:"url_scheme,omitempty"`
	ATTR_URL_FULL                     *AttributeConfiguration `json:"url_full,omitempty"`
	ATTR_USER_AGENT_ORIGINAL          *AttributeConfiguration `json:"user_agent_original,omitempty"`
}

type MetricsConfiguration struct {
	METRIC_COUNTER_CREDENTIALS_REQUEST *MetricConfiguration `json:"fga_client_credentials_request,omitempty"`
	METRIC_HISTOGRAM_REQUEST_DURATION  *MetricConfiguration `json:"fga_client_request_duration,omitempty"`
	METRIC_HISTOGRAM_QUERY_DURATION    *MetricConfiguration `json:"fga_client_query_duration,omitempty"`
}

type Configuration struct {
	Metrics *MetricsConfiguration `json:"metrics,omitempty"`
}
