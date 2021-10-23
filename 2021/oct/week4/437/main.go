package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	parents := map[*TreeNode]*TreeNode{}
	result := int(0)
	helper(root, parents, targetSum, &result)
	return result
}

func helper(root *TreeNode, parents map[*TreeNode]*TreeNode, targetSum int, result *int) {
	if root == nil {
		return
	}

	parent := root
	sum := int(0)
	for parent != nil {
		sum += parent.Val
		parent = parents[parent]
		if sum == targetSum {
			*result++
		}
		// fmt.Printf("root = %v, sum = %v\n", root, sum)
	}

	if root.Left != nil {
		parents[root.Left] = root
		helper(root.Left, parents, targetSum, result)
	}

	if root.Right != nil {
		parents[root.Right] = root
		helper(root.Right, parents, targetSum, result)
	}
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 10}
    root.Left = &TreeNode{Val: 5}
    root.Right = &TreeNode{Val: -3}
    root.Left.Left = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 2}
    root.Right.Right = &TreeNode{Val: 11}
    root.Left.Left.Left = &TreeNode{Val: 3}
    root.Left.Left.Right = &TreeNode{Val: -2}
    root.Left.Right.Right = &TreeNode{Val: 1}
	return root
}

func makeTree2() *TreeNode {
    root := &TreeNode{Val: 5}
    root.Left = &TreeNode{Val: 4}
    root.Right = &TreeNode{Val: 8}
    root.Left.Left = &TreeNode{Val: 11}
    root.Right.Left = &TreeNode{Val: 13}
    root.Right.Right = &TreeNode{Val: 4}
    root.Left.Left.Left = &TreeNode{Val: 7}
    root.Left.Left.Right = &TreeNode{Val: 2}
    root.Right.Right.Left = &TreeNode{Val: 5}
    root.Right.Right.Right = &TreeNode{Val: 1}
    return root
}

func makeTree3() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Left = &TreeNode{Val: 4}
    root.Left.Right = &TreeNode{Val: 5}
    return root
}

func makeTree4() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    return root
}

func main() {
	// result: 3
	// root := makeTree1()
	// targetSum := int(8)

	// result: 3
	// root := makeTree2()
	// targetSum := int(22)

	// result: 2
	// root := makeTree3()
	// targetSum := int(7)

	// result: 0
	root := makeTree4()
	targetSum := int(4)

	// result: 
	// root := makeTree()
	// targetSum := int(0)

	result := pathSum(root, targetSum)
	fmt.Printf("result = %v\n", result)
}

