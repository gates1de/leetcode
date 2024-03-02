package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }

    result := int(0)
    leftDepth := inorder(0, root.Left, &result)
    rightDepth := inorder(0, root.Right, &result)
    if result < leftDepth + rightDepth {
        result = leftDepth + rightDepth
    }
    return result
}

func inorder(depth int, root *TreeNode, result *int) int {
    if root == nil {
        return 0
    }

    leftDepth := inorder(depth, root.Left, result)
    rightDepth := inorder(depth, root.Right, result)

    if *result < leftDepth + rightDepth {
        *result = leftDepth + rightDepth
    }

    if leftDepth > rightDepth {
        return leftDepth + 1
    }

    return rightDepth + 1
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 5}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 3
	// root := makeTree1()

	// result: 1
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := diameterOfBinaryTree(root)
	fmt.Printf("result = %v\n", result)
}

