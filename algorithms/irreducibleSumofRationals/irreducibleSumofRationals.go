/*
 date: 2019-04-16 09:44:56
	You will have a list of rationals in the form
	lst = [ [numer_1, denom_1] , ... , [numer_n, denom_n] ]
	or
	lst = [ (numer_1, denom_1) , ... , (numer_n, denom_n) ]
	where all numbers are positive integers. You have to produce their sum N / D in an irreducible form: this means that N and D have only 1 as a common divisor.
	Return the result in the form:
	[N, D] in Ruby, Crystal, Python, Clojure, JS, CS, PHP, Julia
	Just "N D" in Haskell, PureScript
	"[N, D]" in Java, CSharp, TS, Scala, PowerShell
	"N/D" in Go, Nim
	{N, D} in C++, Elixir
	"{N, D}" in C
	Some((N, D)) in Rust
	Some "N D" in F#, Ocaml
	c(N, D) in R
	(N, D) in Swift
	If the result is an integer (D evenly divides N) return:
	an integer in Ruby, Crystal, Elixir, Clojure, Python, JS, CS, PHP, R, Julia
	Just "n" (Haskell, PureScript)
	"n" Java, CSharp, TS, Scala, PowerShell, Go, Nim
	{n, 1} in C++
	"{n, 1}" in C
	Some((n, 1)) in Rust
	Some "n" in F#, Ocaml,
	(n, 1) in Swift
	If the input list is empty, return nil/None/null/Nothing (or {0, 1} in C++; "{0, 1}" in C) (or "0" in Scala, PowerShell, Go, Nim)
	Example:
[ [1, 2], [1, 3], [1, 4] ]  -->  [13, 12]

    1/2  +  1/3  +  1/4     =      13/12
See sample tests for more examples and the form of results.
*/
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func SumFracts(arr [][]int) string {
	if len(arr) == 0 {
		return "0"
	}
	fmt.Println(arr)
	var a1 ,lcmNum int = arr[0][1],0
	for i :=1 ; i<len(arr);i++ {
		lcmNum= lcm(a1,arr[i][1])
		a1 = lcmNum
	}
	var n int = 0
	for _, a := range arr {
		n+=lcmNum/a[1]*a[0]
	}
	var nlcmNum = gcd(n,lcmNum)
	if n > nlcmNum && lcmNum > nlcmNum{
		n = n/nlcmNum
		lcmNum = lcmNum /nlcmNum
	}
	if n % lcmNum == 0 {
		return strconv.FormatInt(int64(n/lcmNum), 10)
	}
	return strconv.FormatInt(int64(n), 10) + "/" + strconv.FormatInt(int64(lcmNum), 10)
}

//最大公约数 欧几里得辗转相除法
// 当处理较大整数时，欧几里得辗转相除法 比 更相减损法 效率更高
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b,a%b)
}

//最小公倍数 ,公式法: 2数的最小公倍数=a*b/2数的最大公约数
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

//
//更相减损法
//更相减损法出自中国古代的数学专著，其原文是
//可半者半之，不可半者，副置分母、子之数，以少减多，更相减损，求其等也。以等数约之。
//—— 《九章算术》
//这段文言文翻译成白话文，包含了下面四个算法步骤
//如果两个数都是偶数，那就先将它们除以二
//若两个数都是奇数，那么就比较两者大小，用大的数字作被减数，小的数字作减数
//将减法得出的结果与减数作比较，重复上面的步骤
//直到最后两个数相等，那么最大公约数就是最后相等的那个数，或最后相等的数乘以每次约去的2最后的乘积
func gcd2(x, y int) int {
	for x != y {
		if x > y {
			x -= y
		} else {
			y -= x
		}
	}
	return x
}

//math.big
func SumFracts2(arr [][]int) string {
	sum := big.NewRat(0,1)

	for _, f := range arr {
		sum.Add(sum, big.NewRat(int64(f[0]), int64(f[1])))
	}
	fmt.Println(sum.RatString())

	if sum.Denom().Cmp(big.NewInt(1)) == 0 {
		return sum.Num().String()
	}
	return sum.String()
}

func main() {
	// "13/12"
	//fmt.Println(SumFracts([][]int{{1, 2}, {1, 3}, {1, 4}}))
	////"2"
	//fmt.Println(SumFracts([][]int{{1, 3}, {5, 3}}))
	//fmt.Println(SumFracts([][]int{{2,7}, {1,3},{1,12}}))
	fmt.Println(SumFracts2([][]int{{69,130},{87,1310},{3,4}}))
	//fmt.Println(SumFracts([][]int{{3086,6179},{3086,3105},{6173,9259}}))
	//fmt.Println(SumFracts([][]int{{12,3},{15,3}}))
	//fmt.Println(gcd(319, 377))
	//fmt.Println(gcd(59, 84))
	//fmt.Println(lcm(21, 12))
	//fmt.Println(strconv.Itoa(1))
}
