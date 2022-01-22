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
		// fmt.Printf("slow = %v, fast = %v\n", slow, fast)
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fast = head
			for slow != fast {
				// fmt.Printf("after has cycle: slow = %v, fast = %v\n", slow, fast)
				slow = slow.Next
				fast = fast.Next
			}
			return fast
		}
	}

	return nil
}

func makeList1() *ListNode {
	root := &ListNode{Val: 3}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next = &ListNode{Val: -4}
	root.Next.Next.Next.Next = root.Next
	return root
}

func makeList2() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = root.Next
	return root
}

func makeList3() *ListNode {
	root := &ListNode{Val: 1}
	return root
}

func makeList4() *ListNode {
	root := &ListNode{Val: 3}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next = &ListNode{Val: -4}
	root.Next.Next.Next.Next = &ListNode{Val: -5}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	root.Next.Next.Next.Next.Next.Next = &ListNode{Val: 100}
	root.Next.Next.Next.Next.Next.Next.Next = root.Next.Next.Next.Next
	return root
}

func main() {
	// result: indexed 1 node
	// head := makeList1()

	// result: indexed 0 node
	// head := makeList2()

	// result: nil
	// head := makeList3()

	// result: indexed 4 node
	head := makeList4()

	// result: 
	// head := makeList()

	result := detectCycle(head)
	fmt.Printf("result = %v\n", result)
}

