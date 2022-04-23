package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type BSTIterator struct {
	Nodes []*TreeNode
	Len int
	Index int
}

func Constructor(root *TreeNode) BSTIterator {
	nodes := []*TreeNode{}
	preorder(root, &nodes)
	return BSTIterator{Nodes: nodes, Len: len(nodes), Index: 0}
}

func (this *BSTIterator) Next() int {
	if !this.HasNext() {
		return -1
	}

	result := this.Nodes[this.Index].Val
	this.Index++
	return result
}

func (this *BSTIterator) HasNext() bool {
	return this.Index < this.Len
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
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 20}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
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
	// root := makeTree1()
	root := makeTree2()

	// root := makeTree()

	obj := Constructor(root)

	// root1 operators
	// result: [null, 3, 7, true, 9, true, 15, true, 20, false]
	// fmt.Printf("obj.Next() = %v\n", obj.Next())
	// fmt.Printf("obj.Next() = %v\n", obj.Next())
	// fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())
	// fmt.Printf("obj.Next() = %v\n", obj.Next())
	// fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())
	// fmt.Printf("obj.Next() = %v\n", obj.Next())
	// fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())
	// fmt.Printf("obj.Next() = %v\n", obj.Next())
	// fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())

	// root2 operators
	// result: [null, true, true, true, 3, 1, 2, false]
	fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())
	fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())
	fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())
	fmt.Printf("obj.Next() = %v\n", obj.Next())
	fmt.Printf("obj.Next() = %v\n", obj.Next())
	fmt.Printf("obj.Next() = %v\n", obj.Next())
	fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())

	// fmt.Printf("obj.Next() = %v\n", obj.Next())
	// fmt.Printf("obj.HasNext() = %v\n", obj.HasNext())
}

