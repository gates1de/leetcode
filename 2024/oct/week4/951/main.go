package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	findCanonicalForm(root1)
	findCanonicalForm(root2)
	return areEquivalent(root1, root2)
}

func findCanonicalForm(root *TreeNode) {
	if root == nil {
		return
	}

	findCanonicalForm(root.Left)
	findCanonicalForm(root.Right)

	if root.Right == nil {
		return
	}

	if root.Left == nil {
		root.Left = root.Right
		root.Right = nil
		return
	}

	left := root.Left
	right := root.Right

	if left.Val > right.Val {
		root.Left = right
		root.Right = left
	}
}

func areEquivalent(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}

	return areEquivalent(root1.Left, root2.Left) && areEquivalent(root1.Right, root2.Right)
}

func makeTree1() (*TreeNode, *TreeNode) {
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Left.Left = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 5}
	root1.Left.Right.Left = &TreeNode{Val: 7}
	root1.Left.Right.Right = &TreeNode{Val: 8}
	root1.Right = &TreeNode{Val: 3}
	root1.Right.Left = &TreeNode{Val: 6}

	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 3}
	root2.Left.Right = &TreeNode{Val: 6}
	root2.Right = &TreeNode{Val: 2}
	root2.Right.Left = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 5}
	root2.Right.Right.Left = &TreeNode{Val: 8}
	root2.Right.Right.Right = &TreeNode{Val: 7}
	return root1, root2
}

func makeTree2() (*TreeNode, *TreeNode) {
	var root1 *TreeNode

	root2 := &TreeNode{Val: 1}
	return root1, root2
}

func makeTree() (*TreeNode, *TreeNode) {
	var root1, root2 *TreeNode
	return root1, root2
}

func main() {
	// result: true
	// root1, root2 := makeTree1()

	// result: false
	root1, root2 := makeTree2()

	// result: 
	// root1, root2 := makeTree()

	result := flipEquiv(root1, root2)
	fmt.Printf("result = %v\n", result)
}

