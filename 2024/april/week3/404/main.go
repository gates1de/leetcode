package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	result := int(0)
	helper(0, false, root, &result)
	return result
}

func helper(current int, isLeft bool, root *TreeNode, result *int) {
    if root == nil {
        return
    }

    helper(current, true, root.Left, result)
    helper(current, false, root.Right, result)
    if isLeft && root.Left == nil && root.Right == nil {
        *result += root.Val
    }
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
    return root
}

func makeTree3() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
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
	// result: 24
    // root := makeTree1()

    // result: 0
    // root := makeTree2()

    // result: 2
    // root := makeTree3()

    // result: 5
    root := makeTree4()

    // result:
    // root := makeTree()

	result := sumOfLeftLeaves(root)
	fmt.Printf("result = %v\n", result)
}

