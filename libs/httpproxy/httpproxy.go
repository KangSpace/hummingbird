//httpProxy get http client or proxyClient
package httpproxy

import (
	"net/http"
	"net/url"
)

//HttpProxy
type HttpProxy struct {
	//proxy server , e.g,: http://127.0.0.1:1880
	Proxy string
}

//get proxy http client
func (downloader *HttpProxy) Client() *http.Client {
	var client *http.Client
	if len(downloader.Proxy) > 0 {
		proxy := func(_ *http.Request) (*url.URL, error) {
			//"http://127.0.0.1:1880"
			return url.Parse(downloader.Proxy)
		}
		transport := &http.Transport{Proxy: proxy}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}
	return client
	//resp, err := client.Get("http://www.google.com")
}
