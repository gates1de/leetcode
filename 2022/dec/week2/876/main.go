package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	nodes := make([]*ListNode, 0, 100)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	return nodes[len(nodes) / 2]
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
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
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
	fmt.Printf("node (%p): %v\n", head, head)
	printList(head.Next)
}

func main() {
	// result: [3,4,5]
	// head := makeList1()

	// result: [4,5,6]
	// head := makeList2()

	// result: [1]
	head := makeList3()

	// result: 
	// head := makeList()

	result := middleNode(head)
	printList(result)
}

