package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }

    root.Left, root.Right = root.Right, root.Left
    root.Left = invertTree(root.Left)
    root.Right = invertTree(root.Right)
    return root
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree3() *TreeNode {
	var root *TreeNode
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
	fmt.Printf("%p: %v\n", root, root)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// result: [4,7,2,9,6,3,1]
	// root := makeTree1()

	// result: [2,3,1]
	// root := makeTree2()

	// result: []
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := invertTree(root)
	printTree(result)
}

