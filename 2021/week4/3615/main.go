package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	fmt.Printf("lists = %v\n", lists)
	var current *ListNode
	if len(lists) == 0 {
		return current
	}
	return min(lists, current)
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

func min(lists []*ListNode, current *ListNode) *ListNode {
	// fmt.Printf("current = %v\n", current)
	var minNode *ListNode
	minNodeIndex := int(0)
	for i, node := range lists {
		// fmt.Printf("node = %v\n", node)
		if node == nil || (minNode != nil && minNode.Val < node.Val) {
			continue
		}
		if current == nil {
			current = &ListNode{Val: node.Val}
			// fmt.Printf("first current = %v\n", current)
			return min(lists, current)
		}
		minNode = node
		minNodeIndex = i
	}
	if minNode == nil {
		return nil
	}
	current = &ListNode{Val: minNode.Val}
	lists[minNodeIndex] = minNode.Next
	current.Next = min(lists, current)
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
	arrLists := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	// arrLists := []int{}
	// arrLists := [][]int{{}}
	lists := make([]*ListNode, len(arrLists))
	for i, s := range arrLists {
		if len(s) == 0 {
			continue
		}
		var node *ListNode
		lists[i] = makeListNode(s, node)
	}
	result := mergeKLists(lists)
	printListNode(result)
	// fmt.Printf("result = %v\n", result)
}

