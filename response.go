package httpclient

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// Response contains the raw http.Response reference OR any error that took
// place while performing the request
type Response struct{ res *http.Response }

// Close closes the response body on the Responses http Response
func (r *Response) Close() error { return r.res.Body.Close() }

// Response returns the http Response reference that is on the Response
func (r *Response) Response() (*http.Response, error) {
	return r.res, nil
}

// Status returns the status message on the Response
func (r *Response) Status() (string, error) {
	return r.res.Status, nil
}

// StatusCode returns the status code found on the Response
func (r *Response) StatusCode() (int, error) {
	return r.res.StatusCode, nil
}

// String attempts to return the decoded response as a string
func (r *Response) String() (string, error) {
	bytes, err := r.Bytes()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Bytes attempts to return the decoded response as bytes
func (r *Response) Bytes() ([]byte, error) {
	return ioutil.ReadAll(r.res.Body)
}

// JSON attempts to JSON decode the response body into the passed interface
func (r *Response) JSON(i interface{}) error {
	return json.NewDecoder(r.res.Body).Decode(i)
}

// XML attempts to XML decode the response body into the passed interface
func (r *Response) XML(i interface{}) error {
	return xml.NewDecoder(r.res.Body).Decode(i)
}
