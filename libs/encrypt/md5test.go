package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
//DO NOT USE THIS , THIS IS ERROR Example
func md5V2(str string) string {
	h := md5.New()
	//ERROR
	return hex.EncodeToString(h.Sum([]byte(str)))
}

func md5V3(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func main() {
	str := "bscdn-c29365958fad1ef2a1afa875a62ff35f/etthd/ysgzzz000080/400.mp41593788484"
	md5Str := md5V(str)
	fmt.Println(md5Str)
	fmt.Println(md5V2(str))
	fmt.Println(md5V3(str))
}
