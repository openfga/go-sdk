package openfga

import "context"

type headersContextKey struct{}

// ContextWithHeaders returns a context carrying the provided headers.
func ContextWithHeaders(ctx context.Context, headers map[string]string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, headersContextKey{}, headers)
}

// HeadersFromContext extracts headers from the context if present.
func HeadersFromContext(ctx context.Context) (map[string]string, bool) {
	headers, ok := ctx.Value(headersContextKey{}).(map[string]string)
	return headers, ok
}
