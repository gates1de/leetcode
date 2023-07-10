package main
import (
	"fmt"
	"math"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depths := make(map[*TreeNode]int)
	depths[root] = 1
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := math.MaxInt32

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		depth := depths[node]

		if node.Left != nil {
			depths[node.Left] = depth + 1
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			depths[node.Right] = depth + 1
			queue = append(queue, node.Right)
		}

		if node.Left == nil && node.Right == nil {
			result = min(result, depth)
		}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 5}
	root.Right.Right.Right.Right = &TreeNode{Val: 6}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 2
	// root := makeTree1()

	// result: 5
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := minDepth(root)
	fmt.Printf("result = %v\n", result)
}

