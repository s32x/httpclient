package httpclient /* import "s32x.com/httpclient" */

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Response contains the raw http.Response reference OR any error that took
// place while performing the request
type Response struct {
	Status int         `json:"status,omitempty"`
	Header http.Header `json:"header,omitempty"`
	Body   []byte      `json:"body,omitempty"`
}

// newResponse creates and returns a new, fully populated Response struct after
// reading the response body
func newResponse(r *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("read http response: %w", err)
	}
	return &Response{Status: r.StatusCode, Header: r.Header, Body: body}, nil
}

// ContentType returns the content-type header found on the response
func (r *Response) ContentType() string { return r.Header.Get("Content-Type") }

// String attempts to return the decoded response as a string
func (r *Response) String() (string, error) { return string(r.Body), nil }

// Bytes attempts to return the decoded response as bytes
func (r *Response) Bytes() []byte { return r.Body }

// JSON attempts to JSON decode the response body into the passed interface
func (r *Response) JSON(out interface{}) error { return json.Unmarshal(r.Body, out) }

// XML attempts to XML decode the response body into the passed interface
func (r *Response) XML(out interface{}) error { return xml.Unmarshal(r.Body, out) }
