/*
 date: 2019-04-15 11:54:00
To participate in a prize draw each one gives his/her firstname.
Each letter of a firstname has a value which is its rank in the English alphabet. A and a have rank 1, B and b rank 2 and so on.
The length of the firstname is added to the sum of these ranks hence a number n. An array of random weights is linked to the firstnames and each n is multiplied by its corresponding weight to get what they call a winning number.
Example: names: COLIN,AMANDBA,AMANDAB,CAROL,PauL,JOSEPH weights: [1, 4, 4, 5, 2, 1]
PAUL -> n = length of firstname + 16 + 1 + 21 + 12 = 4 + 50 -> 54 The weight associated with PAUL is 2 so Paul's winning number is 54 * 2 = 108.
Now one can sort the firstnames in decreasing order of the winning numbers. When two people have the same winning number sort them alphabetically by their firstnames.
#Task:
parameters: st a string of firstnames, we an array of weights, n a rank
return: the firstname of the participant whose rank is n (ranks are numbered from 1)
#Example: names: COLIN,AMANDBA,AMANDAB,CAROL,PauL,JOSEPH weights: [1, 4, 4, 5, 2, 1] n: 4
The function should return: PauL
Note:
If st is empty return "No participants".
If n is greater than the number of participants then return "Not enough participants".
See Examples Test Cases for more examples.
FUNDAMENTALSSTRINGSSORTINGALGORITHMS
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

type namePoint struct {
	name  string
	point int
}
type NamePoints []*namePoint

func (nps NamePoints) Len() int { return len(nps) }
func (nps NamePoints) Swap(i, j int) { nps[i], nps[j] = nps[j], nps[i] }
func (nps NamePoints) Less(i, j int) bool { return nps[i].point > nps[j].point || (nps[i].point == nps[j].point &&  nps[i].name < nps[j].name)}

func NthRank(st string, we []int, n int) string {
	sts := strings.Split(st, ",")
	if len(st) == 0 {
		return "No participants"
	}
	if len(sts) != len(we) || len(sts) < n {
		return "Not enough participants"
	}
	var namePoints = make([]*namePoint, 0)
	//var nameMap = make(map[string]int)
	for i, s := range sts {
		name := strings.ToUpper(s)
		point := len(name)
		for _, n := range name {
			point += int(n) - 64
		}
		point = point * we[i]
		namePoints = append(namePoints, &namePoint{name: s, point: point})
	}
	sort.Sort(NamePoints(namePoints))
	if n > 0 {
		n = n - 1
	}
	return namePoints[n].name
}

func main() {
	st := "Elijah,Chloe,Elizabeth,Matthew,Natalie,Jayden"
	we := []int{1, 3, 5, 5, 3, 6}
	n := 2
	fmt.Println(NthRank(st, we, n))

	var names = []string{"COLIN", "AMANDBA", "AMANDAB", "CAROL", "PauL", "JOSEPH"}
	var weights = []int{1, 4, 4, 5, 2, 1}
	fmt.Println(NthRank(strings.Join(names, ","), weights, 10))
	fmt.Println("a" > "b")
}
