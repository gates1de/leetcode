package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	valuesToRemove := make(map[int]bool)
	for _, num := range nums {
		valuesToRemove[num] = true
	}

	for head != nil && valuesToRemove[head.Val] {
		head = head.Next
	}

	if head == nil {
		return nil
	}

	current := head
	for current != nil && current.Next != nil {
		if valuesToRemove[current.Next.Val] {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func makeList1() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	return head
}

func makeList() *ListNode {
	var head *ListNode
	return head
}

func printList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head)
	printList(head.Next)
}

func main() {
	// result: [4,5]
	// nums := []int{1,2,3}
	// head := makeList1()

	// result: [2,2,2]
	// nums := []int{1}
	// head := makeList2()

	// result: [1,2,3,4]
	nums := []int{5}
	head := makeList3()

	// result: []
	// nums := []int{}
	// head := makeList()

	result := modifiedList(nums, head)
	printList(result)
}

