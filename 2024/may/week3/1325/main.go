package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	stack := make([]*TreeNode, 0)
	current := root
	var lastRightNode *TreeNode

	for len(stack) > 0 || current != nil {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack) - 1]

		if current.Right != lastRightNode && current.Right != nil {
			current = current.Right
			continue
		}

		stack = stack[:len(stack) - 1]

		if current.Right == nil && current.Left == nil && current.Val == target {
			if len(stack) == 0 {
				return nil
			}

			var parent *TreeNode
			if len(stack) > 0 {
				parent = stack[len(stack) - 1]
			}

			if parent != nil {
				if parent.Left == current {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
			}
		}

		lastRightNode = current
		current = nil
	}

	return root
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 4}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Left.Left = &TreeNode{Val: 2}
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
	// result: [1,null,3,null,4]
	// root := makeTree1()
	// target := int(2)

	// result: [1,3,null,null,2]
	// root := makeTree2()
	// target := int(3)

	// result: [1]
	root := makeTree3()
	target := int(2)

	// result: []
	// root := makeTree()
	// target := int(0)

	result := removeLeafNodes(root, target)
	printTree(result)
}

