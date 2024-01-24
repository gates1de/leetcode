package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	result := int(0)
	preorder(root, 0, &result)
	return result
}

func preorder(node *TreeNode, path int, result *int) {
	if node == nil {
		return
	}

	path = path ^ (1 << node.Val)
	if node.Left == nil && node.Right == nil {
		if (path & (path - 1)) == 0 {
			*result++
		}
	}

	preorder(node.Left, path, result)
	preorder(node.Right, path, result)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 1}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Left.Right.Right = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 9}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 2
	// root := makeTree1()

	// result: 1
	// root := makeTree2()

	// result: 1
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := pseudoPalindromicPaths(root)
	fmt.Printf("result = %v\n", result)
}

