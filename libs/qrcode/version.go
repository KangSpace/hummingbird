//
// author: Kang
// date:  2017-05-25
package qrcode

import "math"

/*
QR码现有40个标准版本，4个微型版本。QR码的数据编码方式有四种：
	数字（Numeric）：0-9
	大写字母和数字（alphanumeric）：0-9，A-Z，空格，$，%，*，+，-，.，/，:
	二进制/字节：通过 ISO/IEC 8859-1 标准编码
	日本汉字/假名：通过 Shift JISJIS X 0208 标准编码
 QR码还有四种容错级别可以选择：
	L（Low）：7%的字码可被修正
	M（Medium）：15%的字码可被修正
	Q（Quartile）：25%的字码可被修正
	H（High）：30%的字码可被修正
	(Wikipedia: QR code, https://en.wikipedia.org/wiki/QR_code)
  Version 1是21 x 21的矩阵，Version 2是 25 x 25的矩阵，Version 3是29的尺寸，每增加一个version，就会增加4的尺寸，
  公式是：(V-1)*4 + 21（V是版本号） 最高Version 40，(40-1)*4+21 = 177，所以最高是177 x 177 的正方形。
*/
type Version struct {
	//版本号,1-40
	Version byte
	//容错级别
	Level ErrorCorrectionLevel

	//每个纠错码块纠错码字数
	NumberOfErrorCorrectionCodewordsPreBlock byte

	//纠错码块(number of block)数量组1,如 version 5-Q 的 2(group1) 2(group2)
	NumberOfErrorCorrectionBlocksInGroup1 byte
	//每个纠错码块数据码字数 组1 (All codewords shall be 8 bits in length)
	NumberOfDataCodewordsPerBlockInGroup1 byte

	//纠错码块(number of block)数量 组2
	NumberOfErrorCorrectionBlocksInGroup2 byte
	//每个纠错码块数据码字数 组2
	NumberOfDataCodewordsPerBlockInGroup2 byte

	//数据码字数量(非总码字Not Total Number of codewords)
	//数据码字数量 = (纠错码块数量组1 * 每个纠错码块数据码字数组1) + (纠错码块数量组2 * 每个纠错码块数据码字数组2)
	//NumberOfDataCodewords byte
	//总数据码数
	//TotalNumberOfCodewords byte
}

//数据码字数量
func (v Version) getNumberOfDataCodewords() int {
	//数据码字数量 = (纠错码块数量组1 * 每个纠错码块数据码字数组1) + (纠错码块数量组2 * 每个纠错码块数据码字数组2)
	group1Data := int(v.NumberOfDataCodewordsPerBlockInGroup1) * int(v.NumberOfErrorCorrectionBlocksInGroup1)
	group2Data := int(v.NumberOfDataCodewordsPerBlockInGroup2) * int(v.NumberOfErrorCorrectionBlocksInGroup2)
	return (group1Data + group2Data)
}

//容错级别
type ErrorCorrectionLevel byte

const (
	//Recovery Capacity 7%
	L ErrorCorrectionLevel = iota
	//Recovery Capacity 15%
	M
	//Recovery Capacity 25%
	Q
	//Recovery Capacity 30%
	H
)

