/*
 date: 2019-04-19 22:15:05
 A natural number is called k-prime if it has exactly k prime factors, counted with multiplicity. A natural number is thus prime if and only if it is 1-prime.

Examples:
k = 2 -> 4, 6, 9, 10, 14, 15, 21, 22, …
k = 3 -> 8, 12, 18, 20, 27, 28, 30, …
k = 5 -> 32, 48, 72, 80, 108, 112, …
Task:

Given an integer k and a list arr of positive integers the function consec_kprimes (or its variants in other languages) returns how many times in the sequence arr numbers come up twice in a row with exactly k prime factors?

Examples:

arr = [10005, 10030, 10026, 10008, 10016, 10028, 10004]
consec_kprimes(4, arr) => 3 because 10005 and 10030 are consecutive 4-primes, 10030 and 10026 too as well as 10028 and 10004 but 10008 and 10016 are 6-primes.

consec_kprimes(4, [10175, 10185, 10180, 10197]) => 3 because 10175-10185 and 10185- 10180 and 10180-10197 are all consecutive 4-primes.
Note:

It could be interesting to begin with: https://www.codewars.com/kata/k-primes
 */
package main

import (
	"fmt"
	"math"
)

func ConsecKprimes(k int, arr []int) int {
	count := 0
	for i,t:=1,arr[0];i<len(arr);i++{
		tk:= kprimes(t)
		if(k == tk && tk == kprimes(arr[i])){
			count++
		}
		t = arr[i]
	}
	return count
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
func main() {
	fmt.Println(ConsecKprimes(2, []int{10081, 10071, 10077, 10065, 10060, 10070, 10086, 10083, 10078, 10076, 10089, 10085, 10063, 10074, 10068, 10073, 10072, 10075})== 2)
	fmt.Println(ConsecKprimes(6, []int{10064})== 0)
	fmt.Println(ConsecKprimes(1, []int{10054, 10039, 10053, 10051, 10047, 10043, 10037, 10034})== 0)
}
