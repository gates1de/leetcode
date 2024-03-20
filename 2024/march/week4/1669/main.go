package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1 == nil && list2 == nil {
		return nil
	}

	var start *ListNode
	end := list1

	for i := 0; i < b; i++ {
		if i == a - 1 {
			start = end
		}

		end = end.Next
	}

	start.Next = list2

	for list2.Next != nil {
		list2 = list2.Next
	}

	list2.Next = end.Next
	end.Next = nil
        
	return list1
}

func makeList1() (*ListNode, *ListNode) {
	list1 := &ListNode{Val: 10}
	list1.Next = &ListNode{Val: 1}
	list1.Next.Next = &ListNode{Val: 13}
	list1.Next.Next.Next = &ListNode{Val: 6}
	list1.Next.Next.Next.Next = &ListNode{Val: 9}
	list1.Next.Next.Next.Next.Next = &ListNode{Val: 5}

	list2 := &ListNode{Val: 1000000}
	list2.Next = &ListNode{Val: 1000001}
	list2.Next.Next = &ListNode{Val: 1000002}

	return list1, list2
}

func makeList2() (*ListNode, *ListNode) {
	list1 := &ListNode{Val: 0}
	list1.Next = &ListNode{Val: 1}
	list1.Next.Next = &ListNode{Val: 2}
	list1.Next.Next.Next = &ListNode{Val: 3}
	list1.Next.Next.Next.Next = &ListNode{Val: 4}
	list1.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	list1.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	list2 := &ListNode{Val: 1000000}
	list2.Next = &ListNode{Val: 1000001}
	list2.Next.Next = &ListNode{Val: 1000002}
	list2.Next.Next.Next = &ListNode{Val: 1000003}
	list2.Next.Next.Next.Next = &ListNode{Val: 1000004}

	return list1, list2
}

func makeList() (*ListNode, *ListNode) {
	var list1, list2 *ListNode
	return list1, list2
}

func printList(head *ListNode) {
	if head == nil {
		return
	}

	fmt.Println(head)
	printList(head.Next)
}

func main() {
	// result: [10,1,13,1000000,1000001,1000002,5]
	// list1, list2 := makeList1()
	// a := int(3)
	// b := int(4)

	// result: [0,1,1000000,1000001,1000002,1000003,1000004,6]
	list1, list2 := makeList2()
	a := int(2)
	b := int(5)

	// result: []
	// list1, list2 := makeList()
	// a := int(0)
	// b := int(0)

	result := mergeInBetween(list1, a, b, list2)
	printList(result)
}

