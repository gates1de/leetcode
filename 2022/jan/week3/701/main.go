package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	target := &TreeNode{Val: val}
	if root == nil {
		return target
	}

	if val < root.Val {
		if root.Left == nil && root.Right == nil {
			root.Left = target
		} else {
			root.Left = insertIntoBST(root.Left, val)
		}
	} else {
		if root.Left == nil && root.Right == nil {
			root.Right = target
		} else {
			root.Right = insertIntoBST(root.Right, val)
		}
	}

	return root
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
    root := &TreeNode{Val: 40}
    root.Left = &TreeNode{Val: 20}
    root.Right = &TreeNode{Val: 60}
    root.Left.Left = &TreeNode{Val: 10}
    root.Left.Right = &TreeNode{Val: 30}
    root.Right.Left = &TreeNode{Val: 50}
    root.Right.Right = &TreeNode{Val: 70}
    return root
}

func makeTree3() *TreeNode {
	var root *TreeNode
    return root
}

func printNode(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("root[%p] = %v\n", root, root)
	printNode(root.Left)
	printNode(root.Right)
}

func main() {
	// result: [4,2,7,1,3,5]
	// root := makeTree1()
	// val := int(5)

	// result: [40,20,60,10,30,50,70]
	// root := makeTree2()
	// val := int(25)

	// result: [0]
	root := makeTree3()
	val := int(0)

	// result: 
	// root := makeTree()
	// val := int(0)

	result := insertIntoBST(root, val)
	printNode(result)
}

