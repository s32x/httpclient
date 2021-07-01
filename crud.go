package httpclient /* import "s32x.com/httpclient" */

import (
	"fmt"
	"net/http"
)

// Postf takes a format and a variadic of arguments and returns a prepopulated
// Post request
func (c *Client) Postf(format string, a ...interface{}) *Request {
	return c.Post(fmt.Sprintf(format, a...))
}

// Post takes a path and returns a prepopulated Post request
func (c *Client) Post(path string) *Request {
	return c.Request(http.MethodPost, path)
}

// Putf takes a format and a variadic of arguments and returns a prepopulated
// Put request
func (c *Client) Putf(format string, a ...interface{}) *Request {
	return c.Put(fmt.Sprintf(format, a...))
}

// Put takes a path and returns a prepopulated Put request
func (c *Client) Put(path string) *Request {
	return c.Request(http.MethodPut, path)
}

// Patchf takes a format and a variadic of arguments and returns a prepopulated
// Patch request
func (c *Client) Patchf(format string, a ...interface{}) *Request {
	return c.Patch(fmt.Sprintf(format, a...))
}

// Patch takes a path and returns a prepopulated Patch request
func (c *Client) Patch(path string) *Request {
	return c.Request(http.MethodPatch, path)
}

// Headf takes a format and a variadic of arguments and returns a prepopulated
// Head request
func (c *Client) Headf(format string, a ...interface{}) *Request {
	return c.Head(fmt.Sprintf(format, a...))
}

// Head takes a path and returns a prepopulated Head request
func (c *Client) Head(path string) *Request {
	return c.Request(http.MethodHead, path)
}

// Getf takes a format and a variadic of arguments and returns a prepopulated
// Get request
func (c *Client) Getf(format string, a ...interface{}) *Request {
	return c.Get(fmt.Sprintf(format, a...))
}

// Get takes a path and returns a prepopulated Get request
func (c *Client) Get(path string) *Request {
	return c.Request(http.MethodGet, path)
}

// Deletef takes a format and a variadic of arguments and returns a prepopulated
// Delete request
func (c *Client) Deletef(format string, a ...interface{}) *Request {
	return c.Delete(fmt.Sprintf(format, a...))
}

// Delete takes a path and returns a prepopulated Delete request
func (c *Client) Delete(path string) *Request {
	return c.Request(http.MethodDelete, path)
}

// Request creates a new Request copying configuration from the base Client
func (c *Client) Request(method, path string) *Request {
	r := &Request{
		client:  c.client,
		cache:   c.cache,
		method:  method,
		baseURL: c.baseURL,
		path:    path,
		headers: []header{},
	}
	for _, h := range c.headers {
		r.headers = append(r.headers, header{key: h.key, value: h.value})
	}
	return r
}
