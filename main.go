package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	x := nums[0]
	for i := 1; i < len(nums); i++ {
		x = x ^ nums[i]
	}
	return x
}

func main() {
	fmt.Println("Hello world!")
	input := []int{1}
	ret := singleNumber(input)
	fmt.Println("Return value of x = ", ret)

}
