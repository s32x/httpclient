package httpclient

import (
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	c := New()
	assert.Equal(t, &Client{
		client:  &http.Client{},
		baseURL: "",
		header:  sync.Map{},
	}, c)

	c = c.WithClient(&http.Client{Timeout: 20 * time.Second})
	assert.Equal(t, &Client{
		client:  &http.Client{Timeout: 20 * time.Second},
		baseURL: "",
		header:  sync.Map{},
	}, c)

	c = c.WithTimeout(time.Second * 10)
	assert.Equal(t, &Client{
		client:  &http.Client{Timeout: 10 * time.Second},
		baseURL: "",
		header:  sync.Map{},
	}, c)

	c = c.WithBaseURL("https://example.com")
	assert.Equal(t, &Client{
		client:  &http.Client{Timeout: 10 * time.Second},
		baseURL: "https://example.com",
		header:  sync.Map{},
	}, c)

	// var m sync.Map
	// m.Store("hkey", "hval")

	// c = c.WithHeader("hkey", "hval")
	// assert.Equal(t, &Client{
	// 	client:         &http.Client{Timeout: 10 * time.Second},
	// 	baseURL:        "https://example.com",
	// 	header:         m,
	// 	expectedStatus: 200,
	// 	retryCount:     0,
	// }, c)

	req := c.Request(http.MethodGet, "/test")
	assert.Equal(t, &Request{
		client:  &http.Client{Timeout: 10 * time.Second},
		method:  "GET",
		baseURL: "https://example.com",
		path:    "/test",
		header:  sync.Map{},
		body:    nil,
		ctx:     nil,
	}, req)

	// Verify generating the request hasn't modified the Client
	assert.Equal(t, &Client{
		client:  &http.Client{Timeout: 10 * time.Second},
		baseURL: "https://example.com",
		header:  sync.Map{},
	}, c)
}
