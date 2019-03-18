//
// author: Kang
// date: 2017-05-27
package qrcode

import (
	"fmt"
	"strings"
)

//digits 0-9; upper case letters A-Z; nine other characters:space, $ % * + - . / :
var characters string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"

//字母数字编码器
type AlphanumericEncoder struct {
}

//编码
func (e AlphanumericEncoder) Encode(content string, ecl ErrorCorrectionLevel) (*BitList, *Version, error) {

	contentLenIsOdd := len(content)%2 == 1
	contentBitCount := (len(content) / 2) * 11
	if contentLenIsOdd {
		contentBitCount += 6
	}
	vi := getSmallestVersion(contentBitCount, ecl, Alphanumeric)
	if vi == nil {
		return nil, nil, TooMuchCharactersError
	}

	res := new(BitList)
	res.AddBits(int(Alphanumeric), 4)
	res.AddBits(len(content), vi.getNumberOfBitsInCharCountIndicator(Alphanumeric))

	encoder := stringToAlphaIdx(content)

	for idx := 0; idx < len(content)/2; idx++ {
		c1 := <-encoder
		c2 := <-encoder
		if c1 < 0 || c2 < 0 {
			return nil, nil, fmt.Errorf("\"%s\" can not be encoded as %s", content, Alphanumeric)
		}
		res.AddBits(c1*45+c2, 11)
	}
	if contentLenIsOdd {
		c := <-encoder
		if c < 0 {
			return nil, nil, fmt.Errorf("\"%s\" can not be encoded as %s", content, Alphanumeric)
		}
		res.AddBits(c, 6)
	}
	addPaddingAndTerminator(res, vi)
	return res, vi, nil
}

func stringToAlphaIdx(content string) <-chan int {
	result := make(chan int)
	go func() {
		for _, r := range content {
			idx := strings.IndexRune(characters, r)
			result <- idx
			if idx < 0 {
				break
			}
		}
		close(result)
	}()

	return result
}
