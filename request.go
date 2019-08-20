package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
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
	defer res.Close()
	return res.Bytes()
}

// JSON is a convenience method that handles executing, defer closing, and
// decoding the JSON body into the passed interface before returning
func (r *Request) JSON(in interface{}) error {
	res, err := r.Do()
	if err != nil {
		return err
	}
	defer res.Close()
	return res.JSON(in)
}

// XML is a convenience method that handles executing, defer closing, and
// decoding the XML body into the passed interface before returning
func (r *Request) XML(in interface{}) error {
	res, err := r.Do()
	if err != nil {
		return err
	}
	defer res.Close()
	return res.XML(in)
}

// Do performs the passed request and returns a populated Response
// NOTE: As with the standard library, when calling Do you must remember to
// close the response body : res.Body.Close()
func (r *Request) Do() (*Response, error) {
	if r.err != nil {
		return nil, r.err
	}

	// Convert the Request to a standard http Request and perform the request
	// with retries, returning the wrapped http.Response
	req, err := r.toHTTPRequest()
	if err != nil {
		return nil, err
	}
	return doRetry(r.client, req, r.expectedStatus, r.retryCount)
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
	r.header.Range(func(key, value interface{}) bool {
		req.Header.Set(key.(string), value.(string))
		return true
	})
	return req, nil
}

// doRetry executes the passed http Request using the passed http Client and
// retries as many times as specified
func doRetry(c *http.Client, r *http.Request, expectedStatus,
	retryCount int) (*Response, error) {
	// Perform the request using the standard library
	res, err := c.Do(r)
	if err != nil {
		return nil, err
	}

	// Retry for the expected status code or return the response
	if expectedStatus > 0 && res.StatusCode != expectedStatus {
		if retryCount > 0 {
			return doRetry(c, r, expectedStatus, retryCount-1)
		}
		return nil, fmt.Errorf("Unexpected status code received : %v",
			res.StatusCode)
	}
	return &Response{res: res}, nil
}
