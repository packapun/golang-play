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

func intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	countMap := make(map[int]int)

	for _, num := range nums1 {
		countMap[num]++
	}

	for _, num := range nums2 {
		if count, exists := countMap[num]; exists && count > 0 {
			result = append(result, num)
			countMap[num]--
		}
	}
	return result
}

func plusOne(digits []int) []int {
	n := len(digits)
	if n == 0 {
		return digits
	}
	i := n - 1
	for i >= 0 && digits[i] == 9 {
		digits[i] = 0
		i -= 1
	}
	if i >= 0 {
		digits[i] += 1
	} else {
		// The digits of the given number are all 9s
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}

func moveZeroes(nums []int) {
	count := 0
	i := 0
	j := 0
	n := len(nums)
	for i < n && j < n {
		if nums[i] == 0 {
			count++
		} else {
			nums[j] = nums[i]
			j += 1
		}

		i += 1
	}
	for count > 0 {
		nums[n-count] = 0
		count--
	}
}

// Rotate 2D array by 90 degrees
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := range len(matrix) {
		reverse(matrix[i])
	}
}

func reverse(nums []int) {
	start := 0
	end := len(nums) - 1

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
