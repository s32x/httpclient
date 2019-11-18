package httpclient

import (
	"net/http"
	"time"
)

// Client is an http.Client wrapper
type Client struct {
	client  *http.Client
	baseURL string
	headers []header
}

// header is a struct that contains a key and a value
type header struct{ key, value string }

// New creates a new Client reference given a client timeout
func New() *Client {
	return &Client{
		client:  &http.Client{},
		headers: []header{},
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
	c.headers = append(c.headers, header{key: key, value: value})
	return c
}
