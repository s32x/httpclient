package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"sync"
)

// Request is a type used for configuring, performing and decoding HTTP
// requests
type Request struct {
	err            error
	client         *http.Client // DO NOT MODIFY THIS CLIENT
	method         string
	baseURL        string
	path           string
	header         sync.Map
	expectedStatus int // The statusCode that is expected for a success
	retryCount     int // Number of times to retry
	body           io.ReadWriter
	ctx            context.Context
}

// Error returns the error stored on the Request
func (r *Request) Error() error { return r.err }

// WithMethod sets the passed method as the method on the Request
func (r *Request) WithMethod(method string) *Request {
	r.method = method
	return r
}

// WithPath sets the passed path as the path on the Request
func (r *Request) WithPath(path string) *Request {
	r.path = path
	return r
}

// WithBytes sets the passed bytes as the body to be used on the Request
func (r *Request) WithBytes(body []byte) *Request {
	r.body = bytes.NewBuffer(body)
	return r
}

// WithString sets the passed string as the body to be used on the Request
func (r *Request) WithString(body string) *Request {
	r.body = bytes.NewBufferString(body)
	return r
}

// WithJSON sets the JSON encoded passed interface as the body to be used on
// the Request
func (r *Request) WithJSON(body interface{}) *Request {
	r = r.WithContentType("application/json")
	r.body = bytes.NewBuffer(nil)
	r.err = json.NewEncoder(r.body).Encode(body)
	return r
}

// WithXML sets the XML encoded passed interface as the body to be used on the
// Request
func (r *Request) WithXML(body interface{}) *Request {
	r = r.WithContentType("application/xml")
	r.body = bytes.NewBuffer(nil)
	r.err = xml.NewEncoder(r.body).Encode(body)
	return r
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
	r.header.Store(key, value)
	return r
}

// WithExpectedStatus sets the desired status-code that will be a success
func (r *Request) WithExpectedStatus(expectedStatusCode int) *Request {
	r.expectedStatus = expectedStatusCode
	return r
}

// WithRetry sets the desired number of retries on the Request
func (r *Request) WithRetry(retryCount int) *Request {
	r.retryCount = retryCount
	return r
}

// Do performs the passed request and returns a populated Response
func (r *Request) Do() (*Response, error) {
	if r.err != nil {
		return nil, r.err
	}

	// Convert the Request to a standard http Request
	req, err := r.toHTTPRequest()
	if err != nil {
		return nil, err
	}

	// Perform the request and return the wrapped Response
	res, err := doRetry(r.client, req, r.expectedStatus, r.retryCount)
	if err != nil {
		return nil, err
	}
	return &Response{res: res}, nil
}

// toHTTPRequest converts a Request to a standard HTTP Request
func (r *Request) toHTTPRequest() (*http.Request, error) {
	if r.err != nil {
		return nil, r.err
	}

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
	r.header.Range(func(key, value interface{}) bool {
		req.Header.Set(key.(string), value.(string))
		return true
	})
	return req, nil
}

// doRetry executes the passed http Request using the passed http Client and
// retries as many times as specified
func doRetry(c *http.Client, r *http.Request, expectedStatus, retryCount int) (*http.Response, error) {
	// Perform the request
	res, err := c.Do(r)
	if err != nil {
		return nil, err
	}

	// Retry for the expected status code or return the response
	if retryCount > 0 && res.StatusCode != expectedStatus {
		return doRetry(c, r, expectedStatus, retryCount-1)
	}
	return res, nil
}
