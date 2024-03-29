//图片处理相关包
//
//author: Kang
//date: 2017-05-31
package myimage

import (
	"fmt"
	"image"
	"image/draw"
	"time"

	"code.google.com/p/graphics-go/graphics"
	"hummingbird/libs/util"
)

//等比例缩放
func Scale(dest draw.Image, src image.Image) error {
	startTime := time.Now()
	err := graphics.Scale(dest, src)
	fmt.Println("Scale used:", util.CostTimeCalc(startTime, time.Now()))
	return err
}
