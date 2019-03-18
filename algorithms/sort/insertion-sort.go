package sort

import "log"

//
//插入排序-正序
//author: kango2gler@gmail.com
//
func InsertionSort(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for ; i >= 0 && a[i] > key; i-- {
			a[i+1] = a[i]
		}
		a[i+1] = key
		log.Println(a)
	}
	return a
}

//
//插入排序-倒序
//author: kango2gler@gmail.com
//
func InsertionSortDesc(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for ; i >= 0 && a[i] < key; i-- {
			a[i+1] = a[i]
		}
		a[i+1] = key
		log.Println(a)
	}
	return a
}
