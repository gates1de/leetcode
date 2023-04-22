package main
import (
	"fmt"
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
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 9}
	root.Left.Left.Left = &TreeNode{Val: 6}
	root.Right.Right.Left = &TreeNode{Val: 7}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 5}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 4
	// root := makeTree1()

	// result: 7
	// root := makeTree2()

	// result: 2
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := widthOfBinaryTree(root)
	fmt.Printf("result = %v\n", result)
}

