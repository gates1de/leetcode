package main
import (
	"./treenode"
	"fmt"
)

func convertBST(root *treenode.TreeNode) *treenode.TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	sum := int(0)
	result := &treenode.TreeNode{Val: root.Val}
	result = traversal(root, result, 1, &sum)
	return result
}

func traversal(node *treenode.TreeNode, result *treenode.TreeNode, index int, sum *int) *treenode.TreeNode {
	if node == nil {
		return nil
	}
	result = &treenode.TreeNode{Val: 0}
	result.Right = traversal(node.Right, result.Right, index * 2 + 1, sum)
	*sum += node.Val
	result.Val = *sum
	result.Left = traversal(node.Left, result.Left, index * 2, sum)
	return result
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

