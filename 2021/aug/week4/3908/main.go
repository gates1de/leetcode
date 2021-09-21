package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := map[int]bool{}
	return helper(root, m, k)
}

func helper(root *TreeNode, m map[int]bool, k int) bool {
	if root == nil {
		return false
	}

	if m[k - root.Val] {
		return true
	}

	m[root.Val] = true

	left := helper(root.Left, m, k)
	if left {
		return true
	}
	return helper(root.Right, m, k)
}

func makeNode1() *TreeNode {
    root := &TreeNode{Val: 5}
    child1 := &TreeNode{Val: 3}
    child2 := &TreeNode{Val: 6}
    child3 := &TreeNode{Val: 2}
    child4 := &TreeNode{Val: 4}
    child5 := &TreeNode{Val: 7}
    root.Left = child1
    root.Right = child2
    child1.Left = child3
    child1.Right = child4
    child2.Right = child5
    return root
}

func makeNode2() *TreeNode {
    root := &TreeNode{Val: 2}
    child1 := &TreeNode{Val: 1}
    child2 := &TreeNode{Val: 3}
	root.Left = child1
	root.Right = child2
    return root
}

func makeNode3() *TreeNode {
    root := &TreeNode{Val: 1}
    return root
}

func main() {
	// result: true
    // root := makeNode1()
	// k := int(7)

	// result: false
    // root := makeNode1()
	// k := int(28)

	// result: true
    // root := makeNode2()
	// k := int(4)

	// result: false
    // root := makeNode2()
	// k := int(1)

	// result: true
    // root := makeNode2()
	// k := int(3)

	// result: false
    // root := makeNode3()
	// k := int(1)

	// result: false
    root := makeNode3()
	k := int(2)

	// result: 
    // root := makeNode()
	// k := int(0)

	result := findTarget(root, k)
	fmt.Printf("result = %v\n", result)
}

