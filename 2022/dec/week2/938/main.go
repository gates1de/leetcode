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
	if root == nil {
		return 0
	}

	val := int(0)
	if low <= root.Val && root.Val <= high {
		val = root.Val
	}
	return val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 18}
	return root
}

func makeTree2() *TreeNode {
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
	// result: 32
	// root := makeTree1()
	// low := int(7)
	// high := int(15)

	// result: 23
	root := makeTree2()
	low := int(6)
	high := int(10)

	// result: 
	// root := makeTree()
	// low := int(0)
	// high := int(0)

	result := rangeSumBST(root, low, high)
	fmt.Printf("result = %v\n", result)
}

