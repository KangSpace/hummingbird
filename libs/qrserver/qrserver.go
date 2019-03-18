// qrcode 处理server :8881
//
// author: Kang
// date: 2017-05-30
// run: qrserver [8881]
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
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
	Code string
	Msg  string
	Data map[string]string
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

//default
func welcomeHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("server started !"))
}

//qr处理
func qrHandle(w http.ResponseWriter, req *http.Request) {
	webserver.SetAccessControlAllowOrgin(w)
	// type :d return data
	data := req.FormValue("data")
	//200
	inputSize := req.FormValue("size")
	size := 200
	if b, _ := regexp.MatchString("^[0-9]{2,3}$", inputSize); b {
		if size_, err := strconv.ParseInt(inputSize, 10, 32); err == nil {
			size = int(size_)
		} else {
			fmt.Println("ParseInt error:", err)
		}
	}
	fmt.Println("size:", size)
	var msg string
	if data != "" {
		fmt.Println("data:", data)
		startTime := time.Now()
		nowTime := strconv.FormatInt(startTime.UnixNano(), 10)
		name := base64.StdEncoding.EncodeToString([]byte("QRCODE_"+nowTime)) + ".png"
		fmt.Println("name:", name, " ,time:", nowTime)
		if err := genQrCode(file.MyFile{path, name}, data, size); err == nil {
			endTime := time.Now()
			fmt.Println("name:", name, " ,cost:", util.CostTimeCalc(startTime, endTime))
			data := map[string]string{
				"name": name,
			}
			obj := ReturnObject{"1", "成功", data}
			fmt.Println("obj:", obj)
			if j, err_ := json.Marshal(obj); err_ == nil {
				w.Write(j)
				return
			} else {
				fmt.Println("qrHandle json parse error")
			}
		} else {
			if qrcode.TooMuchCharactersError == err {
				msg = err.Error()
			}
			fmt.Println("genQrCode error:", err)
		}
		w.WriteHeader(500)
	} else {
		w.WriteHeader(500)
		msg = "qrHandle no data"
		fmt.Println("no data")
	}
	w.Write([]byte(msg))
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
