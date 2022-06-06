package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	m := map[*ListNode]bool{}
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if m[headB] {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

func makeList1() (*ListNode, *ListNode) {
	a := &ListNode{Val: 4}
	a.Next = &ListNode{Val: 1}
	a.Next.Next = &ListNode{Val: 8}
	a.Next.Next.Next = &ListNode{Val: 4}
	a.Next.Next.Next.Next = &ListNode{Val: 5}

	b := &ListNode{Val: 5}
	b.Next = &ListNode{Val: 6}
	b.Next.Next = &ListNode{Val: 1}
	b.Next.Next.Next = a.Next.Next
	b.Next.Next.Next.Next = a.Next.Next.Next
	b.Next.Next.Next.Next.Next = a.Next.Next.Next.Next
	return a, b
}

func makeList2() (*ListNode, *ListNode) {
	a := &ListNode{Val: 1}
	a.Next = &ListNode{Val: 9}
	a.Next.Next = &ListNode{Val: 1}
	a.Next.Next.Next = &ListNode{Val: 2}
	a.Next.Next.Next.Next = &ListNode{Val: 4}

	b := &ListNode{Val: 3}
	b.Next = a.Next.Next.Next
	b.Next.Next = a.Next.Next.Next.Next
	return a, b
}

func makeList3() (*ListNode, *ListNode) {
	a := &ListNode{Val: 2}
	a.Next = &ListNode{Val: 6}
	a.Next.Next = &ListNode{Val: 4}

	b := &ListNode{Val: 1}
	b.Next = &ListNode{Val: 5}
	return a, b
}

func printList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Printf("%p: %v\n", head, head)
	printList(head.Next)
}

func main() {
	// result: 8
	// headA, headB := makeList1()

	// result: 2
	// headA, headB := makeList2()

	// result: nil
	headA, headB := makeList3()

	// result: 
	// headA, headB := makeList()

	result := getIntersectionNode(headA, headB)
	printList(result)
}

