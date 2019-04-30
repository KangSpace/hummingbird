/*
 date: 2019-04-29 17:02:15
Instructions
Output
The action of a Caesar cipher is to replace each plaintext letter with a different one a fixed number of places up or down the alphabet.

This program performs a variation of the Caesar shift. The shift increases by 1 for each character (on each iteration).

If the shift is initially 1, the first character of the message to be encoded will be shifted by 1, the second character will be shifted by 2, etc...

Coding: Parameters and return of function "movingShift"
param s: a string to be coded

param shift: an integer giving the initial shift

The function "movingShift" first codes the entire string and then returns an array of strings containing the coded string in 5 parts (five parts because, to avoid more risks, the coded message will be given to five runners, one piece for each runner).

If possible the message will be equally divided by message length between the five runners. If this is not possible, parts 1 to 5 will have subsequently non-increasing lengths, such that parts 1 to 4 are at least as long as when evenly divided, but at most 1 longer. If the last part is the empty string this empty string must be shown in the resulting array.

For example, if the coded message has a length of 17 the five parts will have lengths of 4, 4, 4, 4, 1. The parts 1, 2, 3, 4 are evenly split and the last part of length 1 is shorter. If the length is 16 the parts will be of lengths 4, 4, 4, 4, 0. Parts 1, 2, 3, 4 are evenly split and the fifth runner will stay at home since his part is the empty string. If the length is 11, equal parts would be of length 2.2, hence parts will be of lengths 3, 3, 3, 2, 0.

You will also implement a "demovingShift" function with two parameters

Decoding: parameters and return of function "demovingShift"
1) an array of strings: s (possibly resulting from "movingShift", with 5 strings)

2) an int shift

"demovingShift" returns a string.

Example:
u = "I should have known that you would have a perfect answer for me!!!"

movingShift(u, 1) returns :

v = ["J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"]

(quotes added in order to see the strings and the spaces, your program won't write these quotes, see Example Test Cases)

and demovingShift(v, 1) returns u.

#Ref:

Caesar Cipher : http://en.wikipedia.org/wiki/Caesar_cipher


Analysis:
	s: character
shift: init move seed
movingShift:
	ts := s + shift
    if [A-Z] + shift > 90 {
      return [A-Z] -> ts = ts - math.floor((ts - 65 ) / 26) * 26
    }
	if [a-z] + shift > 122 {
      return [a-z] -> ts = ts - math.floor((ts - 97 ) / 26) * 26
    }
demovingShift:
	ts := s - shift
    if [A-Z] - shift < 65 {
      return [A-Z] -> ts = ts + math.Ceil((65 - ts ) / 26) * 26
    }
	if [a-z] + shift < 97 {
      return [a-z] -> ts = ts + math.Ceil((97 - ts) / 26) * 26
    }
*/
package main

import (
	"fmt"
	"math"
	"strings"
)

func MovingShift(s string, shift int) []string {
	perLen := int(math.Ceil(float64(len(s))/float64(5)))
	var result []string
	str := ""
	for i, s_ := range s[:] {
		str1 := string(s_)
		if (s_ >= 65 && s_ <= 90) || (s_ >= 97 && s_ <= 122) {
			ts := s_ + int32(shift)
			if s_ >= 65 && s_<=90 && ts >90{
				ts = ts - (ts - 65 ) / 26 * 26
			}else if s_ >= 97 && s_<=122 && ts >122{
				ts = ts - (ts - 97 ) / 26 * 26
			}
			str1 = string(ts)
		}
		str += str1
		shift++
		if (i > 0 && (i+1)%perLen == 0) || (i == len(s)-1 && len(result) < 5) {
			result = append(result, str)
			str = ""
		}
	}
	for len(result) < 5 {
		result = append(result, "")
	}
	return result
}
func DemovingShift(arr []string, shift int) string {
	str := ""
	for _, v := range arr {
		if len(v) == 0 {
			break
		}
		for _, s_ := range v[:] {
			ts := s_
			if (s_ >= 65 && s_ <= 90) || (s_ >= 97 && s_ <= 122) {
				ts = ts - int32(shift)
				if s_ >= 65 && s_<=90 && ts <65 {
					ts = ts + int32(math.Ceil(float64(65 - ts)/ float64(26)) * 26)
				}else if s_ >= 97 && s_<=122 && ts < 97{
					ts = ts + int32(math.Ceil(float64(97 - ts)/ float64(26)) * 26)
				}
			}
			str += string(ts)
			shift++
		}
	}
	return str
}
func main() {
	//var sol1 = []string{"T p", "oc ", "iwl", "yo", ""}
	//var u1 = "S mkx bocod"
	//fmt.Println(MovingShift(u1,1))
	//fmt.Println(DemovingShift(sol1,1))

	//var sol1 = []string{"T p", "oc ", "iwl", "yo", ""}
	//<[]string | len:5, cap:5>: ["Y ONDIQZF!", " hu Azpucl", "r! vca qqn", "fukc mldl ", "gr eqqi;"]
	var u1 = "O CAPTAIN! my Captain! our fearful trip is done;"
	fmt.Println(MovingShift(u1,10))
	var u2 = " uoxIirmoveNreefckgieaoiEcooqo"
	fmt.Println(MovingShift(u2,2))
	//fmt.Println(DemovingShift(sol1,1))

	var u = "I should have known that you would have a perfect answer for me!!!"
	var sol = []string{"J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"}
	fmt.Println(MovingShift(u, 1))
	fmt.Println(strings.Join(MovingShift(u, 1), ",") == strings.Join(sol, ","))
	fmt.Println(DemovingShift(sol, 1))
	fmt.Println(DemovingShift(sol, 1) == u)

}
