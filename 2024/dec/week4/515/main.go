package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func largestValues(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    queue := make([]*TreeNode, 0)
    levels := make(map[*TreeNode]int)
    result := make([]int, 0)
    levels[root] = 0
    queue = append(queue, root)

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        level := levels[node]

        if level >= len(result) {
            result = append(result, node.Val)
        } else {
            result[level] = max(result[level], node.Val)
        }

        if node.Left != nil {
            queue = append(queue, node.Left)
            levels[node.Left] = level + 1
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
            levels[node.Right] = level + 1
        }
    }

    return result
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
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
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: [1,3,9]
	// root := makeTree1()

	// result: [1,3]
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := largestValues(root)
	fmt.Printf("result = %v\n", result)
}

