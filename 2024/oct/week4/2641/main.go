package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	levelSums := make([]int, 100000)
	levelSums = calculateLevelSum(root, 0, levelSums)
	replace(root, 0, 0, levelSums)
	return root
}

func calculateLevelSum(node *TreeNode, level int, levelSums []int) []int {
	if node == nil {
		return levelSums
	}

	levelSums[level] += node.Val
	levelSums = calculateLevelSum(node.Left, level + 1, levelSums)
	levelSums = calculateLevelSum(node.Right, level + 1, levelSums)
	return levelSums
}

func replace(node *TreeNode, siblingSum int, level int, levelSums []int) {
	if node == nil {
		return
	}

	left := int(0)
	right := int(0)

	if node.Left != nil {
		left = node.Left.Val
	}
	if node.Right != nil {
		right = node.Right.Val
	}

	if level == 0 || level == 1 {
		node.Val = 0
	} else {
		node.Val = levelSums[level] - node.Val - siblingSum
	}

	replace(node.Left, right, level + 1, levelSums)
	replace(node.Right, left, level + 1, levelSums)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
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
	// result: [0,0,0,7,7,null,11]
	// root := makeTree1()

	// result: [0,0,0]
	root := makeTree2()

	// result: []
	// root := makeTree()

	result := replaceValueInTree(root)
	printTree(result)
}

