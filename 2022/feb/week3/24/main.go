package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

// Do not use array
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    d := &ListNode{Next: head}
    n1, n2, prev := head, head.Next, d
    for n1 != nil && n2 != nil {
        tmp := n2.Next
        n2.Next, n1.Next, prev.Next = n1, tmp, n2

        if tmp == nil {
            n1, n2 = nil, nil
            continue
        }
        prev, n1, n2 = n1, tmp, tmp.Next
    }
    return d.Next
}

// Use array
func mySolution(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nodes := []*ListNode{}
	node := head
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}

	var result *ListNode
	for i := 0; i < len(nodes); i++ {
		if i % 2 != 0 {
			continue
		}

		if i + 1 < len(nodes) {
			nodes[i + 1].Next = nil
			result = insert(result, nodes[i + 1])
		}
		nodes[i].Next = nil
		result = insert(result, nodes[i])
	}

	return result
}

func insert(head *ListNode, target *ListNode) *ListNode {
	last := head
	if head == nil {
		return target
	}

	for last.Next != nil {
		last = last.Next
	}
	last.Next = target
	return head
}

func makeList1() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	return head
}

func makeList2() *ListNode {
	var head *ListNode
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
	return head
}

func makeList4() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	return head
}

func printList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head)
	printList(head.Next)
}

func main() {
	// result: [2,1,4,3]
	// head := makeList1()

	// result: []
	// head := makeList2()

	// result: [1]
	// head := makeList3()

	// result: [2,1,4,3,5]
	head := makeList4()

	// result: 
	// head := makeList()

	result := swapPairs(head)
	printList(result)
}

