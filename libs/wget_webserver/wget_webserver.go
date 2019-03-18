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
package main

import (
	"20dot.com/hummingbird/libs/httpproxy"
	"20dot.com/hummingbird/libs/instagram"
	"20dot.com/hummingbird/libs/trycatch"
	"20dot.com/hummingbird/libs/webserver"
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"net/url"
	"path/filepath"
)

//default port 8882
var addr = ":8882"

//return obj ,fields first character must uppercase
type ReturnObject struct {
	Code string
	Msg  string
	Data map[string]string
}

// http port
var port = "8882"

// page banner, default: KangSpace.org
var banner = "KangSpace.org"

//  file save path,default: /usr/soft/public/
var path = "/usr/soft/public/"

// file download path,default: https://dl.kangspace.org/dl/
var downloadPath = "https://dl.kangspace.org/dl/"

// proxy
var proxy = ""

func main() {
	args := os.Args
	fmt.Println("args:", args)
	port_ := flag.String("p", port, "http port ,default: 8882")
	path_ := flag.String("path", path, "file save path,default: /usr/soft/public/")
	downloadPath_ := flag.String("dp", downloadPath, "download path,default: https://dl.kangspace.org/dl/")
	banner_ := flag.String("b", banner, "page banner, default: KangSpace.org")
	proxy_ := flag.String("proxy", proxy, "http proxy server")
	flag.Parse()
	port = *port_
	path = *path_
	downloadPath = *downloadPath_
	banner = *banner_
	proxy = *proxy_
	/*
		if len(args) > 1 {
			if b, _ := regexp.MatchString("^[0-9]{2,6}", args[1]); b {
				addr = ":" + args[1]
			}
		}
	*/
	addr = ":" + port
	fmt.Println("addr:", addr)
	server := &webserver.WebServer{Addr: addr}
	handles := http.NewServeMux()
	handles.Handle("/", http.HandlerFunc(welcomeHandle))
	//wget download
	handles.Handle("/wget", http.HandlerFunc(wgetHandle))
	//instagram shared resource download
	handles.Handle("/ins-get", http.HandlerFunc((&instagram.InsResourceDownloader{Proxy:proxy}).InsGetHttpHandle))
	handles.Handle("/http-get", http.HandlerFunc(httpGetHandle))
	server.Run(handles)
}

//default
func welcomeHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	str := "<pre>wget server started !</pre>" +
		"<pre>Usage:" +
		"<p style='color:red'>" + req.Host + "/wget?url=xxx</p>" +//+ req.RequestURI
		"<p style='color:red'>" + req.Host + "/http-get?url=xxx</span></p>" +
		"<p style='color:red'>" + req.Host + "/ins-get?url=xxx</span></p>" +
		"</pre>"
	w.Write([]byte(str))
}

//wget处理
func wgetHandle(w http.ResponseWriter, req *http.Request) {
	webserver.SetAccessControlAllowOrgin(w)
	trycatch.Trycatch(func() {
		// wget url
		url := req.FormValue("url")
		wget_exec_command(w, url)
	}, func(e interface{}) {
		w.Write([]byte("error:" + path))
	})
}

//httpGetHandle : http resource handle
func httpGetHandle(w http.ResponseWriter, req *http.Request) {
	webserver.SetAccessControlAllowOrgin(w)
	url_ := req.FormValue("url")
	targetUrl, err := url.Parse(url_)
	if len(strings.Trim(url_, " ")) < 1 || (strings.Index(url_, "http") != 0 && strings.Index(url_, "https") != 0) || err != nil {
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		str := "<pre>Usage:<p style='color:red'>" + req.Host + "/http-get?url=http://xxxxx/xxx</p></pre>"
		w.Write([]byte(str))
		return
	}
	_,fileName := filepath.Split(targetUrl.Path)
	trycatch.Trycatch(func() {
		// wget url
		resp, err := (&(httpproxy.HttpProxy{proxy})).Client().Get(url_)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("file not found:" + url_+"\n"))
			w.Write([]byte("error:"))
			fmt.Fprintln(w,err)
			return
		}

		defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		readReader := bufio.NewReader(resp.Body)
		//resp.Header.
		for k, v := range resp.Header {
			w.Header().Set(k, strings.Join(v, ","))
		}
		if len(fileName) > 0 {
			w.Header().Set("Content-disposition","filename="+fileName)
		}
		for {
			var bytes = make([]byte, 1024)
			n, err := readReader.Read(bytes)
			w.Write(bytes[:n])
			if io.EOF == err {
				break
			}
		}
	}, func(e interface{}) {
		//w.Write([]byte("error:" + path))
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w,e)
	})
}

