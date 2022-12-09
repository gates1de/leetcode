package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
    if root == nil {
        return 0
    }

    result := int(0)
    dfs(root, root.Val, root.Val, &result)
    return result
}

func dfs(root *TreeNode, anc int, des int, result *int) {
    if root == nil {
        return
    }

    maxDiff := abs(anc, des)
    aDiff := abs(anc, root.Val)
    dDiff := abs(des, root.Val)

    if maxDiff > aDiff && maxDiff > dDiff {
        dfs(root.Left, anc, des, result)
        dfs(root.Right, anc, des, result)
        return
    }

    if aDiff > dDiff {
        maxDiff = abs(anc, root.Val)
        if *result < maxDiff {
            *result = maxDiff
        }
        dfs(root.Left, anc, root.Val, result)
        dfs(root.Right, anc, root.Val, result)
    } else {
        maxDiff = abs(des, root.Val)
        if *result < maxDiff {
            *result = maxDiff
        }
        dfs(root.Left, des, root.Val, result)
        dfs(root.Right, des, root.Val, result)
    }
}

func abs(a int, b int) int {
    if b > a {
        return b - a
    }
    return a - b
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 8}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 4}
	root.Left.Right.Right = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 14}
	root.Right.Right.Left = &TreeNode{Val: 13}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 0}
	root.Right.Right.Left = &TreeNode{Val: 3}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 7
	// root := makeTree1()

	// result: 3
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := maxAncestorDiff(root)
	fmt.Printf("result = %v\n", result)
}

