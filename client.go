package httpclient

import (
	"net/http"
	"sync"
	"time"
)

// Client is an http.Client wrapper
type Client struct {
	client         *http.Client
	baseURL        string
	header         sync.Map
	expectedStatus int
	retryCount     int
}

// New creates a new Client reference given a client timeout
func New() *Client {
	return &Client{
		client:         &http.Client{},
		header:         sync.Map{},
		expectedStatus: http.StatusOK, // Default to expect 200 status codes
	}
}

// WithClient sets the http client on the Client
func (c *Client) WithClient(client *http.Client) *Client {
	c.client = client
	return c
}

// WithTimeout sets the timeout on the http client
func (c *Client) WithTimeout(timeout time.Duration) *Client {
	c.client.Timeout = timeout
	return c
}

// WithBaseURL sets the baseURL on the Client
func (c *Client) WithBaseURL(url string) *Client {
	c.baseURL = url
	return c
}

// WithHeader sets the headers on the Client
func (c *Client) WithHeader(key, value string) *Client {
	c.header.Store(key, value)
	return c
}

// WithExpectedStatus sets the desired status-code that will be a success on
// the Client
func (c *Client) WithExpectedStatus(expectedStatusCode int) *Client {
	c.expectedStatus = expectedStatusCode
	return c
}

// WithRetry sets the desired number of retries on the Client
func (c *Client) WithRetry(retryCount int) *Client {
	c.retryCount = retryCount
	return c
}

// Request creates a new Request copying configuration from the base Client
func (c *Client) Request() *Request {
	r := &Request{
		client:         c.client,
		baseURL:        c.baseURL,
		header:         sync.Map{},
		expectedStatus: c.expectedStatus,
		retryCount:     c.retryCount,
	}
	c.header.Range(func(key, val interface{}) bool {
		r.header.Store(key, val)
		return true
	})
	return r
}
