package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	sums := [][]int{}
	helper(0, root, &sums)

	result := make([]float64, len(sums))
	for i, countSum := range sums {
		if countSum[0] == 0 {
			result[i] = float64(0)
		} else {
			result[i] = float64(countSum[1]) / float64(countSum[0])
		}
	}

	return result
}

func helper(level int, root *TreeNode, sums *[][]int) {
	if root == nil {
		return
	}

	if level >= len(*sums) {
		*sums = append(*sums, []int{1, root.Val})
	} else {
		(*sums)[level][0]++
		(*sums)[level][1] += root.Val
	}

	helper(level + 1, root.Left, sums)
	helper(level + 1, root.Right, sums)
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
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Left.Left = &TreeNode{Val: 15}
	root.Left.Right = &TreeNode{Val: 7}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	return root
}

func makeTree4() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 13}
	root.Right.Left = &TreeNode{Val: 14}
	root.Right.Right = &TreeNode{Val: 5}
	root.Right.Left.Left = &TreeNode{Val: 9}
	root.Right.Left.Right = &TreeNode{Val: 10}
	root.Right.Right.Left = &TreeNode{Val: 15}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: [3.00000,14.50000,11.00000]
	// root := makeTree1()

	// result: [3.00000,14.50000,11.00000]
	// root := makeTree2()

	// result: [1.00000]
	// root := makeTree3()

	// result: [2,13,9.5,11.33333]
	root := makeTree4()

	// result: 
	// root := makeTree()

	result := averageOfLevels(root)
	fmt.Printf("result = %v\n", result)
}

