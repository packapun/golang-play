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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to simplify edge cases
	dummy := &ListNode{}
	current := dummy

	// Compare nodes from both lists and link the smaller one
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Attach remaining nodes (at most one list will have remaining nodes)
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Return the head (skip dummy node)
	return dummy.Next
}

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false 
    }
    p1 := head
    p2 := head 
    hasLoop := false
    for p2 != nil && p1 != nil && p1.Next != nil {
        p1 = p1.Next
        p1 = p1.Next
        p2 = p2.Next
        
        if p1 == p2 {
            hasLoop = true
            break 
        }
        
    }
    return hasLoop
}

 type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}


func maxDepth(root *TreeNode) int {
    return calculateDepth(root) 
}

func calculateDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    return 1 + max(calculateDepth(root.Left), calculateDepth(root.Right))
}

func main() {
	fmt.Println("Hello world! This is the main driver program")
}
