package main
import (
	"fmt"
	"sort"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func recoverTree(root *TreeNode)  {
	nodes := []*TreeNode{}
	preorder(root, &nodes)
	sortedNodes := make([]*TreeNode, len(nodes))
	copy(sortedNodes, nodes)
	sort.Slice(sortedNodes, func(a, b int) bool {
		return sortedNodes[a].Val < sortedNodes[b].Val
	})

	var target1 *TreeNode
	var target2 *TreeNode
	for i := 0; i < len(nodes); i++ {
		n := nodes[i]
		sn := sortedNodes[i]
		if n.Val != sn.Val {
			if target1 == nil {
				target1 = n
			} else {
				target2 = n
				break
			}
		}
	}

	target1.Val, target2.Val = target2.Val, target1.Val
}

func preorder(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}

	preorder(root.Left, nodes)
	*nodes = append(*nodes, root)
	preorder(root.Right, nodes)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 2}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Left.Left = &TreeNode{Val: 5}
	root.Right.Right.Left = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 9}
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	printTree(root.Left)
	fmt.Printf("%p: %v\n", root, root)
	printTree(root.Right)
}

func main() {
	// result: [3,1,null,null,2]
	// root := makeTree1()

	// result: [2,1,4,null,null,3]
	// root := makeTree2()

	// result: [5,3,6,2,4,null,8,1,null,null,null,7,9]
	root := makeTree3()

	// result: 
	// root := makeTree()

	recoverTree(root)
	printTree(root)
}

