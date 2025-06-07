package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	profit := 0

	// Key insight: To maximize profit, we should capture every upward price movement
	// This is equivalent to buying at every local minimum and selling at every local maximum
	for i := 1; i < len(prices); i++ {
		// If price increased from yesterday, we profit from that increase
		// (buy yesterday, sell today)
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

func rotateRight(input []int) []int {
	n := len(input)
	if n < 2 {
		return input
	}
	x := input[n-1]
	for j := n - 1; j > 0; j-- {
		input[j] = input[j-1]
	}
	input[0] = x
	return input
}

func rotate(nums []int, k int) {
	for _ = range k {
		nums = rotateRight(nums)
	}
}

func main() {
	fmt.Println("Hello world! This is the main driver program")

	testCases := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{6, 5, 4, 3, 2},
		{1},
		{0, 0, 0, 0, 1},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test case %d : %v \n", i+1, testCase)
		rotate(testCase, 3)
		fmt.Printf("Test case %d : %v \n", i+1, testCase)
	}

}
