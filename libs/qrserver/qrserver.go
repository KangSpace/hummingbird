// qrcode 处理server :8881
//
// author: Kang
// date: 2017-05-30
// run: qrserver [8881]
// file will be save on /usr/qrcode/ dictionary
package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/KangSpace/gqrcode"
	"github.com/KangSpace/gqrcode/core/output"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"hummingbird/libs/file"
	"hummingbird/libs/util"
	"hummingbird/libs/webserver"
)

//default port 8881
var addr = ":8881"

//return obj ,fields first character must uppercase
type ReturnObject struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data map[string]string `json:"data"`
}

func main() {
	args := os.Args
	fmt.Println("args:", args)
	if len(args) > 1 {
		if b, _ := regexp.MatchString("^[0-9]{2,6}", args[1]); b {
			addr = ":" + args[1]
		}
	}

	server := &webserver.WebServer{Addr: addr}
	handles := http.NewServeMux()
	handles.Handle("/", http.HandlerFunc(welcomeHandle))
	handles.Handle("/qrcode", http.HandlerFunc(qrHandle))
	server.Run(handles)
}

var path = os.Getenv("HOME") + "/usr/qrcode"

const (
	ServiceUsageQrCode = "/qrcode \t e.g.:/qrcode?data=test&s=200&rt=json \n" +
		"         \t data: data for need to generate qrcode ,max length is 7089 ,refer to https://en.wikipedia.org/wiki/QR_code \n" +
		"         \t s: size for qrcode image ,int value,default is 200 \n" +
		"         \t rt: json/png ,result type,defalut is png   \n"
	QRCodeMaxLength = 7089
)

//default
func welcomeHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("server started !\n\n"))
	w.Write([]byte("service list :\n"))
	w.Write([]byte(ServiceUsageQrCode))
}

//qr处理
func qrHandle(w http.ResponseWriter, req *http.Request) {
	//set CROS
	webserver.SetAccessControlAllowOrgin(w)
	// type :d return data
	data := req.FormValue("data")
	//default is 200
	inputSize := req.FormValue("s")
	//default is png
	returnType := req.FormValue("rt")
	size := 400
	isJsonResp := returnType == "json"
	if !isJsonResp && len(strings.Trim(returnType, " ")) < 1 {
		returnType = "png"
	}
	if b, _ := regexp.MatchString("^[0-9]{2,4}$", inputSize); b {
		if size_, err := strconv.ParseInt(inputSize, 10, 32); err == nil {
			size = int(size_)
		} else {
			fmt.Println("ParseInt error:", err)
		}
	}
	fmt.Println("size:", size, " , returnType:", returnType)
	var msg string
	if len(data) > 0 && len(data) <= QRCodeMaxLength {
		fmt.Println("data:", data)
		startTime := time.Now()
		nowTime := strconv.FormatInt(startTime.UnixNano(), 10)
		name := strings.Replace(base64.StdEncoding.EncodeToString([]byte("QRCODE_SIZE_"+strconv.Itoa(size)+"_"+nowTime)), "=", "", -1) + ".png"
		fmt.Println("name:", name, " ,time:", nowTime)
		if returnType == "json" {
			destFile := file.MyFile{Path: path, FileName: name}
			if err := genQrCode(destFile, data, size); err == nil {
				endTime := time.Now()
				fmt.Println("name:", name, " ,cost:", util.CostTimeCalc(startTime, endTime))
				if err := writeHandle(w, returnType, name, destFile.FilePath()); err == nil {
					return
				} else {
					msg = err.Error()
				}
			} else {
				fmt.Println("genQrCode error:", err)
			}
		} else {
			if err := writeImageHandle(w, data, size); err == nil {
				return
			} else {
				msg = err.Error()
			}
		}
	}
	if len(data) < 1 {
		msg = "error: data is null!\n\n"
		fmt.Println("no data")
	}
	if len(data) > QRCodeMaxLength {
		msg = "error: data max length is " + strconv.Itoa(QRCodeMaxLength) + "!\n\n"
		fmt.Println("no data")
	}
	webserver.SetContentTypeHtml(w)
	w.WriteHeader(500)
	w.Write([]byte(msg))
	writeUsage(w)
}

func writeUsage(w http.ResponseWriter) {
	w.Write([]byte("Usage :\n"))
	w.Write([]byte(ServiceUsageQrCode))
}

//输出处理
//w response
//rt returnType: json/png
func writeHandle(w http.ResponseWriter, rt string, name string, filepath string) error {
	if rt == "json" {
		webserver.SetContentTypeJson(w)
		data := map[string]string{
			"name": name,
		}
		obj := ReturnObject{"1", "success", data}
		fmt.Println("obj:", obj)
		if j, err_ := json.Marshal(obj); err_ == nil {
			w.Write(j)
			return nil
		}
		return errors.New("qrHandle json parse error")
	} else {
		webserver.SetContentTypePng(w)
		if bytes, err := ioutil.ReadFile(filepath); err == nil {
			w.Write(bytes)
			return nil
		}
		return errors.New("qrcode file not exists! ")
	}
	return errors.New("not support operation! ")
}

func writeImageHandle(w http.ResponseWriter, data string, size int) error {
	webserver.SetContentTypePng(w)
	return genQrCodeToWriter(w, data, size)
}

func mkdirIfAbsent(f file.MyFile, run func(onError func(error))) error {
	if err := file.CreateDir(f); err == nil {
		Try(func() {
			run(func(runErr error) {
				err = runErr
			})
		}, func(e interface{}) error {
			defer os.Remove(f.FilePath())
			fmt.Println("ERR- Remove:", f.FilePath())
			return err
		})
	} else {
		return err
	}
	return nil
}

//生产qrcode图片
//name: qr图片名称
//path: qr图片路径
func genQrCode(f file.MyFile, content string, size int) error {
	return mkdirIfAbsent(f, func(onError func(error)) {
		qrcode, _ := gqrcode.NewQRCodeAutoQuiet(content)
		err := qrcode.Encode(output.NewPNGOutput(size), f.FilePath())
		log.Println(f.FilePath())
		if err != nil {
			onError(err)
		}
	})
}
func genQrCodeToWriter(w io.Writer, content string, size int) error {
	qrcode, _ := gqrcode.NewQRCodeAutoQuiet(content)
	err := qrcode.EncodeToWriter(output.NewPNGOutput(size), w)
	return err
}

//实现 try catch 例子
func Try(fun func(), handler func(interface{}) error) error {
	var err error
	defer func() {
		if err_ := recover(); err_ != nil {
			fmt.Println("ERR- Try recover Err:", err_)
			err = handler(err_)
		}
	}()
	fun()
	if err != nil {
		fmt.Println("ERR- Try Err:", err)
	}
	return err
}
