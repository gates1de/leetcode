package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		newRoot := &TreeNode{Val: val}
		newRoot.Left = root
		return newRoot
	}

	stack := []*TreeNode{root}
	parents := map[*TreeNode]*TreeNode{}
	depths := map[*TreeNode]int{}
	depths[root] = 1

	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]

		if node.Left != nil {
			stack = append(stack, node.Left)
			depths[node.Left] = depths[node] + 1
			parents[node.Left] = node
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			depths[node.Right] = depths[node] + 1
			parents[node.Right] = node
		}
	}

	for node, dep := range depths {
		if dep == depth - 1 {
			nodeLeft := node.Left
			nodeRight := node.Right
			node.Left = &TreeNode{Val: val}
			node.Right = &TreeNode{Val: val}
			node.Left.Left = nodeLeft
			node.Right.Right = nodeRight
		}
	}

	return root
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 5}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 1}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func printTree(root *TreeNode) {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		fmt.Printf("node (%p) = %v\n", node, node)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
}

func main() {
	// result: [4,1,1,2,null,null,6,3,1,5]
	// root := makeTree1()
	// val := int(1)
	// depth := int(2)

	// result: [4,2,null,1,1,3,null,null,1]
	// root := makeTree2()
	// val := int(1)
	// depth := int(3)

	// result: [1,2,3,4,null,null,null,5,5]
	// root := makeTree3()
	// val := int(5)
	// depth := int(4)

	// result: [4,2,6,1,1,1,1,3,null,null,1,5]
	root := makeTree1()
	val := int(1)
	depth := int(3)

	// result: 
	// root := makeTree()
	// val := int(0)
	// depth := int(0)

	result := addOneRow(root, val, depth)
	printTree(result)
}


