package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	n := len(queries)
	currentMaxHeight := int(0)
	maxHeightAfterRemoval := make([]int, 100001)
	result := make([]int, n)

	traverseLeftToRight(root, 0, &currentMaxHeight, &maxHeightAfterRemoval)
	currentMaxHeight = 0
	traverseRightToLeft(root, 0, &currentMaxHeight, &maxHeightAfterRemoval)

	for i, _ := range result {
		result[i] = maxHeightAfterRemoval[queries[i]]
	}

	return result
}

func traverseLeftToRight(node *TreeNode, currentHeight int, currentMaxHeight *int, maxHeightAfterRemoval *[]int) {
	if node == nil {
		return
	}

	(*maxHeightAfterRemoval)[node.Val] = *currentMaxHeight
	*currentMaxHeight = max(*currentMaxHeight, currentHeight)

	traverseLeftToRight(node.Left, currentHeight + 1, currentMaxHeight, maxHeightAfterRemoval)
	traverseLeftToRight(node.Right, currentHeight + 1, currentMaxHeight, maxHeightAfterRemoval)
}

func traverseRightToLeft(node *TreeNode, currentHeight int, currentMaxHeight *int, maxHeightAfterRemoval *[]int) {
	if node == nil {
		return
	}

	(*maxHeightAfterRemoval)[node.Val] = max((*maxHeightAfterRemoval)[node.Val], *currentMaxHeight)
	*currentMaxHeight = max(*currentMaxHeight, currentHeight)

	traverseRightToLeft(node.Right, currentHeight + 1, currentMaxHeight, maxHeightAfterRemoval)
	traverseRightToLeft(node.Left, currentHeight + 1, currentMaxHeight, maxHeightAfterRemoval)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Right = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: [2]
	// root := makeTree1()
	// queries := []int{4}

	// result: [3,2,3,2]
	root := makeTree2()
	queries := []int{3,2,4,8}

	// result: []
	// root := makeTree()
	// queries := []int{}

	result := treeQueries(root, queries)
	fmt.Printf("result = %v\n", result)
}

