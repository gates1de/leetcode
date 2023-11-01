package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func findMode(root *TreeNode) []int {
	m := make(map[int]int)
	maxCount := int(0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		m[node.Val]++
		maxCount = max(maxCount, m[node.Val])


		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	result := make([]int, 0, len(m))
	for val, count := range m {
		if maxCount == count {
			result = append(result, val)
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 2}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 0}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: [2]
	// root := makeTree1()

	// result: [0]
	root := makeTree2()

	// result: []
	// root := makeTree()

	result := findMode(root)
	fmt.Printf("result = %v\n", result)
}

