package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := int(0)
	postOrder(root, &sum)
	return root
}

func postOrder(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	postOrder(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	postOrder(root.Left, sum)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 8}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 1}
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
	// result: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
	// root := makeTree1()

	// result: [1,null,1]
	root := makeTree2()

	// result: []
	// root := makeTree()

	result := bstToGst(root)
	printTree(result)
}

