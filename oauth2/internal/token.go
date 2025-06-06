// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"mime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/openfga/go-sdk/internal/utils/retryutils"
	"github.com/openfga/go-sdk/telemetry"
)

type RequestConfig struct {
	RetryParams retryutils.RetryParams

	Debug bool
}

func (rc *RequestConfig) New(config RequestConfig) error {
	retryParams, err := retryutils.NewRetryParams(&config.RetryParams)
	if err != nil {
		return err
	}

	rc.RetryParams = *retryParams
	rc.Debug = config.Debug

	return nil
}

// Token represents the credentials used to authorize
// the requests to access protected resources on the OAuth 2.0
// provider's backend.
//
// This type is a mirror of oauth2.Token and exists to break
// an otherwise-circular dependency. Other internal packages
// should convert this Token into an oauth2.Token before use.
type Token struct {
	// AccessToken is the token that authorizes and authenticates
	// the requests.
	AccessToken string

	// TokenType is the type of token.
	// The Type method returns either this or "Bearer", the default.
	TokenType string

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string

	// Expiry is the optional expiration time of the access token.
	//
	// If zero, TokenSource implementations will reuse the same
	// token forever and RefreshToken or equivalent
	// mechanisms for that TokenSource will not be used.
	Expiry time.Time

	// Raw optionally contains extra metadata from the server
	// when updating a token.
	Raw interface{}
}

// tokenJSON is the struct representing the HTTP response from OAuth2
// providers returning a token in JSON form.
type tokenJSON struct {
	AccessToken  string         `json:"access_token"`
	TokenType    string         `json:"token_type"`
	RefreshToken string         `json:"refresh_token"`
	ExpiresIn    expirationTime `json:"expires_in"` // at least PayPal returns string, while most return number
}

func (e *tokenJSON) expiry() (t time.Time) {
	if v := e.ExpiresIn; v != 0 {
		return time.Now().Add(time.Duration(v) * time.Second)
	}
	return
}

type expirationTime int32

func (e *expirationTime) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || string(b) == "null" {
		return nil
	}
	var n json.Number
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}
	i, err := n.Int64()
	if err != nil {
		return err
	}
	if i > math.MaxInt32 {
		i = math.MaxInt32
	}
	*e = expirationTime(i)
	return nil
}

// RegisterBrokenAuthHeaderProvider previously did something. It is now a no-op.
//
// Deprecated: this function no longer does anything. Caller code that
// wants to avoid potential extra HTTP requests made during
// auto-probing of the provider's auth style should set
// Endpoint.AuthStyle.
func RegisterBrokenAuthHeaderProvider(tokenURL string) {}

// AuthStyle is a copy of the golang.org/x/oauth2 package's AuthStyle type.
type AuthStyle int

const (
	AuthStyleUnknown  AuthStyle = 0
	AuthStyleInParams AuthStyle = 1
	AuthStyleInHeader AuthStyle = 2
)

// authStyleCache is the set of tokenURLs we've successfully used via
// RetrieveToken and which style auth we ended up using.
// It's called a cache, but it doesn't (yet?) shrink. It's expected that
// the set of OAuth2 servers a program contacts over time is fixed and
// small.
var authStyleCache struct {
	sync.Mutex
	m map[string]AuthStyle // keyed by tokenURL
}

// ResetAuthCache resets the global authentication style cache used
// for AuthStyleUnknown token requests.
func ResetAuthCache() {
	authStyleCache.Lock()
	defer authStyleCache.Unlock()
	authStyleCache.m = nil
}

// lookupAuthStyle reports which auth style we last used with tokenURL
// when calling RetrieveToken and whether we have ever done so.
func lookupAuthStyle(tokenURL string) (style AuthStyle, ok bool) {
	authStyleCache.Lock()
	defer authStyleCache.Unlock()
	style, ok = authStyleCache.m[tokenURL]
	return
}

// setAuthStyle adds an entry to authStyleCache, documented above.
func setAuthStyle(tokenURL string, v AuthStyle) {
	authStyleCache.Lock()
	defer authStyleCache.Unlock()
	if authStyleCache.m == nil {
		authStyleCache.m = make(map[string]AuthStyle)
	}
	authStyleCache.m[tokenURL] = v
}

// newTokenRequest returns a new *http.Request to retrieve a new token
// from tokenURL using the provided clientID, clientSecret, and POST
// body parameters.
//
// inParams is whether the clientID & clientSecret should be encoded
// as the POST body. An 'inParams' value of true means to send it in
// the POST body (along with any values in v); false means to send it
// in the Authorization header.
func newTokenRequest(tokenURL, clientID, clientSecret string, v url.Values, authStyle AuthStyle) (*http.Request, error) {
	if authStyle == AuthStyleInParams {
		v = cloneURLValues(v)
		if clientID != "" {
			v.Set("client_id", clientID)
		}
		if clientSecret != "" {
			v.Set("client_secret", clientSecret)
		}
	}
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if authStyle == AuthStyleInHeader {
		req.SetBasicAuth(url.QueryEscape(clientID), url.QueryEscape(clientSecret))
	}
	return req, nil
}

func cloneURLValues(v url.Values) url.Values {
	v2 := make(url.Values, len(v))
	for k, vv := range v {
		v2[k] = append([]string(nil), vv...)
	}
	return v2
}

