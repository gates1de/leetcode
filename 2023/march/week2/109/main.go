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

func sortedListToBST(head *ListNode) *TreeNode {
	vals := make([]int, 0, 20000)
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	var root *TreeNode
	return helper(root, vals)
}

func helper(root *TreeNode, vals []int) *TreeNode {
	n := len(vals)
	if n == 0 {
		return root
	} else if n == 1 {
		return &TreeNode{Val: vals[0]}
	}

	root = &TreeNode{Val: vals[n / 2]}
	lefts := vals[:n / 2]
	rights := vals[(n / 2) + 1:]

	root.Left = helper(root.Left, lefts)
	root.Right = helper(root.Right, rights)
	return root
}

func makeList1() *ListNode {
	head := &ListNode{Val: -10}
	head.Next = &ListNode{Val: -3}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next = &ListNode{Val: 9}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 0}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	return head
}

func makeList() *ListNode {
	var head *ListNode
	return head
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	// result: [0,-3,9,-10,null,5]
	// head := makeList1()

	// result: []
	// head := makeList()

	// result: [3,2,4,0]
	head := makeList2()

	// result: 
	// head := makeList()

	result := sortedListToBST(head)
	printTree(result)
}

