package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
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
	head := &ListNode{Val: 1}
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
	// result: [2,1,4,3]
	// head := makeList1()

	// result: [1]
	head := makeList2()

	// result: 
	// head := makeList()

	result := swapPairs(head)
	printList(result)
}

