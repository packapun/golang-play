package main

import (
	"fmt"
)

func runTestCases() {
	testCases := [][]int{
		{1, 2, 3, 4, 4},
		{1, 2, 3, 4, 5, 5, 5},
		{6, 5, 4, 3, 2, 1, 2, 3, 4, 1, 2},
		{1},
		{0, 0, 0, 0, 1},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test case %d : %v \n", i+1, testCase)
		fmt.Println(testCase)
	}
}

func validateRow(rowNumber int, board [][]byte) bool {
	if rowNumber >= len(board) {
		return false
	}
	hashMap := make(map[byte]bool)
	fmt.Println("Inspecting row", string(board[rowNumber]))
	for _, v := range board[rowNumber] {
		if isEmptyValue(v) {
			continue
		}
		if _, ok := hashMap[v]; ok {
			return false
		} else {
			hashMap[v] = true
		}
	}
	return true
}

func validateColumn(colNumber int, board [][]byte) bool {
	if colNumber >= len(board) {
		return false
	}
	column := []byte{}
	hashMap := make(map[byte]bool)

	for _, row := range board {
		column = append(column, row[colNumber])
	}

	fmt.Println("Inspecting column", string(column))
	for _, v := range column {
		if isEmptyValue(v) {
			continue
		}
		if _, ok := hashMap[v]; ok {
			return false
		} else {

			hashMap[v] = true
		}
	}
	return true
}

func isEmptyValue(val byte) bool {
	return val == '.'
}

func isValidSudoku(board [][]byte) bool {
	size := len(board)
	if size == 0 {
		return false
	}

	for i := range size {
		if !validateRow(i, board) || !validateColumn(i, board) {
			return false
		}
	}
	subgridSize := size / 3
	if subgridSize > 1 {

		for k := range size {
			subgrid := [][]byte{}
			// fmt.Println("K = ", k)
			for i := range subgridSize {
				subgrid = append(subgrid, make([]byte, subgridSize))
				for j := range subgridSize {
					valX := i + 3*(k%3)
					valY := j + 3*(k/3)
					// fmt.Printf("Value of (x,y) = (%d, %d)\n ", valX, valY)
					subgrid[i][j] = board[valX][valY]
				}
			}

			for _, row := range subgrid {
				fmt.Println(string(row))
			}
			fmt.Println("---")

			if !isValidSudoku(subgrid) {
				return false
			}
		}

	}
	return true
}

func main() {
	fmt.Println("Hello world! This is the main driver program")
	board := [][]byte{
		{'1', '.', '.', '.', '.', '.', '.', '.', '4'},
		{'2', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'3', '.', '7', '.', '.', '.', '.', '.', '.'},

		{'4', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '8', '.'},
		{'6', '.', '.', '.', '.', '.', '.', '.', '.'},

		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'8', '.', '.', '.', '.', '.', '.', '.', '9'},
		{'9', '.', '.', '.', '1', '.', '1', '.', '2'},
	}

	fmt.Println(isValidSudoku(board))
}
