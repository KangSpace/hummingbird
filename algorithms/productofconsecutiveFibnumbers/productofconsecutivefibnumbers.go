/*
date: 2019-04-15 09:27:07
The Fibonacci numbers are the numbers in the following integer sequence (Fn):
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, ...
such as
F(n) = F(n-1) + F(n-2) with F(0) = 0 and F(1) = 1.
Given a number, say prod (for product), we search two Fibonacci numbers F(n) and F(n+1) verifying
F(n) * F(n+1) = prod.
Your function productFib takes an integer (prod) and returns an array:
[F(n), F(n+1), true] or {F(n), F(n+1), 1} or (F(n), F(n+1), True)
depending on the language if F(n) * F(n+1) = prod.
If you don't find two consecutive F(m) verifying F(m) * F(m+1) = prodyou will return
[F(m), F(m+1), false] or {F(n), F(n+1), 0} or (F(n), F(n+1), False)
F(m) being the smallest one such as F(m) * F(m+1) > prod.

Examples
productFib(714) # should return (21, 34, true),
                # since F(8) = 21, F(9) = 34 and 714 = 21 * 34
productFib(800) # should return (34, 55, false),
                # since F(8) = 21, F(9) = 34, F(10) = 55 and 21 * 34 < 800 < 34 * 55
Notes: Not useful here but we can tell how to choose the number n up to which to go: we can use the "golden ratio" phi which is (1 + sqrt(5))/2 knowing that F(n) is asymptotic to: phi^n / sqrt(5). That gives a possible upper bound to n.
You can see examples in "Example test".
References
http://en.wikipedia.org/wiki/Fibonacci_number

http://oeis.org/A000045
*/
package main

import (
	"fmt"
	"math"
	"time"
)

func Fib(x uint64) uint64 {
	if x == 0 || x == 1 {
		return x
	}
	return Fib(x-1) + Fib(x-2)
}

func ProductFib(prod uint64) [3]uint64 {
	var a, b ,isFound uint64 = 0, 1,0
	for prod > a*b {
		a, b = b, a+b
	}
	if prod == a*b {
		isFound = 1
	}
	return [3]uint64{a, b,isFound}
}
func main() {
	fmt.Println(math.Pow((math.Sqrt(5)+1)/2, float64(1)) / math.Sqrt(5) * float64(1))
	a := time.Now().Nanosecond()
	fmt.Println(a)
	fmt.Println(Fib(10))
	fmt.Println(ProductFib(4895))
	fmt.Println(ProductFib(5895))
	fmt.Println(ProductFib(74049690))
	fmt.Println(ProductFib(84049690))
	fmt.Println(time.Now().Nanosecond() - a)
}
