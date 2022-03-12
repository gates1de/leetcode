package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list2 == nil {
		return list1
	} else if list1 == nil {
		return list2
	}

	var result *ListNode
	if list1.Val <= list2.Val {
		result = insert(result, list1, true)
		list1 = list1.Next
	} else {
		result = insert(result, list2, true)
		list2 = list2.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			result = insert(result, list1, true)
			list1 = list1.Next
		} else {
			result = insert(result, list2, true)
			list2 = list2.Next
		}
	}

	if list1 != nil {
		result = insert(result, list1, false)
	} else {
		result = insert(result, list2, false)
	}
	return result
}

func insert(head *ListNode, target *ListNode, isNew bool) *ListNode {
	last := head
	if isNew {
		target = &ListNode{Val: target.Val}
	}
	if head == nil {
		return target
	}

	for last.Next != nil {
		last = last.Next
	}
	last.Next = target
	return head
}

func makeLists1() (*ListNode, *ListNode) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}
	node1.Next = node2
	node1.Next.Next = node3

	node4 := &ListNode{Val: 1}
	node5 := &ListNode{Val: 3}
	node6 := &ListNode{Val: 4}
	node4.Next = node5
	node4.Next.Next = node6

	return node1, node4
}

func makeLists2() (*ListNode, *ListNode) {
	return nil, nil
}

func makeLists3() (*ListNode, *ListNode) {
	return nil, &ListNode{Val: 0}
}

func printList(node *ListNode) {
	if node == nil {
		fmt.Println("end")
		return
	}
	fmt.Println(node)
	printList(node.Next)
}

func main() {
	// result: [1,1,2,3,4,4]
	// list1, list2 := makeLists1()

	// result: []
	// list1, list2 := makeLists2()

	// result: [0]
	list1, list2 := makeLists3()

	// result: 
	// list1, list2 := makeLists()

	result := mergeTwoLists(list1, list2)
	printList(result)
}

