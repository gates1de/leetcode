package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := int(0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *int) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftSum, leftChildrenCount := dfs(root.Left, result)
	rightSum, rightChildrenCount := dfs(root.Right, result)

	sum := root.Val + leftSum + rightSum
	childrenCount := 1 + leftChildrenCount + rightChildrenCount

	if root.Val == sum / childrenCount {
		*result++
	}

	return sum, childrenCount
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 8}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 6}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 5
	// root := makeTree1()

	// result: 1
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := averageOfSubtree(root)
	fmt.Printf("result = %v\n", result)
}

