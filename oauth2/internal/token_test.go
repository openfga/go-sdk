// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"

	"github.com/openfga/go-sdk/internal/utils/retryutils"
)

const (
	testMaxRetry    = 2
	testMinWaitInMs = 20
	testURL         = "https://auth.fga.example"
)

var testTokenRequestConfig = RequestConfig{
	RetryParams: retryutils.RetryParams{
		MaxRetry:    testMaxRetry,
		MinWaitInMs: testMinWaitInMs,
	},
	Debug: false,
}

func TestRetrieveToken_InParams(t *testing.T) {
	ResetAuthCache()
	const clientID = "client-id"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.FormValue("client_id"), clientID; got != want {
			t.Errorf("client_id = %q; want %q", got, want)
		}
		if got, want := r.FormValue("client_secret"), ""; got != want {
			t.Errorf("client_secret = %q; want empty", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"access_token": "ACCESS_TOKEN", "token_type": "bearer"}`)
	}))
	defer ts.Close()
	_, err := RetrieveToken(context.Background(), clientID, "", ts.URL, url.Values{}, AuthStyleInParams, RequestConfig{
		RetryParams: retryutils.RetryParams{
			MaxRetry:    1,
			MinWaitInMs: testMinWaitInMs,
		},
	})
	if err != nil {
		t.Errorf("RetrieveToken = %v; want no error", err)
	}
}

func TestRetrieveTokenWithContexts(t *testing.T) {
	ResetAuthCache()
	const clientID = "client-id"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"access_token": "ACCESS_TOKEN", "token_type": "bearer"}`)
	}))
	defer ts.Close()

	_, err := RetrieveToken(context.Background(), clientID, "", ts.URL, url.Values{}, AuthStyleInParams, testTokenRequestConfig)
	if err != nil {
		t.Errorf("RetrieveToken (with background context) = %v; want no error", err)
	}

	retrieved := make(chan struct{})
	cancellingts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-retrieved
	}))
	defer cancellingts.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = RetrieveToken(ctx, clientID, "", cancellingts.URL, url.Values{}, AuthStyleInParams, testTokenRequestConfig)
	close(retrieved)
	if err == nil {
		t.Errorf("RetrieveToken (with cancelled context) = nil; want error")
	}
}

// Test the retry logic where request is successful at the end
func TestRetrieveTokenWithContextsRetry(t *testing.T) {
	ResetAuthCache()
	const clientID = "client-id"

	type JSONResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}
	expectedResponse := JSONResponse{
		AccessToken: "ACCESS_TOKEN",
		TokenType:   "bearer",
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	firstMock := httpmock.NewStringResponder(429, "")
	secondMock, _ := httpmock.NewJsonResponder(200, expectedResponse)

	httpmock.RegisterResponder("POST", testURL,
		firstMock.Then(firstMock).Then(firstMock).Then(secondMock),
	)

	_, err := RetrieveToken(context.Background(), clientID, "", testURL, url.Values{}, AuthStyleInParams, RequestConfig{
		RetryParams: retryutils.RetryParams{
			MaxRetry:    3,
			MinWaitInMs: testMinWaitInMs,
		},
	})
	if err != nil {
		t.Errorf("RetrieveToken (with background context) = %v; want no error", err)
	}
}

// Test the retry logic where request is successful at the end
func TestRetrieveTokenWithContextsRetryMaxExceeded(t *testing.T) {
	ResetAuthCache()
	const clientID = "client-id"

	type JSONResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}
	expectedResponse := JSONResponse{
		AccessToken: "ACCESS_TOKEN",
		TokenType:   "bearer",
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	firstMock := httpmock.NewStringResponder(429, "")
	secondMock, _ := httpmock.NewJsonResponder(200, expectedResponse)

	httpmock.RegisterResponder("POST", testURL,
		firstMock.Then(firstMock).Then(firstMock).Then(secondMock),
	)

	_, err := RetrieveToken(context.Background(), clientID, "", testURL, url.Values{}, AuthStyleInParams, RequestConfig{
		RetryParams: retryutils.RetryParams{
			MaxRetry:    2,
			MinWaitInMs: testMinWaitInMs,
		},
	})
	if err == nil {
		t.Errorf("RetrieveToken (with background context); expected error after exceeding max retries, got none")
	}
}

