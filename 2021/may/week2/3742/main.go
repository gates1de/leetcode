package main
import (
	"fmt"
	treenode "treenode/treenode"
)

func flatten(root *TreeNode)  {
    helper(root)
}

func helper (root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    left := helper(root.Left)
    right := helper(root.Right)
    if left == nil && right == nil {
        return root;
    } else if left != nil {
        temp := root.Right
        root.Right = root.Left
        root.Left = nil
        left.Right = temp
        if right == nil {
            return left
        }
    }
    return right
}

// Slow & Use more memory
func mySolution(root *treenode.TreeNode)  {
	// printNode(root)
	if root == nil {
		return
	}
	nodes := []*treenode.TreeNode{}
	nodes = preOrder(root, nodes)
	fmt.Printf("nodes = %v\n", nodes)
	for i, node := range nodes {
		if i + 1 == len(nodes) {
			break
		}
		node.Right = nodes[i + 1]
		nodes[i] = node
	}
	*root = *nodes[0]
	printNode(root)
}

func preOrder(root *treenode.TreeNode, nodes []*treenode.TreeNode) []*treenode.TreeNode {
	if root == nil {
		return nodes
	}
	nodes = append(nodes, &treenode.TreeNode{Val: root.Val})
	nodes = preOrder(root.Left, nodes)
	nodes = preOrder(root.Right, nodes)
	return nodes
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
	// result: [1, null, 2, null, 3, null, 4, null, 5, null, 6]
	rootNums := []int{1, 2, 5, 3, 4, -101, 6}

	root := treenode.Maketree(rootNums)
	flatten(root)
	fmt.Printf("result = %v\n", root)
}

