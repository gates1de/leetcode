package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	nodes := map[int]int{}
	result := int(0)
	helper(nodes, root, &result)
	return result
}

func helper(nodes map[int]int, root *TreeNode, result *int) {
	if root == nil {
		return
	}

	newNodes := map[int]int{}
	for key, val := range nodes {
		newNodes[key] = val
	}

	if _, ok := newNodes[root.Val]; ok {
		newNodes[root.Val]--
	} else {
		newNodes[root.Val]++
	}

	if newNodes[root.Val] == 0 {
		delete(newNodes, root.Val)
	}

	if root.Left == nil && root.Right == nil {
		if len(newNodes) <= 1 {
			*result++
		}
		return
	}

	helper(newNodes, root.Left, result)
	helper(newNodes, root.Right, result)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 1}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Left.Right.Right = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 9}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 2
	// root := makeTree1()

	// result: 1
	// root := makeTree2()

	// result: 1
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := pseudoPalindromicPaths(root)
	fmt.Printf("result = %v\n", result)
}

