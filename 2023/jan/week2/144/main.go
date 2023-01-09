package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 101)
	preorder(root, &result)
	return result
}

func preorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	preorder(root.Left, result)
	preorder(root.Right, result)
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

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 7}
	root.Left.Right.Left = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 18}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: [1,2,3]
	// root := makeTree1()

	// result: [1]
	// root := makeTree2()

	// result: [10,5,3,1,7,6,15,13,18]
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := preorderTraversal(root)
	fmt.Printf("result = %v\n", result)
}

