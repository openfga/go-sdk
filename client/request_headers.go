package client

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const (
	clientMethodHeader        = "X-OpenFGA-Client-Method"
	clientBulkRequestIDHeader = "X-OpenFGA-Client-Bulk-Request-Id"
)

func addClientMethodHeader(options *RequestOptions, method string) {
	if options.Headers == nil {
		options.Headers = make(map[string]string)
	} else {
		options.Headers = cloneHeaders(options.Headers)
	}
	options.Headers[clientMethodHeader] = method
}

func addBulkRequestIDHeader(options *RequestOptions) error {
	if options.Headers == nil {
		options.Headers = make(map[string]string)
	} else {
		options.Headers = cloneHeaders(options.Headers)
	}
	if options.Headers[clientBulkRequestIDHeader] != "" {
		return nil
	}
	id, err := newBulkRequestID()
	if err != nil {
		return err
	}
	options.Headers[clientBulkRequestIDHeader] = id
	return nil
}

func cloneHeaders(headers map[string]string) map[string]string {
	cloned := make(map[string]string, len(headers))
	for key, value := range headers {
		cloned[key] = value
	}
	return cloned
}

func newBulkRequestID() (string, error) {
	var id [16]byte
	if _, err := rand.Read(id[:]); err != nil {
		return "", fmt.Errorf("failed to generate bulk request ID: %w", err)
	}
	id[6] = (id[6] & 0x0f) | 0x40
	id[8] = (id[8] & 0x3f) | 0x80
	return hex.EncodeToString(id[0:4]) + "-" +
		hex.EncodeToString(id[4:6]) + "-" +
		hex.EncodeToString(id[6:8]) + "-" +
		hex.EncodeToString(id[8:10]) + "-" +
		hex.EncodeToString(id[10:16]), nil
}
