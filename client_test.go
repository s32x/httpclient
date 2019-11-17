package httpclient

import (
	"net/http"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Client
	}{
		{
			name: "new",
			want: &Client{
				client: &http.Client{},
				header: sync.Map{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_WithClient(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		header  sync.Map
	}
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Client
	}{
		{
			name: "client with 20 second timeout",
			fields: fields{
				client: &http.Client{},
				header: sync.Map{},
			},
			args: args{
				client: &http.Client{Timeout: 20 * time.Second},
			},
			want: &Client{
				client: &http.Client{Timeout: 20 * time.Second},
				header: sync.Map{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				header:  tt.fields.header,
			}
			if got := c.WithClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.WithClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_WithTimeout(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		header  sync.Map
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Client
	}{
		{
			name: "5 second timeout",
			fields: fields{
				client: &http.Client{},
				header: sync.Map{},
			},
			args: args{
				timeout: 5 * time.Second,
			},
			want: &Client{
				client: &http.Client{Timeout: 5 * time.Second},
				header: sync.Map{},
			},
		},
		{
			name: "30 second timeout",
			fields: fields{
				client: &http.Client{},
				header: sync.Map{},
			},
			args: args{
				timeout: 30 * time.Second,
			},
			want: &Client{
				client: &http.Client{Timeout: 30 * time.Second},
				header: sync.Map{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				header:  tt.fields.header,
			}
			if got := c.WithTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.WithTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_WithBaseURL(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		header  sync.Map
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Client
	}{
		{
			name: "example base URL",
			fields: fields{
				client: &http.Client{},
				header: sync.Map{},
			},
			args: args{
				url: "https://example.com",
			},
			want: &Client{
				client:  &http.Client{},
				baseURL: "https://example.com",
				header:  sync.Map{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				header:  tt.fields.header,
			}
			if got := c.WithBaseURL(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.WithBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestClient_WithHeader(t *testing.T) {
// 	type fields struct {
// 		client  *http.Client
// 		baseURL string
// 		header  sync.Map
// 	}
// 	type args struct {
// 		key   string
// 		value string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   *Client
// 	}{
// 		{
// 			name: "example base URL",
// 			fields: fields{
// 				client: &http.Client{},
// 				header: sync.Map{},
// 			},
// 			args: args{
// 				key:   "some_header_key",
// 				value: "some_header_value",
// 			},
// 			want: &Client{
// 				client: &http.Client{},
// 				header: sync.Map{},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := &Client{
// 				client:  tt.fields.client,
// 				baseURL: tt.fields.baseURL,
// 				header:  tt.fields.header,
// 			}
// 			if got := c.WithHeader(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Client.WithHeader() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
