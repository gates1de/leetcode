package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(nil, preorder, inorder)
}

func helper(root *TreeNode, preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 1 && len(inorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	headVal := preorder[0]
	if root == nil {
		root = &TreeNode{Val: headVal}
	}

	midIndex := int(-1)
	for i, num := range inorder {
		if headVal == num {
			midIndex = i
			break
		}
	}

	if midIndex == -1 {
		return nil
	}

	inorderLeft := inorder[:midIndex]
	inorderRight := inorder[midIndex + 1:]
	preorderLeft := make([]int, len(preorder))
	preorderRight := make([]int, len(preorder))
	for _, num := range inorderLeft {
		index := findIndex(num, preorder)
		if index >= 0 {
			preorderLeft[index] = num
		}
	}
	for _, num := range inorderRight {
		index := findIndex(num, preorder)
		if index >= 0 {
			preorderRight[index] = num
		}
	}

	preorderLeft = preorderLeft[1:len(inorderLeft) + 1]
	preorderRight = preorderRight[len(preorderRight) - len(inorderRight):]

	if len(preorderLeft) > 0 && len(inorderLeft) > 0 {
		root.Left = helper(root.Left, preorderLeft, inorderLeft)
	}
	if len(preorderRight) > 0 && len(inorderRight) > 0 {
		root.Right = helper(root.Right, preorderRight, inorderRight)
	}
	return root
}

func findIndex(target int, nums []int) int {
	for i, num := range nums {
		if target == num {
			return i
		}
	}
	return -1
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("root (%p) = %v\n", root, root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// result: [3,9,20,null,null,15,7]
	// preorder := []int{3,9,20,15,7}
	// inorder := []int{9,3,15,20,7}

	// result: [-1]
	// preorder := []int{-1}
	// inorder := []int{-1}

	// result: [1,null,2]
	preorder := []int{1, 2}
	inorder := []int{2, 1}

	// result: 
	// preorder := []int{}
	// inorder := []int{}

	result := buildTree(preorder, inorder)
	printTree(result)
}