func RetrieveToken(ctx context.Context, clientID, clientSecret, tokenURL string, v url.Values, authStyle AuthStyle, config RequestConfig) (*Token, error) {
	needsAuthStyleProbe := authStyle == 0
	if needsAuthStyleProbe {
		if style, ok := lookupAuthStyle(tokenURL); ok {
			authStyle = style
			needsAuthStyleProbe = false
		} else {
			authStyle = AuthStyleInHeader // the first way we'll try
		}
	}
	req, err := newTokenRequest(tokenURL, clientID, clientSecret, v, authStyle)
	if err != nil {
		return nil, err
	}
	token, err := doTokenRoundTrip(ctx, req, config)
	if err != nil && needsAuthStyleProbe {
		// If we get an error, assume the server wants the
		// clientID & clientSecret in a different form.
		// See https://code.google.com/p/goauth2/issues/detail?id=31 for background.
		// In summary:
		// - Reddit only accepts client secret in the Authorization header
		// - Dropbox accepts either it in URL param or Auth header, but not both.
		// - Google only accepts URL param (not spec compliant?), not Auth header
		// - Stripe only accepts client secret in Auth header with Bearer method, not Basic
		//
		// We used to maintain a big table in this code of all the sites and which way
		// they went, but maintaining it didn't scale & got annoying.
		// So just try both ways.
		authStyle = AuthStyleInParams // the second way we'll try
		req, _ = newTokenRequest(tokenURL, clientID, clientSecret, v, authStyle)
		token, err = doTokenRoundTrip(ctx, req, config)
	}
	if needsAuthStyleProbe && err == nil {
		setAuthStyle(tokenURL, authStyle)
	}
	// Don't overwrite `RefreshToken` with an empty value
	// if this was a token refreshing request.
	if token != nil && token.RefreshToken == "" {
		token.RefreshToken = v.Get("refresh_token")
	}
	return token, err
}

func singleTokenRoundTrip(ctx context.Context, req *http.Request) (*Token, error) {
	r, err := ContextClient(ctx).Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
	_ = r.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
	}
	if code := r.StatusCode; code < 200 || code > 299 {
		return nil, &RetrieveError{
			Response: r,
			Body:     body,
		}
	}

	var token *Token
	content, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	switch content {
	case "application/x-www-form-urlencoded", "text/plain":
		vals, err := url.ParseQuery(string(body))
		if err != nil {
			return nil, err
		}
		token = &Token{
			AccessToken:  vals.Get("access_token"),
			TokenType:    vals.Get("token_type"),
			RefreshToken: vals.Get("refresh_token"),
			Raw:          vals,
		}
		e := vals.Get("expires_in")
		expires, _ := strconv.Atoi(e)
		if expires != 0 {
			token.Expiry = time.Now().Add(time.Duration(expires) * time.Second)
		}
	default:
		var tj tokenJSON
		if err = json.Unmarshal(body, &tj); err != nil {
			return nil, err
		}
		token = &Token{
			AccessToken:  tj.AccessToken,
			TokenType:    tj.TokenType,
			RefreshToken: tj.RefreshToken,
			Expiry:       tj.expiry(),
			Raw:          make(map[string]interface{}),
		}
		_ = json.Unmarshal(body, &token.Raw) // no error checks for optional fields
	}
	if token.AccessToken == "" {
		return nil, errors.New("oauth2: server response missing access_token")
	}
	return token, nil
}

func doTokenRoundTrip(ctx context.Context, req *http.Request, config RequestConfig) (*Token, error) {
	const operationName = "TokenExchange"
	var token *Token
	var err error

	maxRetry := config.RetryParams.MaxRetry
	minWaitInMs := config.RetryParams.MinWaitInMs
	debug := config.Debug

	for i := 0; i <= maxRetry; i++ {
		token, err = singleTokenRoundTrip(ctx, req)
		if err == nil {
			if otel := telemetry.Extract(ctx); otel != nil {
				attrs := make(map[*telemetry.Attribute]string)
				attrs[telemetry.HTTPRequestResendCount] = strconv.Itoa(i)
				_, _ = otel.Metrics.CredentialsRequest(1, attrs)
			}
			return token, err
		}

		// If this was a request error, then check if it was a 429 or 5xx error and retry.
		// if this is a network error, then retry.
		// We do not want to retry any other error
		shouldRetry := true
		responseHeaders := http.Header{}
		var rErr *RetrieveError
		errorStyle := "network"
		if errors.As(err, &rErr) {
			statusCode := rErr.Response.StatusCode
			if statusCode == http.StatusTooManyRequests || (statusCode >= http.StatusInternalServerError && statusCode != http.StatusNotImplemented) {
				shouldRetry = true
				responseHeaders = rErr.Response.Header
				errorStyle = "api"
			} else {
				shouldRetry = false
			}
		}

		if shouldRetry {
			timeToWait := retryutils.GetTimeToWait(i, maxRetry, minWaitInMs, responseHeaders, operationName)
			if timeToWait > 0 {
				if debug {
					log.Printf("\nWaiting %v to retry %v (%v %v) due to %s error (error=%v) on attempt %v\n", timeToWait, operationName, req.Method, req.URL, errorStyle, err, i)
				}

				time.Sleep(timeToWait)
				continue
			}
		}

		break
	}

	return nil, err
}

type RetrieveError struct {
	Response *http.Response
	Body     []byte
}

func (r *RetrieveError) Error() string {
	return fmt.Sprintf("oauth2: cannot fetch token: %v\nResponse: %s", r.Response.Status, r.Body)
}
