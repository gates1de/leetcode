package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	result := int(0)
	helper(root, low, high, &result)
	return result
}

func helper(root *TreeNode, low int, high int, result *int) {
	if root == nil {
		return
	}

	if low <= root.Val && root.Val <= high {
		*result += root.Val
	}

	helper(root.Left, low, high, result)
	helper(root.Right, low, high, result)
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 10}
    root.Left = &TreeNode{Val: 5}
    root.Right = &TreeNode{Val: 15}
    root.Left.Left = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 7}
    root.Right.Right = &TreeNode{Val: 18}
    return root
}

func makeTree2() *TreeNode {
    root := &TreeNode{Val: 10}
    root.Left = &TreeNode{Val: 5}
    root.Right = &TreeNode{Val: 15}
    root.Left.Left = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 7}
    root.Right.Left = &TreeNode{Val: 13}
    root.Right.Right = &TreeNode{Val: 18}
    root.Left.Left.Left = &TreeNode{Val: 1}
    root.Left.Right.Left = &TreeNode{Val: 6}
    return root
}

func makeTree3() *TreeNode {
    root := &TreeNode{Val: 6}
    return root
}

func main() {
	// result: 32
	// root := makeTree1()
	// low := int(7)
	// high := int(15)

	// result: 23
	// root := makeTree2()
	// low := int(6)
	// high := int(10)

	// result: 0
	root := makeTree3()
	low := int(2)
	high := int(4)

	// result: 
	// root := makeTree()
	// low := int(0)
	// high := int(0)

	result := rangeSumBST(root, low, high)
	fmt.Printf("result = %v\n", result)
}

