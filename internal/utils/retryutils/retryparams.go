package retryutils

import "fmt"

// RetryParams configures configuration for retry in case of HTTP too many request
type RetryParams struct {
	MaxRetry    int `json:"maxRetry,omitempty"`
	MinWaitInMs int `json:"minWaitInMs,omitempty"`
}

var defaultRetryParams = RetryParams{
	MaxRetry:    defaultMaxRetry,
	MinWaitInMs: defaultMinWaitInMs,
}

func (r *RetryParams) Validate() error {
	if r.MaxRetry < 0 || r.MaxRetry > retryMaxAllowedNumber {
		return fmt.Errorf("maxRetry must be between 0 and %d", retryMaxAllowedNumber)

	}

	if r.MinWaitInMs <= 0 {
		return fmt.Errorf("maxRetry must be greater than 0")
	}

	return nil
}

func GetRetryParamsOrDefault(r *RetryParams) RetryParams {
	if r == nil {
		return defaultRetryParams
	}

	return *r
}

func NewRetryParams(retryParams *RetryParams) (*RetryParams, error) {
	if retryParams == nil {
		return &defaultRetryParams, nil
	}

	if err := retryParams.Validate(); err != nil {
		return &defaultRetryParams, err
	}

	return retryParams, nil
}
