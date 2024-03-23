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

func makeList() *ListNode {
	var head *ListNode
	return head
}

func main() {
	// result: true
	// head := makeList1()

	// result: false
	head := makeList2()

	// result: 
	// head := makeList()

	result := isPalindrome(head)
	fmt.Printf("result = %v\n", result)
}

