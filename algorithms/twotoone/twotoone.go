/*
date: 2019-04-08
Take 2 strings s1 and s2 including only letters from ato z. Return a new sorted string, the longest possible, containing distinct letters,

each taken only once - coming from s1 or s2.
Examples:
a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
longest(a, b) -> "abcdefklmopqwxy"

a = "abcdefghijklmnopqrstuvwxyz"
longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)


func TwoToOne(a string , b string) string{
	arr := strings.Split(a+b,"")
	sort.Strings(arr)
	result :=""
	for _,t := range arr{
		temp := string(t)
		if !strings.Contains(result,temp){
			result += temp
		}
	}
	return result
}

//2019-04-08 11:05:18
func TwoToOne2(a string , b string) string{
	arrstr := strIntoArr("",a)
	arrstr = strIntoArr(arrstr,b)
	arr := str2arr(arrstr)
	sort.Strings(arr)
	return strings.Join(arr,"")
}

func strIntoArr(arrstr,a string) string{
	for i,t:=range a{
		fmt.Println(i,",",t,",",string(t))
		//arr = append(arr,string(t))
		temp := string(t)
		if !strings.Contains(arrstr,temp) {
			arrstr+=temp
		}
	}
	return arrstr
}

func str2arr(a string) []string{
	arr := make([]string,0,1)
	for _,t:=range a {
		arr = append(arr,string(t))
	}
	return arr
}

func main() {
	fmt.Println(strings.Split("abcsd",""))
	fmt.Print(TwoToOne("acb","abcefd"))
}
