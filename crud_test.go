package httpclient

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_Postf(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "POST",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "POST",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Postf(tt.args.format, tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Postf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Post(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
	}
	type args struct {
		path string
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "POST",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "POST",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Post(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Post() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Putf(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "PUT",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "PUT",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Putf(tt.args.format, tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Putf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Put(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
	}
	type args struct {
		path string
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "PUT",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "PUT",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Put(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Patchf(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "PATCH",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "PATCH",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Patchf(tt.args.format, tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Patchf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Patch(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
	}
	type args struct {
		path string
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "PATCH",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "PATCH",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Patch(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Patch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Headf(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "HEAD",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "HEAD",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Headf(tt.args.format, tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Headf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Head(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
	}
	type args struct {
		path string
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "HEAD",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "HEAD",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Head(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Head() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Getf(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "GET",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "GET",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Getf(tt.args.format, tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Getf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Get(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
	}
	type args struct {
		path string
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "GET",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "GET",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Get(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Deletef(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "DELETE",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				format: "/%s/test/%d",
				a:      []interface{}{"asdf", 9},
			},
			want: &Request{
				client:  &http.Client{},
				method:  "DELETE",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Deletef(tt.args.format, tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Deletef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Delete(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
	}
	type args struct {
		path string
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
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "DELETE",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				path: "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "DELETE",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Delete(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Request(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseURL string
		headers []header
	}
	type args struct {
		method string
		path   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "POST basic",
			fields: fields{
				client:  &http.Client{},
				headers: []header{},
			},
			args: args{
				method: "POST",
				path:   "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "POST",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
		{
			name: "PUT basic with header",
			fields: fields{
				client: &http.Client{},
				headers: []header{
					header{key: "some_key", value: "some_value"},
				},
			},
			args: args{
				method: "POST",
				path:   "/asdf/test/9",
			},
			want: &Request{
				client: &http.Client{},
				method: "POST",
				path:   "/asdf/test/9",
				headers: []header{
					header{key: "some_key", value: "some_value"},
				},
			},
		},
		{
			name: "GET with baseurl",
			fields: fields{
				client:  &http.Client{},
				baseURL: "https://example.com",
				headers: []header{},
			},
			args: args{
				method: "GET",
				path:   "/asdf/test/9",
			},
			want: &Request{
				client:  &http.Client{},
				method:  "GET",
				baseURL: "https://example.com",
				path:    "/asdf/test/9",
				headers: []header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				client:  tt.fields.client,
				baseURL: tt.fields.baseURL,
				headers: tt.fields.headers,
			}
			if got := c.Request(tt.args.method, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Request() = %v, want %v", got, tt.want)
			}
		})
	}
}
