package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	m := map[int][2]int{}
	var parent *TreeNode
	helper(root, parent, 0, m)
	return m[x][0] != m[y][0] && m[x][1] == m[y][1]
}

func helper(root *TreeNode, parent *TreeNode, depth int, m map[int][2]int) {
	if root == nil {
		return
	}

	if parent != nil {
		m[root.Val] = [2]int{parent.Val, depth}
	}
	helper(root.Left, root, depth + 1, m)
	helper(root.Right, root, depth + 1, m)
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Left = &TreeNode{Val: 4}
    return root
}

func makeTree2() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 4}
    root.Right.Right = &TreeNode{Val: 5}
    return root
}

func makeTree3() *TreeNode {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 4}
    return root
}

func makeTree4() *TreeNode {
    root := &TreeNode{Val: 10}
    root.Left = &TreeNode{Val: 5}
    root.Right = &TreeNode{Val: 7}
    root.Left.Left = &TreeNode{Val: 4}
    root.Left.Right = &TreeNode{Val: 6}
    root.Right.Right = &TreeNode{Val: 11}
    root.Left.Left.Left = &TreeNode{Val: 3}
    root.Left.Left.Right = &TreeNode{Val: 8}
    root.Left.Right.Right = &TreeNode{Val: 1}
	return root
}

func makeTree5() *TreeNode {
    root := &TreeNode{Val: 5}
    root.Left = &TreeNode{Val: 4}
    root.Right = &TreeNode{Val: 8}
    root.Left.Left = &TreeNode{Val: 11}
    root.Right.Left = &TreeNode{Val: 13}
    root.Right.Right = &TreeNode{Val: 9}
    root.Left.Left.Left = &TreeNode{Val: 7}
    root.Left.Left.Right = &TreeNode{Val: 2}
    root.Right.Right.Left = &TreeNode{Val: 5}
    root.Right.Right.Right = &TreeNode{Val: 1}
    return root
}

func main() {
	// result: false
	// root := makeTree1()
	// x := int(4)
	// y := int(3)

	// result: true
	// root := makeTree2()
	// x := int(5)
	// y := int(4)

	// result: false
	// root := makeTree3()
	// x := int(2)
	// y := int(3)

	// result: true
	// root := makeTree4()
	// x := int(3)
	// y := int(1)

	// result: false
	root := makeTree5()
	x := int(5)
	y := int(1)

	// result: 
	// root := makeTree()
	// x := int(0)
	// y := int(0)

	result := isCousins(root, x, y)
	fmt.Printf("result = %v\n", result)
}

