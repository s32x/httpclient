package httpclient

import "net/http"

// Post takes a path and returns a prepopulated Post request
func (c *Client) Post(path string) *Request {
	return newRequest(c, http.MethodPost, path)
}

// Put takes a path and returns a prepopulated Put request
func (c *Client) Put(path string) *Request {
	return newRequest(c, http.MethodPut, path)
}

// Patch takes a path and returns a prepopulated Patch request
func (c *Client) Patch(path string) *Request {
	return newRequest(c, http.MethodPatch, path)
}

// Head takes a path and returns a prepopulated Head request
func (c *Client) Head(path string) *Request {
	return newRequest(c, http.MethodHead, path)
}

// Get takes a path and returns a prepopulated Get request
func (c *Client) Get(path string) *Request {
	return newRequest(c, http.MethodGet, path)
}

// Delete takes a path and returns a prepopulated Delete request
func (c *Client) Delete(path string) *Request {
	return newRequest(c, http.MethodDelete, path)
}
