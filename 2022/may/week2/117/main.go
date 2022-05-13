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
	m := map[int][]*Node{}
	m[0] = []*Node{root}
	helper(1, root, m)

	level := int(0)
	for true {
		nodes := m[level]
		if len(nodes) == 0 {
			break
		}

		top := nodes[0]
		for i := 1; i < len(nodes); i++ {
			insert(top, nodes[i])
		}

		level++
	}

	return m[0][0]
}

func helper(level int, root *Node, m map[int][]*Node) {
	if root == nil {
		return
	}

	if m[level] == nil {
		m[level] = []*Node{}
	}

	nodes := []*Node{}
	if root.Left != nil {
		nodes = append(nodes, root.Left)
		m[level] = append(m[level], root.Left)
	}
	if root.Right != nil {
		nodes = append(nodes, root.Right)
		m[level] = append(m[level], root.Right)
	}

	for _, node := range nodes {
		helper(level + 1, node, m)
	}
}

func insert(root *Node, target *Node) *Node {
	if root == nil {
		return target
	}

	last := root
	for last.Next != nil {
		last = last.Next
	}
	last.Next = target

	return root
}

func makeTree1() *Node {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Right = &Node{Val: 7}
	return root
}

func makeTree2() *Node {
	var root *Node
	return root
}

func printTree(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// result: [1,#,2,3,#,4,5,7,#]
	// root := makeTree1()

	// result: []
	// root := makeTree2()

	// result: 
	// root := makeTree()

	result := connect(root)
	printTree(result)
}

