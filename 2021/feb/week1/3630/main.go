package main
import (
	"./treenode"
	"fmt"
)

func rightSideView(root *treenode.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// printNode(root)
	// fmt.Printf("root = %v, left = %v, right = %v\n", root, root.Left, root.Right)

	result := []int{}
	for root != nil {
		result = append(result, root.Val)
		// fmt.Printf("root = %v\n", root)
		if root.Right != nil {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return result
}

func printNode(root *treenode.TreeNode) {
    if root == nil {
        return
    }
    fmt.Printf("node = %v\n", root)
    printNode(root.Left)
    printNode(root.Right)
}

func main() {
	// rootNums := []int{1, 2, 3, -1, 5, -1, 4}
	// rootNums := []int{1, -1, 3}
	// rootNums := []int{}
	// rootNums := []int{1, 2}
	rootNums := []int{1, 2, 3, 4}

	root := treenode.Maketree(rootNums)
	result := rightSideView(root)
	fmt.Printf("result = %v\n", result)
}

