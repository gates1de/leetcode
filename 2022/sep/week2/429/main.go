package main
import (
	"fmt"
)

type Node struct {
    Val int
    Children []*Node
}

func levelOrder(root *Node) [][]int {
	levels := map[*Node]int{}
	queue := []*Node{root}
	result := [][]int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nodeLevel := levels[node]

		if len(result) <= nodeLevel {
			result = append(result, []int{})
		}
		result[nodeLevel] = append(result[nodeLevel], node.Val)

		if node.Children == nil || len(node.Children) == 0 {
			continue
		}

		for _, n := range node.Children {
			nextLevel := nodeLevel + 1
			levels[n] = nextLevel
			queue = append(queue, n)
		}
	}

	return result
}

func makeNode1() *Node {
	root := &Node{Val: 1}
	root.Children = []*Node{&Node{Val: 3}, &Node{Val: 2}, &Node{Val: 4}}
	root.Children[0].Children = []*Node{&Node{Val: 5}, &Node{Val: 6}}
	return root
}

func makeNode() *Node {
	var root *Node
	return root
}

func main() {
	// result: [[1],[3,2,4],[5,6]]
	root := makeNode1()

	// result: 
	// root := makeNode()

	result := levelOrder(root)
	fmt.Printf("result = %v\n", result)
}

