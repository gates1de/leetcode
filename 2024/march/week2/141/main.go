package main
import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    slow := head
    fast := head.Next
    for slow.Next != nil && fast != nil && fast.Next != nil {
        if slow == fast {
            return true
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return false
}

func makeList1() *ListNode {
   head := &ListNode{Val: 3}
   head.Next = &ListNode{Val: 2}
   head.Next.Next = &ListNode{Val: 0}
   head.Next.Next.Next = &ListNode{Val: -4}
   head.Next.Next.Next.Next = head.Next
   return head
}

func makeList2() *ListNode {
   head := &ListNode{Val: 1}
   head.Next = &ListNode{Val: 2}
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

func main() {
	// results: true
	// head := makeList1()

	// results: true
	// head := makeList2()

	// results: false
	head := makeList3()

	// results: 
	// head := makeList()

	result := hasCycle(head)
	fmt.Printf("result = %v\n", result)
}

