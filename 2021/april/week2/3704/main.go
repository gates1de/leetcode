package main
import (
	treenode "treenode/treenode"
	"fmt"
)

func deepestLeavesSum(root *treenode.TreeNode) int {
	sum := int(0)
	rootNums := root.GenerateArray()
	nodeCount := len(rootNums)
	deepestDepth := int(0)
	for nodeCount > 0 {
		deepestDepth++
		nodeCount /= 2
	}
	// fmt.Printf("rootNums = %v, deepestDepth = %v\n", rootNums, deepestDepth)
	dfs(root, &sum, 1, deepestDepth)
	return sum
}

func dfs(node *treenode.TreeNode, sum *int, depth int, deepestDepth int) {
	fmt.Printf("node = %v\n", node)
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil && depth == deepestDepth {
		*sum += node.Val
		fmt.Printf("sum = %v\n", *sum)
		return
	}
	dfs(node.Left, sum, depth + 1, deepestDepth)
	dfs(node.Right, sum, depth + 1, deepestDepth)
}

func main() {
	// result: 15
	rootNums := []int{1, 2, 3, 4, 5, 0, 6, 7, 0, 0, 0, 0, 8}

	// result:
	// rootNums := []int{}

	root := treenode.Maketree(rootNums)
	result := deepestLeavesSum(root)
	fmt.Printf("result = %v\n", result)
}

