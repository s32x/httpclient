package httpclient

import (
	"io/ioutil"
	"net/http"
)

// Response is a basic HTTP response struct containing
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

// NewResponse creates a new http Response
func NewResponse(res *http.Response) *Response {
	// Create a map of all headers
	headers := make(map[string]string)
	for k, v := range res.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	// Decode the body into a slice of bytes
	body, _ := ioutil.ReadAll(res.Body)

	// Return the fully constructed Response
	return &Response{
		StatusCode: res.StatusCode,
		Headers:    headers,
		Body:       body,
	}
}
