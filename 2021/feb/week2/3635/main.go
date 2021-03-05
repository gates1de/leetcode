package main
import (
	"fmt"
)

type Node struct {
    Val int
    Next *Node
    Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	result := &Node{Val: head.Val}
	allNodes := []*Node{}

	firstHead := head
	copyHead := head
	for copyHead != nil {
		allNodes = append(allNodes, &Node{Val: copyHead.Val})
		copyHead = copyHead.Next
	}
	// fmt.Printf("copy = %v, head = %v\n", copyHead, head)

	i := int(0)
	resultHead := allNodes[0]
	for head != nil {
		result = allNodes[i]
		if i + 1 < len(allNodes) {
			result.Next = allNodes[i + 1]
		}
		if head.Random != nil {
			j := int(0)
			copyHead = firstHead
			for copyHead != nil {
				if copyHead == head.Random {
					result.Random = allNodes[j]
					break
				}
				copyHead = copyHead.Next
				j++
			}
		}
		// fmt.Printf("head = %v, result = %v\n", head, result)
		result = result.Next
		head = head.Next
		i++
	}
	result = resultHead

	return result
}

func printNode(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("head = %v\n", head)
	printNode(head.Next)
}

func main() {
	head := &Node{Val: 7}
	node1 := &Node{Val: 13}
	node2 := &Node{Val: 11}
	node3 := &Node{Val: 10}
	node4 := &Node{Val: 1}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	head.Random = nil
	node1.Random = head
	node2.Random = node4
	node3.Random = node2
	node4.Random = head
	result := copyRandomList(head)
	printNode(result)
	// fmt.Printf("result = %v\n", result)
}

