package httpclient /* import "s32x.com/httpclient" */

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/cenkalti/backoff/v4"
	"s32x.com/httpclient/cache"
)

// Request is a type used for configuring, performing and decoding HTTP
// requests
type Request struct {
	err            error
	client         *http.Client // DO NOT MODIFY THIS CLIENT
	cache          cache.Cacher
	method         string
	baseURL        string
	path           string
	headers        []header
	expectedStatus int // The statusCode that is expected for a success
	retryCount     int // Number of times to retry
	body           io.ReadWriter
	ctx            context.Context
}

// WithBody sets the body on the request with the passed io.ReadWriter
func (r *Request) WithBody(body io.ReadWriter) *Request {
	r.body = body
	return r
}

// WithBytes sets the passed bytes as the body to be used on the Request
func (r *Request) WithBytes(body []byte) *Request {
	return r.WithBody(bytes.NewBuffer(body))
}

// WithString sets the passed string as the body to be used on the Request
func (r *Request) WithString(body string) *Request {
	return r.WithBody(bytes.NewBufferString(body))
}

// WithForm encodes and sets the passed url.Values as the body to be used on
// the Request
func (r *Request) WithForm(data url.Values) *Request {
	return r.WithBody(bytes.NewBufferString(data.Encode())).
		WithContentType("application/x-www-form-urlencoded")
}

// WithJSON sets the JSON encoded passed interface as the body to be used on
// the Request
func (r *Request) WithJSON(body interface{}) *Request {
	r.body = bytes.NewBuffer(nil)
	r.err = json.NewEncoder(r.body).Encode(body)
	return r.WithContentType("application/json")
}

// WithXML sets the XML encoded passed interface as the body to be used on the
// Request
func (r *Request) WithXML(body interface{}) *Request {
	r.body = bytes.NewBuffer(nil)
	r.err = xml.NewEncoder(r.body).Encode(body)
	return r.WithContentType("application/xml")
}

// WithContext sets the context on the Request
func (r *Request) WithContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

// WithContentType sets the content-type that will be set in the headers on the
// Request
func (r *Request) WithContentType(typ string) *Request {
	return r.WithHeader("Content-Type", typ)
}

// WithHeader sets a header that will be used on the Request
func (r *Request) WithHeader(key, value string) *Request {
	r.headers = append(r.headers, header{key: key, value: value})
	return r
}

// WithExpectedStatus sets the desired status-code that will be a success. If
// the expected status code isn't received an error will be returned or the
// request will be retried if a count has been set with WithRetry(...)
func (r *Request) WithExpectedStatus(expectedStatusCode int) *Request {
	r.expectedStatus = expectedStatusCode
	return r
}

// WithRetry sets the desired number of retries on the Request
// Note: In order to trigger retries you must set an expected status code with
// the WithExpectedStatus(...) method
func (r *Request) WithRetry(retryCount int) *Request {
	r.retryCount = retryCount
	return r
}

