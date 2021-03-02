package main
import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

// l1 に格納していく
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l2 == nil {
		return l1
	}

	if l1 == nil || l1.Val > l2.Val {
		var insertNode *ListNode
		insertNode = &ListNode{}
		insertNode.Val = l2.Val
		insertNode.Next = l1
		return mergeTwoLists(insertNode, l2.Next)
	}

	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}

func generateListNodeFromSlice(s []int, current *ListNode) *ListNode {
	if len(s) == 0 {
		return current
	}

	if current == nil {
		current = &ListNode{Val: s[0], Next: nil}
		return generateListNodeFromSlice(s[1:], current)
	}
	current.Next = generateListNodeFromSlice(s[1:], &ListNode{Val: s[0], Next: nil})
	return current
}

func printListNode(l *ListNode) {
	if l == nil {
		return
	}
	fmt.Printf("l.Val = %v\n", l.Val)
	printListNode(l.Next)
}

func main() {
	var l1 *ListNode
	var l2 *ListNode
	s1 := []int{1, 2, 4, 8, 9}
	s2 := []int{1, 3, 4, 7, 8}
	l1 = generateListNodeFromSlice(s1, l1)
	l2 = generateListNodeFromSlice(s2, l2)
	result := mergeTwoLists(l1, l2)
	printListNode(result)
}

