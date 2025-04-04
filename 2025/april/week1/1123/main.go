package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	result, _ := dfs(root)
	return result
}

func dfs(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	leftRoot, leftVal := dfs(root.Left)
	rightRoot, rightVal := dfs(root.Right)

	if leftVal > rightVal {
		return leftRoot, leftVal + 1
	}

	if leftVal < rightVal {
		return rightRoot, rightVal + 1
	}

	return root, leftVal + 1
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// result: [2,7,4]
	// root := makeTree1()

	// result: [1]
	// root := makeTree2()

	// result: [2]
	root := makeTree3()

	// result: []
	// root := makeTree()

	result := lcaDeepestLeaves(root)
	printTree(result)
}

