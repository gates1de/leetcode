package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	result := int(0)
	traverse(root, start, &result)
	return result
}

func traverse(root *TreeNode, start int, maxDistance *int) int {
	result := int(0)
	if root == nil {
		return result
	}

	leftDepth := traverse(root.Left, start, maxDistance)
	rightDepth := traverse(root.Right, start, maxDistance)

    if root.Val == start {
        *maxDistance = max(leftDepth, rightDepth)
		result = -1
    } else if leftDepth >= 0 && rightDepth >= 0 {
        result = max(leftDepth, rightDepth) + 1
    } else {
		distance := abs(leftDepth) + abs(rightDepth)
        *maxDistance = max(*maxDistance, distance)
        result = min(leftDepth, rightDepth) - 1
    }

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Right.Left = &TreeNode{Val: 9}
	root.Left.Right.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 6}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 4
	// root := makeTree1()
	// start := int(3)

	// result: 0
	root := makeTree2()
	start := int(1)

	// result: 
	// root := makeTree()
	// start := int(0)

	result := amountOfTime(root, start)
	fmt.Printf("result = %v\n", result)
}

