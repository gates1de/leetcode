package main

import (
	"fmt"
)

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(maxDepth(root.Left) - maxDepth(root.Right)) <= 1 &&
				isBalanced(root.Left) &&
				isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
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

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	root.Left.Right.Left = &TreeNode{Val: 4}
	root.Left.Right.Right = &TreeNode{Val: 4}
	return root
}

func makeTree3() *TreeNode {
	var root *TreeNode
	return root
}

func makeTree() *TreeNode {
	root := &TreeNode{Val: 0}
	return root
}

func printNode(root *TreeNode) {
	if root == nil {
			return
	}
	fmt.Printf("node %p = %v\n", root, root)
	printNode(root.Left)
	printNode(root.Right)
}

func main() {
	// result: true
	// root := makeTree1()

	// result: false
	// root := makeTree2()

	// result: true
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := isBalanced(root)
	fmt.Printf("result = %v\n", result)
}

