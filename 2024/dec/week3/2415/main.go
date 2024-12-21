package main
import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	dfs(root.Left, root.Right, 0)
	return root
}

func dfs(left *TreeNode, right *TreeNode, level int) {
	if left == nil || right == nil {
		return
	}

	if level % 2 == 0 {
		left.Val, right.Val = right.Val, left.Val
	}

	dfs(left.Left, right.Right, level + 1)
	dfs(left.Right, right.Left, level + 1)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 8}
	root.Left.Right = &TreeNode{Val: 13}
	root.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 21}
	root.Right.Right = &TreeNode{Val: 34}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 13}
	root.Right = &TreeNode{Val: 11}
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
	// result: [2,5,3,8,13,21,34]
	// root := makeTree1()

	// result: [7,11,13]
	root := makeTree2()

	// result: []
	// root := makeTree()

	result := reverseOddLevels(root)
	printTree(result)
}

