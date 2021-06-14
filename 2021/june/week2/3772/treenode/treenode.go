package treenode
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func Maketree(list []int) *TreeNode {
    if len(list) == 0 {
		return nil
	}

	parent := &TreeNode{Val: list[0]}
	for i, val := range list {
        insert(parent, i + 1, 1, val)
	}
    return parent
}


func insert(parent *TreeNode, i int, parentIndex int, val int) bool {
	fmt.Printf("i = %v, parentIndex = %v, val = %v\n", i, parentIndex, val)
    if parent == nil || val < 0 {
		return false
	}

    if parent.Left == nil {
		parent.Left = &TreeNode{Val: val}
        return true
	}
    if parent.Right == nil {
		parent.Right = &TreeNode{Val: val}
        return true
	}

	isInserted := insert(parent.Left, i, parentIndex * 2, val)
    if isInserted {
        return true
	}

    return insert(parent.Right, i, parentIndex * 2 + 1, val)
}

func (node *TreeNode) Swap() {
	node.Left, node.Right = node.Right, node.Left
}

func (node *TreeNode) GenerateArray() []int {
	result := []int{}
	dfs(node, &result, 1)
	return result
}

func dfs(node *TreeNode, result *[]int, nodeIndex int) {
	if node == nil {
		return
	}
	for nodeIndex - 1 > len(*result) {
		*result = append(*result, -1)
	}
	*result = append(*result, node.Val)
	dfs(node.Left, result, 2 * nodeIndex)
	dfs(node.Right, result, 2 * nodeIndex + 1)
}

