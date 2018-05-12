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
func (c *Client) SetCache(cacheExp, cacheCleanup time.Duration) *Client {
	c.Cache = cache.New(cacheExp, cacheCleanup)
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

// Head performs a HEAD request using the passed path
func (c *Client) Head(path string) error {
	// Execute the request and return the response
	_, err := c.bytes(http.MethodHead, path, nil)
	return err
}

// GetBytes performs a GET request using the passed path
func (c *Client) GetBytes(path string) ([]byte, error) {
	// Execute the request and return the response
	return c.bytes(http.MethodGet, path, nil)
}

// GetString performs a GET request and returns the response
// as a string
func (c *Client) GetString(path string) (string, error) {
	// Retrieve the bytes and decode the response
	body, err := c.GetBytes(path)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// GetJSON performs a basic http GET request and decodes the JSON
// response into the out interface
func (c *Client) GetJSON(path string, out interface{}) error {
	// Retrieve the bytes and decode the response
	body, err := c.GetBytes(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, out)
}

// Delete performs a DELETE request using the passed path
func (c *Client) Delete(path string) error {
	// Execute the request and return the response
	_, err := c.bytes(http.MethodDelete, path, nil)
	return err
}

// bytes executes the passed request using the Client
// http.Client, returning all the bytes read from the response
func (c *Client) bytes(method, path string, in interface{}) ([]byte, error) {
	// Assemble the BaseURL + Path url
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
		if bIface, ok := c.Cache.Get(url); ok {
			if bytes, ok := bIface.([]byte); ok {
				return bytes, nil
			}
		}
	}

	// Generate the request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Set all headers
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	// Execute the passed request
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
		return bytes, fmt.Errorf("Non 200 status code : %s", res.Status)
	}

	// Decode and return the bytes
	return bytes, nil
}
