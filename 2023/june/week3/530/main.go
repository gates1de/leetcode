package main
import (
	"fmt"
	"math"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	sortedNums := make([]int, 0, 10000)
	inorder(root, &sortedNums)
	if len(sortedNums) <= 1 {
		return 0
	}

	result := math.MaxInt32
	pre := sortedNums[0]
	for _, val := range sortedNums[1:] {
		diff := abs(pre, val)
		result = min(result, diff)
		pre = val
	}

	return result
}

func inorder(root *TreeNode, sortedNums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, sortedNums)
	*sortedNums = append(*sortedNums, root.Val)
	inorder(root.Right, sortedNums)
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
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
	root.Right.Right = &TreeNode{Val: 12}
	root.Right.Left = &TreeNode{Val: 49}
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
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := getMinimumDifference(root)
	fmt.Printf("result = %v\n", result)
}

