package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := make(map[int]*TreeNode, len(descriptions)*2)
	children := make(map[int]bool, len(descriptions))

	for _, d := range descriptions {
		parentVal := d[0]
		childVal := d[1]
		isLeft := d[2] == 1

		parent, exists := nodeMap[parentVal]
		if !exists {
			parent = &TreeNode{Val: parentVal}
			nodeMap[parentVal] = parent
		}
		child, exists := nodeMap[childVal]
		if !exists {
			child = &TreeNode{Val: childVal}
			nodeMap[childVal] = child
		}

		if isLeft {
			parent.Left = child
		} else {
			parent.Right = child
		}

		children[childVal] = true
	}

	for _, node := range nodeMap {
		if !children[node.Val] {
			return node
		}
	}

	return nil
}

func printList(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root)
	printList(root.Left)
	printList(root.Right)
}

func main() {
	// result: [50,20,80,15,17,19]
	// descriptions := [][]int{{20,15,1},{20,17,0},{50,20,1},{50,80,0},{80,19,1}}

	// result: [1,2,null,null,3,4]
	descriptions := [][]int{{1,2,1},{2,3,0},{3,4,1}}

	// result: []
	// descriptions := [][]int{}

	result := createBinaryTree(descriptions)
	printList(result)
}
