// 常量、变量
// author: Kang
// date:  2017-05-26
package qrcode

import "errors"

const (
	//模式指示符位长
	ModeIndicatorBitLen byte = 4
)

//编码类型常量,值为模式指示器 10进制值
const (
	//digits 0-9 0001 1
	Numeric EncoderType = 1 << iota
	//digits 0-9; upper case letters A-Z; nine other characters:space, $ % * + - . / :
	// 0010 2
	Alphanumeric
	// (JIS 8-bit character set (Latin and Kana) in accordance with JIS X 0201
	// 0100 4
	EightBitByte
	//汉字字符（根据JIS X 0208附件1 Shift编码的Shift JIS字符集 表示。请注意，QR码中的汉字字符可以有值8140HEX -9FFCHEX和E040HEX
	// -EBBFHEX，可压缩成13位）
	// 1000 8
	KanjiByte
	//同 EightBitByte
	Unicode EncoderType = 4

	// 0111
	ECI EncoderType = 7
	// 0011
	StructureAppend EncoderType = 3
	// 0101
	FNC1 EncoderType = 5
	// 1001
	FNC1_2P EncoderType = 9
	// 0000
	Terminator = 0
	//未实现的其他类型
	Unknow
)

//编码类型,正则
var (
	NumericRegExpType      = "^[0-9]+$"
	AlphaNumericRegExpType = "^[0-9A-Z]+$"
)

//模式指示器 ModeIndicator
//四位标识符指示下一个数据序列被编码的模式
/*
type ModeIndicator byte
const (
	//Extended Channel Interpretation (ECI) Mode
	//扩展通道解释
	ECIModeIndictor ModeIndicator = "0111"
	NumericModeIndictor ModeIndicator = "0001"
	AlphanumericModeIndictor ModeIndicator = "0010"
	EightBitByteModeIndictor ModeIndicator = "0100"
	KanjiModeIndictor ModeIndicator = "1000"
	StructuredAppendModeIndictor ModeIndicator = "0011"
	FNC1ModeIndictor ModeIndicator = "0101"
	FNC1_2ModeIndictor ModeIndicator = "1001"
	//End Of Message
	TerminatorModeIndictor ModeIndicator = "0000"
)
*/
//Representation of data
const (
	//亮块 0
	LIGHT byte = iota
	//黑块 1
	DARK
)

//
//字符技术指示器在一个模式下定义数据串长度的位序列
var (
	NumericModel_1_to_9   byte = 10
	NumericModel_10_to_26 byte = 12
	NumericModel_27_to_40 byte = 14

	AlphanumericMode_1_to_9   byte = 9
	AlphanumericMode_10_to_26 byte = 11
	AlphanumericMode_27_to_40 byte = 13

	EightBitByteMode_1_to_9   byte = 8
	EightBitByteMode_10_to_26 byte = 16
	EightBitByteMode_27_to_40 byte = 16

	KanjiMode_1_to_9   byte = 8
	KanjiMode_10_to_26 byte = 10
	KanjiMode_27_to_40 byte = 12
)

//错误定义
var (
	//未知字符集错误
	UnknowCharacterSetError error = errors.New("Unknow character set to encode!")
	TooMuchCharactersError  error = errors.New("Too much characters to encode!")
	//编码错误
	EncodeError error = errors.New("Encode error!")
	//输入参数为空错误
	ArguementsNullError error = errors.New("arguements is null!")
	//输出png错误
	PNGEncodeError error = errors.New("PNG encode error!")
)
