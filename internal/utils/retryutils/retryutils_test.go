package retryutils

import (
	"net/http"
	"testing"
	"time"
)

func TestParseRetryAfterHeaderValue(t *testing.T) {
	t.Run("Valid integer seconds", func(t *testing.T) {
		headers := http.Header{}
		headers.Set(RetryAfterHeaderName, "120")
		duration := ParseRetryAfterHeaderValue(headers, RetryAfterHeaderName)
		if duration != 120*time.Second {
			t.Fatalf("Expected 120 seconds, got %v", duration)
		}
	})

	t.Run("Valid date format", func(t *testing.T) {
		headers := http.Header{}
		futureTime := time.Now().Add(2 * time.Hour).UTC().Format(http.TimeFormat)
		headers.Set(RetryAfterHeaderName, futureTime)
		duration := ParseRetryAfterHeaderValue(headers, RetryAfterHeaderName)
		if duration <= 0 {
			t.Fatalf("Expected positive 2h duration, got %v", duration)
		}
	})

	t.Run("Invalid header value", func(t *testing.T) {
		headers := http.Header{}
		headers.Set(RetryAfterHeaderName, "invalid")
		duration := ParseRetryAfterHeaderValue(headers, RetryAfterHeaderName)
		if duration != 0 {
			t.Fatalf("Expected 0 duration, got %v", duration)
		}
	})

	t.Run("Empty header value", func(t *testing.T) {
		headers := http.Header{}
		duration := ParseRetryAfterHeaderValue(headers, RetryAfterHeaderName)
		if duration != 0 {
			t.Fatalf("Expected 0 duration, got %v", duration)
		}
	})
}

func TestParseRetryHeaderValue(t *testing.T) {
	t.Run("Retry-After header present", func(t *testing.T) {
		headers := http.Header{}
		headers.Set(RetryAfterHeaderName, "120")
		duration := parseRetryHeaderValue(headers)
		if duration != 120*time.Second {
			t.Fatalf("Expected 120 seconds, got %v", duration)
		}
	})

	t.Run("X-RateLimit-Reset header present", func(t *testing.T) {
		headers := http.Header{}
		headers.Set(RateLimitResetHeaderName, "120")
		duration := parseRetryHeaderValue(headers)
		if duration != 120*time.Second {
			t.Fatalf("Expected 120 seconds, got %v", duration)
		}
	})

	t.Run("X-Rate-Limit-Reset header present", func(t *testing.T) {
		headers := http.Header{}
		headers.Set(RateLimitResetAltHeaderName, "120")
		duration := parseRetryHeaderValue(headers)
		if duration != 120*time.Second {
			t.Fatalf("Expected 120 seconds, got %v", duration)
		}
	})

	t.Run("No retry headers present", func(t *testing.T) {
		headers := http.Header{}
		duration := parseRetryHeaderValue(headers)
		if duration != 0 {
			t.Fatalf("Expected 0 duration, got %v", duration)
		}
	})
}

func TestGetTimeToWait(t *testing.T) {
	t.Run("Exceed max retry count", func(t *testing.T) {
		duration := GetTimeToWait(5, 3, 100, http.Header{}, "")
		if duration != 0 {
			t.Fatalf("Expected 0 duration, got %v", duration)
		}
	})

	t.Run("Retry-After header present", func(t *testing.T) {
		headers := http.Header{}
		headers.Set(RetryAfterHeaderName, "120")
		duration := GetTimeToWait(1, 3, 100, headers, "")
		if duration != 120*time.Second {
			t.Fatalf("Expected 120 seconds, got %v", duration)
		}
	})

	t.Run("Random time within max backoff", func(t *testing.T) {
		duration := GetTimeToWait(1, 3, 100, http.Header{}, "Check")
		if duration <= 0 || duration > time.Duration(MaxBackoffTimeInSec)*time.Second {
			t.Fatalf("Expected duration between 0 and %v, got %v", time.Duration(MaxBackoffTimeInSec)*time.Second, duration)
		}
	})
}
