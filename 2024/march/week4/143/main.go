package main
import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode)  {
    h := head
    nodes := []*ListNode{}
    for h != nil {
        nodes = append(nodes, h)
        h = h.Next
    }

    n := len(nodes)
    result := head
    i := int(0)
    start := int(1)
    end := n - 1
    for i < n {
        var node *ListNode
        if i % 2 != 0 {
            node = nodes[start]
            start++
        } else {
            node = nodes[end]
            end--
        }

        result.Next = node
        result = result.Next
        i++
    }
    result.Next = nil
    head = result
}

func makeList1() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	return head
}

func makeList() *ListNode {
	var head *ListNode
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
	// result: [1,4,2,3]
	// head := makeList1()

	// result: [1,5,2,4,3]
	head := makeList2()

	// result: []
	// head := makeList()

	reorderList(head)
	printList(head)
}