var allVersions = []*Version{
	&Version{1, L, 7, 1, 19, 0, 0},
	&Version{1, M, 10, 1, 16, 0, 0},
	&Version{1, Q, 13, 1, 13, 0, 0},
	&Version{1, H, 17, 1, 9, 0, 0},
	&Version{2, L, 10, 1, 34, 0, 0},
	&Version{2, M, 16, 1, 28, 0, 0},
	&Version{2, Q, 22, 1, 22, 0, 0},
	&Version{2, H, 28, 1, 16, 0, 0},
	&Version{3, L, 15, 1, 55, 0, 0},
	&Version{3, M, 26, 1, 44, 0, 0},
	&Version{3, Q, 18, 2, 17, 0, 0},
	&Version{3, H, 22, 2, 13, 0, 0},
	&Version{4, L, 20, 1, 80, 0, 0},
	&Version{4, M, 18, 2, 32, 0, 0},
	&Version{4, Q, 26, 2, 24, 0, 0},
	&Version{4, H, 16, 4, 9, 0, 0},
	&Version{5, L, 26, 1, 108, 0, 0},
	&Version{5, M, 24, 2, 43, 0, 0},
	&Version{5, Q, 18, 2, 15, 2, 16},
	&Version{5, H, 22, 2, 11, 2, 12},
	&Version{6, L, 18, 2, 68, 0, 0},
	&Version{6, M, 16, 4, 27, 0, 0},
	&Version{6, Q, 24, 4, 19, 0, 0},
	&Version{6, H, 28, 4, 15, 0, 0},
	&Version{7, L, 20, 2, 78, 0, 0},
	&Version{7, M, 18, 4, 31, 0, 0},
	&Version{7, Q, 18, 2, 14, 4, 15},
	&Version{7, H, 26, 4, 13, 1, 14},
	&Version{8, L, 24, 2, 97, 0, 0},
	&Version{8, M, 22, 2, 38, 2, 39},
	&Version{8, Q, 22, 4, 18, 2, 19},
	&Version{8, H, 26, 4, 14, 2, 15},
	&Version{9, L, 30, 2, 116, 0, 0},
	&Version{9, M, 22, 3, 36, 2, 37},
	&Version{9, Q, 20, 4, 16, 4, 17},
	&Version{9, H, 24, 4, 12, 4, 13},
	&Version{10, L, 18, 2, 68, 2, 69},
	&Version{10, M, 26, 4, 43, 1, 44},
	&Version{10, Q, 24, 6, 19, 2, 20},
	&Version{10, H, 28, 6, 15, 2, 16},
	&Version{11, L, 20, 4, 81, 0, 0},
	&Version{11, M, 30, 1, 50, 4, 51},
	&Version{11, Q, 28, 4, 22, 4, 23},
	&Version{11, H, 24, 3, 12, 8, 13},
	&Version{12, L, 24, 2, 92, 2, 93},
	&Version{12, M, 22, 6, 36, 2, 37},
	&Version{12, Q, 26, 4, 20, 6, 21},
	&Version{12, H, 28, 7, 14, 4, 15},
	&Version{13, L, 26, 4, 107, 0, 0},
	&Version{13, M, 22, 8, 37, 1, 38},
	&Version{13, Q, 24, 8, 20, 4, 21},
	&Version{13, H, 22, 12, 11, 4, 12},
	&Version{14, L, 30, 3, 115, 1, 116},
	&Version{14, M, 24, 4, 40, 5, 41},
	&Version{14, Q, 20, 11, 16, 5, 17},
	&Version{14, H, 24, 11, 12, 5, 13},
	&Version{15, L, 22, 5, 87, 1, 88},
	&Version{15, M, 24, 5, 41, 5, 42},
	&Version{15, Q, 30, 5, 24, 7, 25},
	&Version{15, H, 24, 11, 12, 7, 13},
	&Version{16, L, 24, 5, 98, 1, 99},
	&Version{16, M, 28, 7, 45, 3, 46},
	&Version{16, Q, 24, 15, 19, 2, 20},
	&Version{16, H, 30, 3, 15, 13, 16},
	&Version{17, L, 28, 1, 107, 5, 108},
	&Version{17, M, 28, 10, 46, 1, 47},
	&Version{17, Q, 28, 1, 22, 15, 23},
	&Version{17, H, 28, 2, 14, 17, 15},
	&Version{18, L, 30, 5, 120, 1, 121},
	&Version{18, M, 26, 9, 43, 4, 44},
	&Version{18, Q, 28, 17, 22, 1, 23},
	&Version{18, H, 28, 2, 14, 19, 15},
	&Version{19, L, 28, 3, 113, 4, 114},
	&Version{19, M, 26, 3, 44, 11, 45},
	&Version{19, Q, 26, 17, 21, 4, 22},
	&Version{19, H, 26, 9, 13, 16, 14},
	&Version{20, L, 28, 3, 107, 5, 108},
	&Version{20, M, 26, 3, 41, 13, 42},
	&Version{20, Q, 30, 15, 24, 5, 25},
	&Version{20, H, 28, 15, 15, 10, 16},
	&Version{21, L, 28, 4, 116, 4, 117},
	&Version{21, M, 26, 17, 42, 0, 0},
	&Version{21, Q, 28, 17, 22, 6, 23},
	&Version{21, H, 30, 19, 16, 6, 17},
	&Version{22, L, 28, 2, 111, 7, 112},
	&Version{22, M, 28, 17, 46, 0, 0},
	&Version{22, Q, 30, 7, 24, 16, 25},
	&Version{22, H, 24, 34, 13, 0, 0},
	&Version{23, L, 30, 4, 121, 5, 122},
	&Version{23, M, 28, 4, 47, 14, 48},
	&Version{23, Q, 30, 11, 24, 14, 25},
	&Version{23, H, 30, 16, 15, 14, 16},
	&Version{24, L, 30, 6, 117, 4, 118},
	&Version{24, M, 28, 6, 45, 14, 46},
	&Version{24, Q, 30, 11, 24, 16, 25},
	&Version{24, H, 30, 30, 16, 2, 17},
	&Version{25, L, 26, 8, 106, 4, 107},
	&Version{25, M, 28, 8, 47, 13, 48},
	&Version{25, Q, 30, 7, 24, 22, 25},
	&Version{25, H, 30, 22, 15, 13, 16},
	&Version{26, L, 28, 10, 114, 2, 115},
	&Version{26, M, 28, 19, 46, 4, 47},
	&Version{26, Q, 28, 28, 22, 6, 23},
	&Version{26, H, 30, 33, 16, 4, 17},
	&Version{27, L, 30, 8, 122, 4, 123},
	&Version{27, M, 28, 22, 45, 3, 46},
	&Version{27, Q, 30, 8, 23, 26, 24},
	&Version{27, H, 30, 12, 15, 28, 16},
	&Version{28, L, 30, 3, 117, 10, 118},
	&Version{28, M, 28, 3, 45, 23, 46},
	&Version{28, Q, 30, 4, 24, 31, 25},
	&Version{28, H, 30, 11, 15, 31, 16},
	&Version{29, L, 30, 7, 116, 7, 117},
	&Version{29, M, 28, 21, 45, 7, 46},
	&Version{29, Q, 30, 1, 23, 37, 24},
	&Version{29, H, 30, 19, 15, 26, 16},
	&Version{30, L, 30, 5, 115, 10, 116},
	&Version{30, M, 28, 19, 47, 10, 48},
	&Version{30, Q, 30, 15, 24, 25, 25},
	&Version{30, H, 30, 23, 15, 25, 16},
	&Version{31, L, 30, 13, 115, 3, 116},
	&Version{31, M, 28, 2, 46, 29, 47},
	&Version{31, Q, 30, 42, 24, 1, 25},
	&Version{31, H, 30, 23, 15, 28, 16},
	&Version{32, L, 30, 17, 115, 0, 0},
	&Version{32, M, 28, 10, 46, 23, 47},
	&Version{32, Q, 30, 10, 24, 35, 25},
	&Version{32, H, 30, 19, 15, 35, 16},
	&Version{33, L, 30, 17, 115, 1, 116},
	&Version{33, M, 28, 14, 46, 21, 47},
	&Version{33, Q, 30, 29, 24, 19, 25},
	&Version{33, H, 30, 11, 15, 46, 16},
	&Version{34, L, 30, 13, 115, 6, 116},
	&Version{34, M, 28, 14, 46, 23, 47},
	&Version{34, Q, 30, 44, 24, 7, 25},
	&Version{34, H, 30, 59, 16, 1, 17},
	&Version{35, L, 30, 12, 121, 7, 122},
	&Version{35, M, 28, 12, 47, 26, 48},
	&Version{35, Q, 30, 39, 24, 14, 25},
	&Version{35, H, 30, 22, 15, 41, 16},
	&Version{36, L, 30, 6, 121, 14, 122},
	&Version{36, M, 28, 6, 47, 34, 48},
	&Version{36, Q, 30, 46, 24, 10, 25},
	&Version{36, H, 30, 2, 15, 64, 16},
	&Version{37, L, 30, 17, 122, 4, 123},
	&Version{37, M, 28, 29, 46, 14, 47},
	&Version{37, Q, 30, 49, 24, 10, 25},
	&Version{37, H, 30, 24, 15, 46, 16},
	&Version{38, L, 30, 4, 122, 18, 123},
	&Version{38, M, 28, 13, 46, 32, 47},
	&Version{38, Q, 30, 48, 24, 14, 25},
	&Version{38, H, 30, 42, 15, 32, 16},
	&Version{39, L, 30, 20, 117, 4, 118},
	&Version{39, M, 28, 40, 47, 7, 48},
	&Version{39, Q, 30, 43, 24, 22, 25},
	&Version{39, H, 30, 10, 15, 67, 16},
	&Version{40, L, 30, 19, 118, 6, 119},
	&Version{40, M, 28, 18, 47, 31, 48},
	&Version{40, Q, 30, 34, 24, 34, 25},
	&Version{40, H, 30, 20, 15, 61, 16},
}

