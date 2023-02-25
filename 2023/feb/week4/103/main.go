package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levelMap := make(map[*TreeNode]int)
	levelMap[root] = 0
	queue := make([]*TreeNode, 0, 2000)
	queue = append(queue, root)
	result := make([][]int, 1, 2000)

	for len(queue) > 0 {
		node := queue[0]
		fmt.Println(node)
		queue = queue[1:]
		level := levelMap[node]

		if level + 1 >= len(result) && (node.Left != nil || node.Right != nil) {
			result = append(result, make([]int, 0))
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			levelMap[node.Left] = level + 1
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			levelMap[node.Right] = level + 1
		}

		if level % 2 == 1 {
			result[level] = append(result[level], 0)
			copy(result[level][1:], result[level])
			result[level][0] = node.Val
		} else {
			result[level] = append(result[level], node.Val)
		}
	}

	return result
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	return root
}

func makeTree3() *TreeNode {
	var root *TreeNode
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: [[3],[20,9],[15,7]]
	// root := makeTree1()

	// result: [[1]]
	// root := makeTree2()

	// result: []
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := zigzagLevelOrder(root)
	fmt.Printf("result = %v\n", result)
}

