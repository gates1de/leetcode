package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	head := &TreeNode{Val: 0}
	head.Right = root
	current := head
	for current.Right != nil {
		if current.Right.Left != nil {
			rightRotate(current, current.Right)
		} else {
			current = current.Right
		}
	}

	nodeCount := int(0)
	current = head.Right
	for current != nil {
		nodeCount++
		current = current.Right
	}

	m := math.Pow(2, math.Floor(math.Log(float64(nodeCount + 1)) / math.Log(2))) - 1
	makeRotations(head, nodeCount - int(m))

	for m > 1 {
		m /= 2
		makeRotations(head, int(m))
	}

	return head.Right
}

func rightRotate(parent *TreeNode, node *TreeNode) {
	tmp := node.Left
	node.Left = tmp.Right
	tmp.Right = node
	parent.Right = tmp
}

func leftRotate(parent *TreeNode, node *TreeNode) {
	tmp := node.Right
	node.Right = tmp.Left
	tmp.Left = node
	parent.Right = tmp
}

func makeRotations(head *TreeNode, count int) {
	current := head
	for i := 0; i < count; i++ {
		tmp := current.Right
		leftRotate(current, tmp)
		current = current.Right
	}
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 3}
	root.Right.Right.Right = &TreeNode{Val: 4}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// result: [2,1,3,null,null,null,4]
	// root := makeTree1()

	// result: [2,1,3]
	root := makeTree2()

	// result: []
	// root := makeTree()


	result := balanceBST(root)
	printTree(result)
}
