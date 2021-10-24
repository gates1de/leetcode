package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func countNodes(root *TreeNode) int {
	nodes := []*TreeNode{root}
	result := int(0)
	for len(nodes) >= 1 {
		node := nodes[0]
		nodes = nodes[1:]
		if node == nil {
			continue
		}

		result++
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}

	return result
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
    root := &TreeNode{Val: 5}
    root.Left = &TreeNode{Val: 6}
    root.Right = &TreeNode{Val: 8}
    root.Left.Left = &TreeNode{Val: 11}
    root.Right.Left = &TreeNode{Val: 13}
    root.Right.Right = &TreeNode{Val: 14}
    root.Left.Left.Left = &TreeNode{Val: 10}
    root.Left.Left.Right = &TreeNode{Val: 12}
    root.Right.Left.Left = &TreeNode{Val: 15}
    return root
}

func makeTree3() *TreeNode {
	var root *TreeNode
    return root
}

func makeTree4() *TreeNode {
    root := &TreeNode{Val: 1}
    return root
}

func main() {
	// result: 6
	// root := makeTree1()

	// result: 9
	// root := makeTree2()

	// result: 0
	// root := makeTree3()

	// result: 1
	root := makeTree4()

	// result: 
	// root := makeTree()

	result := countNodes(root)
	fmt.Printf("result = %v\n", result)
}

