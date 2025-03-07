package retryutils

import "testing"

func TestValidateMaxRetry(t *testing.T) {
	t.Run("MaxRetry within valid range", func(t *testing.T) {
		params := RetryParams{MaxRetry: 2, MinWaitInMs: 100}
		err := params.Validate()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("MaxRetry below valid range", func(t *testing.T) {
		params := RetryParams{MaxRetry: -1, MinWaitInMs: 100}
		err := params.Validate()
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
	})

	t.Run("MaxRetry above valid range", func(t *testing.T) {
		params := RetryParams{MaxRetry: 15 + 1, MinWaitInMs: 100}
		err := params.Validate()
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
	})
}

func TestValidateMinWaitInMs(t *testing.T) {
	t.Run("MinWaitInMs greater than 0", func(t *testing.T) {
		params := RetryParams{MaxRetry: 2, MinWaitInMs: 100}
		err := params.Validate()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("MinWaitInMs equal to 0", func(t *testing.T) {
		params := RetryParams{MaxRetry: 2, MinWaitInMs: 0}
		err := params.Validate()
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
	})

	t.Run("MinWaitInMs less than 0", func(t *testing.T) {
		params := RetryParams{MaxRetry: 2, MinWaitInMs: -1}
		err := params.Validate()
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
	})
}

func TestNewRetryParamsValidation(t *testing.T) {
	t.Run("Nil input returns default", func(t *testing.T) {
		got, err := NewRetryParams(nil)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if *got != defaultRetryParams {
			t.Fatalf("Expected %v, got %v", defaultRetryParams, got)
		}
	})

	t.Run("Valid input returns input", func(t *testing.T) {
		params := &RetryParams{MaxRetry: 2, MinWaitInMs: 100}
		got, err := NewRetryParams(params)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if *got != *params {
			t.Fatalf("Expected %v, got %v", params, got)
		}
	})

	t.Run("Invalid input returns default with error", func(t *testing.T) {
		params := &RetryParams{MaxRetry: -1, MinWaitInMs: 100}
		got, err := NewRetryParams(params)
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
		if *got != defaultRetryParams {
			t.Fatalf("Expected %v, got %v", defaultRetryParams, got)
		}
	})
}
