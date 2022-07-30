package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pParents := map[*TreeNode]*TreeNode{}
	qParents := map[*TreeNode]*TreeNode{}

	memoParents(root, p, pParents)
	memoParents(root, q, qParents)

	pVisited := map[*TreeNode]bool{}
	pVisited[p] = true
	pTmp := pParents[p]
	for pTmp != nil {
		pVisited[pTmp] = true
		pTmp = pParents[pTmp]
	}

	qTmp := q
	for qTmp != nil {
		if pVisited[qTmp] {
			return qTmp
		}
		qTmp = qParents[qTmp]
	}

	return root
}

func memoParents(root *TreeNode, target *TreeNode, parents map[*TreeNode]*TreeNode) {
	if root == nil || root == target {
		return
	}

	if root.Left != nil {
		parents[root.Left] = root
		memoParents(root.Left, target, parents)
	}
	if root.Right != nil {
		parents[root.Right] = root
		memoParents(root.Right, target, parents)
	}
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: -1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: -2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 8}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 3
	// root := makeTree1()
	// p := root.Left
	// q := root.Right

	// result: 5
	// root := makeTree1()
	// p := root.Left
	// q := root.Left.Right.Right

	// result: 1
	// root := makeTree2()
	// p := root
	// q := root.Left

	// result: -1
	// root := makeTree3()
	// p := root.Right
	// q := root.Left.Left

	// result: 0
	root := makeTree3()
	p := root.Left.Left
	q := root.Left

	// result: 
	// root := makeTree()
	// p := root.Left
	// q := root.Right

	result := lowestCommonAncestor(root, p, q)
	fmt.Printf("result = %v\n", result)
}

