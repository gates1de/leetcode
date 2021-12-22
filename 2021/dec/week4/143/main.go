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
		// fmt.Printf("node = %v\n", node)

		result.Next = node
		result = result.Next
		i++
	}
	result.Next = nil
	head = result
}

func makeList1() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	return root
}

func makeList2() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	return root
}

func makeList3() *ListNode {
	root := &ListNode{Val: 0}
	return root
}

func makeList4() *ListNode {
	root := &ListNode{Val: 3}
	root.Next = &ListNode{Val: 2}
	return root
}

func printListNode(l *ListNode) {
    if l == nil {
        return
    }
    fmt.Printf("l = %v\n", l)
    printListNode(l.Next)
}

func main() {
	// result: [1,4,2,3]
	// head := makeList1()

	// result: [1,5,2,4,3]
	// head := makeList2()

	// result: [0]
	// head := makeList3()

	// result: [3,2]
	head := makeList4()

	// result: 
	// head := makeList()

	reorderList(head)
	printListNode(head)
}

