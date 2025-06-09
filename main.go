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

func main() {
	fmt.Println("Hello world! This is the main driver program")
}
