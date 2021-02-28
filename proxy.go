package httpclient /* import "s32x.com/httpclient" */

import (
	"net/http"
	"net/url"

	"h12.io/socks"
)

// Transports is a map of proxy TransportFuncs keyed by their protocol
var Transports = map[string]TransportFunc{
	"http": func(addr string) *http.Transport {
		u, _ := url.Parse("http://" + addr)
		return &http.Transport{Proxy: http.ProxyURL(u)}
	},
	"https": func(addr string) *http.Transport {
		u, _ := url.Parse("http://" + addr)
		return &http.Transport{Proxy: http.ProxyURL(u)}
	},
	"socks4": func(addr string) *http.Transport {
		return &http.Transport{Dial: socks.Dial("socks4://" + addr)}
	},
	"socks5": func(addr string) *http.Transport {
		u, _ := url.Parse("socks5://" + addr)
		return &http.Transport{Proxy: http.ProxyURL(u)}
	},
}

// TransportFunc takes an address to a proxy server and returns a fully
// populated http Transport
type TransportFunc func(addr string) *http.Transport

// WithProxy sets a proxy on the Client
func (c *Client) WithProxy(protocol, addr string) *Client {
	if transportFunc, ok := Transports[protocol]; ok {
		return c.WithTransport(transportFunc(addr))
	}
	return c
}
