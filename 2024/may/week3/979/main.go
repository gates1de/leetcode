package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	moves := int(0)
	dfs(root, &moves)
	return moves
}

func dfs(root *TreeNode, moves *int) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left, moves)
	right := dfs(root.Right, moves)

	*moves += abs(left) + abs(right)

	return root.Val - 1 + left + right
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 0}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 0}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 2
	// root := makeTree1()

	// result: 3
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := distributeCoins(root)
	fmt.Printf("result = %v\n", result)
}

