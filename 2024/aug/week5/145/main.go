package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	helper(root, &result)
	return result
}

func helper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, result)
	helper(root.Right, result)
	*result = append(*result, root.Val)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
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
	// result: [3,2,1]
	root := makeTree1()

	// result: []
	// root := makeTree2()

	// result: []
	// root := makeTree()

	result := postorderTraversal(root)
	fmt.Printf("result = %v\n", result)
}

