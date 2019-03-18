package util

import (
	"log"
	"testing"
)

func TestCost(t *testing.T) {

	c := Cost(func() {
		calculateNum()
	})
	log.Println("耗费时间:", c)
}

func calculateNum(args ...interface{}) {
	c := 0
	for i := 1; i <= 10000; i++ {
		c += i
	}
	// time.Sleep(10 * time.Second)
	log.Println(c)
}
