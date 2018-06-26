package httpclient

// httpclient is a convenience package for executing HTTP
// requests. It's safe in that it always closes response
// bodies and returns byte slices, strings or decodes
// responses into interfaces

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// Client is an http.Client wrapper
type Client struct {
	Client  *http.Client
	BaseURL string
	Headers map[string]string
}

// DefaultClient is a default Client for using without
// having to declare a Client
var DefaultClient = NewBaseClient()

// NewBaseClient creates a new Client reference given a
// client timeout
func NewBaseClient() *Client {
	return &Client{Client: &http.Client{}}
}

// SetTimeout sets the timeout on the httpclients client
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.Client.Timeout = timeout
	return c
}

// SetBaseURL sets the baseURL on the Client which will
// be used on all subsequent requests
func (c *Client) SetBaseURL(url string) *Client {
	c.BaseURL = url
	return c
}

// SetHeaders sets the headers on the Client which will
// be used on all subsequent requests
func (c *Client) SetHeaders(headers map[string]string) *Client {
	c.Headers = headers
	return c
}

// Do performs the request and returns a fully populated
// Response
func (c *Client) Do(method, path string,
	headers map[string]string, body []byte) (*Response, error) {
	// Build the full request URL
	url := c.BaseURL + path

	// Encode the body if one was passed
	var b io.ReadWriter
	if body != nil {
		b = bytes.NewBuffer(body)
	}

	// Generate a new request using the new URL
	r, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, err
	}

	// Add any client and passed headers to the new request
	if c.Headers != nil {
		for k, v := range c.Headers {
			r.Header.Set(k, v)
		}
	}
	if headers != nil {
		for k, v := range headers {
			r.Header.Set(k, v)
		}
	}

	// Execute the fully constructed request
	res, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response into a Response and return
	return NewResponse(res), nil
}
