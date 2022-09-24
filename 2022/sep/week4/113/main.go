package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 1}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	helper(0, []int{}, root, targetSum, &result)
	return result
}

func helper(currentSum int, paths []int, root *TreeNode, targetSum int, result *[][]int) {
	if root == nil {
		return
	}

	currentSum += root.Val
	newPaths := make([]int, len(paths))
	copy(newPaths, paths)
	newPaths = append(newPaths, root.Val)
	if currentSum == targetSum && root.Left == nil && root.Right == nil {
		*result = append(*result, newPaths)
		return
	}

	helper(currentSum, newPaths, root.Left, targetSum, result)
	helper(currentSum, newPaths, root.Right, targetSum, result)
}

func main() {
	// result: [[5,4,11,2],[5,8,4,5]]
	// root := makeTree1()
	// targetSum := int(22)

	// result: []
	// root := makeTree2()
	// targetSum := int(5)

	// result: []
	root := makeTree3()
	targetSum := int(0)

	// result: 
	// root := makeTree()
	// targetSum := int(0)

	result := pathSum(root, targetSum)
	fmt.Printf("result = %v\n", result)
}

