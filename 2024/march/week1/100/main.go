package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }

    if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
        return false
    }

    isSameLeft := isSameTree(p.Left, q.Left)
    if !isSameLeft {
        return false
    }

    return isSameTree(p.Right, q.Right)
}

func makeTree1() (*TreeNode, *TreeNode) {
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 3}

	q := &TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 2}
	q.Right = &TreeNode{Val: 3}

	return p, q
}

func makeTree2() (*TreeNode, *TreeNode) {
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}

	q := &TreeNode{Val: 1}
	q.Right = &TreeNode{Val: 2}

	return p, q
}

func makeTree3() (*TreeNode, *TreeNode) {
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 1}

	q := &TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 1}
	q.Right = &TreeNode{Val: 2}

	return p, q
}

func makeTree() (*TreeNode, *TreeNode) {
	var p, q *TreeNode
	return p, q
}

func main() {
	// result: true
	// p, q := makeTree1()

	// result: false
	// p, q := makeTree2()

	// result: false
	p, q := makeTree3()

	// result: 
	// p, q := makeTree()

	result := isSameTree(p, q)
	fmt.Printf("result = %v\n", result)
}

