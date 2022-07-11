package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	bfs(root, &result)
	return result
}

func bfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	levelMap := map[*TreeNode]int{}
	added := map[int]bool{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		level := levelMap[node]
		if !added[level] {
			*result = append(*result, node.Val)
			added[level] = true
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
			levelMap[node.Right] = level + 1
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			levelMap[node.Left] = level + 1
		}
	}
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree3() *TreeNode {
	var root *TreeNode
	return root
}

func makeTree4() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Left = &TreeNode{Val: 1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: [1,3,4]
	// root := makeTree1()

	// result: [1,3]
	// root := makeTree2()

	// result: []
	// root := makeTree3()

	// result: [1,3,4,1]
	root := makeTree4()

	// result: 
	// root := makeTree()

	result := rightSideView(root)
	fmt.Printf("result = %v\n", result)
}

