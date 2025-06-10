package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 || head == nil {
		return head
	}

	size := 0
	p := head

	for p != nil {
		size++
		p = p.Next
	}

	if n > size {
		return head
	}

	positionToDelete := size - n
	p = head
	i := 0

	for i < positionToDelete-1 && p != nil {
		p = p.Next
		i++
	}

	// Delete p.Next
	temp := p.Next
	p.Next = p.Next.Next
	temp.Next = nil

	return head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	fmt.Println("Hello world! This is the main driver program")
}
