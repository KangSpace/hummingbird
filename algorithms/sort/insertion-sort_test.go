package sort

import (
	"log"
	"testing"
)

//
//插入排序-正序
//author: kango2gler@gmail.com
//
func TestInsertionSort(t *testing.T) {
	log.Println("=========")
	a := []int{5, 2, 4, 6, 1, 3}
	a = InsertionSort(a)
	log.Println(a)
}

//
//插入排序-倒序
//author: kango2gler@gmail.com
//
func TestInsertionSortDesc(t *testing.T) {
	log.Println("=========")
	a := []int{5, 2, 4, 6, 1, 3}
	a = InsertionSortDesc(a)
	log.Println(a)
}
