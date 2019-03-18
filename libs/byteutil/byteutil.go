// byte处理
// author: Kang
// date:
package byteutil

import (
	"encoding/base64"
	"log"
)

//解码base64为字符串
func Base64Decode(base64Str string) string {
	log.Println("base64Decode base64Str:", base64Str)
	if decodeBytes, err := base64.StdEncoding.DecodeString(base64Str); err == nil {
		decodeStr := string(decodeBytes)
		log.Print("解码base64字符串为:", decodeStr)
		return decodeStr
	} else {
		log.Fatal("解码base64字符串错误:", err)
	}
	return ""
}

//byte数组为字符串
func ByteToString(bytes []byte) string {
	return string(bytes)
}
