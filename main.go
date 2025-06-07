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
func main() {
	fmt.Println("Hello world! This is the main driver program")
	prices := []int{1, 1, 1, 1, 1, 1, 13}
	fmt.Println(maxProfit(prices))

}
