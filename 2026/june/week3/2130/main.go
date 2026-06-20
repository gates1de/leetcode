package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := reverseList(slow)
	result := int(0)
	for first := head; second != nil; first, second = first.Next, second.Next {
		sum := first.Val + second.Val
		if sum > result {
			result = sum
		}
	}

	return result
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func makeList1() *ListNode {
	head := &ListNode{Val: 5}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 100000}
	return head
}

func makeList() *ListNode {
	var head *ListNode
	return head
}

func main() {
	// result: 6
	// head := makeList1()

	// result: 7
	// head := makeList2()

	// result: 100001
	head := makeList3()

	// result: 
	// head := makeList()

	result := pairSum(head)
	fmt.Printf("result = %v\n", result)
}
