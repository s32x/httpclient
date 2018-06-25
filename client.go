package httpclient

// httpclient is a convenience package for executing HTTP
// requests. It's safe in that it always closes response
// bodies and returns byte slices, strings or decodes
// responses into interfaces

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	cache "github.com/patrickmn/go-cache"
)

// Client is an http.Client wrapper
type Client struct {
	Client  *http.Client
	Cache   *cache.Cache
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

// SetCache sets the cache on the Client which will be
// used on all subsequent requests
func (c *Client) SetCache(expiration, cleanup time.Duration) *Client {
	c.Cache = cache.New(expiration, cleanup)
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

// bytes executes the passed request using the Client
// http.Client, returning all the bytes read from the response
func (c *Client) bytes(method, path string,
	headers map[string]string, in interface{}) ([]byte, error) {
	// Assemble the full request URL
	url := c.BaseURL + path

	// Marshal a request body if one exists
	var body io.ReadWriter
	if in != nil {
		if err := json.NewEncoder(body).Encode(in); err != nil {
			return nil, err
		}
	}

	// Return cached content
	if method == http.MethodGet && c.Cache != nil {
		if b, ok := c.Cache.Get(url); ok {
			return b.([]byte), nil
		}
	}

	// Generate a new http Request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Set all headers
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	// Execute the newly generated request
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the body
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Store the new bytes response in cache
	if method == http.MethodGet && c.Cache != nil {
		c.Cache.SetDefault(url, bytes)
	}

	// Check the status code for an OK
	if res.StatusCode >= 400 {
		return bytes, fmt.Errorf("400+ status code received : %s", res.Status)
	}

	// Decode and return the bytes
	return bytes, nil
}
