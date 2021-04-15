package main
import (
	"fmt"
)

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
        return nil
    }

	list := []*ListNode{}
	node := head
	for node != nil {
		if node.Val < x {
			list = append(list, node)
		}
		node = node.Next
	}

	node = head
	for node != nil {
		if x <= node.Val {
			list = append(list, node)
		}
		node = node.Next
	}

	list[len(list) - 1].Next = nil

	var result *ListNode
	var top *ListNode
	for _, n := range list {
		if result == nil {
			result = n
			top = n
		} else {
			result.Next = n
			result = result.Next
		}
	}

	result = top
	// printListNode(result)

	return result
}

type ListNode struct {
    Val int
    Next *ListNode
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
	// result: [1, 2, 2, 4, 3, 5]
	// nums := []int{1, 4, 3, 2, 5, 2}
	// x := int(3)

	// result: [1, 2]
	nums := []int{2, 1}
	x := int(2)

	// result: 
	// nums := []int{}
	// x := int()

	var head *ListNode
	head = makeListNode(nums, head)
	result := partition(head, x)
	fmt.Printf("result = %v\n", result)
}

