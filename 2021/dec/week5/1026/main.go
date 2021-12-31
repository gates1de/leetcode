package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := int(0)
	helper(root, root.Val, root.Val, &result)
	return result
}

func helper(root *TreeNode, anc int, des int, result *int) {
	// fmt.Printf("anc = %v, des = %v, current = %v\n", anc, des, *result)
	if root == nil {
		return
	}

	maxDiff := abs(anc, des)
	aDiff := abs(anc, root.Val)
	dDiff := abs(des, root.Val)

	if maxDiff > aDiff && maxDiff > dDiff {
		helper(root.Left, anc, des, result)
		helper(root.Right, anc, des, result)
		return
	}

	if aDiff > dDiff {
		maxDiff = abs(anc, root.Val)
		if *result < maxDiff {
			*result = maxDiff
		}
		helper(root.Left, anc, root.Val, result)
		helper(root.Right, anc, root.Val, result)
	} else {
		maxDiff = abs(des, root.Val)
		if *result < maxDiff {
			*result = maxDiff
		}
		helper(root.Left, des, root.Val, result)
		helper(root.Right, des, root.Val, result)
	}
}

func abs(a int, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 8}
    root.Left = &TreeNode{Val: 3}
    root.Right = &TreeNode{Val: 10}
    root.Left.Left = &TreeNode{Val: 1}
    root.Left.Right = &TreeNode{Val: 6}
    root.Right.Right = &TreeNode{Val: 14}
    root.Left.Right.Left = &TreeNode{Val: 4}
    root.Left.Right.Right = &TreeNode{Val: 7}
    root.Right.Right.Left = &TreeNode{Val: 13}
    return root
}

func makeTree2() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Right = &TreeNode{Val: 2}
    root.Right.Right = &TreeNode{Val: 0}
    root.Right.Right.Left = &TreeNode{Val: 3}
    return root
}

func makeTree3() *TreeNode {
    root := &TreeNode{Val: 6}
    return root
}

func main() {
	// result: 7
	// root := makeTree1()

	// result: 
	// root := makeTree2()

	// result: 0
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := maxAncestorDiff(root)
	fmt.Printf("result = %v\n", result)
}

