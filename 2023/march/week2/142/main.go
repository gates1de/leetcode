package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

    slow := head
    fast := head
    for slow != nil {
        if fast == nil || fast.Next == nil {
            return nil
        }

        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            fast = head
            for slow != fast {
                slow = slow.Next
                fast = fast.Next
            }
            return fast
        }
    }

	return nil
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
	head.Next.Next = head
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
	// result: 1
	// head := makeList1()

	// result: 0
	// head := makeList2()

	// result: -1
	head := makeList3()

	// result: 
	// head := makeList()

	result := detectCycle(head)
	fmt.Printf("result = %v\n", result)
}

