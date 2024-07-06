package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	var node *ListNode
	var result *ListNode
	sum := int(0)

	for head != nil {
		if head.Val > 0 {
			sum += head.Val
		} else {
			if node == nil {
				node = &ListNode{Val: sum}
				result = node
			} else {
				node.Next = &ListNode{Val: sum}
				node = node.Next
			}

			sum = 0
		}

		head = head.Next
	}

	return result.Next
}

func makeList1() *ListNode {
	head := &ListNode{Val: 0}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 0}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
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
	// result: [4,11]
	// head := makeList1()

	// result: [1,3,4]
	head := makeList2()

	// result: []
	// head := makeList()

	result := mergeNodes(head)
	printList(result)
}

