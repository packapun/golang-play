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

func RemoveDuplicates(nums []int) int {
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

func Reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
