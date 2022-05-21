package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	// key: level, val: sum
	m := map[int]int{}
	helper(0, root, m)
	return m[len(m) - 1]
}

func helper(level int, root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}

	m[level] += root.Val
	helper(level + 1, root.Left, m)
	helper(level + 1, root.Right, m)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 8}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 3}
	root.Left.Left.Left = &TreeNode{Val: 9}
	root.Left.Right.Left = &TreeNode{Val: 1}
	root.Left.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 5}
	return root
}


func makeTree3() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 15
	// root := makeTree1()

	// result: 19
	// root := makeTree2()

	// result: 0
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := deepestLeavesSum(root)
	fmt.Printf("result = %v\n", result)
}

