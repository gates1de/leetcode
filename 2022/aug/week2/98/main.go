package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nums := []int{}
	preorder(root, &nums)

	pre := nums[0]
	for _, num := range nums[1:] {
		if num <= pre {
			return false
		}
		pre = num
	}

	return true
}

func preorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	preorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	preorder(root.Right, nums)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 12}
	return root
}

func makeTree4() *TreeNode {
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 15}
	root.Left.Left = &TreeNode{Val: 8}
	root.Left.Right = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 12}
	root.Right.Right = &TreeNode{Val: 18}
	return root
}

func makeTree5() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: true
	// root := makeTree1()

	// result: false
	// root := makeTree2()

	// result: true
	// root := makeTree3()

	// result: false
	// root := makeTree4()

	// result: false
	root := makeTree5()

	// result: 
	// root := makeTree()

	result := isValidBST(root)
	fmt.Printf("result = %v\n", result)
}

