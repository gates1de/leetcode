package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	result := int(0)
	for head != nil {
		result *= 2
		if head.Val == 1 {
			result += 1
		}
		head = head.Next
	}
	return result
}

func makeList1() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 0}
	root.Next.Next = &ListNode{Val: 1}
	return root
}

func makeList2() *ListNode {
	root := &ListNode{Val: 0}
	return root
}

func makeList3() *ListNode {
	root := &ListNode{Val: 1}
	return root
}

func makeList4() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 0}
	root.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	return root
}

func makeList5() *ListNode {
	root := &ListNode{Val: 0}
	root.Next = &ListNode{Val: 0}
	return root
}

func main() {
	// result: 5
	// head := makeList1()

	// result: 0
	// head := makeList2()

	// result: 1
	// head := makeList3()

	// result: 18880
	// head := makeList4()

	// result: 0
	head := makeList5()

	// result: 
	// head := makeList()

	result := getDecimalValue(head)
	fmt.Printf("result = %v\n", result)
}

