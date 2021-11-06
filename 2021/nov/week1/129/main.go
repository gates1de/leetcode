package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := int(0)
	helper(0, root, &result)
	return result
}

func helper(current int, root *TreeNode, result *int) {
	if root == nil {
		return
	}

	current = current * 10 + root.Val
	// fmt.Printf("current = %v\n", current)

	helper(current, root.Left, result)
	helper(current, root.Right, result)
	if root.Left == nil && root.Right == nil {
		*result += current
	}
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    return root
}

func makeTree2() *TreeNode {
    root := &TreeNode{Val: 4}
    root.Left = &TreeNode{Val: 9}
    root.Right = &TreeNode{Val: 0}
    root.Left.Left = &TreeNode{Val: 5}
    root.Left.Right = &TreeNode{Val: 1}
    return root
}

func makeTree3() *TreeNode {
    root := &TreeNode{Val: 1}
    return root
}

func makeTree4() *TreeNode {
    root := &TreeNode{Val: 5}
    root.Left = &TreeNode{Val: 6}
    root.Right = &TreeNode{Val: 8}
    root.Left.Left = &TreeNode{Val: 1}
    root.Right.Left = &TreeNode{Val: 3}
    root.Right.Right = &TreeNode{Val: 4}
    root.Left.Left.Left = &TreeNode{Val: 0}
    root.Left.Left.Right = &TreeNode{Val: 2}
    root.Right.Left.Left = &TreeNode{Val: 5}
    return root
}

func main() {
	// result: 25
	// root := makeTree1()

	// result: 1026
	// root := makeTree2()

	// result: 1
	// root := makeTree3()

	// result: 17641
	root := makeTree4()

	// result: 
	// root := makeTree()

	result := sumNumbers(root)
	fmt.Printf("result = %v\n", result)
}