// String is a convenience method that handles executing, defer closing, and
// decoding the body into a string before returning
func (r *Request) String() (string, error) {
	bytes, err := r.Bytes()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Bytes is a convenience method that handles executing, defer closing, and
// decoding the body into a slice of bytes before returning
func (r *Request) Bytes() ([]byte, error) {
	res, err := r.Do()
	if err != nil {
		return nil, err
	}
	if r.expectedStatus > 0 && res.Status != r.expectedStatus {
		return nil, fmt.Errorf("unexpected status code: %v", res.Status)
	}
	return res.Bytes(), nil
}

// JSON is a convenience method that handles executing, defer closing, and
// decoding the JSON body into the passed interface before returning
func (r *Request) JSON(out interface{}) error {
	res, err := r.Do()
	if err != nil {
		return err
	}
	if r.expectedStatus > 0 && res.Status != r.expectedStatus {
		return fmt.Errorf("unexpected status code: %v", res.Status)
	}
	return res.JSON(out)
}

// JSONWithError is identical to the JSON(...) method but also takes an errOut
// interface for when the status code isn't expected. In this case the response
// body will be decoded into the errOut interface and the boolean (expected)
// will return false
func (r *Request) JSONWithError(out interface{}, errOut interface{}) (bool, error) {
	res, err := r.Do()
	if err != nil {
		return false, err
	}
	if r.expectedStatus > 0 && res.Status != r.expectedStatus {
		return false, res.JSON(errOut)
	}
	return true, res.JSON(out)
}

// XML is a convenience method that handles executing, defer closing, and
// decoding the XML body into the passed interface before returning
func (r *Request) XML(out interface{}) error {
	res, err := r.Do()
	if err != nil {
		return err
	}
	if r.expectedStatus > 0 && res.Status != r.expectedStatus {
		return fmt.Errorf("unexpected status code: %v", res.Status)
	}
	return res.XML(out)
}

// XMLWithError is identical to the XML(...) method but also takes an errOut
// interface for when the status code isn't expected. In this case the response
// body will be decoded into the errOut interface and the boolean (expected)
// will return false
func (r *Request) XMLWithError(out interface{}, errOut interface{}) (bool, error) {
	res, err := r.Do()
	if err != nil {
		return false, err
	}
	if r.expectedStatus > 0 && res.Status != r.expectedStatus {
		return false, res.XML(errOut)
	}
	return true, res.XML(out)
}

// Error performs the request and returns any errors that result from the Do
func (r *Request) Error() error {
	res, err := r.Do()
	if err != nil {
		return err
	}
	if r.expectedStatus > 0 && res.Status != r.expectedStatus {
		return fmt.Errorf("unexpected status code: %v", res.Status)
	}
	return nil
}

// Do performs the base request and returns a populated Response
// NOTE: As with the standard library, when calling Do you must remember to
// close the response body : res.Body.Close()
func (r *Request) Do() (*Response, error) {
	if r.err != nil {
		return nil, r.err
	}

	// Check any existing cache for a cached response
	cacheKey := fmt.Sprintf("%s:%s%s", r.method, r.baseURL, r.path)
	if r.cache != nil && r.method == http.MethodGet {
		b, err := r.cache.Get(cacheKey)
		if err == nil && b != nil {
			var res Response
			if err := json.Unmarshal(b, &res); err != nil {
				return nil, fmt.Errorf("cache json unmarshal: %w", err)
			}
			return &res, nil
		}
		if !errors.Is(err, cache.ErrKeyNotFound) {
			return nil, fmt.Errorf("cache get (%s): %w", cacheKey, err)
		}
	}

	// Convert the Request to a standard http.Request
	req, err := r.toHTTPRequest()
	if err != nil {
		return nil, fmt.Errorf("to http request: %w", err)
	}

	// Perform the request with retries, returning the wrapped http.Response
	res, err := doRetry(r.client, req, r.expectedStatus, r.retryCount)
	if err != nil {
		return nil, fmt.Errorf("do retry: %w", err)
	}

	// Store the res in the cache if a cache is configured and it's a
	// successful response
	if r.cache != nil && r.method == http.MethodGet && res.Status < 300 {
		resBytes, err := json.Marshal(res)
		if err != nil {
			return nil, fmt.Errorf("json encode for cache: %w", err)
		}
		if err := r.cache.Set(cacheKey, resBytes); err != nil {
			return nil, fmt.Errorf("cache set: %w", err)
		}
	}
	return res, nil
}

// toHTTPRequest converts a Request to a standard HTTP Request. It assumes
// there is no error on the request.
func (r *Request) toHTTPRequest() (*http.Request, error) {
	// Generate a new http Request using client and passed Request
	req, err := http.NewRequest(r.method, r.baseURL+r.path, r.body)
	if err != nil {
		return nil, err
	}

	// Apply a context if one is set on the Request
	if r.ctx != nil {
		req = req.WithContext(r.ctx)
	}

	// Apply all headers from both the client and the Request
	for _, h := range r.headers {
		req.Header.Set(h.key, h.value)
	}
	return req, nil
}

// doRetry executes the passed http Request using the passed http Client and
// retries as many times as specified
func doRetry(c *http.Client, r *http.Request, expectedStatus, retryCount int) (*Response, error) {
	// Create a ticker that will execute the exponential backoff algorithm
	ticker := backoff.NewTicker(backoff.NewExponentialBackOff())

	// Continuously retry HTTP requests
	tries := 0
	for range ticker.C {
		// Increment the tries value to indicate which try num we're on
		tries++

		// Execute the request
		r, err := do(c, r, expectedStatus)
		if err == nil {
			return r, nil
		}

		// Stop the ticker and break out of the tick loop
		if retryCount <= tries {
			ticker.Stop()
			break
		}
	}
	return nil, fmt.Errorf("request failed after %v retries", retryCount)
}

// do attempts to execute the passed request on the passed client and returns
// the resulting bytes and a status code
func do(c *http.Client, r *http.Request, expectedStatus int) (*Response, error) {
	// Perform the request using the standard library
	res, err := c.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer res.Body.Close()

	// If the status code isn't what we expect, return an error
	if expectedStatus > 0 && expectedStatus != res.StatusCode {
		return nil, fmt.Errorf("unexpected status code: %s", res.Status)
	}
	return newResponse(res)
}
