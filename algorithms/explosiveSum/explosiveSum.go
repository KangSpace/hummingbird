/*
 date: 2019-04-26 17:48:50
How many ways can you make the sum of a number?
From wikipedia: https://en.wikipedia.org/wiki/Partition_(number_theory)#

In number theory and combinatorics, a partition of a positive integer n, also called an integer partition, is a way of writing n as a sum of positive integers. Two sums that differ only in the order of their summands are considered the same partition. If order matters, the sum becomes a composition. For example, 4 can be partitioned in five distinct ways:

4
3 + 1
2 + 2
2 + 1 + 1
1 + 1 + 1 + 1
Examples
Basic
sum(1) // 1
sum(2) // 2  -> 1+1 , 2
sum(3) // 3 -> 1+1+1, 1+2, 3
sum(4) // 5 -> 1+1+1+1, 1+1+2, 1+3, 2+2, 4
sum(5) // 7 -> 1+1+1+1+1, 1+1+1+2, 1+1+3, 1+2+2, 1+4, 5, 2+3

sum(10) // 42
Explosive
sum(50) // 204226
sum(80) // 15796476
sum(100) // 190569292
See here for more examples.
*/
package main

import "fmt"

func Sum(n int) int{
	return len(enums(n))
}

//整数分解
func enums(n int) [][]int {
	enums_ := make([][]int, 0)
	ns := make([]int, n+1)
	ns[1] = n
	k := 1
	for k != 0 {
		x := ns[k-1] +1
		y := ns[k] -1
		k -= 1
		for x <= y {
			ns[k] = x
			y -= x
			k += 1
		}
		ns[k] = x + y
		subEnums := make([]int,0)
		for i :=0 ; i <= k ;i++{
			subEnums = append(subEnums,ns[i])
		}
		enums_ = append(enums_,subEnums)
	}
	return enums_
}
func main() {
	fmt.Println(Sum(1)== 1)
	fmt.Println(Sum(2)== 2)
	fmt.Println(Sum(3)== 3)
	fmt.Println(enums(4))
	fmt.Println(enums(5))
	fmt.Println(Sum(5)== 7)
	fmt.Println(Sum(10)==42)
}
