package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	var result *TreeNode
	result = helper(result, preorder)
	return result
}

func helper(root *TreeNode, preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return root
	}

	root = &TreeNode{Val: preorder[0]}
	preorder = preorder[1:]
	left := preorder
	right := []int{}
	for i, num := range preorder {
		if root.Val < num {
			left = preorder[:i]
			right = preorder[i:]
			break
		}
	}
	root.Left = helper(root.Left, left)
	root.Right = helper(root.Right, right)

	return root
}

func printPreorder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("root = %v\n", root)
	printPreorder(root.Left)
	printPreorder(root.Right)
}

func main() {
	// result: [8,5,10,1,7,null,12]
	// preorder := []int{8, 5, 1, 7, 10, 12}

	// result: [1,null,3]
	// preorder := []int{1, 3}

	// result: [4,2]
	preorder := []int{4, 2}

	// result: 
	// preorder := []int{}

	result := bstFromPreorder(preorder)
	printPreorder(result)
}

