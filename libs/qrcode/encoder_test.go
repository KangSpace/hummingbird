//
// author: Kang
// date:
package qrcode

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"
)

func ExampleNumericEncode() {
	f, _ := os.Create("D:/1/2/3/qrcode_numeric.png")
	defer f.Close()

	qrcode, err := Encode("111111111111111111", L, Numeric)
	if err != nil {
		fmt.Println(err)
	} else {
		qrcode, _, err = Scale(qrcode, 300, 300, false)
		if err == nil {
			png.Encode(f, qrcode)
		} else {
			fmt.Println(err)
		}

	}
}

func ExampleAlphaNumericEncode() {
	f, _ := os.Create("D:/1/2/3/qrcode_alphanumeric.png")
	defer f.Close()

	qrcode, err := Encode("123ABDCD", L, Alphanumeric)
	if err != nil {
		fmt.Println(err)
	} else {
		qrcode, _, err = Scale(qrcode, 300, 300, false)
		if err == nil {
			png.Encode(f, qrcode)
		} else {
			fmt.Println(err)
		}

	}
}

func ExampleUnicodeEncodeEncode() {
	f, _ := os.Create("D:/1/2/3/qrcode_unicode.png")
	defer f.Close()

	qrcode, err := Encode("http:你好我，中文www.20.com/?ABCDabcd!@#$%^&*()_+=-|[]{};:'\",.<>/\\`~*", L, Unicode)
	if err != nil {
		fmt.Println(err)
	} else {
		qrcode, _, err = Scale(qrcode, 300, 300, false)
		if err == nil {
			png.Encode(f, qrcode)
		} else {
			fmt.Println(err)
		}

	}
}

// 将二维码中间涂黑
func drowBlackToCenter() {
	f, err := os.OpenFile("D:/1/2/3/qrcode_unicode.png", os.O_RDWR, 0666)
	f2, err := os.OpenFile("D:/1/2/3/2.png", os.O_CREATE|os.O_RDWR, 0666)
	defer f.Close()
	defer f2.Close()
	if err != nil {
		log.Fatal(err)
	}

	var img image.Image
	if img, err = png.Decode(f); err == nil {
		//创建一个图像
		m := image.NewRGBA(image.Rect(0, 0, 300, 300)) //*NRGBA (image.Image interface)
		// 填充白色,并把其写入到m
		draw.Draw(m, m.Bounds(), img, image.ZP, draw.Src)
		draw.Draw(m, image.Rect(100, 100, 150, 150), image.NewUniform(color.Black), image.Point{100, 100}, draw.Src)
		if err := png.Encode(f2, m); err == nil {
			fmt.Println("save successfully")
		} else {
			fmt.Println("2:", err)
		}
	} else {
		fmt.Println("3:", err)
	}

}

func TEST_EncodeToPng() {
	var qrCode = &QRCode{300, "abc", nil}
	f, err := os.OpenFile("D:/1/2/3/qrcode_abc.png", os.O_CREATE|os.O_RDWR, 0666)
	if err == nil {
		if err = EncodeToPng(qrCode, f); err == nil {
			fmt.Println("save successfull")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

//二维码测试
func Test_Qncode(t *testing.T) {
	//ExampleUnicodeEncodeEncode()
	//drowBlackToCenter()
	TEST_EncodeToPng()
}
