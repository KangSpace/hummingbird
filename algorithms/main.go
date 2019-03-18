package algorithms

import (
	"fmt"
	"log"
)

const c1 = '1'

type o1 struct {
	name string
	val  interface{}
}

func init() {
	var io1 = o1{name: "o01", val: "o01 val"}
	log.Println(io1)
	log.Println(" init A")
}

func A() {
	log.Println(" test A")
}

func TwoSum(nums []int, target int) []int {
	var result = make([]int, 2)
	fmt.Println("nums=", len(nums))
	for idx, i := range nums {
		for jdx, j := range nums {
			if i+j == target {
				result[0] = idx
				result[1] = jdx
				return result
			}

		}
	}
	return nil
}
