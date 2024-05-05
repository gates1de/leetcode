package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteNode(node *ListNode) {
	if node == nil {
		return
	}

	next := node.Next
	if next == nil {
		node = nil
	} else {
		*node = *next
	}
}

func helper(head *ListNode, node *ListNode) {
	for head != nil {
		if head == node {
			deleteNode(node)
		}
		head = head.Next
	}
}

func makeList1() *ListNode {
	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next = &ListNode{Val: 9}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
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
	// result: [4,1,9]
	// head := makeList1()
	// node := head.Next.Next

	// result: [4,5,9]
	// head := makeList1()
	// node := head.Next

	// result: [2,3]
	head := makeList2()
	node := head

	// result: 
	// head := makeList()
	// node := head

	helper(head, node)
	printList(head)
}