func TestRetrieveTokenWithContextsFailure(t *testing.T) {
	ResetAuthCache()
	const clientID = "client-id"

	type JSONResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	firstMock := httpmock.NewStringResponder(429, "")

	httpmock.RegisterResponder("POST", testURL,
		firstMock,
	)

	_, err := RetrieveToken(context.Background(), clientID, "", testURL, url.Values{}, AuthStyleInParams, testTokenRequestConfig)
	if err == nil {
		t.Errorf("Expect error to be returned when oauth server fails consistently")
	}
}

// Test that a 401 returns an error straight away
func TestRetrieveTokenWithUnauthorizedErrors(t *testing.T) {
	ResetAuthCache()
	const clientID = "client-id"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	firstMock := httpmock.NewStringResponder(401, "")

	httpmock.RegisterResponder("POST", testURL,
		firstMock,
	)

	// Set AuthStyleInParams to avoid making a discovery request
	_, err := RetrieveToken(context.Background(), clientID, "", testURL, url.Values{}, AuthStyleInParams, testTokenRequestConfig)
	if err == nil {
		t.Errorf("Expect error to be returned when oauth server fails consistently")
	}

	if httpmock.GetTotalCallCount() != 1 {
		t.Errorf("Expected request to be called once and not be retried on a 401, it was called %v times", httpmock.GetTotalCallCount())
	}
}

// Test that a 500 retries
func TestRetrieveTokenWithServerErrorRetries(t *testing.T) {
	ResetAuthCache()
	const clientID = "client-id"

	type JSONResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}
	expectedResponse := JSONResponse{
		AccessToken: "ACCESS_TOKEN",
		TokenType:   "bearer",
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	firstMock := httpmock.NewStringResponder(500, "")
	secondMock, _ := httpmock.NewJsonResponder(200, expectedResponse)

	httpmock.RegisterResponder("POST", testURL,
		firstMock.Then(firstMock).Then(secondMock),
	)

	// Set AuthStyleInParams to avoid making a discovery request
	_, err := RetrieveToken(context.Background(), clientID, "", testURL, url.Values{}, AuthStyleInParams, testTokenRequestConfig)
	if err != nil {
		t.Errorf("RetrieveToken (with background context) = %v; want no error", err)
	}
}

// Test that constant 500s errors
func TestRetrieveTokenWithServerErrorEventuallyErrors(t *testing.T) {
	ResetAuthCache()
	const clientID = "client-id"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	firstMock := httpmock.NewStringResponder(500, "")

	httpmock.RegisterResponder("POST", testURL,
		firstMock,
	)

	// Set AuthStyleInParams to avoid making a discovery request
	_, err := RetrieveToken(context.Background(), clientID, "", testURL, url.Values{}, AuthStyleInParams, testTokenRequestConfig)
	if err == nil {
		t.Errorf("Expect error to be returned when oauth server fails consistently")
	}

	// Expect the request to be called once and then retried testMaxRetry times
	if httpmock.GetTotalCallCount() != testMaxRetry+1 {
		t.Errorf("Expected request to be retried %v times on a 500, it was retried %v", testMaxRetry, httpmock.GetTotalCallCount()-1)
	}
}

func TestExpiresInUpperBound(t *testing.T) {
	var e expirationTime
	if err := e.UnmarshalJSON([]byte(fmt.Sprint(int64(math.MaxInt32) + 1))); err != nil {
		t.Fatal(err)
	}
	const want = math.MaxInt32
	if e != want {
		t.Errorf("expiration time = %v; want %v", e, want)
	}
}
