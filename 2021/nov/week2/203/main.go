package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	return helper(head, val)
}

func helper(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return head
	}

	head.Next = helper(head.Next, val)
	return head
}

func makeList1() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 6}
	root.Next.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	root.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	return root
}

func makeList2() *ListNode {
	var root *ListNode
	return root
}

func makeList3() *ListNode {
	root := &ListNode{Val: 7}
	root.Next = &ListNode{Val: 7}
	root.Next.Next = &ListNode{Val: 7}
	root.Next.Next.Next = &ListNode{Val: 7}
	return root
}

func makeList4() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 7}
	root.Next.Next = &ListNode{Val: 7}
	root.Next.Next.Next = &ListNode{Val: 7}
	return root
}

func printNode(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Printf("val = %v\n", head.Val)
	printNode(head.Next)
}

func main() {
	// result: [1,2,3,4,5]
	// head := makeList1()
	// val := int(6)

	// result: []
	// head := makeList2()
	// val := int(1)

	// result: []
	// head := makeList3()
	// val := int(7)

	// result: 
	// head := makeList()
	// val := int(0)

	// result: [1,7,7,7]
	head := makeList4()
	val := int(6)

	result := removeElements(head, val)
	printNode(result)
}

