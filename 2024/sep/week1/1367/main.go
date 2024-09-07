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

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	if head.Val != root.Val {
		return false
	}

	return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
}

func makeList1() *ListNode {
	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 8}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 6}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next = &ListNode{Val: 8}
	return head
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Left.Left = &TreeNode{Val: 6}
	root.Right.Left.Right = &TreeNode{Val: 8}
	root.Right.Left.Right.Left = &TreeNode{Val: 1}
	root.Right.Left.Right.Right = &TreeNode{Val: 3}
	return root
}

func makeList() *ListNode {
	var head *ListNode
	return head
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: true
	// head := makeList1()
	// root := makeTree1()

	// result: true
	// head := makeList2()
	// root := makeTree1()

	// result: false
	head := makeList3()
	root := makeTree1()

	// result: 
	// head := makeList()
	// root := makeTree()

	result := isSubPath(head, root)
	fmt.Printf("result = %v\n", result)
}

