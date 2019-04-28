/*
 date: 2019-04-28 15:20:01
In this kata you have to correctly return who is the "survivor", ie: the last element of a Josephus permutation.

Basically you have to assume that n people are put into a circle and that they are eliminated in steps of k elements, like this:

josephus_survivor(7,3) => means 7 people in a circle;
one every 3 is eliminated until one remains
[1,2,3,4,5,6,7] - initial sequence
[1,2,4,5,6,7] => 3 is counted out
[1,2,4,5,7] => 6 is counted out
[1,4,5,7] => 2 is counted out
[1,4,5] => 7 is counted out
[1,4] => 5 is counted out
[4] => 1 counted out, 4 is the last element - the survivor!
The above link about the "base" kata description will give you a more thorough insight about the origin of this kind of permutation, but basically that's all that there is to know to solve this kata.

Notes and tips: using the solution to the other kata to check your function may be helpful, but as much larger numbers will be used, using an array/list to compute the number of the survivor may be too slow; you may assume that both n and k will always be >=1.
*/
package main

import "fmt"

func JosephusSurvivor2(n, k int) int {
	if n == 1 {
		return 1
	}
	return (JosephusSurvivor(n-1, k)+k-1)%n + 1
}

func JosephusSurvivor(n, k int) int {
	items := make([]int, n)
	for i := 0; i < n; i++ {
		items[i] = i + 1
	}
	counter := 0
	idx := 0
	for len(items) != 1 {
		if idx >= len(items) {
			idx = 0
		}
		if (counter+1)%k == 0 {
			items = append(items[:idx], items[idx+1:]...)
			counter = 0
		} else {
			counter += 1
			idx += 1
		}
	}
	return items[0]
}

func main() {
	fmt.Println(JosephusSurvivor(7, 3))
	fmt.Println(JosephusSurvivor(11, 19))
	fmt.Println(JosephusSurvivor(40, 3))
	fmt.Println(JosephusSurvivor(14, 2))
	fmt.Println(JosephusSurvivor(100, 1))
	//fmt.Println(JosephusSurvivor(7, 3) == 4)
	//fmt.Println(JosephusSurvivor(11, 19) == 10)
	//fmt.Println(JosephusSurvivor(40, 3) == 28)
	//fmt.Println(JosephusSurvivor(14, 2) == 13)
	//fmt.Println(JosephusSurvivor(100, 1) == 100)
}
