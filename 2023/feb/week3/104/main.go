package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := int(1)
	dfs(1, root.Left, &result)
	dfs(1, root.Right, &result)
	return result
}

func dfs(depth int, root *TreeNode, result *int) {
	if root == nil {
		if depth > *result {
			*result = depth
		}
		return
	}
	dfs(depth + 1, root.Left, result)
	dfs(depth + 1, root.Right, result)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 3
	// root := makeTree1()

	// result: 2
	// root := makeTree2()

	// result: 1
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := maxDepth(root)
	fmt.Printf("result = %v\n", result)
}

