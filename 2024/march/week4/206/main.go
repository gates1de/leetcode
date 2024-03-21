package main
import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    current := head
    var prev *ListNode
    var next *ListNode
    for current != nil {
        next = current.Next
        current.Next = prev
        prev = current
        current = next
    }
    head = prev
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
	// result: [5,4,3,2,1]
	// head := makeList1()

	// result: [2,1]
	head := makeList2()

	// result: 
	// head := makeList()

	result := reverseList(head)
	printList(result)
}

