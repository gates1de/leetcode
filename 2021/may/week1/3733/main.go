package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedListToBST(head * ListNode) *TreeNode {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return helper(nums)
}

func helper(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return &TreeNode{
        Val: nums[len(nums) / 2],
        Left: helper(nums[:len(nums) / 2]),
        Right: helper(nums[len(nums) / 2 + 1:]),
    }
}

// Wrong Answer
func ngSolution(head *ListNode) *TreeNode {
	// printListNode(head)
	var tree *TreeNode
	for head != nil {
		tree = insert(tree, head.Val)
		fmt.Printf("head.Val = %v\n", head.Val)
		head = head.Next
	}
	printTreeNode(tree)
	return nil
}

func insert(parent *TreeNode, val int) *TreeNode {
	if parent == nil {
		parent = &TreeNode{Val: val}
		return parent
	}

	if val < parent.Val {
		parent.Left = insert(parent.Left, val)
		parent = rotateRight(parent)
	} else {
		parent.Right = insert(parent.Right, val)
		parent = rotateLeft(parent)
	}
	return parent
}

func rotateLeft(root *TreeNode) *TreeNode {
	tmp := root.Right
	root.Right = tmp.Left
	tmp.Left = root
	return root
}

func rotateRight(root *TreeNode) *TreeNode {
	tmp := root.Left
	root.Left = tmp.Right
	tmp.Right = root
	return root
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func max (a, b int) int {
	if b > a {
		return b
	}
	return a
}

////////////////////////////////////////////////////

func makeListNode(s []int, current *ListNode) *ListNode {
    if len(s) == 0 {
        if current == nil {
            current = &ListNode{Val: 0}
        }
        return current
    }

    if current == nil {
        current = &ListNode{Val: s[0]}
        return makeListNode(s[1:], current)
    }
    current.Next = makeListNode(s[1:], &ListNode{Val: s[0]})
    return current
}

func printListNode(l *ListNode) {
    if l == nil {
        return
    }
    fmt.Printf("l = %v\n", l)
    printListNode(l.Next)
}

func printTreeNode(t *TreeNode) {
    if t == nil {
        return
    }
    fmt.Printf("t = %v\n", t)
    printTreeNode(t.Left)
    printTreeNode(t.Right)
}

func main() {
	// result: [0, -3, 9, -10, null, 5]
	headNums := []int{-10, -3, 0, 5, 9}

	// result: 
	// headNums := []int{}

	var head *ListNode
	head = makeListNode(headNums, head)
	result := sortedListToBST(head)
	fmt.Printf("result = %v\n", result)
}

