package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}

	targetSum -= root.Val
	if hasPathSum(root.Left, targetSum) {
		return true
	}
	return hasPathSum(root.Right, targetSum)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree3() *TreeNode {
	var root *TreeNode
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: true
	// root := makeTree1()
	// targetSum := int(22)

	// result: false
	// root := makeTree2()
	// targetSum := int(5)

	// result: false
	root := makeTree3()
	targetSum := int(0)

	// result: 
	// root := makeTree()
	// targetSum := int(0)

	result := hasPathSum(root, targetSum)
	fmt.Printf("result = %v\n", result)
}

