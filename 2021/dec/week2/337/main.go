package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(false, root)
}

func helper(isRobbed bool, root *TreeNode) int {
	if root == nil {
		return 0
	}

	if isRobbed {
		left := helper(false, root.Left)
		right := helper(false, root.Right)
		return left + right
	}

	val := root.Val

	rob := val
	rob += helper(true, root.Left)
	rob += helper(true, root.Right)

	doNotRob := int(0)
	doNotRob += helper(false, root.Left)
	doNotRob += helper(false, root.Right)

	if rob < doNotRob {
		return doNotRob
	}

	return rob
}

func makeTree1() *TreeNode {
    root := &TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 3}
    root.Right.Right = &TreeNode{Val: 1}
    return root
}

func makeTree2() *TreeNode {
    root := &TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 4}
    root.Right = &TreeNode{Val: 5}
    root.Left.Left = &TreeNode{Val: 1}
    root.Left.Right = &TreeNode{Val: 3}
    root.Right.Right = &TreeNode{Val: 1}
    return root
}

func makeTree3() *TreeNode {
    root := &TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 4}
    root.Right = &TreeNode{Val: 5}
    root.Left.Left = &TreeNode{Val: 2}
    root.Left.Right = &TreeNode{Val: 3}
    root.Right.Right = &TreeNode{Val: 1}
    return root
}

func makeTree4() *TreeNode {
    root := &TreeNode{Val: 1}
    return root
}

func main() {
	// result: 7
	// root := makeTree1()

	// result: 9
	// root := makeTree2()

	// result: 10
	// root := makeTree3()

	// result: 1
	root := makeTree4()

	// result: 
	// root := makeTree()

	result := rob(root)
	fmt.Printf("result = %v\n", result)
}

