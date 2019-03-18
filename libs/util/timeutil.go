// author: kango2gler@gmail.com
// date: 2018-06-12 23:53:36

// 时间操作工具类
//
// func:
//	timeutil.Cost(func(){ ...})//获取当前执行所花费的时间
package util

import (
	"log"
	"strconv"
	"time"
)

//calculate cost time
func Cost(run func()) time.Duration {
	start := time.Now()
	run()
	cost := time.Since(start)
	log.Printf("cost %s", cost)
	return cost
}

//计算时长
func CostTimeCalc(startNanoTime time.Time, endNanoTime time.Time) string {
	return strconv.FormatFloat(float64(endNanoTime.UnixNano()-startNanoTime.UnixNano())/1000/1000/1000, 'f', 10, 64) + "s"
}
