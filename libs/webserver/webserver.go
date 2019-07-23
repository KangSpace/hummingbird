// Web Http Server
package webserver

import (
	"fmt"
	"net/http"
)

type WebServer struct {
	//value is ":port"
	Addr string
}

const (
	ContentTypeHTML  = "text/html;charset=utf-8"
	ContentTypePlain = "text/plain;charset=utf-8"
	ContentTypeJSON  = "application/json;charset=utf-8"
	ContentTypePNG   = "image/png"
)

// 创建 http Server
// 	handles : *http.serveMux
// 	example:
//		addr := ":8881"
//	  server := &webserver.WebServer{Addr: addr}
//		handles := http.NewServeMux();
//		handles.Handle("/", http.HandlerFunc(welcomeHandle))
//		handles.Handle("/qrcode", http.HandlerFunc(qrHandle))
//		server.Run(handles)
func (w *WebServer) Run(handlers *http.ServeMux) {
	fmt.Println("Host:", w.Addr)
	fmt.Println("# URIs:", handlers)
	if err := http.ListenAndServe(w.Addr, handlers); err != nil {
		fmt.Println("ListenAndServe ERROR:", err)
	} else {
		fmt.Println("Host:", "server started")
	}
}

//设置相应头
func SetResponseHeader(w http.ResponseWriter, key string, value string) {
	w.Header().Set(key, value)
}

//设置请求头
func SetRequestHeader(w http.ResponseWriter, key string, value string) {
	w.Header().Set(key, value)
}

//设置跨域允许访问
func SetAccessControlAllowOrgin(w http.ResponseWriter) {
	SetResponseHeader(w, "Access-Control-Allow-Origin", "*")
}

//设置返回类型为JSON
func SetContentTypeJson(w http.ResponseWriter) {
	SetResponseHeader(w, "Content-Type", ContentTypeJSON)
}

//设置返回类型为HTML
func SetContentTypeHtml(w http.ResponseWriter) {
	SetResponseHeader(w, "Content-Type", ContentTypeHTML)
}

//设置返回类型为PNG
func SetContentTypePng(w http.ResponseWriter) {
	SetResponseHeader(w, "Content-Type", ContentTypePNG)
}
