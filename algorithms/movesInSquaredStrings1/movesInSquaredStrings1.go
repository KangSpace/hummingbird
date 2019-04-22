/*
 date: 2019-04-19 21:22:22
 This kata is the first of a sequence of four about "Squared Strings".

You are given a string of n lines, each substring being n characters long: For example:

s = "abcd\nefgh\nijkl\nmnop"

We will study some transformations of this square of strings.

Vertical mirror: vert_mirror (or vertMirror or vert-mirror)
vert_mirror(s) => "dcba\nhgfe\nlkji\nponm"
Horizontal mirror: hor_mirror (or horMirror or hor-mirror)
hor_mirror(s) => "mnop\nijkl\nefgh\nabcd"
or printed:

vertical mirror   |horizontal mirror
abcd --> dcba     |abcd --> mnop
efgh     hgfe     |efgh     ijkl
ijkl     lkji     |ijkl     efgh
mnop     ponm     |mnop     abcd
#Task:

Write these two functions
and

high-order function oper(fct, s) where

fct is the function of one variable f to apply to the string s (fct will be one of vertMirror, horMirror)
#Examples:

s = "abcd\nefgh\nijkl\nmnop"
oper(vert_mirror, s) => "dcba\nhgfe\nlkji\nponm"
oper(hor_mirror, s) => "mnop\nijkl\nefgh\nabcd"
Note:
The form of the parameter fct in oper changes according to the language. You can see each form according to the language in "Sample Tests".

Bash Note:
The input strings are separated by , instead of \n. The ouput strings should be separated by \r instead of \n. See "Sample Tests".

Forthcoming katas will study other tranformations.
 */
package main

import (
	"fmt"
	"strings"
)

func VertMirror(s string) string {
	r:=make([]string,0)
	for _,ss:= range strings.Split(s,"\n"){
		r = append(r,strings.Join(reverse(strings.Split(ss,"")),""))
	}
	return strings.Join(r,"\n")
}
func HorMirror(s string) string {
	return strings.Join(reverse(strings.Split(s,"\n")),"\n")
}
type FParam func(string) string
func Oper(f FParam, x string) string {
	return f(x)
}

func reverse(ss []string) []string{
	for i,j:=0,len(ss)-1;i<j;i,j=i+1,j-1{
		ss[i],ss[j]= ss[j],ss[i]
	}
	return ss
}

func main() {
	//fmt.Println(Oper(VertMirror,"hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu"))
	//fmt.Println(Oper(HorMirror,"lVHt\nJVhv\nCSbg\nyeCt"))
	//fmt.Println(Oper(VertMirror,"abcd\nefgh\nijkl\nmnop"))
	//fmt.Println(Oper(HorMirror,"abcd\nefgh\nijkl\nmnop"))

	fmt.Println(Oper(VertMirror,"hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu")=="QHdgSh\noaMDnH\nXxNNlC\nHxxvRi\nAvVTqb\nuRySvw")
	fmt.Println(Oper(HorMirror,"lVHt\nJVhv\nCSbg\nyeCt")== "yeCt\nCSbg\nJVhv\nlVHt")
	fmt.Println(Oper(VertMirror,"abcd\nefgh\nijkl\nmnop") == "dcba\nhgfe\nlkji\nponm")
	fmt.Println(Oper(HorMirror,"abcd\nefgh\nijkl\nmnop") == "mnop\nijkl\nefgh\nabcd")
}
