package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	result := int(-1)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	levels := make(map[*TreeNode]int)
	levels[root] = 1
	currentLevel := int(0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			continue
		}

		level := levels[node]
		if level > currentLevel {
			result = node.Val
			currentLevel = level
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			levels[node.Left] = level + 1
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			levels[node.Right] = level + 1
		}
	}

	return result
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}
	root.Right.Left.Left = &TreeNode{Val: 7}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 1
	// root := makeTree1()

	// result: 7
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := findBottomLeftValue(root)
	fmt.Printf("result = %v\n", result)
}

