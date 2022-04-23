package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    
    nodes := []*TreeNode{}
    helper(root, &nodes)
	var result *TreeNode
	for _, node := range nodes {
		result = insertRight(result, node)
	}

    return result
}

func helper(root *TreeNode, nodes *[]*TreeNode) {
    if root == nil {
        return
    }
    
    helper(root.Left, nodes)
    *nodes = append(*nodes, root)
    helper(root.Right, nodes)
}

func insertRight(root *TreeNode, target *TreeNode) *TreeNode {
    if root == nil {
		target.Left = nil
		target.Right = nil
        return target
    }
    
    if target == nil {
        return root
    }
    
	target.Left = nil
	target.Right = nil
	right := root
    for right.Right != nil {
        right = right.Right
    }
	right.Right = target
	return root
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Left.Left = &TreeNode{Val: 1}
	root.Right.Right.Left = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 9}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 7}
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	printTree(root.Left)
	fmt.Printf("%p: %v\n", root, root)
	printTree(root.Right)
}

func main() {
	// result: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9] 
	// root := makeTree1()

	// result: [1,null,5,null,7]
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := increasingBST(root)
	printTree(result)
}

