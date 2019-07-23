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
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hummingbird/libs/file"
	"github.com/hummingbird/libs/qrcode"
	"github.com/hummingbird/libs/util"
	"github.com/hummingbird/libs/webserver"
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

var path = "/usr/qrcode/"

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
	size := 200
	if len(strings.Trim(returnType, " ")) < 1 {
		returnType = "png"
	}
	if b, _ := regexp.MatchString("^[0-9]{2,4}$", inputSize); b {
		if size_, err := strconv.ParseInt(inputSize, 10, 32); err == nil {
			size = int(size_)
		} else {
			fmt.Println("ParseInt error:", err)
		}
	}
	fmt.Println("size:", size)
	var msg string
	if len(data) > 0 && len(data) <= QRCodeMaxLength {
		fmt.Println("data:", data)
		startTime := time.Now()
		nowTime := strconv.FormatInt(startTime.UnixNano(), 10)
		name := strings.Replace(base64.StdEncoding.EncodeToString([]byte("QRCODE_SIZE_"+strconv.Itoa(size)+"_"+nowTime)), "=", "", -1) + ".png"
		fmt.Println("name:", name, " ,time:", nowTime)
		if err := genQrCode(file.MyFile{path, name}, data, size); err == nil {
			endTime := time.Now()
			fmt.Println("name:", name, " ,cost:", util.CostTimeCalc(startTime, endTime))
			if err := writeHandle(w, returnType, name, path+name); err == nil {
				return
			} else {
				msg = err.Error()
			}
		} else {
			if qrcode.TooMuchCharactersError == err {
				msg = err.Error()
			}
			fmt.Println("genQrCode error:", err)
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
	writeUseage(w)
}

func writeUseage(w http.ResponseWriter) {
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
		obj := ReturnObject{"1", "sucsses", data}
		fmt.Println("obj:", obj)
		if j, err_ := json.Marshal(obj); err_ == nil {
			w.Write(j)
			return nil
		}
		return errors.New("qrHandle json parse error")
	}
	if rt == "png" {
		webserver.SetContentTypePng(w)
		if bytes, err := ioutil.ReadFile(filepath); err == nil {
			w.Write(bytes)
			return nil
		}
		return errors.New("qrcode file not exists")
	}
	return errors.New("not support opeation")
}

//生产qrcode图片
//name: qr图片名称
//path: qr图片路径
func genQrCode(f file.MyFile, content string, size int) error {
	if err := file.CreateDir(f); err == nil {
		if file, err := os.OpenFile(f.FilePath(), os.O_CREATE|os.O_RDWR, 0666); err == nil {
			defer file.Close()
			Try(func() {
				if err = qrcode.EncodeToPng((qrcode.NewQrCode(size, content)), file); err != nil {
					fmt.Println("ERROR genQrCode qrcode.EncodeToPng error:", err)
					file.Close()
					defer os.Remove(f.FilePath())
				}
			}, func(e interface{}) error {
				file.Close()
				defer os.Remove(f.FilePath())
				fmt.Println("ERR- Remove:", f.FilePath())
				err = qrcode.EncodeError
				return err
			})
			return err

		} else {
			return err
		}
	} else {
		return err
	}
	return nil
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
