/*
 date: 2019-04-19 22:42:28
 A natural number is called k-prime if it has exactly k prime factors, counted with multiplicity. For example:

k = 2  -->  4, 6, 9, 10, 14, 15, 21, 22, ...
k = 3  -->  8, 12, 18, 20, 27, 28, 30, ...
k = 5  -->  32, 48, 72, 80, 108, 112, ...
A natural number is thus prime if and only if it is 1-prime.

Task:
Complete the function count_Kprimes (or countKprimes, count-K-primes, kPrimes) which is given parameters k, start, end (or nd) and returns an array (or a list or a string depending on the language - see "Solution" and "Sample Tests") of the k-primes between start (inclusive) and end (inclusive).

Example:
countKprimes(5, 500, 600) --> [500, 520, 552, 567, 588, 592, 594]
Notes:

The first function would have been better named: findKprimes or kPrimes :-)
In C some helper functions are given (see declarations in 'Solution').
For Go: nil slice is expected when there are no k-primes between start and end.
Second Task (puzzle):
Given positive integers s, a, b, c where a is1-prime, b is 3-prime, c is 7-prime, find the total number of solutions where a + b + c = s. Call this function puzzle(s).

Examples:
puzzle(138)  -->  1  because [2 + 8 + 128] is the only solution
puzzle(143)  -->  2  because [3 + 12 + 128] and [7 + 8 + 128] are the solutions
 */
package main

import (
	"fmt"
	"math"
)

func CountKprimes(k, start, nd int)  []int {
	r:=make([]int,0)
	for ;start<=nd;start++ {
		if k == kprimes(start) {
			r = append(r,start)
		}
	}
	if len(r) == 0 {
		return nil
	}
	return r
}
func kprimes(i int)int{
	k:=0
	for j:=2;j<=i;j++{
		if(i % j == 0 && primes(j)){
			i/=j
			k++
			j = 1
		}
	}
	return k
}
func primes(n int) bool{
	if n<=3{
		return n>1;
	}
	for i:=2;i<=int(math.Sqrt(float64(n)));i++{
		if n % i == 0{
			return false;
		}
	}
	return true
}
func Puzzle(s int) int {
	// your code
	//s = primes(n) + kprimes(3)+kprimes(7);
	// 7-primes
	count := 0
	is7 :=CountKprimes(7,2,s)
	is3 :=CountKprimes(3,2,s)
	is1 :=CountKprimes(1,2,s)
	for _,i := range is1{
		for _,i2 := range is3{
			for _,i3 := range is7{
				if i + i2 + i3 == s {
					count ++
				}
			}
		}
	}
	return count
}

func main() {
	//[]int{1020, 1026, 1032, 1044, 1050, 1053, 1064,   1072, 1092, 1100}
	fmt.Println(CountKprimes( 5, 1000, 1100))
	////, []int{500, 520, 552, 567, 588, 592, 594}
	fmt.Println(CountKprimes( 5, 500, 600))
	////nil
	fmt.Println(CountKprimes( 12, 100000, 100100))
	fmt.Println(kprimes(8))
	//1
	fmt.Println(Puzzle(138))
	//2
	fmt.Println(Puzzle(143))
	fmt.Println(nil == []int{})
}
