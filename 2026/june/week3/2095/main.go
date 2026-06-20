package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var prev *ListNode
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next
	return head
}

func makeList1() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 2}
	head.Next = &ListNode{Val: 1}
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
	// result: [1,3,4,1,2,6]
	// head := makeList1()

	// result: [1,2,4]
	// head := makeList2()

	// result: [2]
	head := makeList3()

	// result: 
	// head := makeList()

	result := deleteMiddle(head)
	printList(result)
}
