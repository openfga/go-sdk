// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
)

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
		io.WriteString(w, `{"access_token": "ACCESS_TOKEN", "token_type": "bearer"}`)
	}))
	defer ts.Close()
	_, err := RetrieveToken(context.Background(), clientID, "", ts.URL, url.Values{}, AuthStyleInParams)
	if err != nil {
		t.Errorf("RetrieveToken = %v; want no error", err)
	}
}

func TestRetrieveTokenWithContexts(t *testing.T) {
	ResetAuthCache()
	const clientID = "client-id"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"access_token": "ACCESS_TOKEN", "token_type": "bearer"}`)
	}))
	defer ts.Close()

	_, err := RetrieveToken(context.Background(), clientID, "", ts.URL, url.Values{}, AuthStyleUnknown)
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
	_, err = RetrieveToken(ctx, clientID, "", cancellingts.URL, url.Values{}, AuthStyleUnknown)
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

	const testURL = "http://testserver.com"
	httpmock.RegisterResponder("POST", testURL,
		firstMock.Then(firstMock).Then(firstMock).Then(secondMock),
	)

	_, err := RetrieveToken(context.Background(), clientID, "", testURL, url.Values{}, AuthStyleUnknown)
	if err != nil {
		t.Errorf("RetrieveToken (with background context) = %v; want no error", err)
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

	const testURL = "http://testserver.com"
	httpmock.RegisterResponder("POST", testURL,
		firstMock,
	)

	_, err := RetrieveToken(context.Background(), clientID, "", testURL, url.Values{}, AuthStyleUnknown)
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

	const testURL = "http://testserver.com"
	httpmock.RegisterResponder("POST", testURL,
		firstMock,
	)

	// Set AuthStyleInHeader to avoid making a discovery request
	_, err := RetrieveToken(context.Background(), clientID, "", testURL, url.Values{}, AuthStyleInHeader)
	if err == nil {
		t.Errorf("Expect error to be returned when oauth server fails consistently")
	}

	log.Print(httpmock.GetTotalCallCount())

	if httpmock.GetTotalCallCount() != 1 {
		t.Errorf("Expected request to not be retried on a 401")
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

	const testURL = "http://testserver.com"
	httpmock.RegisterResponder("POST", testURL,
		firstMock.Then(firstMock).Then(firstMock).Then(secondMock),
	)

	_, err := RetrieveToken(context.Background(), clientID, "", testURL, url.Values{}, AuthStyleUnknown)
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

	const testURL = "http://testserver.com"
	httpmock.RegisterResponder("POST", testURL,
		firstMock,
	)

	// Set AuthStyleInHeader to avoid making a discovery request
	_, err := RetrieveToken(context.Background(), clientID, "", testURL, url.Values{}, AuthStyleInHeader)
	if err == nil {
		t.Errorf("Expect error to be returned when oauth server fails consistently")
	}

	log.Print(httpmock.GetTotalCallCount())

	if httpmock.GetTotalCallCount() != cMaxRetry {
		t.Errorf("Expected request to not be retried on a 401")
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
