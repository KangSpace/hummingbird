//
// author: Kang
// date:  2017-05-27
package qrcode

//(JIS 8-bit character set (Latin and Kana) in accordance with JIS X 0201
//var characters string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"

//8位字节数据编码器
type EightBitByteEncoder struct {
}

//编码
func (encoder EightBitByteEncoder) Encode(content string, ecl ErrorCorrectionLevel) (*BitList, *Version, error) {
	data := []byte(content)
	vi := getSmallestVersion(len(data)*8, ecl, EightBitByte)
	if vi == nil {
		return nil, nil, TooMuchCharactersError
	}
	// It's not correct to add the unicode bytes to the result directly but most readers can't handle the
	// required ECI header...
	res := new(BitList)
	res.AddBits(int(EightBitByte), 4)
	res.AddBits(len(content), vi.getNumberOfBitsInCharCountIndicator(EightBitByte))
	for _, b := range data {
		res.AddByte(b)
	}
	addPaddingAndTerminator(res, vi)
	return res, vi, nil
}
