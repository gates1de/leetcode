package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := removeNodes(head.Next)
	if head.Val < next.Val {
		return next
	}

	head.Next = next
	return head
}

func makeList1() *ListNode {
	head := &ListNode{Val: 5}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 13}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 8}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 1}
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
	// result: [13,8]
	// head := makeList1()

	// result: [1,1,1,1]
	head := makeList2()

	// result: []
	// head := makeList()

	result := removeNodes(head)
	printList(result)
}

