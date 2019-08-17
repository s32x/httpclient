package httpclient

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Response contains the raw http.Response reference OR any error that took
// place while performing the request
type Response struct {
	err error
	res *http.Response
}

// Response returns the http Response reference that is on the Response
func (r *Response) Response() (*http.Response, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.res, nil
}

// Status returns the status message on the Response
func (r *Response) Status() (string, error) {
	if r.err != nil {
		return "", r.err
	}
	return r.res.Status, nil
}

// StatusCode returns the status code found on the Response
func (r *Response) StatusCode() (int, error) {
	if r.err != nil {
		return 0, r.err
	}
	return r.res.StatusCode, nil
}

// ExpectStatus takes an expected status code and sets an error on the Response
// if the expected code isn't what is actually received
func (r *Response) ExpectStatus(expected int) *Response {
	if code := r.res.StatusCode; code != expected {
		r.err = fmt.Errorf("Unexpected status code : %v", code)
		return r
	}
	return r
}

// Bytes attempts to return the decoded response as bytes
func (r *Response) Bytes() ([]byte, error) {
	if r.err != nil {
		return nil, r.err
	}
	defer r.res.Body.Close()
	return ioutil.ReadAll(r.res.Body)
}

// String attempts to return the decoded response as a string
func (r *Response) String() (string, error) {
	if r.err != nil {
		return "", r.err
	}
	bytes, err := r.Bytes()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// JSON attempts to JSON decode the response body into the passed interface
func (r *Response) JSON(i interface{}) error {
	if r.err != nil {
		return r.err
	}
	defer r.res.Body.Close()
	return json.NewDecoder(r.res.Body).Decode(i)
}

// XML attempts to XML decode the response body into the passed interface
func (r *Response) XML(i interface{}) error {
	if r.err != nil {
		return r.err
	}
	defer r.res.Body.Close()
	return xml.NewDecoder(r.res.Body).Decode(i)
}
