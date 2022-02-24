package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    left, right := split(head)
    left, right = sortList(left), sortList(right)
    return merge(left, right)
}

func split(head *ListNode) (*ListNode, *ListNode) {
    var prev *ListNode
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    prev.Next = nil
    return head, slow
}

func merge(left *ListNode, right *ListNode) *ListNode {
    var head *ListNode
    var curr *ListNode

    if left.Val < right.Val {
        head = left
        left = left.Next
    } else {
        head = right
        right = right.Next
    }
    curr = head

    for left != nil && right != nil {
        if left.Val < right.Val {
            curr.Next = left
            left = left.Next
        } else {
            curr.Next = right
            right = right.Next
        }
        curr = curr.Next
    }

    if left != nil {
        curr.Next = left
    } else if right != nil {
        curr.Next = right
    }

    return head
}

// Slow
func mySolution(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var result *ListNode
	for head != nil {
		result = insert(head.Val, result)
		head = head.Next
	}
	return result
}

func insert(val int, head *ListNode) *ListNode {
	if head == nil {
		return &ListNode{Val: val}
	}
	if val <= head.Val {
		tmp := head
		head = &ListNode{Val: val, Next: tmp}
		return head
	}

	h := head
	for head != nil {
		if val >= head.Val {
			if head.Next != nil && val >= head.Next.Val {
				head = head.Next
				continue
			}
			tmp := head.Next
			head.Next = &ListNode{Val: val, Next: tmp}
			break
		}

		head = head.Next
	}

	return h
}

func makeList1() *ListNode {
	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 3}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: -1}
	head.Next = &ListNode{Val: 5}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 0}
	return head
}

func makeList3() *ListNode {
	var head *ListNode
	return head
}

func printList(head *ListNode) {
	fmt.Println("printList")
	for head != nil {
		fmt.Printf("head = %v\n", head)
		head = head.Next
	}
}

func main() {
	// result: [1,2,3,4]
	// head := makeList1()

	// result: [-1,0,3,4,5]
	// head := makeList2()

	// result: []
	head := makeList3()

	// result: 
	// head := makeList()

	result := sortList(head)
	printList(result)
}

