package main
import (
	"fmt"
)

type Node struct {
    Val int
    Prev *Node
    Next *Node
    Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	// fmt.Printf("root (%p) = %v\n", root, root)
	if root.Child != nil {
		tmp := root.Next
		root.Child.Prev = root
		root.Next = flatten(root.Child)
		root.Child = nil
		addTail(root, tmp)
		// fmt.Printf("-----------------------\n")
		// printNode(root)
	} else {
		root.Next = flatten(root.Next)
	}
	// fmt.Printf("after root = %v\n", root)

	return root
}

func addTail(root *Node, target *Node) {
	if root == nil || target == nil {
		return
	}

	if root.Next == nil {
		root.Next = target
		target.Prev = root
		return
	}

	addTail(root.Next, target)
}

func printNode(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("node (%p) = %v\n", root, root)
	printNode(root.Next)
	printNode(root.Child)
}

func makeNode1() *Node {
	root := &Node{Val: 1}
	root.Next = &Node{Val: 2, Prev: root}
	root.Next.Next = &Node{Val: 3, Prev: root.Next}
	root.Next.Next.Next = &Node{Val: 4, Prev: root.Next.Next}
	root.Next.Next.Next.Next = &Node{Val: 5, Prev: root.Next.Next.Next}
	root.Next.Next.Next.Next.Next = &Node{Val: 6, Prev: root.Next.Next.Next.Next}

	root.Next.Next.Child = &Node{Val: 7}
	root.Next.Next.Child.Next = &Node{Val: 8, Prev: root.Next.Next.Child}
	root.Next.Next.Child.Next.Next = &Node{Val: 9, Prev: root.Next.Next.Child.Next}
	root.Next.Next.Child.Next.Next.Next = &Node{Val: 10, Prev: root.Next.Next.Child.Next.Next}

	root.Next.Next.Child.Next.Child = &Node{Val: 11}
	root.Next.Next.Child.Next.Child.Next = &Node{Val: 12, Prev: root.Next.Next.Child.Next.Child}

	return root
}

func makeNode2() *Node {
	root := &Node{Val: 1}
	root.Next = &Node{Val: 2, Prev: root}

	root.Child = &Node{Val: 3}

	return root
}

func makeNode3() *Node {
	var root *Node
	return root
}

func main() {
	// result: [1,2,3,7,8,11,12,9,10,4,5,6]
	root := makeNode1()

	// result: [1,3,2]
	// root := makeNode2()

	// result: []
	// root := makeNode3()

	// result: 
	// root := makeNode()

	// result: 
	// root := makeNode()

	result := flatten(root)
	printNode(result)
}

