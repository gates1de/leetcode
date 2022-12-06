package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddHead := head
	evenHead := head.Next

	oddNode := oddHead
	evenNode := evenHead

	head = head.Next
	head = head.Next
	isOdd := true
	for head != nil {
		if isOdd {
			oddNode.Next = head
			oddNode = oddNode.Next
		} else {
			evenNode.Next = head
			evenNode = evenNode.Next
		}
		head = head.Next
		isOdd = !isOdd
	}

	evenNode.Next = nil
	if oddNode == nil {
		oddNode = evenHead
	} else {
		oddNode.Next = evenHead
	}
	return oddHead
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
	head := &ListNode{Val: 2}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
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

func printList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Printf("node (%p): %v\n", head, head)
	printList(head.Next)
}

func main() {
	// result: [1,3,5,2,4]
	// head := makeList1()

	// result: [2,3,6,7,1,5,4]
	// head := makeList2()

	// result: [1]
	head := makeList3()

	// result: 
	// head := makeList()

	result := oddEvenList(head)
	printList(result)
}

