package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	var result *ListNode
	for head != nil {
		result = insert(head.Val, result)
		head = head.Next
	}
	return result
}

func insert(val int, head *ListNode) *ListNode {
	if head == nil {
		return &ListNode{Val: val, Next: nil}
	}
	if head.Next == nil {
		target := &ListNode{Val: val, Next: nil}
		if head.Val < val {
			head.Next = target
			return head
		}
		target.Next = head
		return target
	}

	if val <= head.Val {
		return &ListNode{Val: val, Next: head}
	}

	tmp := head
	pre := head
	for tmp != nil && tmp.Val < val {
		pre = tmp
		tmp = tmp.Next
	}
	target := &ListNode{Val: val, Next: tmp}
	if pre != tmp {
		pre.Next = target
	}
	return head
}

func makeList1() *ListNode {
	root := &ListNode{Val: 4}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next = &ListNode{Val: 3}
	return root
}

func makeList2() *ListNode {
	root := &ListNode{Val: -1}
	root.Next = &ListNode{Val: 5}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 0}
	return root
}

func makeList3() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 0}
	root.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	return root
}

func makeList4() *ListNode {
	root := &ListNode{Val: 0}
	return root
}

func makeList5() *ListNode {
	root := &ListNode{Val: 3}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 4}
	return root
}

func printListNode(l *ListNode) {
    if l == nil {
        return
    }
    fmt.Printf("l = %v\n", l)
    printListNode(l.Next)
}

func main() {
	// result: [1,2,3,4]
	// head := makeList1()

	// result: [-1,0,3,4,5]
	// head := makeList2()

	// result: [0,0,0,0,0,0,0,0,0,0,1,1,1,1,1]
	// head := makeList3()

	// result: [0]
	// head := makeList4()

	// result: [2,3,4]
	head := makeList5()

	// result: 
	// head := makeList()

	result := insertionSortList(head)
	printListNode(result)
}

