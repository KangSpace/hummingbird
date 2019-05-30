package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type DownloadRange struct {
	Range int
}

func Download() {
	var url = "https://cdn5.hd.etiantian.net/143dfd3dfb6516f13c325bd961bcff6d/5CEFAE99/etthd/wkczsx000676/400.mp4"
	fileName := "D:\\tmp\\400.mp4"
	file := getFile(fileName)
	defer file.Close()
	fileInfo, _ := file.Stat()
	rangeStart := fileInfo.Size()
	defaultRangeRound := int64(100000)
	//contentAllLen := int64(1)
	for {
		if readCloser, contentAllLen_, fetchLen := fetchBytes(url, strconv.FormatInt(rangeStart,10)+"-"+strconv.FormatInt(defaultRangeRound,10)); readCloser != nil {
			rangeStart += fetchLen
			defaultRangeRound += rangeStart
			//contentAllLen = contentAllLen_
			writeFile(file, readCloser)
			if rangeStart >= contentAllLen_{
				break
			}
		}else{
			fmt.Println("handle error!")
			break
		}
	}
	fmt.Println("handle end!")

}

func getFile(fileName string) *os.File {
	//如果文件不存在则创建
	if file, err := os.OpenFile(fileName,  os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModeType); err!=nil{
		fmt.Println(err)
	}else{
		return file
	}
	return nil
}

func writeFile(file *os.File, readCloser io.ReadCloser) int {
	if readCloser != nil {
		defer readCloser.Close()
		if b, e := ioutil.ReadAll(readCloser); e == nil {
			file.Write(b)
			file.Sync()
		}
	}
	return 0
}

func fetchBytes(url, rangeRound string) (io.ReadCloser, int64, int64) {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Range", "bytes="+rangeRound)
	response, _ := http.DefaultClient.Do(request)

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	if response.StatusCode == 206 {
		contentRange := response.Header.Get("Content-Range")
		contentAllLen, _ := strconv.ParseInt(strings.Split(contentRange, "/")[1], 10, 64)
		fmt.Println("server file accept range feature!!!")
		if response.ContentLength > 0 {
			return response.Body, contentAllLen, response.ContentLength
		}
	}
	return nil, 0, 0
}

func main() {
	Download()
}