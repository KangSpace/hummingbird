package qrcode

import (
	"image"
	"image/color"
	"math"
)

type QRCode struct {
	//尺寸
	dimension int
	//内容
	content string
	//数据
	data *BitList
}

func (qr *QRCode) Dimension() int {
	return qr.dimension
}

func (qr *QRCode) Content() string {
	return qr.content
}

func (qr *QRCode) Metadata() Metadata {
	return Metadata{"QR Code", 2}
}

func (qr *QRCode) ColorModel() color.Model {
	return color.Gray16Model
}

func (qr *QRCode) Bounds() image.Rectangle {
	return image.Rect(0, 0, qr.dimension, qr.dimension)
}

func (qr *QRCode) At(x, y int) color.Color {
	if qr.Get(x, y) {
		return color.Black
	}
	return color.White
}

func (qr *QRCode) Get(x, y int) bool {
	return qr.data.GetBit(x*qr.dimension + y)
}

func (qr *QRCode) Set(x, y int, val bool) {
	qr.data.SetBit(x*qr.dimension+y, val)
}

func (qr *QRCode) calcPenalty() uint {
	return qr.calcPenaltyRule1() + qr.calcPenaltyRule2() + qr.calcPenaltyRule3() + qr.calcPenaltyRule4()
}

func (qr *QRCode) calcPenaltyRule1() uint {
	var result uint
	for x := 0; x < qr.dimension; x++ {
		checkForX := false
		var cntX uint
		checkForY := false
		var cntY uint

		for y := 0; y < qr.dimension; y++ {
			if qr.Get(x, y) == checkForX {
				cntX++
			} else {
				checkForX = !checkForX
				if cntX >= 5 {
					result += cntX - 2
				}
				cntX = 1
			}

			if qr.Get(y, x) == checkForY {
				cntY++
			} else {
				checkForY = !checkForY
				if cntY >= 5 {
					result += cntY - 2
				}
				cntY = 1
			}
		}

		if cntX >= 5 {
			result += cntX - 2
		}
		if cntY >= 5 {
			result += cntY - 2
		}
	}

	return result
}

func (qr *QRCode) calcPenaltyRule2() uint {
	var result uint
	for x := 0; x < qr.dimension-1; x++ {
		for y := 0; y < qr.dimension-1; y++ {
			check := qr.Get(x, y)
			if qr.Get(x, y+1) == check && qr.Get(x+1, y) == check && qr.Get(x+1, y+1) == check {
				result += 3
			}
		}
	}
	return result
}

func (qr *QRCode) calcPenaltyRule3() uint {
	pattern1 := []bool{true, false, true, true, true, false, true, false, false, false, false}
	pattern2 := []bool{false, false, false, false, true, false, true, true, true, false, true}

	var result uint
	for x := 0; x <= qr.dimension-len(pattern1); x++ {
		for y := 0; y < qr.dimension; y++ {
			pattern1XFound := true
			pattern2XFound := true
			pattern1YFound := true
			pattern2YFound := true

			for i := 0; i < len(pattern1); i++ {
				iv := qr.Get(x+i, y)
				if iv != pattern1[i] {
					pattern1XFound = false
				}
				if iv != pattern2[i] {
					pattern2XFound = false
				}
				iv = qr.Get(y, x+i)
				if iv != pattern1[i] {
					pattern1YFound = false
				}
				if iv != pattern2[i] {
					pattern2YFound = false
				}
			}
			if pattern1XFound || pattern2XFound {
				result += 40
			}
			if pattern1YFound || pattern2YFound {
				result += 40
			}
		}
	}

	return result
}

func (qr *QRCode) calcPenaltyRule4() uint {
	totalNum := qr.data.Len()
	trueCnt := 0
	for i := 0; i < totalNum; i++ {
		if qr.data.GetBit(i) {
			trueCnt++
		}
	}
	percDark := float64(trueCnt) * 100 / float64(totalNum)
	floor := math.Abs(math.Floor(percDark/5) - 10)
	ceil := math.Abs(math.Ceil(percDark/5) - 10)
	return uint(math.Min(floor, ceil) * 10)
}

func newQrcode(dim int) *QRCode {
	res := new(QRCode)
	res.dimension = dim
	res.data = NewBitList(dim * dim)
	return res
}

func NewQrCode(dimension int, content string) *QRCode {
	return &QRCode{dimension, content, nil}
}
