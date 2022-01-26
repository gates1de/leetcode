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

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	result := []int{}
	helper(root1, &result)
	helper(root2, &result)
	sort.Ints(result)
	return result
}

func helper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	helper(root.Left, result)
	helper(root.Right, result)
}

func makeTree1() (*TreeNode, *TreeNode) {
    root1 := &TreeNode{Val: 2}
    root1.Left = &TreeNode{Val: 1}
    root1.Right = &TreeNode{Val: 4}

    root2 := &TreeNode{Val: 1}
    root2.Left = &TreeNode{Val: 0}
    root2.Right = &TreeNode{Val: 3}
    return root1, root2
}

func makeTree2() (*TreeNode, *TreeNode) {
    root1 := &TreeNode{Val: 1}
    root1.Right = &TreeNode{Val: 8}

    root2 := &TreeNode{Val: 8}
    root2.Left = &TreeNode{Val: 1}
    return root1, root2
}

func makeTree3() (*TreeNode, *TreeNode) {
	var root1 *TreeNode
	var root2 *TreeNode
    return root1, root2
}

func main() {
	// result: [0,1,1,2,3,4]
	// root1, root2 := makeTree1()

	// result: [1,1,8,8]
	// root1, root2 := makeTree2()

	// result: []
	root1, root2 := makeTree3()

	// result: 
	// root1, root2 := makeTree()

	result := getAllElements(root1, root2)
	fmt.Printf("result = %v\n", result)
}

