package main
import (
	"fmt"
	"sort"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

const SHIFT = int(20)
const MASK = int(0xFFFFF)

func minimumOperations(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := int(0)

	for len(queue) > 0 {
		level := len(queue)
		nodes := make([]int, level)

		for i, _ := range nodes {
			node := queue[0]
			queue = queue[1:]
			nodes[i] = (node.Val << SHIFT) + i

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		sort.Ints(nodes)

		for i := 0; i < level; i++ {
			originalPosition := nodes[i] & MASK
			if originalPosition != i {
				nodes[i], nodes[originalPosition] = nodes[originalPosition], nodes[i]
				i--
				result++
			}
		}
	}

	return result
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 8}
	root.Right.Left.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 5}
	root.Right.Right.Left = &TreeNode{Val: 10}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 3
	// root := makeTree1()

	// result: 3
	// root := makeTree2()

	// result: 0
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := minimumOperations(root)
	fmt.Printf("result = %v\n", result)
}

