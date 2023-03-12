package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode
	for true {
		minIndex := int(0)
		var minNode *ListNode

		for i, node := range lists {
			if node == nil {
				continue
			}
			if minNode == nil || node.Val < minNode.Val {
				minNode = node
				minIndex = i
			}
		}

		if minNode == nil {
			break
		}

		result = insert(result, minNode)
		lists[minIndex] = minNode.Next
	}

	return result
}

func insert(head *ListNode, node *ListNode) *ListNode {
	if node == nil {
		return head
	}

	if head == nil {
		return &ListNode{Val: node.Val}
	}

	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &ListNode{Val: node.Val}
	return head
}

func makeLists1() []*ListNode {
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 4} 
	head1.Next.Next = &ListNode{Val: 5} 

	head2 := &ListNode{Val: 1}
	head2.Next = &ListNode{Val: 3} 
	head2.Next.Next = &ListNode{Val: 4} 

	head3 := &ListNode{Val: 2}
	head3.Next = &ListNode{Val: 6} 
	return []*ListNode{head1, head2, head3}
}

func makeLists2() []*ListNode {
	var head *ListNode
	return []*ListNode{head}
}

func makeLists() []*ListNode {
	return []*ListNode{}
}

func printList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head)
	printList(head.Next)
}

func main() {
	// result: [1,1,2,3,4,4,5,6]
	// lists := makeLists1()

	// result: []
	// lists := makeLists2()

	// result: []
	lists := makeLists()

	result := mergeKLists(lists)
	printList(result)
}

