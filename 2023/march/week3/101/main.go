package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type LabeledNode struct {
	Node *TreeNode
	Depth int
	IsLeft bool
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

    return node1.Val == node2.Val &&
		isMirror(node1.Left, node2.Right) &&
		isMirror(node1.Right, node2.Left)
}

// Fast but more memory
func mySolution(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := make([]*LabeledNode, 0, 1000)
	preorder(root.Left, 1, true, &queue)
	fmt.Println(queue)

	if len(queue) == 0 && root.Right != nil {
		return false
	}

	reversePreorder(root.Right, 1, false, &queue)
	fmt.Println(queue)
	return len(queue) == 0
}

func preorder(root *TreeNode, depth int, isLeft bool, queue *[]*LabeledNode) {
	if root == nil {
		return
	}

	node := &LabeledNode{Node: root, Depth: depth, IsLeft: isLeft}
	*queue = append(*queue, node)

	preorder(root.Left, depth + 1, true, queue)
	preorder(root.Right, depth + 1, false, queue)
}

func reversePreorder(root *TreeNode, depth int, isLeft bool, queue *[]*LabeledNode) {
	if root == nil {
		return
	}

	if len(*queue) == 0 {
		node := &LabeledNode{Node: root, Depth: depth, IsLeft: isLeft}
		*queue = append(*queue, node)
		return
	}

	symmetric := (*queue)[0]

	if (root.Val != symmetric.Node.Val) ||
		(depth != symmetric.Depth) ||
		isLeft == symmetric.IsLeft {
		node := &LabeledNode{Node: root, Depth: depth, IsLeft: isLeft}
		*queue = append(*queue, node)
		return
	}

	*queue = (*queue)[1:]
	reversePreorder(root.Right, depth + 1, false, queue)
	reversePreorder(root.Left, depth + 1, true, queue)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 3}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	return root
}

func makeTree4() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}
	return root
}

func makeTree5() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: true
	// root := makeTree1()

	// result: false
	// root := makeTree2()

	// result: false
	// root := makeTree4()

	// result: false
	root := makeTree5()

	// result: 
	// root := makeTree()

	result := isSymmetric(root)
	fmt.Printf("result = %v\n", result)
}

