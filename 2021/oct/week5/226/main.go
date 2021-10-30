package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	return root
}

func printBfs(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("node = %v\n", node)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 4}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 7}
    root.Left.Left = &TreeNode{Val: 1}
    root.Left.Right = &TreeNode{Val: 3}
    root.Right.Left = &TreeNode{Val: 6}
    root.Right.Right = &TreeNode{Val: 9}
    return root
}

func makeTree2() *TreeNode {
    root := &TreeNode{Val: 2}
    root.Left = &TreeNode{Val: 1}
    root.Right = &TreeNode{Val: 3}
    return root
}

func makeTree3() *TreeNode {
    var root *TreeNode
    return root
}

func makeTree4() *TreeNode {
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

func main() {
	// result: [4,7,2,9,6,3,1]
	// root := makeTree1()

	// result: [2,3,1]
	// root := makeTree2()

	// result: []
	// root := makeTree3()

	// result: [5,8,6,14,13,null,11,null,null,null,15,12,10]
	root := makeTree4()

	// result: 
	// root := makeTree()

	result := invertTree(root)
	printBfs(result)
}

