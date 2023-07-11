package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    m := make(map[*TreeNode]string)
    dfs(root, m, "")
    result := make([]int, 0)
    for key, val := range m {
        if route(val, m[target]) == k {
            result = append(result, key.Val)
        }
    }
    return result
}

func route(a string, b string) int {
    i := int(0)
    for i < len(a) && i < len(b) {
        if a[i] == b[i] {
            i++
        } else {
            break
        }
    }
    return len(a) - i + len(b) - i
}

func dfs(root *TreeNode, m map[*TreeNode]string, path string) {
    m[root] = path
    if root.Left != nil {
        dfs(root.Left, m, path + "0")
    }
    if root.Right != nil {
        dfs(root.Right, m, path + "1")
    }
}

func makeTree1() (*TreeNode, *TreeNode) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	target := root.Left
	return root, target
}

func makeTree2() (*TreeNode, *TreeNode) {
	root := &TreeNode{Val: 1}
	target := root
	return root, target
}

func makeTree() (*TreeNode, *TreeNode) {
	var root *TreeNode
	var target *TreeNode
	return root, target
}

func main() {
	// result: [7,4,1]
	// root, target := makeTree1()
	// k := int(2)

	// result: []
	root, target := makeTree2()
	k := int(3)

	// result: 
	// root, target := makeTree()
	// k := int(0)

	result := distanceK(root, target, k)
	fmt.Printf("result = %v\n", result)
}

