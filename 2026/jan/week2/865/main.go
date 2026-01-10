package main

import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Result struct {
	Node *TreeNode
	Dist int
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	return dfs(root).Node
}

func dfs(root *TreeNode) Result {
	if root == nil {
		return Result{Node: nil, Dist: 0}
	}

	l := dfs(root.Left)
	r := dfs(root.Right)
	if l.Dist > r.Dist {
		return Result{Node: l.Node, Dist: l.Dist + 1}
	}
	if l.Dist < r.Dist {
		return Result{Node: r.Node, Dist: r.Dist + 1}
	}

	return Result{Node: root, Dist: l.Dist + 1}
}

func printNode(root *TreeNode) {
    if root == nil {
        return
    }
    fmt.Printf("node %p = %v\n", root, root)
    printNode(root.Left)
    printNode(root.Right)
}

func makeTree1() *TreeNode {
  root := &TreeNode{Val: 3}
  root.Left = &TreeNode{Val: 5}
  root.Right = &TreeNode{Val: 1}
  root.Left.Left = &TreeNode{Val: 6}
  root.Left.Right = &TreeNode{Val: 2}
  root.Right.Left = &TreeNode{Val: 0}
  root.Right.Right = &TreeNode{Val: 8}
  root.Left.Right.Left = &TreeNode{Val: 7}
  root.Left.Right.Right = &TreeNode{Val: 4}
  return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	return root
}

func makeTree() *TreeNode {
    root := &TreeNode{Val: 0}
    return root
}

func main() {
	// result: [2,7,4]
	// root := makeTree1()

	// result: [1]
	// root := makeTree2()

	// result: [2]
	root := makeTree3()

	// result: []
	// root := makeTree()

	result := subtreeWithAllDeepest(root)
	printNode(result)
}

