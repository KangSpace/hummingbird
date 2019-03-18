//
// author: Kang
// date:  2017-05-27
package qrcode

import (
	"fmt"
	"strconv"
)

//digits 0-9
//var characters string = "0123456789"

//数字编码器
type NumericEncoder struct {
}

//数据编码
//1.输入数据串分为三位数组，每组转换为10位二进制数当量。如果输入数字的数量不是三位数，则最后一位或两位数字被转换为4或7位。
//2.将二进制数据连接
//  将模式指示符和前缀字符计数指示器 添加到二进制数据前,数字模式中的字符计数指示符具有10,12或14位
//  如果用户没有指定要使用的符号版本，选择将适应数据的最小版本
//  对于任何数量的数据字符，数字模式中位流的长度由以下公式给出：(4+C 每个mode都需要)
//        B = 4 + C + 10（D DIV 3）+ R
//		B =比特流中的比特数
// 		C =字符计数指示符中的位数（表3）
// 		D =输入数据字符数
// 			（D MOD 3）= 0时，R = 0
// 			如果（D MOD 3）= 1，则R = 4
// 			如果（D MOD 3）= 2，则R = 7
func (encoder NumericEncoder) Encode(content string, ecl ErrorCorrectionLevel) (*BitList, *Version, error) {
	//校验内容是否有效

	//内容长度
	contentLen := len(content)
	//分组bit 数(不含4 + C),分组数x每组10位bit
	contentBitCount := contentLen / 3 * 10
	switch contentLen % 3 {
	case 1:
		contentBitCount += 4
	case 2:
		contentBitCount += 7
	}
	//获取适应数据的最小版本信息
	version := getSmallestVersion(contentBitCount, ecl, Numeric)
	if version == nil {
		return nil, nil, TooMuchCharactersError
	}
	//连接二进制数据, 模式指示符(0001) + 字符计数指示符数量 (各版本n位)+ 数据
	allData := new(BitList)
	//添加模式指示符
	allData.AddBits(int(Numeric), ModeIndicatorBitLen)
	//添加字符计数指示符
	allData.AddBits(contentLen, version.getNumberOfBitsInCharCountIndicator(Numeric))
	//添加数据字符
	//将输入字符分为3组数字组
	for i := 0; i < contentLen; i += 3 {
		//档次循环数字组
		var currGroup string
		//最后一次循环
		if i+3 <= contentLen {
			currGroup = content[i : i+3]
		} else {
			currGroup = content[i:]
		}
		//获取字符串的10进制数
		number, err := strconv.Atoi(currGroup)
		if err != nil {
			return nil, nil, fmt.Errorf("\"%s\" can not be encoded as %s", content, Numeric)
		}
		//获取number的bit count
		var numberOfCurrGroupBit byte
		switch len(currGroup) % 3 {
		case 0:
			numberOfCurrGroupBit = 10
		case 1:
			numberOfCurrGroupBit = 4
		case 2:
			numberOfCurrGroupBit = 7
		}
		allData.AddBits(number, numberOfCurrGroupBit)
	}
	addPaddingAndTerminator(allData, version)
	//添加最终消息
	return allData, version, nil

}
