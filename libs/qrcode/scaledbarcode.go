package qrcode

import (
	"errors"
	"fmt"
	"image"
	"image/color"
)

type wrapFunc func(x, y int) color.Color

type scaledBarCode struct {
	wrapped     BarCode
	wrapperFunc wrapFunc
	rect        image.Rectangle
}

//目标图片大小
type DestImgSize struct {
	width  int
	height int
}

type intCSscaledBC struct {
	scaledBarCode
}

func (bc *scaledBarCode) Content() string {
	return bc.wrapped.Content()
}

func (bc *scaledBarCode) Metadata() Metadata {
	return bc.wrapped.Metadata()
}

func (bc *scaledBarCode) ColorModel() color.Model {
	return bc.wrapped.ColorModel()
}

func (bc *scaledBarCode) Bounds() image.Rectangle {
	return bc.rect
}

func (bc *scaledBarCode) At(x, y int) color.Color {
	return bc.wrapperFunc(x, y)
}

func (bc *intCSscaledBC) CheckSum() int {
	if cs, ok := bc.wrapped.(BarCodeIntCS); ok {
		return cs.CheckSum()
	}
	return 0
}

// Scale returns a resized barcode with the given width and height.
// DestImgSize: 生成二维码后宽高
func Scale(bc BarCode, width, height int, isFill bool) (BarCode, *DestImgSize, error) {
	switch bc.Metadata().Dimensions {
	case 1:
		return scale1DCode(bc, width, height)
	case 2:
		return scale2DCode(bc, width, height, isFill)
	}

	return nil, nil, errors.New("unsupported barcode format")
}

func newScaledBC(wrapped BarCode, wrapperFunc wrapFunc, rect image.Rectangle) BarCode {
	result := &scaledBarCode{
		wrapped:     wrapped,
		wrapperFunc: wrapperFunc,
		rect:        rect,
	}

	if _, ok := wrapped.(BarCodeIntCS); ok {
		return &intCSscaledBC{*result}
	}
	return result
}

//缩放二维码，返回输入的width宽度数
//isFill 是否填充二维码边框
func scale2DCode(bc BarCode, width, height int, isFill bool) (BarCode, *DestImgSize, error) {
	orgBounds := bc.Bounds()
	orgWidth := orgBounds.Max.X - orgBounds.Min.X
	orgHeight := orgBounds.Max.Y - orgBounds.Min.Y
	fmt.Println("scale2DCode orginBounds:", orgBounds, " ,width:", width, " ,height:", height)
	modSeed := width / orgWidth
	//若除不尽,则将width,height补充除尽  && float64(mod)/float64(orgWidth)>=0.5
	//有边框时,根据输入的width,height获取最大缩放范围的二维码,若有多余,则补充为边框
	//无边框时,返回最大课缩放范围尺寸
	if mod := width % orgWidth; mod != 0 {
		if !isFill {
			width = orgWidth * modSeed
		}
	}
	height = width
	factor := modSeed
	fmt.Println("factor:", factor, " ,width:", width, " ,height:", height, " ,orgWidth:", orgWidth, " ,orgHeight:", orgHeight)
	if factor <= 0 {
		return nil, nil, fmt.Errorf("can not scale barcode to an image smaller than %dx%d", orgWidth, orgHeight)
	}

	offsetX := (width - (orgWidth * factor)) / 2
	offsetY := (height - (orgHeight * factor)) / 2

	wrap := func(x, y int) color.Color {
		if x < offsetX || y < offsetY {
			return color.White
		}
		x = (x - offsetX) / factor
		y = (y - offsetY) / factor
		if x >= orgWidth || y >= orgHeight {
			return color.White
		}
		return bc.At(x, y)
	}
	return newScaledBC(
		bc,
		wrap,
		image.Rect(0, 0, width, height),
	), &DestImgSize{width, height}, nil
}

func scale1DCode(bc BarCode, width, height int) (BarCode, *DestImgSize, error) {
	orgBounds := bc.Bounds()
	orgWidth := orgBounds.Max.X - orgBounds.Min.X
	factor := int(float64(width) / float64(orgWidth))

	if factor <= 0 {
		return nil, nil, fmt.Errorf("can not scale barcode to an image smaller than %dx1", orgWidth)
	}
	offsetX := (width - (orgWidth * factor)) / 2

	wrap := func(x, y int) color.Color {
		if x < offsetX {
			return color.White
		}
		x = (x - offsetX) / factor

		if x >= orgWidth {
			return color.White
		}
		return bc.At(x, 0)
	}

	return newScaledBC(
		bc,
		wrap,
		image.Rect(0, 0, width, height),
	), nil, nil
}
