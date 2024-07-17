package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	result := make([]*TreeNode, 0)
	if root == nil {
		return result
	}

	deleteSet := make(map[int]bool)
	for _, d := range to_delete {
		deleteSet[d] = true
	}


	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)

			if deleteSet[node.Left.Val] {
				node.Left = nil
			}
		}

		if node.Right != nil {
			queue = append(queue, node.Right)

			if deleteSet[node.Right.Val] {
				node.Right = nil
			}
		}

		if deleteSet[node.Val] {
			if node.Left != nil {
				result = append(result, node.Left)
			}
			if node.Right != nil {
				result = append(result, node.Right)
			}
		}
	}

	if !deleteSet[root.Val] {
		result = append(result, root)
	}

	return result
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 4}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: [[1,2,null,4],[6],[7]]
	root := makeTree1()
	to_delete := []int{3,5}

	// result: [[1,2,4]]
	// root := makeTree2()
	// to_delete := []int{3}

	// result: []
	// root := makeTree()
	// to_delete := []int{}

	result := delNodes(root, to_delete)
	fmt.Printf("result = %v\n", result)
}

