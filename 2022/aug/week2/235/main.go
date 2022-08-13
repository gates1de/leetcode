package main
import (
	"fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}
	root.Left.Right.Left = &TreeNode{Val: 3}
	root.Left.Right.Right = &TreeNode{Val: 5}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 6
	// root := makeTree1()
	// p := root.Left
	// q := root.Right

	// result: 2
	// root := makeTree1()
	// p := root.Left
	// q := root.Left.Right

	// result: 2
	root := makeTree2()
	p := root
	q := root.Left

	// result: 
	// root := makeTree()
	// p := root.Left
	// q := root.Right

	result := lowestCommonAncestor(root, p, q)
	fmt.Printf("result = %v\n", result)
}

