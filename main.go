package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	input := nums
	i := 0
	j := 0
	// Loop starting from index [0, n - 2]
	for j < len(nums) {
		if j+1 < len(nums) && input[j] == input[j+1] {
			j += 1
			continue
		}
		input[i] = input[j]
		i += 1
		j += 1
	}
	return i
}

func main() {
	fmt.Println("Hello world! This is the main driver program")
	input := []int{1}
	fmt.Println(removeDuplicates(input))

}
