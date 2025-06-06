package utility

func SingleNumber(nums []int) int {
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
