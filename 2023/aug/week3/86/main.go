package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	before := &ListNode{}
	after := &ListNode{}
    beforeCursor := before
    afterCursor := after

    for head != nil {
        if head.Val < x {
            beforeCursor.Next = head
            beforeCursor = beforeCursor.Next
        } else {
            afterCursor.Next = head
            afterCursor = afterCursor.Next
        }
        head = head.Next
    }

    afterCursor.Next = nil
    beforeCursor.Next = after.Next

    return before.Next
}

func makeList1() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 2}
	head.Next = &ListNode{Val: 1}
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
	// result: [1,2,2,4,3,5]
	// head := makeList1()
	// x := int(3)

	// result: [1,2]
	head := makeList2()
	x := int(2)

	// result: 
	// head := makeList()
	// x := int(0)

	result := partition(head, x)
	printList(result)
}

