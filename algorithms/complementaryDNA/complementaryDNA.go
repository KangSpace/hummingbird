/*
 date: 2019-04-15 11:13:01
Deoxyribonucleic acid (DNA) is a chemical found in the nucleus of cells and carries the "instructions" for the development and functioning of living organisms.
If you want to know more http://en.wikipedia.org/wiki/DNA
In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G". You have function with one side of the DNA (string, except for Haskell); you need to get the other complementary side. DNA strand is never empty or there is no DNA at all (again, except for Haskell).
More similar exercise are found here http://rosalind.info/problems/list-view/ (source)
DNA_strand ("ATTGC") # return "TAACG"
DNA_strand ("GTAT") # return "CATA"
 */
package main

import (
	"fmt"
	"strings"
)

func DNAStrand(dna string) string {
	var dnaCharMap = map[string]string{"A":"T","C":"G","T":"A","G":"C"}
	a:=""
	for _,s:= range dna[:]{
		a+=dnaCharMap[string(s)]
	}
	return a
}

var dnsReplacer *strings.Replacer = strings.NewReplacer("A","T","C","G","T","A","G","C")
func DNAStrand2(dna string) string {
	return dnsReplacer.Replace(dna)
}

func main() {
	fmt.Println(DNAStrand("AAAA"))
	fmt.Println(DNAStrand("ATTGC"))
	fmt.Println(DNAStrand("GTAT"))
	fmt.Println(DNAStrand2("GTAT"))
}