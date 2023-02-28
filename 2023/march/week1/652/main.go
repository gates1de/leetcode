package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	hashTable := make(map[string][]*TreeNode)
	result := make([]*TreeNode, 0)
	preorder(root, hashTable, &result)
	return result
}

func preorder(root *TreeNode, hashTable map[string][]*TreeNode, result *[]*TreeNode) string {
	if root == nil {
		return ""
	}

	hash := fmt.Sprintf("%v", root.Val)
	leftHash := "l" + preorder(root.Left, hashTable, result)
	rightHash := "r" + preorder(root.Right, hashTable, result)
	if leftHash != "l" && rightHash != "r" {
		hash += leftHash + rightHash
	} else if leftHash != "l" {
		hash += leftHash
	} else if rightHash != "r" {
		hash += rightHash
	}

	if hashTable[hash] == nil {
		hashTable[hash] = []*TreeNode{root}
	} else {
		hashTable[hash] = append(hashTable[hash], root)
	}

	if len(hashTable[hash]) == 2 {
		*result = append(*result, hashTable[hash][0])
	}

	return fmt.Sprintf("%v", root.Val) + leftHash + rightHash
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Left.Left = &TreeNode{Val: 4}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: [[2,4],[4]]
	// root := makeTree1()

	// result: [[1]]
	// root := makeTree2()

	// result: [[2,3],[3]]
	root := makeTree3()

	// result: 
	// root := makeTree()

	result := findDuplicateSubtrees(root)
	fmt.Printf("result = %v\n", result)
}

