package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	carry := twice(head)
	if carry != 0 {
		head = &ListNode{Val: carry, Next: head}
	}

	return head
}

func twice(head *ListNode) int {
	if head == nil {
		return 0
	}

	doubledVal := head.Val * 2 + twice(head.Next)
	head.Val = doubledVal % 10
	return doubledVal / 10
}

func makeList1() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 8}
	head.Next.Next = &ListNode{Val: 9}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 9}
	head.Next = &ListNode{Val: 9}
	head.Next.Next = &ListNode{Val: 9}
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
	// result: [3,7,8]
	head := makeList1()

	// result: [1,9,9,8]
	// head := makeList2()

	// result: []
	// head := makeList()

	result := doubleIt(head)
	printList(result)
}

