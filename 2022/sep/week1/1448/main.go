package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	result := int(0)
	helper(-10001, root, &result)
	return result
}

func helper(max int, root *TreeNode, result *int) {
	if root == nil {
		return
	}

	if max <= root.Val {
		max = root.Val
		*result++
	}

	helper(max, root.Left, result)
	helper(max, root.Right, result)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 5}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	return root
}

func makeTree4() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 13}
	root.Right.Left = &TreeNode{Val: 14}
	root.Right.Right = &TreeNode{Val: 5}
	root.Right.Left.Left = &TreeNode{Val: 9}
	root.Right.Left.Right = &TreeNode{Val: 10}
	root.Right.Right.Left = &TreeNode{Val: 8}
	root.Right.Right.Left = &TreeNode{Val: 15}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 4
	// root := makeTree1()

	// result: 3
	// root := makeTree2()

	// result: 1
	// root := makeTree3()

	// result: 4
	root := makeTree4()

	// result: 
	// root := makeTree()

	result := goodNodes(root)
	fmt.Printf("result = %v\n", result)
}

