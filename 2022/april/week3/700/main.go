package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == val {
		return root
	}

	left := searchBST(root.Left, val)
	if left != nil {
		return left
	}
	return searchBST(root.Right, val)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 0}
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
	// result: [2,1,3]
	// root := makeTree1()
	// val := int(2)

	// result: []
	// root := makeTree1()
	// val := int(5)

	// result: [0]
	// root := makeTree2()
	// val := int(0)

	// result: []
	root := makeTree2()
	val := int(1)

	// result: 
	// root := makeTree()
	// val := int(0)

	result := searchBST(root, val)
	printTree(result)
}

