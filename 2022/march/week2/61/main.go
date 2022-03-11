package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	tmp := head
	count := int(0)
	for tmp != nil {
		tmp = tmp.Next
		count++
	}

	k %= count

	for i := 0; i < k; i++ {
		head = rotate(head)
	}

	return head
}

func rotate(head *ListNode) *ListNode {
	tmp := head
	var pre *ListNode
	for head.Next != nil {
		pre = head
		head = head.Next
	}
	if pre != nil {
		pre.Next = nil
	}
	head.Next = tmp
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
	head := &ListNode{Val: 0}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	return head
}

func makeList3() *ListNode {
	return nil
}

func makeList4() *ListNode {
	head := &ListNode{Val: 1}
	return head
}

func printList(head *ListNode) {
	if head == nil {
		fmt.Println("end")
		return
	}
	fmt.Println(head)
	printList(head.Next)
}

func main() {
	// result: [4,5,1,2,3]
	// head := makeList1()
	// k := int(2)

	// result: [2,0,1]
	// head := makeList2()
	// k := int(4)

	// result: []
	// head := makeList3()
	// k := int(0)

	// result: [1]
	head := makeList4()
	k := int(8)

	// result: 
	// head := makeList()
	// k := int(0)

	result := rotateRight(head, k)
	printList(result)
}

