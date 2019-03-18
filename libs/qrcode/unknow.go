//
// author: Kang
// date:  2017-05-27
package qrcode

import "fmt"

//未实现的解码器
//var characters string = ""

//8位字节数据编码器
type UnknowEncoder struct {
}

//编码
func (encoder UnknowEncoder) Encode(content string, ecl ErrorCorrectionLevel) (*BitList, *Version, error) {
	fmt.Println("Unknow character set to encode!")
	return nil, nil, UnknowCharacterSetError
}
