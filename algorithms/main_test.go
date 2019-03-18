package algorithms

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	A()
	t.Fail()
}

func TestTwoSum(t *testing.T) {
	var arr = []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(TwoSum(arr, target))
}
