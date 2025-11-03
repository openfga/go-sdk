package retryutils

import (
	_math "math"
	_rand "math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/openfga/go-sdk/internal/constants"
)

const (
	RetryAfterHeaderName        = constants.RetryAfterHeaderName
	RateLimitResetHeaderName    = constants.RateLimitResetHeaderName
	RateLimitResetAltHeaderName = constants.RateLimitResetAltHeaderName

	RetryHeaderMaxAllowableDurationInSec = constants.RetryHeaderMaxAllowableDurationInSec
	MaxBackoffTimeInSec                  = constants.MaxBackoffTimeInSec

	retryMaxAllowedNumber = constants.RetryMaxAllowedNumber
	defaultMaxRetry       = constants.DefaultMaxRetry
	defaultMinWaitInMs    = constants.DefaultMinWaitInMs
)

// randomTime provides a randomized time
func randomTime(loopCount int, minWaitInMs int) time.Duration {
	if minWaitInMs <= 0 {
		// This is protected against in the defaults, but we should still check in case parameters are passed without the NewRetryParams function
		minWaitInMs = defaultMinWaitInMs
	}

	minTimeToWait := int(_math.Pow(2, float64(loopCount))) * minWaitInMs
	maxTimeToWait := int(_math.Pow(2, float64(loopCount+1))) * minWaitInMs
	return time.Duration(_rand.Intn(maxTimeToWait-minTimeToWait+1)+minTimeToWait) * time.Millisecond
}

// ParseRetryAfterHeaderValue parses the Retry-After header value to time.Duration
func ParseRetryAfterHeaderValue(headers http.Header, headerName string) time.Duration {
	retryAfter := headers.Get(headerName)
	if retryAfter == "" {
		return 0
	}

	// Try to parse as an integer (seconds)
	if seconds, err := strconv.Atoi(retryAfter); err == nil {
		return time.Duration(seconds) * time.Second
	}

	// Try to parse as a date
	if date, err := http.ParseTime(retryAfter); err == nil {
		return time.Until(date)
	}

	return 0
}

// parseRetryHeaderValue parses several possible retry after header value to time.Duration
// starts with Retry-After, then X-RateLimit-Reset, then X-Rate-Limit-Reset
func parseRetryHeaderValue(headers http.Header) time.Duration {
	// if retryAfter is greater than 0 and less than the max backoff time, return retryAfter
	timeToWait := ParseRetryAfterHeaderValue(headers, RetryAfterHeaderName)
	if timeToWait > 0 && timeToWait < RetryHeaderMaxAllowableDurationInSec*time.Second {
		return timeToWait
	}

	// if X-Rate-Limit-Reset is greater than 0 and less than the max backoff time, return retryAfter
	timeToWait = ParseRetryAfterHeaderValue(headers, RateLimitResetHeaderName)
	if timeToWait > 0 && timeToWait < RetryHeaderMaxAllowableDurationInSec*time.Second {
		return timeToWait
	}

	// if X-RateLimit-Reset is greater than 0 and less than the max backoff time, return retryAfter
	timeToWait = ParseRetryAfterHeaderValue(headers, RateLimitResetAltHeaderName)
	if timeToWait > 0 && timeToWait < RetryHeaderMaxAllowableDurationInSec*time.Second {
		return timeToWait
	}

	return 0
}

// GetTimeToWait returns the time to wait for the next retry
// returns 0 if no retry should be attempted
// loopCount: the current loop count
// maxRetry: the maximum number of retries
// minWaitInMs: the minimum wait time in milliseconds
// headers: the headers from the response
// operationName: the operation name, currently unused
func GetTimeToWait(loopCount, maxRetry, minWaitInMs int, headers http.Header, _ string) time.Duration {
	// if the loop count is greater than the max retry, return 0
	if loopCount >= maxRetry {
		return 0
	}

	timeToWait := parseRetryHeaderValue(headers)
	// if timeToWait is greater than 0 that means it's valid and we can use it
	if timeToWait > 0 {
		return timeToWait
	}

	timeToWait = randomTime(loopCount, minWaitInMs)
	if timeToWait > 0 && timeToWait < time.Duration(MaxBackoffTimeInSec)*time.Second {
		return timeToWait
	}

	return time.Duration(MaxBackoffTimeInSec) * time.Second
}
