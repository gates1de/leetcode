package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	lefts := []*ListNode{}
	rights := []*ListNode{}
	tmp := head
	for tmp != nil {
		if tmp.Val < x {
			lefts = append(lefts, tmp)
		} else {
			rights = append(rights, tmp)
		}
		tmp = tmp.Next
	}

	var result *ListNode
	var resultHead *ListNode
	for len(lefts) > 0 {
		var targetNode *ListNode
		targetNode = lefts[0]
		lefts = lefts[1:]

		if result == nil {
			result = targetNode
			resultHead = result
		} else {
			result.Next = targetNode
			result = result.Next
		}
		result.Next = nil
	}

	for len(rights) > 0 {
		var targetNode *ListNode
		targetNode = rights[0]
		rights = rights[1:]

		if result == nil {
			result = targetNode
			resultHead = result
		} else {
			result.Next = targetNode
			result = result.Next
		}
		result.Next = nil
	}

	return resultHead
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
	head := &ListNode{Val: 0}
	return head
}

func printList(head *ListNode) {
	if head == nil {
		fmt.Println("end")
		return
	}
	fmt.Printf("%p: %v\n", head, head)
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

