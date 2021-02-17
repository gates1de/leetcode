package main
import (
	"fmt"
	"./treenode"
)

func verticalTraversal(root *treenode.TreeNode) [][]int {
	return [][]int{}
}

func main() {
	rootNums := []int{3, 9, 20, -1, -1, 15, 7}
	root := treenode.Maketree(rootNums)
	result := verticalTraversal(root)
	fmt.Printf("result = %v\n", result)
}

