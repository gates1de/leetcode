package main
import (
	"fmt"
	"reflect"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}

	leaves1 := make([]int, 0, 200)
	dfs(root1, &leaves1)
	leaves2 := make([]int, 0, 200)
	dfs(root2, &leaves2)

	return reflect.DeepEqual(leaves1, leaves2)
}

func dfs(root *TreeNode, leaves *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
	}
	if root.Left != nil {
		dfs(root.Left, leaves)
	}
	if root.Right != nil {
		dfs(root.Right, leaves)
	}
}

func makeTree1() (*TreeNode, *TreeNode) {
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 5}
	root1.Left.Left = &TreeNode{Val: 6}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Left.Right.Left = &TreeNode{Val: 7}
	root1.Left.Right.Right = &TreeNode{Val: 4}
	root1.Right = &TreeNode{Val: 1}
	root1.Right.Left = &TreeNode{Val: 9}
	root1.Right.Right = &TreeNode{Val: 8}

	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 5}
	root2.Left.Left = &TreeNode{Val: 6}
	root2.Left.Right = &TreeNode{Val: 7}
	root2.Right = &TreeNode{Val: 1}
	root2.Right.Left = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 2}
	root2.Right.Right.Left = &TreeNode{Val: 9}
	root2.Right.Right.Right = &TreeNode{Val: 8}
	return root1, root2
}

func makeTree2() (*TreeNode, *TreeNode) {
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}

	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 3}
	root2.Right = &TreeNode{Val: 2}
	return root1, root2
}

func makeTree() (*TreeNode, *TreeNode) {
	var root1 *TreeNode
	var root2 *TreeNode
	return root1, root2
}

func main() {
	// result: true
	// root1, root2 := makeTree1()

	// result: false
	root1, root2 := makeTree2()

	// result: 
	// root1, root2 := makeTree()

	result := leafSimilar(root1, root2)
	fmt.Printf("result = %v\n", result)
}

