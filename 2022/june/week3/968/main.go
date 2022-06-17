package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
    result := int(0)
    covered := map[*TreeNode]bool{}
    covered[nil] = true

    dfs(root, nil, covered, &result)

    return result
}

func dfs(node, parent *TreeNode, covered map[*TreeNode]bool, result *int) {
    if node != nil {
        dfs(node.Left, node, covered, result)
        dfs(node.Right, node, covered, result)

        if (parent == nil && !covered[node]) ||
            !covered[node.Left] ||
            !covered[node.Right] {
            *result++
            covered[node] = true
            covered[parent] = true
            covered[node.Left] = true
            covered[node.Right] = true
        }
    }
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 0}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Left.Left = &TreeNode{Val: 0}
	root.Left.Left.Left.Right = &TreeNode{Val: 0}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 0}
	return root
}

func makeTree4() *TreeNode {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 0}
	root.Left.Right.Right = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 0}
	root.Right.Left = &TreeNode{Val: 0}
	return root
}

func main() {
	// result: 1
	// root := makeTree1()

	// result: 2
	// root := makeTree2()

	// result: 1
	// root := makeTree3()

	// result: 
	root := makeTree4()

	// result: 
	// root := makeTree()

	result := minCameraCover(root)
	fmt.Printf("result = %v\n", result)
}

