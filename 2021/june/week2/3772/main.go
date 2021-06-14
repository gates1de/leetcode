package main
import (
	"fmt"
	treenode "treenode/treenode"
)

func buildTree(preorder []int, inorder []int) *treenode.TreeNode {
	preIndex := int(0)
	return build(preorder, inorder, &preIndex, 0, len(preorder) - 1)
}

func build(preorder []int, inorder []int, preIndex *int, inStart int, inEnd int) *treenode.TreeNode {
	if inStart > inEnd {
		return nil
	}

	node := &treenode.TreeNode{Val: preorder[*preIndex]}
	*preIndex++

	if inStart == inEnd {
		return node
	}

	inIndex := findIndex(node.Val, inorder, inStart, inEnd)
	if inIndex == -1 {
		return node
	}

	node.Left = build(preorder, inorder, preIndex, inStart, inIndex - 1)
	node.Right = build(preorder, inorder, preIndex, inIndex + 1, inEnd)
	return node
}

func findIndex(target int, inorder []int, start int, end int) int {
	for i := start; i <= end; i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}

func printNode(root *treenode.TreeNode) {
    if root == nil {
		return
    }
    fmt.Printf("node %p = %v\n", root, root)
    printNode(root.Left)
    printNode(root.Right)
}

func main() {
	// result: [3,9,20,null,null,15,7]
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	// result: [-1]
	// preorder := []int{-1}
	// inorder := []int{-1}

	// result: []
	// preorder := []int{}
	// inorder := []int{}

	result := buildTree(preorder, inorder)
	printNode(result)
	// fmt.Printf("result = %v\n", result)
}

