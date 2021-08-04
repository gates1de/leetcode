package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	printNode(root)
	result := [][]int{}
	helper([]int{}, root, targetSum, &result)
	return result
}

func helper(current []int, root *TreeNode, targetSum int, result *[][]int) {
	if root == nil {
		return
	}

	current = append(current, root.Val)
	if root.Left == nil && root.Right == nil && sum(current) == targetSum {
		*result = append(*result, current)
		return
	}
	copyCurrent := make([]int, len(current))
	copy(copyCurrent, current)
	helper(copyCurrent, root.Left, targetSum, result)
	helper(copyCurrent, root.Right, targetSum, result)
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func printNode(root *TreeNode) {
    if root == nil {
        return
    }
    fmt.Printf("node %p = %v\n", root, root)
    printNode(root.Left)
    printNode(root.Right)
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

func main() {
	// result: [[5,4,11,2],[5,8,4,5]]
	root := makeTree1()
	target := int(22)

	// result: []
	// root := makeTree2()
	// target := int(5)

	// result: 
	// root := makeTree()
	// target := int(0)

	result := pathSum(root, target)
	fmt.Printf("result = %v\n", result)
}

