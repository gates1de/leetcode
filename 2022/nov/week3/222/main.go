package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := int(0)
	queue := make([]*TreeNode, 0, 50000)
	queue = append(queue, root)

	for len(queue) > 0 {
		result++
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 6
	// root := makeTree1()

	// result: 1
	// root := makeTree2()

	// result: 0
	root := makeTree()

	// result: 
	// root := makeTree()

	result := countNodes(root)
	fmt.Printf("result = %v\n", result)
}

