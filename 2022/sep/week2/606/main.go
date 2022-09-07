package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func tree2str(root *TreeNode) string {
	result := ""
	if root == nil {
		return result
	}

	helper(root, &result)
	return result[1:len(result) - 1]
}

func helper(root *TreeNode, result *string) {
	if root == nil {
		return
	}

	*result += fmt.Sprintf("(%v", root.Val)
	if root.Left != nil {
		helper(root.Left, result)
	} else {
		if root.Right != nil {
			*result += "()"
		}
	}
	helper(root.Right, result)
	*result += ")"
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 8}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: "1(2(4))(3)"
	// root := makeTree1()

	// result: "1(2()(4))(3)"
	// root := makeTree2()

	// result: 
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := tree2str(root)
	fmt.Printf("result = %v\n", result)
}

