package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	nodes := []*ListNode{}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	return nodes[len(nodes) / 2]
}

func makeList1() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	return root
}

func makeList2() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	return root
}

func makeList3() *ListNode {
	root := &ListNode{Val: 1}
	return root
}

func makeList4() *ListNode {
	root := &ListNode{Val: 3}
	root.Next = &ListNode{Val: 2}
	return root
}

func printListNode(l *ListNode) {
    if l == nil {
        return
    }
    fmt.Printf("l = %v\n", l)
    printListNode(l.Next)
}

func main() {
	// result: [3,4,5]
	// head := makeList1()

	// result: [4,5,6]
	// head := makeList2()

	// result: [1]
	// head := makeList3()

	// result: [2]
	head := makeList4()

	// result: 
	// head := makeList()

	result := middleNode(head)
	printListNode(result)
}

