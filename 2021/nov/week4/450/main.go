package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {
		if root.Left != nil && root.Right != nil {
			var min *TreeNode
			min = findMinNode(root.Right)
			// fmt.Printf("min = %v (%p)\n", min, min)
			root.Val = min.Val
			root.Right = deleteNode(root.Right, min.Val)
		} else if root.Left != nil && root.Right == nil {
			root = root.Left
		} else if root.Left == nil && root.Right != nil {
			root = root.Right
		} else {
			root = nil
		}

		return root
	}

	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}

func findMinNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	min := root
	for min.Left != nil {
		min = min.Left
	}
	return min
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 5}
    root.Left = &TreeNode{Val: 3}
    root.Right = &TreeNode{Val: 6}
    root.Left.Left = &TreeNode{Val: 2}
    root.Left.Right = &TreeNode{Val: 4}
    root.Right.Right = &TreeNode{Val: 7}
    return root
}

func makeTree2() *TreeNode {
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
	// result: [5,4,6,2,null,null,7]
	// root := makeTree1()
	// key := int(3)

	// result: [5,3,6,2,4,null,7]
	// root := makeTree1()
	// key := int(0)

	// result: []
	root := makeTree2()
	key := int(0)

	// result: 
	// root := makeTree()
	// key := int(0)

	result := deleteNode(root, key)
	printNode(result)
}

