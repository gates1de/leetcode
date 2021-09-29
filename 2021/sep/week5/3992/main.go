package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	result := make([]*ListNode, k)
	nodeCount := int(0)
	listNodeCounter(head, &nodeCount)
	remain := nodeCount
	resultIndex := int(0)
	for remain > 0 {
		// fmt.Printf("head = %v, remain = %v\n", head, remain)
		count := nodeCount / k
		if resultIndex < nodeCount % k {
			count++
		}
		result[resultIndex] = head
		var h *ListNode
		for i := 0; i < count; i++ {
			new := &ListNode{Val: head.Val}
			if i == 0 {
				result[resultIndex] = new
				h = result[resultIndex]
			} else {
				result[resultIndex].Next = new
				result[resultIndex] = result[resultIndex].Next
			}
			head = head.Next
		}
		result[resultIndex] = h
		remain -= count
		resultIndex++
	}
	return result
}

func listNodeCounter(head *ListNode, count *int) {
	if head == nil {
		return
	}
	*count++
	listNodeCounter(head.Next, count)
}

func makeListNode(s []int, current *ListNode) *ListNode {
    if len(s) == 0 {
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
	// result: [[1],[2],[3],[],[]]
	// nums := []int{1, 2, 3}
	// k := int(5)

	// result: [[1,2,3,4],[5,6,7],[8,9,10]]
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// k := int(3)

	// result: [[]]
	nums := []int{}
	k := int(1)

	// result: 
	// nums := []int{}
	// k := int(0)

	var head *ListNode
	head = makeListNode(nums, head)
	result := splitListToParts(head, k)
	for _, node := range result {
		fmt.Printf("-------------------------\n")
		printListNode(node)
	}
	// fmt.Printf("result = %v\n", result)
}

