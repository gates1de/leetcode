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

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	levels := map[*TreeNode]int{}
	sums := map[int]int{}
	queue := make([]*TreeNode, 0, 10000)
	queue = append(queue, root)
	levels[root] = 1
	sums[1] = root.Val

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		level := levels[node]
		sums[level] += node.Val

		if node.Left != nil {
			queue = append(queue, node.Left)
			levels[node.Left] = level + 1
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			levels[node.Right] = level + 1
		}
	}

	result := int(1)
	maxSum := math.MinInt32
	for level, sum := range sums {
		if maxSum < sum {
			result = level
			maxSum = sum
		}
	}
	return result
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: -8}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 989}
	root.Right = &TreeNode{Val: 10250}
	root.Right.Left = &TreeNode{Val: 98693}
	root.Right.Right = &TreeNode{Val: -89388}
	root.Right.Right.Right = &TreeNode{Val: -32127}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 2
	// root := makeTree1()

	// result: 2
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := maxLevelSum(root)
	fmt.Printf("result = %v\n", result)
}

