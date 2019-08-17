package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"sync"
)

type Request struct {
	client *Client

	err    error
	method string
	path   string
	body   io.ReadWriter
	ctx    context.Context
	header sync.Map
}

func newRequest(client *Client, method, path string) *Request {
	return &Request{
		client: client,
		method: method,
		path:   path,
		header: sync.Map{},
	}
}

func (r *Request) WithBytes(body []byte) *Request {
	r.body = bytes.NewBuffer(body)
	return r
}

func (r *Request) WithString(body string) *Request {
	r.body = bytes.NewBufferString(body)
	return r
}

func (r *Request) WithJSON(body interface{}) *Request {
	r = r.WithContentType("application/json")
	r.body = bytes.NewBuffer(nil)
	r.err = json.NewEncoder(r.body).Encode(body)
	return r
}

func (r *Request) WithXML(body interface{}) *Request {
	r = r.WithContentType("application/xml")
	r.body = bytes.NewBuffer(nil)
	r.err = xml.NewEncoder(r.body).Encode(body)
	return r
}

func (r *Request) WithContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

func (r *Request) WithHeader(key, value string) *Request {
	r.header.Store(key, value)
	return r
}

func (r *Request) WithContentType(typ string) *Request {
	return r.WithHeader("Content-Type", typ)
}

// Do performs the passed request and returns a populated Response
func (r *Request) Do() *Response {
	if r.err != nil {
		return &Response{err: r.err}
	}

	// Generate a new http Request using client and passed Request
	req, err := http.NewRequest(r.method, r.client.baseURL+r.path, r.body)
	if err != nil {
		return &Response{err: err}
	}

	// Apply a context if one is set on the Request
	if r.ctx != nil {
		req = req.WithContext(r.ctx)
	}

	// Apply all headers from both the client and the Request
	r.client.header.Range(func(key, value interface{}) bool {
		req.Header.Set(key.(string), value.(string))
		return true
	})
	r.header.Range(func(key, value interface{}) bool {
		req.Header.Set(key.(string), value.(string))
		return true
	})

	// Perform the request and return the wrapped Response
	res, err := r.client.client.Do(req)
	if err != nil {
		return &Response{err: err}
	}
	return &Response{res: res}
}
