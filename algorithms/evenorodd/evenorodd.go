/*
 date: 2019-04-10 11:22:50
 Create a function (or write a script in Shell) that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
 */
package main

import "fmt"

func EvenOrOdd(num int) string{
	if num%2== 0 {
		return "Even"
	}else{
		return "Odd"
	}
}
func EvenOrOdd2(num int) string{
	return [2] string{"Even", "Odd"}[num & 1]
}

func main() {
	num1 := 1
	num2 := 2
	num3 := 3
	num4 := 4
	fmt.Println(EvenOrOdd(num1))
	fmt.Println(EvenOrOdd(num2))
	fmt.Println(EvenOrOdd(num3))
	fmt.Println(EvenOrOdd(num4))
}
