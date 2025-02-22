package main
import (
	"fmt"
	"strconv"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func recoverFromPreorder(traversal string) *TreeNode {
	index := int(0)
	return helper(&index, traversal, 0)
}

func helper(index *int, traversal string, depth int) *TreeNode {
	if *index >= len(traversal) {
		return nil
	}

	dashCount := int(0)
	for (*index + dashCount) < len(traversal) && traversal[*index + dashCount] == '-' {
		dashCount++
	}

	if dashCount != depth {
		return nil
	}

	*index += dashCount
	value := int(0)
	for true {
		if *index >= len(traversal) {
			break
		}

		num, err := strconv.Atoi(string(traversal[*index]))
		if err == nil {
			value = value * 10 + num
			*index++
		} else {
			break
		}
	}

	node := &TreeNode{Val: value}
	node.Left = helper(index, traversal, depth + 1)
	node.Right = helper(index, traversal, depth + 1)
	return node
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// result: [1,2,5,3,4,6,7]
	// traversal := "1-2--3--4-5--6--7"

	// result: [1,2,5,3,null,6,null,4,null,7]
	// traversal := "1-2--3---4-5--6---7"

	// result: [1,401,null,349,88,90]
	traversal := "1-401--349---90--88"

	// result: []
	// traversal := ""

	result := recoverFromPreorder(traversal)
	printTree(result)
}

