package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world! This is the main driver program")

	randomNumbers := []int{3, 1, 5, 2}
	sort.Ints(randomNumbers)
	fmt.Println(randomNumbers)

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
