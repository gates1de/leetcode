package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return []int{-1, -1}
	}

	prev := head
	current := head.Next
	index := 1
	firstCriticalIndex := -1
	previousCriticalIndex := -1
	minDistance := 0

	for current.Next != nil {
		if (current.Val < prev.Val && current.Val < current.Next.Val) ||
			(current.Val > prev.Val && current.Val > current.Next.Val) {
			if firstCriticalIndex == -1 {
				firstCriticalIndex = index
			} else {
				distance := index - previousCriticalIndex
				if minDistance == 0 || distance < minDistance {
					minDistance = distance
				}
			}
			previousCriticalIndex = index
		}

		index++
		prev = current
		current = current.Next
	}

	if minDistance == 0 {
		return []int{-1, -1}
	}

	return []int{minDistance, previousCriticalIndex - firstCriticalIndex}
}

func makeList1() *ListNode {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 1}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 5}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
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
	// result: [-1,-1]
	// head := makeList1()

	// result: [1,3]
	// head := makeList2()

	// result: [3,3]
	head := makeList3()

	// result: []
	// head := makeList()

	result := nodesBetweenCriticalPoints(head)
	fmt.Printf("result = %v\n", result)
}
