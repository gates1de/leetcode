package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func isPalindrome(head *ListNode) bool {    
	if head == nil {
		return false
	}

	var tail *ListNode
	tmp := head
	for tmp != nil {
		tail = insert(tail, tmp)
		tmp = tmp.Next
	}

	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}

	return true
}

func insert(head *ListNode, target *ListNode) *ListNode {
	if target == nil {
		return head
	}
	if head == nil {
		return &ListNode{Val: target.Val}
	}

	result := &ListNode{Val: target.Val}
	result.Next = head
	return result
}

func makeList1() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next = &ListNode{Val: 1}
	return head
}

func makeList4() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next = &ListNode{Val: 2}
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
	fmt.Printf("head (%p) = %v\n", head, head)
	printList(head.Next)
}

func main() {
	// result: true
	// head := makeList1()

	// result: false
	// head := makeList2()

	// result: true
	// head := makeList3()

	// result: false
	head := makeList4()

	// result: 
	// head := makeList()

	result := isPalindrome(head)
	fmt.Printf("result = %v\n", result)
}

