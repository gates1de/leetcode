package main
import (
	"fmt"
	"strings"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func smallestFromLeaf(root *TreeNode) string {
	result := ""
	dfs(root, "", &result)
	return result
}

func dfs(root *TreeNode, currentString string, result *string) {
	if root == nil {
		return
	}

	currentString = string(byte(root.Val + 'a')) + currentString

	if root.Left == nil && root.Right == nil {
		if len(*result) == 0 || strings.Compare(*result, currentString) > 0 {
			*result = currentString
		}
	}

	if root.Left != nil {
		dfs(root.Left, currentString, result)
	}

	if root.Right != nil {
		dfs(root.Right, currentString, result)
	}
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 25}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 2}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 1}
	root.Left.Right.Left = &TreeNode{Val: 0}
	root.Right.Left = &TreeNode{Val: 0}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: "dba"
	// root := makeTree1()

	// result: "adz"
	// root := makeTree2()

	// result: "abc"
	root := makeTree3()

	// result: ""
	// root := makeTree()

	result := smallestFromLeaf(root)
	fmt.Printf("result = %v\n", result)
}

