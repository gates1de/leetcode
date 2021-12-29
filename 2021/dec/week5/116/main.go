package main
import (
	"fmt"
)

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

func connect(root *Node) *Node {
	bfs(root)
	return root
}

func bfs(root *Node) {
	if root == nil {
		return
	}

	m := map[int][]*Node{}
	level := int(0)
	m[level] = []*Node{root}

	for len(m) > 0 {
		nodes := m[level]
		delete(m, level)
		level++

		var head *Node

		for _, node := range nodes {
			if head == nil {
				head = node
			} else {
				last := head
				for last.Next != nil {
					last = last.Next
				}
				last.Next = node
			}

			if node.Left != nil {
				if m[level] == nil {
					m[level] = []*Node{}
				}
				m[level] = append(m[level], node.Left)
			}

			if node.Right != nil {
				if m[level] == nil {
					m[level] = []*Node{}
				}
				m[level] = append(m[level], node.Right)
			}
		}
	}
}

func printNode(root *Node) {
	if root == nil {
		return
	}

	nodes := []*Node{root}
	for len(nodes) > 0 {
		node := nodes[0]
		if len(nodes) <= 1 {
			nodes = []*Node{}
		} else {
			nodes = nodes[1:]
		}

		fmt.Printf("node = %v\n", node)

		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}

		for node != nil {
			fmt.Printf("node.Next = %v\n", node.Next)
			node = node.Next
		}
	}
}

func makeNode1() *Node {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Left = &Node{Val: 6}
	root.Right.Right = &Node{Val: 7}
	return root
}

func makeNode2() *Node {
	var root *Node
	return root
}

func makeNode3() *Node {
	root := &Node{Val: 0}
	return root
}

func makeNode4() *Node {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	return root
}

func main() {
	// result: [1,#,2,3,#,4,5,6,7,#]
	// root := makeNode1()

	// result: []
	// root := makeNode2()

	// result: [0]
	// root := makeNode3()

	// result: 
	root := makeNode4()

	// result: 
	// root := makeNode()

	result := connect(root)
	// fmt.Printf("result = %v\n", result)
	printNode(result)
}

