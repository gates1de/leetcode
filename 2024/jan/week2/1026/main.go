package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	result := int(0)
	if root == nil {
		return result
	}

	minVal := root.Val
	maxVal := root.Val

	calcDiff(root, &result, minVal, maxVal)
	return result
}

func calcDiff(root *TreeNode, diff *int, minVal int, maxVal int) {
	if root == nil {
		return
	}

	*diff = max(*diff, max(abs(minVal - root.Val), abs(maxVal - root.Val)))
	minVal = min(minVal, root.Val)
	maxVal = max(maxVal, root.Val)
	calcDiff(root.Left, diff, minVal, maxVal)
	calcDiff(root.Right, diff, minVal, maxVal)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 8}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 4}
	root.Left.Right.Right = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 14}
	root.Right.Right.Left = &TreeNode{Val: 13}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 0}
	root.Right.Right.Left = &TreeNode{Val: 3}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 7
	// root := makeTree1()

	// result: 3
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := maxAncestorDiff(root)
	fmt.Printf("result = %v\n", result)
}

