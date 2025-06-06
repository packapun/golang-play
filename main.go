package main

import (
	"fmt"

	"github.com/packapun/golang-play/utility"
)

func main() {
	fmt.Println("Hello world!")
	input := []int{1}
	ret := utility.SingleNumber(input)
	fmt.Println("Return value of x = ", ret)

}
