package httpclient

import "time"

// DefaultClient is a default Client for using without
// having to declare a Client
var DefaultClient = NewBaseClient(time.Second * 30)

// Head calls Head using the DefaultClient
func Head(url string) error {
	return DefaultClient.Head(url)
}

// GetBytes calls GetBytes using the DefaultClient
func GetBytes(url string) ([]byte, error) {
	return DefaultClient.GetBytes(url)
}

// GetString calls GetString using the DefaultClient
func GetString(url string) (string, error) {
	return DefaultClient.GetString(url)
}

// GetJSON calls GetJSON using the DefaultClient
func GetJSON(url string, out interface{}) error {
	return DefaultClient.GetJSON(url, out)
}

// Delete calls Delete using the DefaultClient
func Delete(url string) error {
	return DefaultClient.Delete(url)
}
