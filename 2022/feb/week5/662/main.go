package main
import (
	"fmt"
	"math"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	result := int(0)
	bfs([]*TreeNode{root}, 0, []int{0}, &result)
	return result
}

func bfs(nodes []*TreeNode, level int, order []int, width *int) {
	if len(nodes) == 0 {
		return
	}

	*width = max(*width, order[len(nodes) - 1] - order[0] + 1)
	nextLevel := []*TreeNode{}
	nextOrder := []int{}
	for i, v := range nodes {
		if v.Left != nil {
			nextLevel = append(nextLevel, v.Left)
			nextOrder = append(nextOrder, order[i] * 2)
		}
		if v.Right != nil {
			nextLevel = append(nextLevel, v.Right)
			nextOrder = append(nextOrder, order[i] * 2 + 1)
		}
	}

	bfs(nextLevel, level + 1, nextOrder, width)
}

func max(values ...int) int {
	maxValue := math.MinInt32
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// Wrong Answer
func ngSolution(root *TreeNode) int {
	levels := map[int][]int{}
	helper(1, 1, root, levels)
	// fmt.Println(levels)

	result := int(0)
	for _, nodes := range levels {
		min := math.MaxInt32
		max := int(-101)
		for _, nodeNum := range nodes {
			if nodeNum < min {
				min = nodeNum
			}
			if nodeNum > max {
				max = nodeNum
			}
		}
		width := max - min + 1

		if width > result {
			result = width
		}
	}

	return result
}

func helper(level int, nodeNum int, root *TreeNode, levels map[int][]int) {
	if root == nil {
		return
	}

	if levels[level] == nil {
		levels[level] = []int{}
	}
	levels[level] = append(levels[level], nodeNum)

	level++
	helper(level, 2 * nodeNum, root.Left, levels)
	helper(level, 2 * nodeNum + 1, root.Right, levels)
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 3}
    root.Right = &TreeNode{Val: 2}
    root.Left.Left = &TreeNode{Val: 5}
    root.Left.Right = &TreeNode{Val: 3}
    root.Right.Right = &TreeNode{Val: 9}
    return root
}

func makeTree2() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 3}
    root.Left.Left = &TreeNode{Val: 5}
    root.Left.Right = &TreeNode{Val: 3}
    return root
}

func makeTree3() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 3}
    root.Right = &TreeNode{Val: 2}
    root.Left.Left = &TreeNode{Val: 5}
    return root
}

func makeTree4() *TreeNode {
    root := &TreeNode{Val: 1}
    return root
}

func makeTree5() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 3}
    root.Right = &TreeNode{Val: 2}
    root.Left.Left = &TreeNode{Val: 5}
    root.Right.Right = &TreeNode{Val: 9}
    root.Left.Left.Left = &TreeNode{Val: 6}
    root.Right.Right.Right = &TreeNode{Val: 7}
    return root
}

func makeTree6() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 1}
    root.Right = &TreeNode{Val: 1}
    root.Left.Left = &TreeNode{Val: 1}
    root.Left.Right = &TreeNode{Val: 1}
    root.Right.Left = &TreeNode{Val: 1}
    root.Right.Right = &TreeNode{Val: 1}
    root.Left.Right.Right = &TreeNode{Val: 1}
    root.Left.Right.Right.Left = &TreeNode{Val: 1}
    root.Left.Right.Right.Right = &TreeNode{Val: 1}
    root.Left.Right.Right.Left.Left = &TreeNode{Val: 1}
    root.Left.Right.Right.Left.Right = &TreeNode{Val: 1}
    root.Left.Right.Right.Right.Left = &TreeNode{Val: 1}
    root.Left.Right.Right.Right.Right = &TreeNode{Val: 1}
    root.Left.Right.Right.Left.Left.Left = &TreeNode{Val: 1}
    root.Left.Right.Right.Left.Right.Left = &TreeNode{Val: 1}
    root.Left.Right.Right.Right.Left.Right = &TreeNode{Val: 1}
    root.Left.Right.Right.Right.Right.Right = &TreeNode{Val: 1}
    return root
}

func main() {
	// result: 4
	// root := makeTree1()

	// result: 2
	// root := makeTree2()

	// result: 2
	// root := makeTree3()

	// result: 1
	// root := makeTree4()

	// result: 8
	// root := makeTree5()

	// result: 8
	root := makeTree6()

	// result: 
	// root := makeTree()

	result := widthOfBinaryTree(root)
	fmt.Printf("result = %v\n", result)
}