func wget_exec_command(w http.ResponseWriter, url string) {
	var styleHtml = "<style type='text/css'>\n" +
		"html,body{margin:0 ;}" +
		"pre{margin:0 20px;}" +
		"#wget_line7{color:#ff0000;}" +
		"\n</style>"
	var bannerHtml = "<div class='nav' style='color: #ffffff;width: auto;height: 40px;vertical-align: middle;background: #3bacf0;align-items: center;padding: 20px;padding-left: 36px;margin-bottom: 10px;box-shadow: 0px 3px 16px 0px rgba(0, 0, 0, 0.1);'>" +
		"<div style='font-size: 1.25rem;line-height: 40px;white-space: nowrap;font-weight: bold;'>" + banner + "</div>" +
		"</div>"
	var consoleDivPrefixHtml = "<div class='console' style='margin: 0 15px;width: calc(100% - 102px);padding: 0 36px;'>" +
		"<div class='console_title' style='font-size: 16px;font-family: Gotham SSm A,Gotham SSm B,Helvetica,Arial,sans-serif;line-height: 2;color: #343434;font-weight: bold;'>Terminal</div>" +
		"<div style='position: relative;overflow-y: auto;overflow-x: hidden;box-sizing: border-box;height: 340px;padding: 19px 27px;background: #000;color: #ccc;font: 12px/16px Menlo,Consolas,Monaco,Lucida Console,Liberation Mono,DejaVu Sans Mono,Bitstream Vera Sans Mono,Courier New,monospace,serif;'>"
	var consoleDivSuffixHtml = "</div></div>"
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Write([]byte(styleHtml))
	w.Write([]byte(bannerHtml))
	w.Write([]byte(consoleDivPrefixHtml))
	//exec.Command("cd ", path_public).Run()
	//cmd := exec.Command("curl","--progress"," -O", url)
	var lastFileSpliterIndex = strings.LastIndex(url, "/")
	var fileName = ""
	// 过滤掉https:// 的长度
	if lastFileSpliterIndex > -1 && lastFileSpliterIndex > 7 &&  lastFileSpliterIndex < len(url) {
		fileName = url[lastFileSpliterIndex+1:]
	}
	if len(fileName) == 0 {
		fileName = "index.html"
	}
	// 短点续传
	cmd := exec.Command("wget", "-c", "-P", path, url)
	errout, err := cmd.StderrPipe()
	//err := cmd.Run()
	err = cmd.Start()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	errorReader := bufio.NewReader(errout)
	var id_wget_line7 = "wget_line7"
	//实时循环读取输出流中的一行内容
	var linenum = 0
	for {
		line, err2 := errorReader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		line = strings.Replace(line, "\n", "", -1)
		if linenum == 7 {
			line = "<div id='" + id_wget_line7 + "'>" + line + "</div>"
			w.Write([]byte(line))
		} else if linenum > 7 {
			//若为换行,不输出
			if len(line) > 0 {
				line = "<script type='text/javascript'>document.getElementById('" + id_wget_line7 + "').innerHTML='" + line + "'</script>"
				w.Write([]byte(line))
			}
		} else {
			line = "<div>" + line + "</div>"
			w.Write([]byte(line))
		}
		linenum++
		f, ok := w.(http.Flusher)
		if ok {
			f.Flush()
		}
	}
	if len(url) > 0 && linenum >= 5 {
		var downloadUrl = downloadPath + fileName
		//输出下载链接
		var downloadUrlHtml = "<div class='downad_link_div' style='margin:35px 0;'>Download: <a target='_blank' href='" + downloadUrl + "'>" + downloadUrl + "</div>"
		w.Write([]byte(downloadUrlHtml))
	} else if len(url) == 0 {
		//输出调用格式
		var usageHtml = "<div class='usage_div' style='margin:35px 0;color:#ff0000;'>Usage: <script type='text/javascript'>var url = location.origin+'/wget?url=example.com/a.xxx';document.write(url);</script></div>"
		w.Write([]byte(usageHtml))
	}
	//输出下载输入框
	var downloadInputHtml = "<div style='position: relative;border-collapse: separate;padding: 30px 0;'>" +
		"<input id='inputurlInp' placeholder='input download url' style='display: inline-block;border-color: #cfdadd;border-radius: 2px;width: 30%;height: 30px;padding: 6px 12px;font-size: 14px;line-height: 1.42857143;color: #555;background-color: #fff;background-image: none;border: 1px solid #cfdadd;box-shadow: none;'>" +
		"<span class='input-group-btn' style='position: relative;font-size: 0;width: 1%;white-space: nowrap;vertical-align: middle;'>" +//else{return 0} ; return 1;
		"<a onclick=\"javascript:var inp = document.getElementById('inputurlInp'); if(inp && inp.value.trim().length){ this.href=location.origin+'/wget?url='+inp.value.trim();};\" _target='blank' type='button' class='btn' style='z-index: 2;margin: 0 10px;display: inline-block;padding: 4px 12px; margin-bottom: 0;font-size: 14px;font-weight: 400;line-height: 1.6;text-align: center;white-space: nowrap;vertical-align: middle;-ms-touch-action: manipulation;touch-action: manipulation;cursor: pointer;-webkit-user-select: none;-moz-user-select: none;-ms-user-select: none;user-select: none;background-image: none;border: 1px solid transparent;border-radius: 4px;background: #3bacf0;color: #fff;'>" +
		"	Download" +
		"	</a>" +
		"		</span>" +
		"	</div>"
	w.Write([]byte(downloadInputHtml))
	w.Write([]byte(consoleDivSuffixHtml))
	cmd.Wait()
}
