package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return root.Val
	}

	nums := make([]int, 1)
	nums[0] = root.Val
	result := int(0)
	preorder(nums, root.Left, &result)
	preorder(nums, root.Right, &result)
	return result
}

func preorder(nums []int, root *TreeNode, result *int) {
	if root == nil {
		return
	}

	newNums := make([]int, len(nums) + 1)
	copy(newNums[1:], nums[0:])
	newNums[0] = root.Val
	if root.Left == nil && root.Right == nil {
		sum := int(0)
		for i, num := range newNums {
			sum += getPlace(i + 1) * num
		}
		*result += sum
		return
	}

	preorder(newNums, root.Left, result)
	preorder(newNums, root.Right, result)
}

func getPlace(digit int) int {
	result := int(1)
	for i := 1; i < digit; i ++ {
		result *= 10
	}
	return result
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: 25
	// root := makeTree1()

	// result: 1026
	root := makeTree2()

	// result: 
	// root := makeTree()

	result := sumNumbers(root)
	fmt.Printf("result = %v\n", result)
}

