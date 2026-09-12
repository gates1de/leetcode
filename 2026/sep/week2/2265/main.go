package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	_, _, result := dfs(root)
	return result
}

func dfs(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0
	}

	leftSum, leftCount, leftResult := dfs(root.Left)
	rightSum, rightCount, rightResult := dfs(root.Right)

	sum := root.Val + leftSum + rightSum
	count := 1 + leftCount + rightCount
	result := leftResult + rightResult

	if root.Val == sum/count {
		result++
	}

	return sum, count, result
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 8}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 6}
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
	// result: 5
	// root := makeTree1()

	// result: 1
	root := makeTree2()

	// result:
	// root := makeTree()

	result := averageOfSubtree(root)
	fmt.Printf("result = %v\n", result)
}
