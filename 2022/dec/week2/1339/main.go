package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	sums := make([]int, 0, 50000)
	dfs(root, &sums)
	allSum := sums[len(sums) - 1]
	result := int(0)
	for _, sum := range sums {
		product := sum * (allSum - sum) % modulo
		if result < product {
			result = product
		}
	}
	return result % modulo
}

func dfs(root *TreeNode, sums *[]int) int {
	if root == nil {
		return 0
	}

	sum := root.Val
	sum += dfs(root.Left, sums)
	sum += dfs(root.Right, sums)
	*sums = append(*sums, sum)
	return sum
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 6}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 110
	// root := makeTree1()

	// result: 90
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := maxProduct(root)
	fmt.Printf("result = %v\n", result)
}

