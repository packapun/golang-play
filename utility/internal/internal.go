package internal

import "sort"

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
