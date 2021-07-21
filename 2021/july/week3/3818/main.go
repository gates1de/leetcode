package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if k <= 1 {
        return head
    }
    nodesLen := int(0)
    for p := head; p != nil; p = p.Next {
        nodesLen++
    }
    rhead := new(ListNode)
    rhead.Next = head
    pre := rhead
    tail := head
    next := head
    for ; nodesLen >= k; nodesLen -= k {
        for i := 0; i < k; i++ {
            temp := next.Next
            next.Next = pre.Next
            pre.Next = next
            next = temp
            tail.Next = next
        }
        pre = tail
        tail = next
    }

    return rhead.Next
}

// Not bad
func mySolution(head *ListNode, k int) *ListNode {
	arr := make([]int, k)
	var result *ListNode
	i := int(0)
	for head != nil {
		arr[i] = head.Val
		i++

		if i == k {
			reverseArray(arr)
			result = makeListNode(arr, result, result != nil)
			i = 0
			arr = make([]int, k)
		}
		head = head.Next
	}
	result = makeListNode(arr[:i], result, result != nil)
	return result
}

func reverseArray(arr []int) {
	for i := 0; i < len(arr) / 2; i++ {
		arr[i], arr[len(arr) - i - 1] = arr[len(arr) - i - 1], arr[i]
	}
}

func makeListNode(s []int, current *ListNode, hasParent bool) *ListNode {
    if len(s) == 0 {
		if hasParent {
			return current
		}
        if current == nil {
            current = &ListNode{Val: 0}
        }
        return current
    }

    if current == nil {
        current = &ListNode{Val: s[0]}
        return makeListNode(s[1:], current, true)
    }
	if current.Next == nil {
		current.Next = makeListNode(s[1:], &ListNode{Val: s[0]}, current != nil)
	} else {
		current.Next = makeListNode(s, current.Next, current != nil)
	}
    return current
}

// func reverseNext(head *ListNode, prev **ListNode) *ListNode {
// 	tmp := head.Next
// 	head.Next = *prev
// 	*prev = head
// 	head = tmp
//     return head
// }

func printListNode(l *ListNode) {
    if l == nil {
        return
    }
    fmt.Printf("l = %v\n", l)
    printListNode(l.Next)
}

func main() {
	// result: [2,1,4,3,5]
    // k := int(2)
    // nums := []int{1, 2, 3, 4, 5}

	// result: [1,2,3,4,5]
    // k := int(1)
    // nums := []int{1, 2, 3, 4, 5}

	// result: [1]
    k := int(1)
    nums := []int{1}

	// result: []
    // k := int(0)
    // nums := []int{}

	var head *ListNode
	head = makeListNode(nums, head, false)
	head = reverseKGroup(head, k)
	printListNode(head)
}

