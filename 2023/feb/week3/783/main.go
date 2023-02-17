package main
import (
	"fmt"
	"sort"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nums := make([]int, 0, 100)
	nums = append(nums, root.Val)
	preorder(root.Left, &nums)
	preorder(root.Right, &nums)

	if len(nums) <= 1 {
		return 0
	}
	sort.Ints(nums)

	result := int(100000)
	for i := 1; i < len(nums); i++ {
		pre := nums[i - 1]
		num := nums[i]
		diff := abs(pre, num)
		if diff < result {
			result = diff
		}
	}
	return result
}

func preorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	*nums = append(*nums, root.Val)
	preorder(root.Left, nums)
	preorder(root.Right, nums)
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 48}
	root.Right.Left = &TreeNode{Val: 12}
	root.Right.Right = &TreeNode{Val: 49}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 0}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 1
	// root := makeTree1()

	// result: 1
	// root := makeTree2()

	// result: 4
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := minDiffInBST(root)
	fmt.Printf("result = %v\n", result)
}

