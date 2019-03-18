// Download Instagram Shared Image/Videos
package instagram

import (
	"testing"
)

func TestMain2(t *testing.T) {
	var url = "https://www.instagram.com/p/Brx04JWHqfS/"
	var proxyServer = "http://127.0.0.1:1880"
	var downloader = &InsResourceDownloader{URL: url, ProxyServer: proxyServer}
	data, _ := downloader.FetchInsShareDataResources()
	t.Log(data.ToJSONString())
}
