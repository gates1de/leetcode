package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	levels := make(map[*TreeNode]int)
	levels[root] = 1
	preLevel := int(-1)
	preVal := int(-1)
	isEven := false

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			continue
		}

		level := levels[node]
		isEven = level % 2 != 0
		if (isEven && node.Val % 2 == 0) || (!isEven && node.Val % 2 == 1) {
			return false
		}

		if preLevel != level {
			preLevel = level
		} else {
			if (isEven && preVal >= node.Val) || (!isEven && preVal <= node.Val) {
				return false
			}
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			levels[node.Left] = level + 1
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			levels[node.Right] = level + 1
		}

		preVal = node.Val
	}

	return true
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 10}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Left = &TreeNode{Val: 12}
	root.Left.Left.Right = &TreeNode{Val: 8}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}
	root.Right.Left.Left = &TreeNode{Val: 6}
	root.Right.Right.Right = &TreeNode{Val: 2}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 7}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 9}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 7}
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

	// result: false
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := isEvenOddTree(root)
	fmt.Printf("result = %v\n", result)
}

