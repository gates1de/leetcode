package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var result *TreeNode
	result = insert(result, nums)
	return result
}

func insert(root *TreeNode, nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return root
	} else if n == 1 {
		return &TreeNode{Val: nums[0]}
	}

	if root == nil {
		root = &TreeNode{Val: nums[n / 2]}
	}

	root.Left = insert(root.Left, nums[:n / 2])
	root.Right = insert(root.Right, nums[n / 2 + 1:])
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("root (%p) = %v\n", root, root)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// result: [0,-3,9,-10,null,5]
	// nums := []int{-10,-3,0,5,9}

	// result: [3,1]
	// nums := []int{1,3}

	// result: [[5,-3,10,-10,0,9]]
	nums := []int{-10,-3,0,5,9,10}

	// result: 
	// nums := []int{}

	result := sortedArrayToBST(nums)
	printTree(result)
}

