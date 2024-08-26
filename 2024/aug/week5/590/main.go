package main
import (
	"fmt"
)

type Node struct {
    Val int
    Children []*Node
}

func postorder(root *Node) []int {
	result := []int{}
	helper(root, &result)
	return result
}

func helper(root *Node, result *[]int) {
	if root == nil {
		return
	}
	for _, child := range root.Children {
		helper(child, result)
	}
	*result = append(*result, root.Val)
}

func makeTree1() *Node {
	root := &Node{Val: 1}
	root.Children = []*Node{&Node{Val: 3}, &Node{Val: 2}, &Node{Val: 4}}
	root.Children[0].Children = []*Node{&Node{Val: 5}, &Node{Val: 6}}
	return root
}

func makeTree2() *Node {
	root := &Node{Val: 1}
	root.Children = []*Node{&Node{Val: 2}, &Node{Val: 3}, &Node{Val: 4}, &Node{Val: 5}}
	root.Children[1].Children = []*Node{&Node{Val: 6}, &Node{Val: 7}}
	root.Children[1].Children[1].Children = []*Node{&Node{Val: 11}}
	root.Children[1].Children[1].Children[0].Children = []*Node{&Node{Val: 14}}
	root.Children[2].Children = []*Node{&Node{Val: 4}}
	root.Children[2].Children[0].Children = []*Node{&Node{Val: 12}}
	root.Children[3].Children = []*Node{&Node{Val: 9}, &Node{Val: 10}}
	root.Children[3].Children[0].Children = []*Node{&Node{Val: 13}}
	return root
}

func makeTree() *Node {
	var root *Node
	return root
}

func main() {
	// result: [5,6,3,2,4,1]
	// root := makeTree1()

	// result: [2,6,14,11,7,3,12,8,4,13,9,10,5,1]
	root := makeTree2()

	// result: []
	// root := makeTree()

	result := postorder(root)
	fmt.Printf("result = %v\n", result)
}

