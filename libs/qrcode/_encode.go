/*
 QR二维码生成器
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
  author: Kang
  date: 2017-05-24
 *
*/
package qrcode

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
)

var (
	//Version 1 : 21 x 21
	QR_VERSION_1 Version = Version{1, 21}
	//Version 2 : 25 x 25
	QR_VERSION_2 Version = Version{1, 25}
	//(V-1)*4 + 21（V是版本号） 最高Version 40
	//Version 40 : 177 x 177 最多可以处理7089位数字
	QR_VERSION_40 Version = Version{1, 177}
)

const (
	BYTE_MODE_PREFIX int = 0x04
)

//QR 版本号
type Version struct {
	//版本号
	number int
	//尺寸
	size int
}

//图片文件
type PicFile struct {
	filepath string
	width    int
	height   int
	file     *os.File
}

func genversion() {

}

/**
 * QR二维码生成流程
 */
func encoderFlow() {
	//1. 数据分析（data analysis）:分析输入数据，根据数据决定要使用的QR码版本、容错级别和编码模式。低版本的QR码无法编码过长的数据，含有非数字字母字符的数据要使用扩展字符编码模式
	//2. 编码数据（data encoding） :根据选择的编码模式，将输入的字符串转换成比特流，插入模式标识码（mode indicator）和终止标识符（terminator），把比特流切分成八比特的字节，加入填充字节来满足标准的数据字码数要求
	//3. 计算容错码（error correction coding）:对步骤二产生的比特流计算容错码，附在比特流之后。高版本的编码方式可能需要将数据流切分成块（block）再分别进行容错码计算。
	//4. 组织数据（structure final message）:根据结构图把步骤三得到的有容错的数据流切分，准备填充。
	//5. 填充（module placement in matrix）:把数据和功能性图样根据标准填充到矩阵中。
	//6. 应用数据掩码（data masking）：应用标准中的八个数据掩码来变换编码区域的数据，选择最优的掩码应用。讲到再展开。
	//7. 填充格式和版本信息（format and version information）：计算格式和版本信息填入矩阵，完成QR码。
}

//初始化QR二维码矩阵
//qrBolckLen :qr version block count
//width :picture width
//height :picture height, equal with width
//w :输出流
func drowInitImage_(qrBlockLen int, width int, height int, w io.Writer) error {
	// 创建一个图像
	m := image.NewRGBA(image.Rect(0, 0, width, width)) //*NRGBA (image.Image interface)
	// 填充白色背景,并把其写入到m
	draw.Draw(m, m.Bounds(), image.NewUniform(color.White), image.ZP, draw.Src)
	// 像素宽度,图片宽/QR矩阵块数
	var pixelWidth = width / qrBlockLen
	//画行
	for i, i_ := 0, 0; i < width; i = i_ {
		i_ = i + pixelWidth
		//画列
		for j, j_ := 0, 0; j < width; j = j_ {
			j_ = j + pixelWidth
			if (i_+j_)/pixelWidth%2 == 0 {
				// 填充黑色,并把其写入到m
				draw.Draw(m, image.Rect(i, j, i_, j_), image.NewUniform(color.Black), image.Point{i, j}, draw.Src)
				// log.Println(i, " ", j, " ", i_, " ", j_, " ")
			}
		}
	}

	//Encode writes the Image m to w in JPEG format.
	return jpeg.Encode(w, m, nil)
}

//将图写到文件中
func drowImgToFile(file PicFile) error {
	return drowInitImage_(8, file.width, file.height, file.file)
}

// 1. 数据分析（data analysis）
func analysisData(string string) Version {
	switch {
	case len([]rune(string)) <= 17:
		return QR_VERSION_1
	default:
		return QR_VERSION_1
	}
}

// 2. 编码数据（data encoding）
func endcodeData() {
	//byte mode
	//	byeStr := "0100"

}

// 3. 计算容错码（error correction coding）
func codingErrorCorrenctionData() {

}

// 5. 填充数据（module placement in matrix）
func fillData() {

}

// 6. 应用数据掩码（data masking）
func maskData() {

}

// 7. 填充格式和版本信息（format and version information）
func fillFmtVersionData() {

}

func main() {
	//drowTest_()
	fmt.Println(BYTE_MODE_PREFIX)
}

// 输出图片测试
func drowTest_() {
	var file_ = PicFile{filepath: "d:/1/2/3/qr2.jpg", width: 240, height: 240}
	var err error
	if file_.file, err = os.OpenFile(file_.filepath, os.O_CREATE, 0666); err != nil {
		log.Fatal("文件 ", file_.filepath, " 创建失败")
	}
	defer (*file_.file).Close()
	if err = drowImgToFile(file_); err != nil {
		log.Fatal("二维码 ", file_.filepath, " 绘制失败")
	}
	log.Println("二维码 ", file_.filepath, " 绘制成功")
}
