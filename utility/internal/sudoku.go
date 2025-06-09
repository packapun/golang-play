package internal

import "fmt"

// Your original code with issues highlighted in comments
func validateRowOriginal(rowNumber int, board [][]byte) bool {
	if rowNumber >= len(board) {
		return false
	}
	hashMap := make(map[byte]bool)
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

func validateColumnOriginal(colNumber int, board [][]byte) bool {
	if colNumber >= len(board) {
		return false
	}
	column := []byte{} // Unnecessary slice creation
	hashMap := make(map[byte]bool)
	for _, row := range board {
		column = append(column, row[colNumber])
	}
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

func isValidSudokuOriginal(board [][]byte) bool {
	size := len(board)
	if size == 0 {
		return false
	}
	for i := range size {
		if !validateRowOriginal(i, board) || !validateColumnOriginal(i, board) {
			return false
		}
	}
	subgridSize := size / 3
	if subgridSize > 1 { // This condition is problematic
		for k := range size {
			subgrid := [][]byte{} // Unnecessary recursion approach
			for i := range subgridSize {
				subgrid = append(subgrid, make([]byte, subgridSize))
				for j := range subgridSize {
					valX := i + 3*(k%3)
					valY := j + 3*(k/3)
					subgrid[i][j] = board[valX][valY]
				}
			}
			if !isValidSudokuOriginal(subgrid) { // Recursive call is overkill
				return false
			}
		}
	}
	return true
}

// IMPROVED VERSIONS

// Cleaner row validation
func validateRow(rowNumber int, board [][]byte) bool {
	if rowNumber >= len(board) {
		return false
	}

	seen := make(map[byte]bool)
	for _, cell := range board[rowNumber] {
		if cell == '.' {
			continue
		}
		if seen[cell] {
			return false
		}
		seen[cell] = true
	}
	return true
}

// Improved column validation (no unnecessary slice)
func validateColumn(colNumber int, board [][]byte) bool {
	if colNumber >= len(board) || len(board) == 0 || colNumber >= len(board[0]) {
		return false
	}

	seen := make(map[byte]bool)
	for row := 0; row < len(board); row++ {
		cell := board[row][colNumber]
		if cell == '.' {
			continue
		}
		if seen[cell] {
			return false
		}
		seen[cell] = true
	}
	return true
}

// Direct 3x3 box validation (no recursion needed)
func validate3x3Box(startRow, startCol int, board [][]byte) bool {
	seen := make(map[byte]bool)

	for row := startRow; row < startRow+3; row++ {
		for col := startCol; col < startCol+3; col++ {
			cell := board[row][col]
			if cell == '.' {
				continue
			}
			if seen[cell] {
				return false
			}
			seen[cell] = true
		}
	}
	return true
}

// Main validation function - cleaner approach
func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 || len(board[0]) != 9 {
		return false
	}

	// Validate all rows and columns
	for i := 0; i < 9; i++ {
		if !validateRow(i, board) || !validateColumn(i, board) {
			return false
		}
	}

	// Validate all 3x3 boxes
	for boxRow := 0; boxRow < 9; boxRow += 3 {
		for boxCol := 0; boxCol < 9; boxCol += 3 {
			if !validate3x3Box(boxRow, boxCol, board) {
				return false
			}
		}
	}

	return true
}

// MOST OPTIMAL: Single-pass solution
func isValidSudokuOptimal(board [][]byte) bool {
	// Use sets to track seen values in rows, columns, and boxes
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	// Initialize maps
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			cell := board[row][col]
			if cell == '.' {
				continue
			}

			// Calculate which 3x3 box this cell belongs to
			boxIndex := (row/3)*3 + col/3

			// Check if number already exists in row, column, or box
			if rows[row][cell] || cols[col][cell] || boxes[boxIndex][cell] {
				return false
			}

			// Mark as seen
			rows[row][cell] = true
			cols[col][cell] = true
			boxes[boxIndex][cell] = true
		}
	}

	return true
}

func main() {
	// Test case 1: Valid Sudoku
	validBoard := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	// Test case 2: Invalid Sudoku (duplicate in row)
	invalidBoard := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, // duplicate 8 in column 0
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println("Testing valid board:")
	fmt.Println("Original approach:", isValidSudokuOriginal(validBoard))
	fmt.Println("Improved approach:", isValidSudoku(validBoard))
	fmt.Println("Optimal approach:", isValidSudokuOptimal(validBoard))

	fmt.Println("\nTesting invalid board:")
	fmt.Println("Original approach:", isValidSudokuOriginal(invalidBoard))
	fmt.Println("Improved approach:", isValidSudoku(invalidBoard))
	fmt.Println("Optimal approach:", isValidSudokuOptimal(invalidBoard))

	// Test edge cases
	fmt.Println("\nTesting edge cases:")
	emptyBoard := [][]byte{}
	fmt.Println("Empty board:", isValidSudoku(emptyBoard))
}
