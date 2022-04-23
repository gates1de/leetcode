package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	return helper(root, &k)
}

func helper(root *TreeNode, k *int) int {
	if root == nil {
		return -1
	}

	result := helper(root.Left, k)
	*k--

	if result != -1 {
		return result
	}

	if *k == 0 {
		return root.Val
	}

	return helper(root.Right, k)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 1}
	return root
}

func main() {
	// result: 1
	// root := makeTree1()
	// k := int(1)

	// result: 3
	root := makeTree2()
	k := int(3)

	// result: 
	// root := makeTree()
	// k := int(1)

	result := kthSmallest(root, k)
	fmt.Printf("result = %v\n", result)
}

