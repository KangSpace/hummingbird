package qrcode

//数字编码器
type UnicodeEncoder struct {
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
func (encoder UnicodeEncoder) Encode(content string, ecl ErrorCorrectionLevel) (*BitList, *Version, error) {
	return EightBitByteEncoder{}.Encode(content, ecl)
}
