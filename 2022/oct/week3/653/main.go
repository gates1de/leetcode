package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	m := map[int]bool{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if m[k - node.Val] {
			return true
		}
		m[node.Val] = true

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return false
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 5}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: true
	// root := makeTree1()
	// k := int(9)

	// result: false
	// root := makeTree1()
	// k := int(28)

	// result: true
	root := makeTree2()
	k := int(9)

	// result: 
	// root := makeTree()
	// k := int(0)

	result := findTarget(root, k)
	fmt.Printf("result = %v\n", result)
}

