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
	return helper(nums)
}

func helper(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return &TreeNode{
        Val: nums[len(nums) / 2],
        Left: helper(nums[:len(nums) / 2]),
        Right: helper(nums[len(nums) / 2 + 1:]),
    }
}

func printTreeNode(t *TreeNode) {
    if t == nil {
        return
    }
    fmt.Printf("t = %v\n", t)
    printTreeNode(t.Left)
    printTreeNode(t.Right)
}

func main() {
	// result: [0,-3,9,-10,null,5]
	// nums := []int{-10, -3, 0, 5, 9}

	// result: [1,3]
	nums := []int{1, 3}

	// result: 
	// nums := []int{}

	result := sortedArrayToBST(nums)
	printTreeNode(result)
}

