//
// author: kango2gler@gmail.com
// date: 2018-12-22
//
// Useage: wget_webserver -p -path -dp -b
// 	   -p  http port ,default: 8882
//	   -path file save path,default: /usr/soft/public/
//	   -dp file download path,default: https://dl.kangspace.org/dl/
//	   -b page banner, default: KangSpace.org
//	   -proxy http proxy, default: nil
// e.g.:wget_webserver -p 8883 -path /usr/soft/ -dp https://dl.kangspace.org/dl/ -b KanfSpace.org
// WebURL :
// 	1. wget to remote server return upload url
// 		http:localhost:8882/wget?url=http://xxx.com/x.x
//	2. download instagram shared resource
// 		http:localhost:8882/ins-get?url=https://www.instagram.com/p/Brx04JWHqfS/
//	3. get remote resource
// 		http:localhost:8882/http-get?url=http://.jpg
//
package webserver

import (
	"fmt"
	"net/url"
	"testing"
	"path/filepath"
)

func TestURL(t *testing.T) {
	//var url_ = "https://scontent-lax3-2.cdninstagram.com/vp/a34102b8f4522dee44e37eb61d55face/5D1E8A42/t51.2885-15/e15/11333435_1431689337140652_526327065_n.jpg?_nc_ht=scontent-lax3-2.cdninstagram.com"
	var url_ = "123.com"
	//var url_ = ""
	targetUrl, err := url.Parse(url_)
	if err != nil{
		fmt.Println(err)
		t.Error(err)
	}
	fmt.Println("targetUrl.Path:",targetUrl.Path)
	fmt.Println("targetUrl.Scheme:",targetUrl.Scheme)
	fmt.Println("targetUrl.ForceQuery:",targetUrl.ForceQuery)
	fmt.Println("targetUrl.Fragment:",targetUrl.Fragment)
	fmt.Println("targetUrl.Opaque:",targetUrl.Opaque)
	fmt.Println("targetUrl.RawPath:",targetUrl.RawPath)
	fmt.Println("targetUrl.RawQuery:",targetUrl.RawQuery)
	fmt.Println("targetUrl:",targetUrl)
	dir,fileName :=filepath.Split(url_)
	fmt.Println(len(fileName))
	if len(fileName) <1{
		fmt.Println("fileName == nil")
	}
	fmt.Println("filepath.Split:",dir,fileName,)
}
