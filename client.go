package httpclient /* import "s32x.com/httpclient" */

import (
	"net/http"
	"time"

	"s32x.com/httpclient/cache"
	"s32x.com/httpclient/cache/cachers"
)

// Client is an http.Client wrapper
type Client struct {
	client  *http.Client
	cache   cache.Cacher
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

// WithTimeout sets the timeout on the http Client
func (c *Client) WithTimeout(timeout time.Duration) *Client {
	c.client.Timeout = timeout
	return c
}

// WithTransport sets teh transport on the http Client
func (c *Client) WithTransport(transport http.RoundTripper) *Client {
	c.client.Transport = transport
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

// WithCache sets a new cacher on the Client
func (c *Client) WithCache(name string, defaultExpiration time.Duration) (*Client, error) {
	var err error
	c.cache, err = cachers.NewBadger(name, defaultExpiration)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Client is a getter that returns a reference to the underlying http Client
func (c *Client) Client() *http.Client { return c.client }
