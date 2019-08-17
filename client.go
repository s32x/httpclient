package httpclient

import (
	"net/http"
	"sync"
)

// Client is an http.Client wrapper
type Client struct {
	client  *http.Client
	baseURL string
	header  sync.Map
}

// New creates a new Client reference given a client timeout
func New() *Client {
	return &Client{client: &http.Client{}, header: sync.Map{}}
}

// WithClient sets the timeout on the Client
func (c *Client) WithClient(client *http.Client) *Client {
	c.client = client
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
