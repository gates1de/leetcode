package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	preIndex := int(0)
	postIndex := int(0)
	return constructTree(&preIndex, &postIndex, preorder, postorder)
}

func constructTree(preIndex *int, postIndex *int, preorder []int, postorder[]int) *TreeNode {
	root := &TreeNode{Val: preorder[*preIndex]}
	*preIndex++

	if root.Val != postorder[*postIndex] {
		root.Left = constructTree(preIndex, postIndex, preorder, postorder)
	}
	if root.Val != postorder[*postIndex] {
		root.Right = constructTree(preIndex, postIndex, preorder, postorder)
	}

	*postIndex++
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
	// result: [1,2,3,4,5,6,7]
	// preorder := []int{1,2,4,5,3,6,7}
	// postorder := []int{4,5,2,6,7,3,1}

	// result: [1]
	preorder := []int{1}
	postorder := []int{1}

	// result: []
	// preorder := []int{}
	// postorder := []int{}

	result := constructFromPrePost(preorder, postorder)
	printTree(result)
}

