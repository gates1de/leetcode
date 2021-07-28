package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	prune(root)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}

func prune(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return nil
		}
		return root
	}

	if root.Left != nil {
		root.Left = prune(root.Left)
	}
	if root.Right != nil {
		root.Right = prune(root.Right)
	}
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
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
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 0}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 1}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 1}
	return root
}

func main() {
	// result: [1,null,0,null,1]
	// root := makeTree1()

	// result: [1,null,1,null,1]
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := pruneTree(root)
	printNode(result)
}

