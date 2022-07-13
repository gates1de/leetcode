package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 2000)
	bfs(root, result)
	deepestLevel := int(0)
	for i, nodes := range result {
		if len(nodes) == 0 {
			deepestLevel = i
			break
		}
	}

	return result[:deepestLevel]
}

func bfs(root *TreeNode, result [][]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	levelMap := map[*TreeNode]int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		level := levelMap[node]
		if result[level] == nil {
			result[level] = []int{}
		}
		result[level] = append(result[level], node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
			levelMap[node.Left] = level + 1
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			levelMap[node.Right] = level + 1
		}
	}
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
	// result: [[3],[9,20],[15,7]]
	// root := makeTree1()

	// result: [[1]]
	// root := makeTree2()

	// result: []
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := levelOrder(root)
	fmt.Printf("result = %v\n", result)
}

