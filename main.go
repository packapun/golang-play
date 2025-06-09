package main

import (
	"fmt"
)

func reverse(x int) int {
	fmt.Println("Reversing", x)
	reversedNum := 0
	isNegative := x < 0
	if isNegative {
		x = -x
	}

	for x > 0 {
		reversedNum = reversedNum*10 + x%10
		if reversedNum > math.MaxInt32 {
			return 0
		}
		x = x / 10
	}
	if isNegative {
		return -reversedNum
	}
	return reversedNum
}

func firstUniqChar(s string) int {
	hashMap := make(map[byte]int)
	stringSlice := []byte(s)
	for _, letter := range stringSlice {
		hashMap[letter]++
	}

	for i, letter := range stringSlice {
		if count, ok := hashMap[letter]; ok && count == 1 {
			return i
		}
	}
	return -1
}


func main() {
	fmt.Println("Hello world! This is the main driver program")
}
