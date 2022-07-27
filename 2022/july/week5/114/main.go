package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	nodes := []*TreeNode{}
	helper(root, &nodes)

	root.Left = nil
	rootHead := root
	for i, node := range nodes {
		if i == 0 {
			continue
		}

		root.Right = node
		root = root.Right
	}

	root = rootHead
}

func helper(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}

	left := root.Left
	right := root.Right
	root.Left = nil
	root.Right = nil
	*nodes = append(*nodes, root)
	helper(left, nodes)
	helper(right, nodes)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 6}
	return root
}

func makeTree2() *TreeNode {
	var root *TreeNode
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 0}
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
	fmt.Printf("root (%p) = %v\n", root, root)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// result: [1,null,2,null,3,null,4,null,5,null,6]
	// root := makeTree1()

	// result: []
	// root := makeTree2()

	// result: [0]
	root := makeTree3()

	// result: 
	// root := makeTree()

	flatten(root)
	printTree(root)
}

