package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// printListNode(head)
	return floydCycleFinding(head)
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

func floydCycleFinding(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for slow != nil {
		fmt.Printf("slow = %v, fast = %v\n", slow, fast)
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func printListNode(l *ListNode) {
    if l == nil {
        return
    }
    fmt.Printf("l = %v\n", l)
    printListNode(l.Next)
}

func main() {
	// list := []int{3, 2, 0, -4}
	// list := []int{1, 2}
	list := []int{1}

	var head *ListNode
	head = makeListNode(list, head)

	// head.Next.Next.Next.Next = head.Next
	// head.Next = head

	result := hasCycle(head)
	fmt.Printf("result = %v\n", result)
}

