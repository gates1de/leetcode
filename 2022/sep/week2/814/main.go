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
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Left != nil || root.Right != nil || root.Val == 1 {
		return root
	}

	return nil
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

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 1}
	root.Left.Left.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 0}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 1}
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
	fmt.Printf("root(%p) = %v\n", root, root)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// result: [1,null,0,null,1]
	// root := makeTree1()

	// result: [1,null,1,null,1]
	// root := makeTree2()

	// result: [1,1,0,1,1,null,1]
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := pruneTree(root)
	printTree(result)
}

