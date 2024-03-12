package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	front := &ListNode{Val: 0, Next: head}
	start := front

	for start != nil {
		prefixSum := int(0)
		end := start.Next

		for end != nil {
			prefixSum += end.Val

			if prefixSum == 0 {
				start.Next = end.Next
			}

			end = end.Next
		}

		start = start.Next
	}

	return front.Next
}

func makeList1() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: -3}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 1}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: -3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: -3}
	head.Next.Next.Next.Next = &ListNode{Val: -2}
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
	// result: [3,1]
	// head := makeList1()

	// result: [1,2,4]
	// head := makeList2()

	// result: [1]
	head := makeList3()

	// result: 
	// head := makeList()

	result := removeZeroSumSublists(head)
	printList(result)
}

