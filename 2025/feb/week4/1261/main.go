package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type FindElements struct {
	Seen map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	seen := make(map[int]bool)
	elements := FindElements{Seen: seen}
	elements.Dfs(root, 0)
	return elements
}

func (this *FindElements) Find(target int) bool {
	return this.Seen[target]
}

func (this *FindElements) Dfs(root *TreeNode, target int) {
	if root == nil {
		return
	}

	this.Seen[target] = true
	this.Dfs(root.Left, target * 2 + 1)
	this.Dfs(root.Right, target * 2 + 2)
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: -1}
	root.Right = &TreeNode{Val: -1}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: -1}
	root.Left = &TreeNode{Val: -1}
	root.Left.Left = &TreeNode{Val: -1}
	root.Left.Right = &TreeNode{Val: -1}
	root.Right = &TreeNode{Val: -1}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: -1}
	root.Right = &TreeNode{Val: -1}
	root.Right.Left = &TreeNode{Val: -1}
	root.Right.Left.Left = &TreeNode{Val: -1}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// result: [null,false,true]
	// root := makeTree1()
	// findNums := []int{1, 2}

	// result: [null,true,true,false]
	// root := makeTree2()
	// findNums := []int{1, 3, 5}

	// result: [null,true,false,false,true]
	root := makeTree3()
	findNums := []int{2, 3, 4, 5}

	// result: []
	// root := makeTree()
	// findNums := []int{}

	obj := Constructor(root)
	for _, num := range findNums {
		fmt.Printf("obj.Find(%d) = %t\n", num, obj.Find(num))
	}
}

