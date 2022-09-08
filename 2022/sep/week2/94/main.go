package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, result)
	*result = append(*result, root.Val)
	inorder(root.Right, result)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	return root
}

func makeTree2() *TreeNode {
	var root *TreeNode
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: [1,3,2]
	// root := makeTree1()

	// result: []
	// root := makeTree2()

	// result: [1]
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := inorderTraversal(root)
	fmt.Printf("result = %v\n", result)
}

