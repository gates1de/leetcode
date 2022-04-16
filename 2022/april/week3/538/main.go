package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	helper(root, 0)
	return root
}

func helper(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}

	sum = helper(root.Right, sum)
	sum += root.Val
	root.Val = sum
	return helper(root.Left, sum)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 3}
	root.Right.Right.Right = &TreeNode{Val: 8}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 1}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 1}
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	printTree(root.Right)
	fmt.Printf("%p: %v\n", root, root)
	printTree(root.Left)
}

func main() {
	// result: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
	// root := makeTree1()

	// result: [1,null,1]
	// root := makeTree2()

	// result: [7,10,4,null,9,null,null,10]
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := convertBST(root)
	printTree(result)
}