//获取适应数据的最小版本
//dataBitsCount (除4+C 外的数据码字数量)
//level 错误纠正级别
//encoderType 编码类型
func getSmallestVersion(dataBitsCount int, level ErrorCorrectionLevel, encoderType EncoderType) *Version {
	dataBitsCount += int(ModeIndicatorBitLen) // + 4 number of bits in mode indicator
	for _, v := range allVersions {
		if v.Level == level {
			//c: 字符计数指示符中的位数（表3）
			c := int(v.getNumberOfBitsInCharCountIndicator(encoderType))
			//每个codeword 是8 bit
			if v.getNumberOfDataCodewords()*8 >= (dataBitsCount + c) {
				return v
			}
		}
	}
	return nil
}

//获取C:字符计数指示符中的位数
func (v *Version) getNumberOfBitsInCharCountIndicator(e EncoderType) byte {
	switch e {
	case Numeric:
		if v.Version < 10 {
			return NumericModel_1_to_9
		}
		if v.Version < 27 {
			return NumericModel_10_to_26
		}
		return NumericModel_27_to_40
	case Alphanumeric:
		if v.Version < 10 {
			return AlphanumericMode_1_to_9
		}
		if v.Version < 27 {
			return AlphanumericMode_10_to_26
		}
		return AlphanumericMode_27_to_40
	case EightBitByte:
		if v.Version < 10 {
			return EightBitByteMode_1_to_9
		}
		if v.Version < 27 {
			return EightBitByteMode_10_to_26
		}
		return EightBitByteMode_27_to_40
	case KanjiByte:
		if v.Version < 10 {
			return KanjiMode_1_to_9
		}
		if v.Version < 27 {
			return KanjiMode_10_to_26
		}
		return KanjiMode_27_to_40
	default:
		return 0
	}
}

func (vi *Version) modulWidth() int {
	return ((int(vi.Version) - 1) * 4) + 21
}

func (vi *Version) alignmentPatternPlacements() []int {
	if vi.Version == 1 {
		return make([]int, 0)
	}

	first := 6
	last := vi.modulWidth() - 7
	space := float64(last - first)
	count := int(math.Ceil(space/28)) + 1

	result := make([]int, count)
	result[0] = first
	result[len(result)-1] = last
	if count > 2 {
		step := int(math.Ceil(float64(last-first) / float64(count-1)))
		if step%2 == 1 {
			frac := float64(last-first) / float64(count-1)
			_, x := math.Modf(frac)
			if x >= 0.5 {
				frac = math.Ceil(frac)
			} else {
				frac = math.Floor(frac)
			}

			if int(frac)%2 == 0 {
				step--
			} else {
				step++
			}
		}
		for i := 1; i <= count-2; i++ {
			result[i] = last - (step * (count - 1 - i))
		}
	}

	return result
}
