package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Val == 0 {
		return false
	} else if root.Val == 1 {
		return true
	}

	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)

	if root.Val == 2 {
		return left || right
	} else if root.Val == 3 {
		return left && right
	}

	return false
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 1}
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
	// result: true
	// root := makeTree1()

	// result: false
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := evaluateTree(root)
	fmt.Printf("result = %v\n", result)
}

