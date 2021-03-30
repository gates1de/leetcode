package main
import (
	"treenode/treenode"
	"fmt"
	"reflect"
)

func flipMatchVoyage(root *treenode.TreeNode, voyage []int) []int {
	path := []int{}
	currentOrder := int(0)
	flippedNodes := []int{}
	dfs(root, voyage, &currentOrder, &flippedNodes, &path)
	if reflect.DeepEqual(path, voyage) {
		return flippedNodes
	}
	return []int{-1}
}

func dfs(root *treenode.TreeNode, voyage []int, currentOrder *int, flippedNodes *[]int, path *[]int) {
	if root == nil {
		return
	}
	*currentOrder++
	*path = append(*path, root.Val)

	if *currentOrder < len(voyage) && root.Left != nil {
		if voyage[len(*path)] == root.Left.Val {
			dfs(root.Left, voyage, currentOrder, flippedNodes, path)
		} else {
			*flippedNodes = append(*flippedNodes, root.Val)
			root.Swap()
			dfs(root.Left, voyage, currentOrder, flippedNodes, path)
		}
	}
	dfs(root.Right, voyage, currentOrder, flippedNodes, path)
}

func main() {
	// rootNums := []int{1, 2}
	// voyage := []int{2, 1}

	// rootNums := []int{1, 2, 3}
	// voyage := []int{1, 3, 2}

	// rootNums := []int{1, 2, 3}
	// voyage := []int{1, 2, 3}

	// rootNums := []int{1, 0, 2}
	// voyage := []int{1, 2}

	// rootNums := []int{1, 2, 0, 3}
	// voyage := []int{1, 3, 2}

	// rootNums := []int{2, 1, 0, 4, 3}
	// voyage := []int{2, 1, 3, 4}

	rootNums := []int{1, 2, 0, 3, 4}
	voyage := []int{1, 2, 4, 3}

	root := treenode.Maketree(rootNums)
	result := flipMatchVoyage(root, voyage)
	fmt.Printf("result = %v\n", result)
}

