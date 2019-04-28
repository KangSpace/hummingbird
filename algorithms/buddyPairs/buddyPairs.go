/*
 date: 2019-04-28 17:14:23
 Buddy pairs
You know what divisors of a number are. The divisors of a positive integer n are said to be proper when you consider only the divisors other than n itself. In the following description, divisors will mean proper divisors. For example for 100 they are 1, 2, 4, 5, 10, 20, 25, and 50.

Let s(n) be the sum of these proper divisors of n. Call buddy two positive integers such that the sum of the proper divisors of each number is one more than the other number:

(n, m) are a pair of buddy if s(m) = n + 1 and s(n) = m + 1

For example 48 & 75 is such a pair:

Divisors of 48 are: 1, 2, 3, 4, 6, 8, 12, 16, 24 --> sum: 76 = 75 + 1
Divisors of 75 are: 1, 3, 5, 15, 25 --> sum: 49 = 48 + 1
Task
Given two positive integers start and limit, the function buddy(start, limit) should return the first pair (n m) of buddy pairs such that n (positive integer) is between start (inclusive) and limit (inclusive); m can be greater than limit and has to be greater than n

If there is no buddy pair satisfying the conditions, then return "Nothing" or (for Go lang) nil

Examples
(depending on the languages)

buddy(10, 50) returns [48, 75]
buddy(48, 50) returns [48, 75]
or
buddy(10, 50) returns "(48 75)"
buddy(48, 50) returns "(48 75)"
Note
for C: The returned string will be free'd.
See more examples in "Sample Tests:" of your language.
*/
package main

import (
	"fmt"
	"github.com/hummingbird/libs/util"
)

func Buddy(start, limit int) []int {
	for i := start; i <= limit; i++ {
		i2 := Sum(arr(i)) - 1
		if i2 > i {
			if Sum(arr(i2))-1 == i {
				return []int{i, i2}
			}
		}
	}
	return nil
}

func Sum(x []int) int {
	sum_ := 0
	for _, i := range x {
		sum_ += i
	}
	return sum_
}

//TODO CALC positive number divisors
func arr(n int) []int {
	arr_ := make([]int, 0)
	for i := 2; i < n/i; i++ {
		if n%i == 0 {
			arr_ = append(arr_, i, n/i)
		}
	}
	return append([]int{1}, arr_...)
}

// the return is `nil` if there is no buddy pair
func Buddy2(start, limit int) []int {
	for i := start; i <= limit; i++ {
		i2 := sumArr(i)
		if sumArr(i2) == i {
			return []int{i, i2}
		}
	}
	return nil
}
func sumArr(n int) int {
	sum := 0
	for i := 2; i < n/i; i++ {
		if n%i == 0 {
			sum += i + n/i
		}
	}
	return sum
}

func main() {
	fmt.Println(arr(48))
	fmt.Println(arr(75))

	fmt.Println(Buddy(10, 50))
	fmt.Println(Buddy(48, 50))
	util.Cost(func() {
		//"[1081184 1331967]"
		fmt.Println(Buddy(1071625, 1103735))
	})

	//"[62744 75495]"
	fmt.Println(Buddy(57345, 90061))
	//"[5775 6128]"
	fmt.Println(Buddy(2693, 7098))
	//"[]"
	fmt.Println(Buddy(6379, 8275))
	//[]
	fmt.Println(Buddy(74226, 79716))
}
