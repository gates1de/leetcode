package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := make([]*TreeNode, 0, 100)
	queue = append(queue, root)
	positionMap := make(map[*TreeNode]int)
	positionMap[root] = 1
	maxPosition := int(1)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		position := positionMap[node]

		if node.Left != nil && node.Right != nil {
			if (node.Left.Val < node.Val || node.Right.Val < node.Val) ||
				node.Left.Val > node.Right.Val {
				return false
			}
		}

		if node.Left != nil {
			positionMap[node.Left] = position * 2
			maxPosition = positionMap[node.Left]
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			positionMap[node.Right] = position * 2 + 1
			maxPosition = positionMap[node.Right]
			queue = append(queue, node.Right)
		}
	}

	if len(positionMap) != maxPosition {
		return false
	}
	return true
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: 9}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}


func main() {
	// result: true
	// root := makeTree1()

	// result: false
	// root := makeTree2()

	// result: true
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := isCompleteTree(root)
	fmt.Printf("result = %v\n", result)
}

