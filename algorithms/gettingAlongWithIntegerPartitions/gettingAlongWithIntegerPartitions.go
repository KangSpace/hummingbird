/*
From wikipedia https://en.wikipedia.org/wiki/Partition_(number_theory)

In number theory and combinatorics, a partition of a positive integer n, also called an integer partition, is a way of writing n as a sum of positive integers. Two sums that differ only in the order of their summands are considered the same partition.

For example, 4 can be partitioned in five distinct ways:

4, 3 + 1, 2 + 2, 2 + 1 + 1, 1 + 1 + 1 + 1.

We can write:

enum(4) -> [[4],[3,1],[2,2],[2,1,1],[1,1,1,1]] and

enum(5) -> [[5],[4,1],[3,2],[3,1,1],[2,2,1],[2,1,1,1],[1,1,1,1,1]].

The number of parts in a partition grows very fast. For n = 50 number of parts is 204226, for 80 it is 15,796,476 It would be too long to tests answers with arrays of such size. So our task is the following:

1 - n being given (n integer, 1 <= n <= 50) calculate enum(n) ie the partition of n. We will obtain something like that:
enum(n) -> [[n],[n-1,1],[n-2,2],...,[1,1,...,1]] (order of array and sub-arrays doesn't matter). This part is not tested.

2 - For each sub-array of enum(n) calculate its product. If n = 5 we'll obtain after removing duplicates and sorting:

prod(5) -> [1,2,3,4,5,6]

prod(8) -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 15, 16, 18]

If n = 40 prod(n) has a length of 2699 hence the tests will not verify such arrays. Instead our task number 3 is:

3 - return the range, the average and the median of prod(n) in the following form (example for n = 5):

"Range: 5 Average: 3.50 Median: 3.50"

Range is an integer, Average and Median are float numbers rounded to two decimal places (".2f" in some languages).

#Notes: Range : difference between the highest and lowest values.

Mean or Average : To calculate mean, add together all of the numbers in a set and then divide the sum by the total count of numbers.

Median : The median is the number separating the higher half of a data sample from the lower half. (https://en.wikipedia.org/wiki/Median)

#Hints: Try to optimize your program to avoid timing out.

Memoization can be helpful but it is not mandatory for being successful.
*/
package main

import (
	"fmt"
	"sort"
)

func Part(n int) string{
	prods := Prod(n)
	return fmt.Sprintf("Range: %d Average: %.2f Median: %.2f", prods[len(prods)-1]-prods[0],float64(sum(prods))/float64(len(prods)),Median(prods))
}

func Median(n []int) float64{
	if len(n) % 2 ==0{
		return float64(n[len(n)/2-1]+n[len(n)/2])/2
	}
	return float64(n[len(n)/2])
}

func Prod(n int) []int {
	enums := enum(n, make(map[int][][]int))
	prods := make(map[int]int, 0)
	for _, i := range enums {
		prod := 1;
		for _, j := range i {
			prod *= j
		}
		prods[prod]=0
	}
	prods_ := make([]int ,0)
	for i,_ := range prods{
		prods_ = append(prods_,i)
	}
    sort.Ints(prods_)
	return prods_
}

func enum(n int, cache map[int][][]int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	enums := [][]int{{n}}
	for m := n; m > 0; m-- {
		r := n - m
		var fr [][]int
		if fr_, ok := cache[r]; ok {
			fr = fr_
		} else {
			fr = enum(r, cache)
			cache[r] = fr
		}
		enums = append(enums, accept(fr, []int{m})...)
	}
	return enums
}
func accept(lst [][]int, n []int) [][]int {
	if len(lst) == 0 {
		return lst
	}
	xs := make([][]int, 0)
	for _, x := range lst {
		if max(x) <= n[0] {
			xs = append(xs, append(n, x...))
		}
	}
	return xs
}
func max(x []int) int {
	max_ := 0
	for _, i := range x {
		if i > max_ {
			max_ = i
		}
	}
	return max_
}
func sum(x []int) int {
	sum_ := 0
	for _, i := range x {
		sum_ += i
	}
	return sum_
}

func partAux(s, k int) [][]int {
	k0 := min(s, k)
	res := [][]int{}
	n := k0;
	var r int
	for ; n > 0; n-- {
		r = s - n
		if r > 0 {
			arr := partAux(r, n)
			for i := 0; i < len(arr); i++ {
				t := arr[i]
				t = append(t, n)
				res = append(res, t)
			}
		} else {
			res = append(res, []int{n})
		}
		fmt.Println(res)
	}
	return res
}
//## GOOD
func ruleAsc(n int) [][]int {
	res := make([][]int, 0)
	a := make([]int, n + 1)
	k := 1
	a[1] = n
	for k != 0 {
		x := a[k - 1] + 1
		y := a[k] - 1
		k -= 1
		for x <= y {
			a[k] = x
			y -= x
			k += 1
		}
		a[k] = x + y
		product := []int{}
		for j := 0; j <= k; j++ {
			product = append(product,a[j])
		}
		res = append(res, product)
	}
	return res
}

func min (x,y int) int{
	if x < y {return x} else {return y}
}


func main() {
	//enum(4) -> [[4],[3,1],[2,2],[2,1,1],[1,1,1,1]] and
	//fmt.Println(partAux(4, 4))
	////enum(5) -> [[5],[4,1],[3,2],[3,1,1],[2,2,1],[2,1,1,1],[1,1,1,1,1]].
	fmt.Println(ruleAsc(5))
	//util.Cost(func() {
	//	fmt.Println(ruleAsc(50))
	//})

	//fmt.Println(partAux(5, 5))
	//fmt.Println(partAux(50, 50))

	//enum(4) -> [[4],[3,1],[2,2],[2,1,1],[1,1,1,1]] and
	//fmt.Println(enum(4, make(map[int][][]int)))
	////enum(5) -> [[5],[4,1],[3,2],[3,1,1],[2,2,1],[2,1,1,1],[1,1,1,1,1]].
	//fmt.Println(enum(5, make(map[int][][]int)))
	//util.Cost(func() {
	//	fmt.Println(enum(50, make(map[int][][]int)))
	//})

	//fmt.Println(prod(5))
	/*prodMap := make(map[int]int)
	prodMap[0]=2
	prodMap[1]=3
	prodMap[2]=1
	prodMap[4]=0
	fmt.Println(reflect.ValueOf(prodMap).MapKeys())
	fmt.Println(prodMap)*/
	//Part(1)
	//Part(2)
	//fmt.Println(Part(43))
	//fmt.Println(Part(1)== "Range: 0 Average: 1.00 Median: 1.00")
	//fmt.Println(Part(2)== "Range: 1 Average: 1.50 Median: 1.50")
	//fmt.Println(Part(3)== "Range: 2 Average: 2.00 Median: 2.00")

}
