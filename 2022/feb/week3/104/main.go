package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := int(0)
	helper(root, 0, &result)
	return result
}

func helper(root *TreeNode, depth int, result *int) {
	if root == nil {
		if depth > *result {
			*result = depth
		}
		return
	}

	helper(root.Left, depth + 1, result)
	helper(root.Right, depth + 1, result)
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 9}
    root.Right = &TreeNode{Val: 20}
    root.Right.Left = &TreeNode{Val: 15}
    root.Right.Right = &TreeNode{Val: 7}
    return root
}

func makeTree2() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Right = &TreeNode{Val: 2}
    return root
}

func makeTree3() *TreeNode {
    root := &TreeNode{Val: 40}
    root.Left = &TreeNode{Val: 20}
    root.Right = &TreeNode{Val: 60}
    root.Left.Left = &TreeNode{Val: 10}
    root.Left.Right = &TreeNode{Val: 30}
    root.Right.Left = &TreeNode{Val: 50}
    root.Right.Right = &TreeNode{Val: 70}
    return root
}

func makeTree4() *TreeNode {
	var root *TreeNode
    return root
}

func main() {
	// result: 3
	// root := makeTree1()

	// result: 2
	// root := makeTree2()

	// result: 3
	// root := makeTree3()

	// result: 0
	root := makeTree4()

	// result: 
	// root := makeTree()

	result := maxDepth(root)
	fmt.Printf("result = %v\n", result)
}

