package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
    if root == nil {
        return 0
    }
    result := int(0)
    dfs(true, 1, root.Left, &result)
    dfs(false, 1, root.Right, &result)
    return result
}

func dfs(isLeft bool, count int, root *TreeNode, result *int) {
    if root == nil {
        return
    }
    *result = max(*result, count)

    if isLeft {
        dfs(false, count + 1, root.Right, result)
        dfs(true, 1, root.Left, result)
    }

    if !isLeft {
        dfs(true, count + 1, root.Left, result)
        dfs(false, 1, root.Right, result)
    }
}

func max(a, b int) int {
    if b > a {
        return b 
    }
    return a
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 1}
	root.Right.Right.Left = &TreeNode{Val: 1}
	root.Right.Right.Right = &TreeNode{Val: 1}
	root.Right.Right.Left.Right = &TreeNode{Val: 1}
	root.Right.Right.Left.Right.Right = &TreeNode{Val: 1}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 1}
	root.Left.Right.Left = &TreeNode{Val: 1}
	root.Left.Right.Right = &TreeNode{Val: 1}
	root.Left.Right.Left.Right = &TreeNode{Val: 1}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 3
	// root := makeTree1()

	// result: 4
	// root := makeTree2()

	// result: 0
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := longestZigZag(root)
	fmt.Printf("result = %v\n", result)
}

