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
	copies := map[*Node]*Node{}
	var pre *Node
	cursor := head
	for cursor != nil {
		if copies[cursor] == nil {
			copies[cursor] = &Node{Val: cursor.Val}
		}

		if cursor.Random != nil {
			if copies[cursor.Random] == nil {
				copies[cursor.Random] = &Node{Val: cursor.Random.Val}
			}
			copies[cursor].Random = copies[cursor.Random]
		}

		if pre != nil {
			pre.Next = copies[cursor]
		}
		pre = copies[cursor]
		cursor = cursor.Next
	}

	// fmt.Println(copies)
	return copies[head]
}

func makeNode1() *Node {
	node1 := &Node{Val: 7}
	node2 := &Node{Val: 13}
	node3 := &Node{Val: 11}
	node4 := &Node{Val: 10}
	node5 := &Node{Val: 1}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node2.Random = node1
	node3.Random = node5
	node4.Random = node3
	node5.Random = node1
	return node1
}

func makeNode2() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}

	node1.Next = node2

	node1.Random = node1
	node2.Random = node1

	return node1
}

func makeNode3() *Node {
	node1 := &Node{Val: 3}
	node2 := &Node{Val: 3}
	node3 := &Node{Val: 3}

	node1.Next = node2
	node2.Next = node3

	node2.Random = node1

	return node1
}

func makeNode4() *Node {
	node1 := &Node{Val: 1}
	return node1
}

func makeNode5() *Node {
	return nil
}

func printNode(node *Node) {
	if node == nil {
		fmt.Println("end")
		return
	}
	fmt.Println(node)
	printNode(node.Next)
}

func main() {
	// result: [[7,null],[13,0],[11,4],[10,2],[1,0]]
	// head := makeNode1()

	// result: [[1,1],[2,1]]
	// head := makeNode2()

	// result: [[3,null],[3,0],[3,null]]
	// head := makeNode3()

	// result: [[1,null]]
	// head := makeNode4()

	// result: []
	head := makeNode5()

	// result: 
	// head := makeNode()

	result := copyRandomList(head)
	printNode(result)
}

