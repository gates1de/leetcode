package main
import (
	"fmt"
)


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		head = deleteRecursive(head, head.Val)
		head = deleteDuplicates(head)
	} else {
		head.Next = deleteDuplicates(head.Next)
	}
	return head
}

func deleteRecursive(head *ListNode, target int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == target {
		return deleteRecursive(head.Next, target)
	}

	return head
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
  list := []int{1, 2, 3, 3, 4, 4, 5}
  // list := []int{1, 1, 1, 2, 3}
  // list := []int{1, 1}
  var head *ListNode
  head = generateListNodeFromSlice(list, head)
  result := deleteDuplicates(head)
  printListNode(result)
}

