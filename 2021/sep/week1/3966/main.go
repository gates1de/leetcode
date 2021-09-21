package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	var next *ListNode
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
		// fmt.Printf("current: %p = %v\n", current, current)
	}
	head = prev
	return head
}

func makeListNode(s []int, current *ListNode) *ListNode {
    if len(s) == 0 {
        if current == nil {
            current = &ListNode{Val: 0}
        }
        return current
    }

    if current == nil {
        current = &ListNode{Val: s[0]}
        return makeListNode(s[1:], current)
    }
    current.Next = makeListNode(s[1:], &ListNode{Val: s[0]})
    return current
}

func printListNode(l *ListNode) {
    if l == nil {
        return
    }
    fmt.Printf("l = %v\n", l)
    printListNode(l.Next)
}

func main() {
	// result: [5,4,3,2,1]
	nums := []int{1, 2, 3, 4, 5}

	// result: [2,1]
	// nums := []int{1, 2}

	// result: 
	// nums := []int{}

	var head *ListNode
	head = makeListNode(nums, head)
	result := reverseList(head)
	printListNode(result)
}

