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

func maxPathSum(root *TreeNode) int {
	result := math.MinInt32
	helper(root, &result)
	return result
}

func helper(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := max(helper(root.Left, result), 0)
	right := max(helper(root.Right, result), 0)
	*result = max(*result, left + right + root.Val)
	return max(left + root.Val, right + root.Val)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: -10}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 6
	// root := makeTree1()

	// result: 42
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := maxPathSum(root)
	fmt.Printf("result = %v\n", result)
}

