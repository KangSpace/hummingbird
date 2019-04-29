/*
 date: 2019-04-29 11:33:07
 Sudoku Background
Sudoku is a game played on a 9x9 grid. The goal of the game is to fill all cells of the grid with digits from 1 to 9, so that each column, each row, and each of the nine 3x3 sub-grids (also known as blocks) contain all of the digits from 1 to 9.
(More info at: http://en.wikipedia.org/wiki/Sudoku)

Sudoku Solution Validator
Write a function validSolution/ValidateSolution/valid_solution() that accepts a 2D array representing a Sudoku board, and returns true if it is a valid solution, or false otherwise. The cells of the sudoku board may also contain 0's, which will represent empty cells. Boards containing one or more zeroes are considered to be invalid solutions.

The board is always 9 cells by 9 cells, and every cell only contains integers from 0 to 9.

Examples
validSolution([
  [5, 3, 4, 6, 7, 8, 9, 1, 2],
  [6, 7, 2, 1, 9, 5, 3, 4, 8],
  [1, 9, 8, 3, 4, 2, 5, 6, 7],
  [8, 5, 9, 7, 6, 1, 4, 2, 3],
  [4, 2, 6, 8, 5, 3, 7, 9, 1],
  [7, 1, 3, 9, 2, 4, 8, 5, 6],
  [9, 6, 1, 5, 3, 7, 2, 8, 4],
  [2, 8, 7, 4, 1, 9, 6, 3, 5],
  [3, 4, 5, 2, 8, 6, 1, 7, 9]
]); // => true
validSolution([
  [5, 3, 4, 6, 7, 8, 9, 1, 2],
  [6, 7, 2, 1, 9, 0, 3, 4, 8],
  [1, 0, 0, 3, 4, 2, 5, 6, 0],
  [8, 5, 9, 7, 6, 1, 0, 2, 0],
  [4, 2, 6, 8, 5, 3, 7, 9, 1],
  [7, 1, 3, 9, 2, 4, 8, 5, 6],
  [9, 0, 1, 5, 3, 7, 2, 1, 4],
  [2, 8, 7, 4, 1, 9, 6, 3, 5],
  [3, 0, 0, 4, 8, 1, 1, 7, 9]
]); // => false
*/
package main

import (
	"fmt"
	"strconv"
)

type SumObj struct {
	Sum   int
	slice []int
}

func ValidateSolution1(m [][]int) bool {
	sum := (1 + 9) * 9 / 2
	vMap := make(map[int]*SumObj)
	v33Map:=make(map[string]*SumObj)
	for i, mm := range m {
		lineSlice := &SumObj{0, []int{}}
		for j, nn := range mm {
			if nn == 0 {
				return false
			}
			v := vMap[j]
			if v == nil {
				v = &SumObj{0, []int{}}
				vMap[j] = v
			}
			v.Sum += nn
			lineSlice.Sum += nn
			var v33Key = strconv.Itoa(i / 3)+"-"+strconv.Itoa(j / 3)
			v33 := v33Map[v33Key]
			if v33 == nil {
				v33 = &SumObj{0, []int{}}
				v33Map[v33Key] = v33
			}
			v33.Sum+=nn
			// line || vertical not match
			if (j == len(mm)-1 && lineSlice.Sum != sum) || //Contain(lineSlice.slice, nn)) ||
				 (i == len(m)-1 && v.Sum != sum) || //Contain(v.slice, nn)) ||
				(i>0 && j>0 && (i+1)%3==0 && (j+1)%3==0 && v33.Sum != sum){
				return false
			}
			v.slice = append(v.slice, nn)
			lineSlice.slice = append(lineSlice.slice, nn)
			v33.slice = append(v33.slice, nn)
		}
	}
	return true
}
func Contain(is []int, i int) bool {
	for _, v := range is {
		if v == i {
			return true
		}
	}
	return false
}

func ValidateSolution(m [][]int) bool {
	sum := (1 + 9) * 9 / 2
	vMap := make(map[int]int)
	v33Map:=make(map[string]int)
	for i, mm := range m {
		lineSum :=0
		for j, nn := range mm {
			if nn == 0 {
				return false
			}
			vMap[j]+=nn
			lineSum += nn
			var v33Key = strconv.Itoa(i / 3)+"-"+strconv.Itoa(j / 3)
			v33Map[v33Key] += nn
			// line || vertical not match
			if (j == len(mm)-1 && lineSum != sum) ||
				 (i == len(m)-1 && vMap[j] != sum) ||
				(i>0 && j>0 && (i+1)%3==0 && (j+1)%3==0 && v33Map[v33Key] != sum){
				return false
			}
		}
	}
	return true
}

func main() {
	var testTrue = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	var testFalse = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 0, 3, 4, 8},
		{1, 0, 0, 3, 4, 2, 5, 6, 0},
		{8, 5, 9, 7, 6, 1, 0, 2, 0},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 0, 1, 5, 3, 7, 2, 1, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 0, 0, 4, 8, 1, 1, 7, 9},
	}
	var test3 = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{2, 3, 1, 5, 6, 4, 8, 9, 7},
		{3, 1, 2, 6, 4, 5, 9, 7, 8},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{5, 6, 4, 8, 9, 7, 2, 3, 1},
		{6, 4, 5, 9, 7, 8, 3, 1, 2},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{8, 9, 7, 2, 3, 1, 5, 6, 4},
		{9, 7, 8, 3, 1, 2, 6, 4, 5}}
	fmt.Println(ValidateSolution(testTrue))
	fmt.Println(ValidateSolution(testFalse))
	fmt.Println(ValidateSolution(test3))
}
