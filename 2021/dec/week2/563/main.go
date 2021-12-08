package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func findTilt(root *TreeNode) int {
	result := int(0)
	helper(root, &result)
	return result
}

func helper(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, result)
	right := helper(root.Right, result)
	*result += abs(left, right)
	return left + right + root.Val
}

func abs(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    return root
}

func makeTree2() *TreeNode {
    root := &TreeNode{Val: 4}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 9}
    root.Left.Left = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 5}
    root.Right.Right = &TreeNode{Val: 7}
    return root
}

func makeTree3() *TreeNode {
    root := &TreeNode{Val: 21}
    root.Left = &TreeNode{Val: 7}
    root.Right = &TreeNode{Val: 14}
    root.Left.Left = &TreeNode{Val: 1}
    root.Left.Right = &TreeNode{Val: 1}
    root.Right.Left = &TreeNode{Val: 2}
    root.Right.Right = &TreeNode{Val: 2}
    root.Left.Left.Left = &TreeNode{Val: 3}
    root.Left.Left.Right = &TreeNode{Val: 3}
    return root
}

func makeTree4() *TreeNode {
    root := &TreeNode{Val: 1}
    return root
}

func makeTree5() *TreeNode {
	var root *TreeNode
    return root
}

func main() {
	// result: 1
	// root := makeTree1()

	// result: 15
	// root := makeTree2()

	// result: 9
	// root := makeTree3()

	// result: 0
	// root := makeTree4()

	// result: 0
	root := makeTree5()

	// result: 
	// root := makeTree()

	result := findTilt(root)
	fmt.Printf("result = %v\n", result)
}

