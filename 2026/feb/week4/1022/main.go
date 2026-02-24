package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	result := int(0)
	helper("", root, &result)
	return result
}

func helper(binary string, root *TreeNode, result *int) {
	if root == nil {
		return
	}

	binary += fmt.Sprintf("%v", root.Val)
	if root.Left == nil && root.Right == nil {
		val, err := strconv.ParseInt(binary, 2, 32)
		if err != nil {
			return
		}
		*result += int(val)
		return
	}

	helper(binary, root.Left, result)
	helper(binary, root.Right, result)
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 0}
    root.Right = &TreeNode{Val: 1}
    root.Left.Left = &TreeNode{Val: 0}
    root.Left.Right = &TreeNode{Val: 1}
    root.Right.Left = &TreeNode{Val: 0}
    root.Right.Right = &TreeNode{Val: 1}
    return root
}

func makeTree2() *TreeNode {
    root := &TreeNode{Val: 0}
    return root
}

func makeTree3() *TreeNode {
    root := &TreeNode{Val: 0}
    root.Left = &TreeNode{Val: 1}
    root.Right = &TreeNode{Val: 1}
    root.Left.Left = &TreeNode{Val: 0}
    root.Left.Right = &TreeNode{Val: 1}
    root.Right.Left = &TreeNode{Val: 0}
    root.Right.Right = &TreeNode{Val: 1}
    root.Left.Left.Left = &TreeNode{Val: 0}
    root.Right.Left.Left = &TreeNode{Val: 1}
    root.Right.Left.Right = &TreeNode{Val: 1}
    root.Right.Left.Left.Left = &TreeNode{Val: 0}
    return root
}

func main() {
	// result: 22
	// root := makeTree1()

	// result: 0
	// root := makeTree2()

	// result: 25
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := sumRootToLeaf(root)
	fmt.Printf("result = %v\n", result)
}

