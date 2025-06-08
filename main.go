package main

import (
	"fmt"
	"sort"
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

func containsDuplicate(nums []int) bool {
	const k int = 2
	sortedNums := nums
	sort.Ints(sortedNums)
	counter := 0
	i := 0
	for i < len(sortedNums)-1 {
		if sortedNums[i] == sortedNums[i+1] {
			counter += 1
		} else {
			counter = 0
		}
		if counter >= k {
			return true
		}
		i += 1
	}
	return false
}

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
		fmt.Println("Has dups =", containsDuplicate(testCase))
		fmt.Println(testCase)
	}
}
