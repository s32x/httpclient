package httpclient

import (
	"net/http"
	"reflect"
	"sync"
	"testing"
)

func TestClient_Postf(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		header  sync.Map
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "basic",
			fields: fields{
				client: &http.Client{},
				header: sync.Map{},
			},
			args: args{
				format: "/%s/test",
				a:      []interface{}{"asdf"},
			},
			want: &Request{
				client: &http.Client{},
				method: "POST",
				path:   "/asdf/test",
				header: sync.Map{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				header:  sync.Map{},
			},
			args: args{
				format: "/%s/test",
				a:      []interface{}{"asdf"},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "POST",
				baseURL: "https://example.com",
				path:    "/asdf/test",
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
			if got := c.Postf(tt.args.format, tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Postf() = %v, want %v", got, tt.want)
			}
		})
	}
}
