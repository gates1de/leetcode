package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		listsLen := len(lists)
		newListsLen := listsLen / 2
		if listsLen % 2 == 1 {
			newListsLen++
		}
		newLists := make([]*ListNode, newListsLen)

		for i := 0; i < listsLen - 1; i += 2 {
			newLists[i / 2] = merge(lists[i], lists[i + 1])
		}

		if listsLen % 2 == 1 {
			newLists[newListsLen - 1] = lists[listsLen - 1]
		}

		// fmt.Println(newLists)

		lists = newLists
	}

	return lists[0]
}

func merge(left *ListNode, right *ListNode) *ListNode {
	var result *ListNode

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Val <= right.Val {
		result = left
		left = left.Next
	} else {
		result = right
		right = right.Next
	}

	node := result
	for left != nil && right != nil {
		if left.Val <= right.Val {
			node.Next = left
			left = left.Next
		} else {
			node.Next = right
			right = right.Next
		}
		node = node.Next
	}

	if left != nil {
		node.Next = left
	}
	if right != nil {
		node.Next = right
	}

	return result
}

func makeLists1() []*ListNode {
	root1 := &ListNode{Val: 1}
	root1.Next = &ListNode{Val: 4}
	root1.Next.Next = &ListNode{Val: 5}

	root2 := &ListNode{Val: 1}
	root2.Next = &ListNode{Val: 3}
	root2.Next.Next = &ListNode{Val: 4}

	root3 := &ListNode{Val: 2}
	root3.Next = &ListNode{Val: 6}

	return []*ListNode{root1, root2, root3}
}

func makeLists2() []*ListNode {
	return []*ListNode{}
}

func makeLists3() []*ListNode {
	var root *ListNode
	return []*ListNode{root}
}

func makeLists4() []*ListNode {
	root1 := &ListNode{Val: 3}

	root2 := &ListNode{Val: 1}

	root3 := &ListNode{Val: 2}

	return []*ListNode{root1, root2, root3}
}

func printListNode(l *ListNode) {
    if l == nil {
        return
    }
    fmt.Printf("l = %v\n", l)
    printListNode(l.Next)
}

func main() {
	// result: [1,1,2,3,4,4,5,6]
	// lists := makeLists1()

	// result: []
	// lists := makeLists2()

	// result: []
	// lists := makeLists3()

	// result: [1,2,3]
	lists := makeLists4()

	// result: 
	// lists := makeLists()

	result := mergeKLists(lists)
	printListNode(result)
}

