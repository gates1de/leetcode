package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	startPath := make([]byte, 0)
	destPath := make([]byte, 0)

	findPath(root, startValue, &startPath)
	findPath(root, destValue, &destPath)

	result := ""
	commonPathLen := int(0)

	for commonPathLen < len(startPath) &&
		commonPathLen < len(destPath) &&
		startPath[commonPathLen] == destPath[commonPathLen] {
		commonPathLen++
	}

	for i := 0; i < len(startPath) - commonPathLen; i++ {
		result += "U"
	}

	for i := commonPathLen; i < len(destPath); i++ {
		result += string(destPath[i])
	}

	return result
}

func findPath(root *TreeNode, target int, path *[]byte) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}

	*path = append(*path, 'L')
	if findPath(root.Left, target, path) {
		return true
	}
	*path = (*path)[:len(*path) - 1]

	*path = append(*path, 'R')
	if findPath(root.Right, target, path) {
		return true
	}
	*path = (*path)[:len(*path) - 1]

	return false
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 4}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: "UURL"
	// root := makeTree1()
	// startValue := int(3)
	// destValue := int(6)

	// result: "L"
	root := makeTree2()
	startValue := int(2)
	destValue := int(1)

	// result: ""
	// root := makeTree()
	// startValue := int(0)
	// destValue := int(0)

	result := getDirections(root, startValue, destValue)
	fmt.Printf("result = %v\n", result)
}

