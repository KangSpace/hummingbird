/*
 date: 2019-04-28 18:48:20
A format for expressing an ordered list of integers is to use a comma separated list of either

individual integers
or a range of integers denoted by the starting integer separated from the end integer in the range by a dash, '-'.
The range includes all integers in the interval including both endpoints. It is not considered a range unless it spans at least 3 numbers. For example ("12, 13, 15-17")
Complete the solution so that it takes a list of integers in increasing order and returns a correctly formatted string in the range format.

Example:

solution([-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]);
// returns "-6,-3-1,3-5,7-11,14,15,17-20"
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func Solution1(list []int) string {
	sort.Ints(list)
	arr := ""
	for x, temp, count, lastNum := 0, 0, 0, 0; x < len(list); x, lastNum, count = x+1, temp, count+1 {
		temp = list[x]
		if x != 0 && temp-lastNum == 1 {
			if arr[len(arr)-1:] != "-" {
				arr = arr[:len(arr)-1] + "-"
			}
			if (x < len(list)-1 && list[x+1]-temp != 1) || x == len(list)-1 {
				if count != 0 && count < 2 {
					arr = arr[:len(arr)-1] + ","
				}
				arr += strconv.Itoa(temp) + ","
				count = 0
			}
		} else {
			arr += strconv.Itoa(temp) + ","
			count = 0
		}
	}
	if len(arr) > 0 {
		return arr[:len(arr)-1]
	}
	return arr
}
func Solution(list []int) (s string) {
	l := len(list) - 1
	for i, j := 0, 0; i < l; i = j {
		s += strconv.Itoa(list[i])
		for j = i; (j < l) && (list[j]+1 == list[j+1]); {
			j++
		}
		if j-i > 1 {
			s += "-"
		} else {
			s += ","
		}
		if i == j {
			j++
		}
	}
	s += strconv.Itoa(list[l])
	return
}
func main() {
	//-6,-3-1,3-5,7-11,14,15,17-20
	fmt.Println(Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
	//"40,44,48,51,52,54,55,58,67,73"
	fmt.Println(Solution([]int{40, 44, 48, 51, 52, 54, 55, 58, 67, 73}))
	//"-25--20,-18,-13,-5--3,2,3"
	fmt.Println(Solution([]int{-25, -24, -23, -22, -21, -20, -18, -13, -5, -4, -3, 2, 3}))

}
