package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

const modulo = int(1e9 + 7)

func maxProduct(root *TreeNode) int {
	// bfs(root)
	sum := int(0)
	m := map[*TreeNode]int{}
	preorder(root, m, &sum)
	result := int(0)
	for _, s1 := range m {
		s2 := sum - s1
		s := (s1 % modulo) * (s2 % modulo)
		if result < s {
			result = s
		}
	}
	return result % modulo
}

func preorder(root *TreeNode, m map[*TreeNode]int, sum *int) int {
    // fmt.Printf("root = %v\n", root)
    if root == nil {
        return 0
    }
	*sum += root.Val
	leftSum := preorder(root.Left, m, sum)
    rightSum := preorder(root.Right, m, sum)
	m[root] = leftSum + rightSum + root.Val
	return leftSum + rightSum + root.Val
	// fmt.Printf("root = %v, leftSum = %v, rightSum = %v\n", root.Val, leftSum, rightSum)
}

func makeNode1() *TreeNode {
    root := &TreeNode{Val: 1}
    child1 := &TreeNode{Val: 2}
    child2 := &TreeNode{Val: 3}
    child3 := &TreeNode{Val: 4}
    child4 := &TreeNode{Val: 5}
    child5 := &TreeNode{Val: 6}
    root.Left = child1
    root.Right = child2
    child1.Left = child3
    child1.Right = child4
    child2.Left = child5
    return root
}

func makeNode2() *TreeNode {
    root := &TreeNode{Val: 1}
    child1 := &TreeNode{Val: 2}
    child2 := &TreeNode{Val: 3}
    child3 := &TreeNode{Val: 4}
    child4 := &TreeNode{Val: 5}
    child5 := &TreeNode{Val: 6}
	root.Right = child1
	child1.Left = child2
	child1.Right = child3
	child3.Left = child4
	child3.Right = child5
    return root
}

func makeNode3() *TreeNode {
    root := &TreeNode{Val: 2}
	child1 := &TreeNode{Val: 3}
	child2 := &TreeNode{Val: 9}
	child3 := &TreeNode{Val: 10}
	child4 := &TreeNode{Val: 7}
	child5 := &TreeNode{Val: 8}
	child6 := &TreeNode{Val: 6}
	child7 := &TreeNode{Val: 5}
	child8 := &TreeNode{Val: 4}
	child9 := &TreeNode{Val: 11}
	child10 := &TreeNode{Val: 1}
	root.Left = child1
	root.Right = child2
	child1.Left = child3
	child1.Right = child4
	child2.Left = child5
	child2.Right = child6
	child3.Left = child7
	child3.Right = child8
	child4.Left = child9
	child4.Right = child10
    return root
}

func makeNode4() *TreeNode {
    root := &TreeNode{Val: 1}
	child1 := &TreeNode{Val: 1}
	root.Left = child1
    return root
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
	// result: 110
    // root := makeNode1()

    // result: 90
    // root := makeNode2()

    // result: 1025
    // root := makeNode3()

    // result: 1
    root := makeNode4()

    // result: 
    // root := makeNode()

	result := maxProduct(root)
	fmt.Printf("result = %v\n", result)
}

