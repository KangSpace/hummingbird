/*
 date: 2019-04-28 13:51:20
Write a function named first_non_repeating_letter that takes a string input, and returns the first character that is not repeated anywhere in the string.

For example, if given the input 'stress', the function should return 't', since the letter t only occurs once in the string, and occurs first in the string.

As an added challenge, upper- and lowercase letters are considered the same character, but the function should return the correct case for the initial letter. For example, the input 'sTreSS' should return 'T'.

If a string contains all repeating characters, it should return an empty string ("") or None -- see sample tests.
*/
package main

import (
	"fmt"
	"strings"
)

func FirstNonRepeating1(str string) string {
	vmap := make(map[string]int, 0)
	for _, v := range str[:] {
		vmv, _ := vmap[string(v)]
		vmv1, _ := vmap[strings.ToUpper(string(v))]
		vmv2, _ := vmap[strings.ToLower(string(v))]
		vs := string(v)
		if vmv1 != 0 {
			vs = strings.ToUpper(string(v))
		} else if vmv2 != 0 {
			vs = strings.ToLower(string(v))
		}
		sv := vmv + vmv1 + vmv2 + 1
		vmap[vs] = sv
	}
	for _, i := range str[:] {
		if v, ok := vmap[string(i)]; ok && v == 1 {
			return string(i)
		}
	}
	return ""
}
func FirstNonRepeating2(str string) string {
	for _, i := range str[:] {
		count := 0
		for _, j := range str[:] {
			if i == j || i-j == 32 || i-j == -32 {
				count++
			}
		}
		if count == 0 {
			return ""
		}
		if count == 1 {
			return string(i)
		}
	}
	return ""
}
func FirstNonRepeating(str string) string {
	for _, i := range str[:] {
		if strings.Count(strings.ToLower(str), strings.ToLower(string(i))) < 2 {
			return string(i)
		}
	}
	return ""
}
func main() {
	//fmt.Println(FirstNonRepeating("sTreSS"))
	//fmt.Println(FirstNonRepeating("sTreSS") == "T")
	//fmt.Println(FirstNonRepeating("sTretreSS"))
	fmt.Println(FirstNonRepeating("hello world, eh?"))
}
