package main
import (
	"./treenode"
	"fmt"
)

func convertBST(root *treenode.TreeNode) *treenode.TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	countMap := map[int]int{}
	sum := int(0)
	maxIndex := int(1)
	traversal(root, 1, &sum, &countMap)
	findMaxIndex(root, 1, &maxIndex)
	fmt.Printf("countMap = %v, maxIndex = %v\n", countMap, maxIndex)
	counts := make([]int, maxIndex)
	for i := 0; i < maxIndex; i++ {
		value, ok := countMap[i + 1]
		if !ok {
			counts[i] = -10001
        } else {
			counts[i] = value
		}
	}
	// fmt.Printf("counts = %v\n", counts)
	return treenode.Maketree(counts)
}

func traversal(node *treenode.TreeNode, index int, sum *int, countMap *map[int]int) {
	if node == nil {
		// fmt.Printf("%v: nil\n", index)
		return
	}
	traversal(node.Right, index * 2 + 1, sum, countMap)
	*sum += node.Val
	// fmt.Printf("%v: sum = %v\n", index, *sum)
	(*countMap)[index] = *sum
	traversal(node.Left, index * 2, sum, countMap)
}

func findMaxIndex(root *treenode.TreeNode, index int, result *int) {
    if root == nil {
        return
    }
	if index > *result {
		*result = index
	}
    findMaxIndex(root.Left, index * 2, result)
    findMaxIndex(root.Right, index * 2 + 1, result)
}

func printNode(root *treenode.TreeNode, index int) {
    if root == nil {
        return
    }
	fmt.Printf("%v: node = %v\n", index, root)
    printNode(root.Left, index * 2)
    printNode(root.Right, index * 2 + 1)
}

func main() {
	// rootNums := []int{4, 1, 6, 0, 2, 5, 7, -10001, -10001, -10001, 3, -10001, -10001, -10001, 8}
	// rootNums := []int{-3, -4, 0, -10001, -10001, -2, 1}
	rootNums := []int{0, -3, 2, -4, -10001, 1}

	root := treenode.Maketree(rootNums)
	// printNode(root, 1)
	result := convertBST(root)
	printNode(result, 1)
	// fmt.Printf("result = %v\n", result)
}

