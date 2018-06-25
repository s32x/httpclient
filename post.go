package httpclient

import (
	"encoding/json"
	"net/http"
)

// PostJSON performs a basic http POST request and decodes the JSON
// response into the out interface
func (c *Client) PostJSON(path string, in, out interface{}) error {
	// Retrieve the bytes and decode the response
	body, err := c.PostBytes(path, nil, in)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, out)
}

// PostBytes performs a POST request using the passed path and body
func (c *Client) PostBytes(path string, headers map[string]string, in interface{}) ([]byte, error) {
	// Execute the request and return the response
	return c.bytes(http.MethodPost, path, headers, in)
}
