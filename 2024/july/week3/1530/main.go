package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	return postOrder(root, distance)[11]
}

func postOrder(node *TreeNode, distance int) [12]int {
	result := [12]int{}
	if node == nil {
		return result
	} else if node.Left == nil && node.Right == nil {
		result[0] = 1
		return result
	}

	lefts := postOrder(node.Left, distance)
	rights := postOrder(node.Right, distance)

	for i := 0; i < 10; i++ {
		result[i + 1] += lefts[i] + rights[i]
	}

	result[11] += lefts[11] + rights[11]

	for d1 := 0; d1 <= distance; d1++ {
		for d2 := 0; d2 <= distance; d2++ {
			if d1 + d2 + 2 <= distance {
				result[11] += lefts[d1] * rights[d2]
			}
		}
	}

	return result
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 3}
	root.Right.Right.Right = &TreeNode{Val: 2}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 1
	// root := makeTree1()
	// distance := int(3)

	// result: 2
	// root := makeTree2()
	// distance := int(3)

	// result: 1
	root := makeTree3()
	distance := int(3)

	// result: 
	// root := makeTree3()
	// distance := int(0)

	result := countPairs(root, distance)
	fmt.Printf("result = %v\n", result)
}

