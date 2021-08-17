package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	result := int(0)
	preorder(root, -10001, &result)
	return result
}

func makeNode1() *TreeNode {
    root := &TreeNode{Val: 3}
    child1 := &TreeNode{Val: 1}
    child2 := &TreeNode{Val: 4}
    child3 := &TreeNode{Val: 3}
    child4 := &TreeNode{Val: 1}
    child5 := &TreeNode{Val: 5}
    root.Left = child1
	root.Right = child2
	child1.Left = child3
	child2.Left = child4
	child2.Right = child5
    return root
}

func makeNode2() *TreeNode {
    root := &TreeNode{Val: 3}
    child1 := &TreeNode{Val: 3}
    child2 := &TreeNode{Val: 4}
    child3 := &TreeNode{Val: 2}
	root.Left = child1
	child1.Left = child2
	child1.Right = child3
	return root
}

func makeNode3() *TreeNode {
    root := &TreeNode{Val: 1}
	return root
}

func preorder(root *TreeNode, currentMax int, result *int) {
	// fmt.Printf("root = %v, currentMax = %v, result = %v\n", root, currentMax, *result)
	if root == nil {
		return
	}
	if currentMax <= root.Val {
		*result++
		currentMax = root.Val
	}
	preorder(root.Left, currentMax, result)
	preorder(root.Right, currentMax, result)
}

func bfs(root *TreeNode) {
     if root == nil {
		 return
     }

     queue := []*TreeNode{root}
     for len(queue) > 0 {
         node := queue[0]
		 if node.Left != nil {
			 queue = append(queue, node.Left)
		 }
		 if node.Right != nil {
			 queue = append(queue, node.Right)
		 }
         queue = queue[1:]
         fmt.Printf("node.Val = %v, queue = %v\n", node.Val, queue)
     }
}

func main() {
    // result: 4
    // root := makeNode1()

    // result: 3
    // root := makeNode2()

    // result: 1
    root := makeNode3()

    // result:
    // root := makeNode()

	result := goodNodes(root)
	fmt.Printf("result = %v\n", result)
}

