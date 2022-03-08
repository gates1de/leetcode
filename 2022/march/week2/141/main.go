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
		// fmt.Println(slow, fast)
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
	head.Next.Next = head.Next
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
	return head
}

func makeList4() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: -1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: -2}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next.Next = &ListNode{Val: -3}
	head.Next.Next.Next.Next.Next.Next = head.Next.Next.Next.Next.Next.Next
	return head
}

func printList(head *ListNode, printed map[*ListNode]bool) {
	if head == nil || printed[head] {
		return
	}
	fmt.Println(head)
	printed[head] = true
	printList(head.Next, printed)
}

func main() {
	// result: true
	// head := makeList1()

	// result: true
	// head := makeList2()

	// result: false
	// head := makeList3()

	// result: true
	head := makeList4()

	// result: 
	// head := makeList()

	// printList(head, map[*ListNode]bool{})

	result := hasCycle(head)
	fmt.Printf("result = %v\n", result)
}

