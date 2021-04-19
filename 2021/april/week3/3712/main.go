package main
import (
	"fmt"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := []*ListNode{}
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	var result *ListNode
	var top *ListNode
	for i, node := range list {
		// fmt.Printf("result = %v, node = %v\n", result, node)
		if i == len(list) - n {
			continue
		}
		result = node
		if top == nil {
			top = result
		}
		// fmt.Printf("result = %v\n", result)
		if i + 1 == len(list) - n {
			if result.Next != nil {
				result.Next = result.Next.Next
			} else {
				result.Next = nil
			}
		}
		result = result.Next
	}
	result = top
	printListNode(result)
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
	// result: [1, 2, 3, 5]
	// nums := []int{1, 2, 3, 4, 5}
	// n := int(2)

	// result: []
	// nums := []int{1}
	// n := int(1)

	// result: [1]
	// nums := []int{1, 2}
	// n := int(1)

	// result: [2]
	nums := []int{1, 2}
	n := int(2)

	// result: 
	// nums := []int{}
	// n := int(0)

	var head *ListNode
	head = makeListNode(nums, head)
	result := removeNthFromEnd(head, n)
	fmt.Printf("result = %v\n", result)
}

