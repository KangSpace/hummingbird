/*
 date: 2019-04-20 13:37:20
 My little sister came back home from school with the following task: given a squared sheet of paper she has to cut it in pieces which, when assembled, give squares the sides of which form an increasing sequence of numbers. At the beginning it was lot of fun but little by little we were tired of seeing the pile of torn paper. So we decided to write a program that could help us and protects trees.

Task
Given a positive integral number n, return a strictly increasing sequence (list/array/string depending on the language) of numbers, so that the sum of the squares is equal to n².

If there are multiple solutions (and there will be), return the result with the largest possible values:

Examples
decompose(11) must return [1,2,4,10]. Note that there are actually two ways to decompose 11², 11² = 121 = 1 + 4 + 16 + 100 = 1² + 2² + 4² + 10² but don't return [2,6,9], since 9 is smaller than 10.

For decompose(50) don't return [1, 1, 4, 9, 49] but [1, 3, 5, 8, 49] since [1, 1, 4, 9, 49] doesn't form a strictly increasing sequence.

Note
Neither [n] nor [1,1,1,…,1] are valid solutions. If no valid solution exists, return nil, null, Nothing, None (depending on the language) or "[]" (C) ,{} (C++), [] (Swift, Go).

The function "decompose" will take a positive integer n and return the decomposition of N = n² as:

[x1 ... xk] or
"x1 ... xk" or
Just [x1 ... xk] or
Some [x1 ... xk] or
{x1 ... xk} or
"[x1,x2, ... ,xk]"
depending on the language (see "Sample tests")

Note for Bash
decompose 50 returns "1,3,5,8,49"
decompose 4  returns "Nothing"
Hint
Very often xk will be n-1.
*/
package main

import (
	"fmt"
	"math"
)

func Decompose(n int64) []int64 {
	return recurse(n*n, n-1)
	//return recurse(n)
}


func recurse(s, i int64) []int64 {
	//fmt.Println("s:", s, ",i:", i)
	if s == 0 {
		return []int64{}
	}
	if s < 0 || i <= 0 {
		return nil
	}
	for ; i > 0; i-- {
		s1 := s-i*i
		s2 := int64(math.Min(float64(i-1), math.Sqrt(float64(s1))))
		sub_ := recurse(s1, s2)
		if sub_ != nil {
			sub_ = append(sub_, i)
			return sub_
		}
	}
	return nil
}
func recurse2(i int64) []int64 {
	r := i * i
	n := i - 1
	b := false
	stack := make([][]int64, 0)
	for n > 0 && !b {
		r2 := r
		n2 := n
		for {
			if r2 == 0 {
				b = true
				break
			}
			if n2 == 0 {
				if len(stack) > 0 {
					r2, n2 = stack[len(stack)-1:][0][0], stack[len(stack)-1:][0][1]
					stack = stack[:len(stack)-1]
					n2 -= 1
				} else {
					break
				}
			}
			if n2 == 0 {
				continue
			}
			stack = append(stack, []int64{r2, n2})
			r2 = r2 - n2*n2
			n2 = int64(math.Min(float64(n2-1), math.Sqrt(float64(r2))))
		}
		if b {
			keys := make([]int64, 0)
			for i := len(stack) - 1; i >= 0; i-- {
				keys = append(keys, stack[i][1])
			}
			return keys
		}
	}
	return nil
}

func main() {
	//[]int64{2, 3, 6}
	fmt.Println(Decompose(7))
	//== []int64{1, 3, 5, 8, 49}
	fmt.Println(Decompose(50))
	//== []int64{3, 4})
	fmt.Println(Decompose(5))
	//== []int64{})
	fmt.Println(Decompose(2))
}
