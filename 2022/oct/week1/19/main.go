package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	var result *ListNode
	for i, node := range nodes {
		if i == len(nodes) - n {
			continue
		}

		result = add(result, node)
	}
	return result
}    

func add(head *ListNode, target *ListNode) *ListNode {
	if target != nil {
		target.Next = nil
	}

	if head == nil {
		return target
	}

	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = target
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
	head.Next.Next = &ListNode{Val: 3}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 1}
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
	fmt.Printf("head (%p) = %v\n", head, head.Val)
	printList(head.Next)
}

func main() {
	// result: [1,2,3,5]
	// head := makeList1()
	// n := int(2)

	// result: [2,3]
	// head := makeList2()
	// n := int(3)

	// result: [3,2]
	head := makeList3()
	n := int(1)

	// result: 
	// head := makeList()
	// n := int(0)

	result := removeNthFromEnd(head, n)
	printList(result)
}

