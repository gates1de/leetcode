package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}

	for root != nil && (root.Val < low || high < root.Val) {
		if root.Val < low {
			root = root.Right
		} else if high < root.Val {
			root = root.Left
		}
	}

	if root == nil {
		return nil
	}

	root.Left = trimLow(root.Left, low)
	root.Right = trimHigh(root.Right, high)
	return root
}

func trimLow(root *TreeNode, low int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		root = trimLow(root.Right, low)
	}

	if root == nil {
		return nil
	}

	root.Left = trimLow(root.Left, low)

	return root
}

func trimHigh(root *TreeNode, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > high {
		root = trimHigh(root.Left, high)
	}

	if root == nil {
		return nil
	}

	root.Right = trimHigh(root.Right, high)

	return root
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 2}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 1}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 9}
	root.Left = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 8}
	root.Right = &TreeNode{Val: 12}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 21}
	root.Right.Right.Left = &TreeNode{Val: 19}
	root.Right.Right.Right = &TreeNode{Val: 25}
	return root
}

func makeTree4() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 1}
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
	// result: [1,null,2]
	// root := makeTree1()
	// low := int(1)
	// high := int(2)

	// result: [3,2,null,1]
	// root := makeTree2()
	// low := int(1)
	// high := int(3)

	// result: [9,null,12,10]
	// root := makeTree3()
	// low := int(9)
	// high := int(15)

	// result: [1]
	root := makeTree4()
	low := int(1)
	high := int(1)

	// result: 
	// root := makeTree()
	// low := int(0)
	// high := int(0)

	result := trimBST(root, low, high)
	printTree(result)
}

