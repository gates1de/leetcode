package main
import (
	"fmt"
)

type Node struct {
    Val int
    Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var result *Node
	queue := []*Node{node}
	m := map[int]*Node{}
	nodesMap := map[int][]*Node{}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if result == nil && m[n.Val] == nil {
			m[n.Val] = &Node{Val: n.Val}
			result = m[n.Val]
		}

		if node.Neighbors == nil {
			continue
		}

		nodes := []*Node{}
		if nodesMap[n.Val] != nil {
			nodes = nodesMap[n.Val]
		}
		for _, neighbor := range n.Neighbors {
			if contains(neighbor, nodes) {
				continue
			}

			queue = append(queue, neighbor)
			if nodesMap[n.Val] == nil {
				nodesMap[n.Val] = []*Node{}
			}
			if m[neighbor.Val] == nil {
				m[neighbor.Val] = &Node{Val: neighbor.Val}
			}
			nodesMap[n.Val] = append(nodesMap[n.Val], m[neighbor.Val])
		}
	}

	// fmt.Println(nodesMap)

	result.Neighbors = nodesMap[1]
	queue = []*Node{result}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if nodesMap[n.Val] == nil {
			continue
		}

		n.Neighbors = nodesMap[n.Val]
		for _, neighbor := range n.Neighbors {
			if contains(n, neighbor.Neighbors) {
				continue
			}
			queue = append(queue, neighbor)
		}
	}
	return result
}

func contains(target *Node, nodes []*Node) bool {
	if target == nil {
		return false
	}

	for _, node := range nodes {
		if node.Val == target.Val {
			return true
		}
	}

	return false
}

func makeGraph1() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}
	return node1
}

func makeGraph2() *Node {
	node1 := &Node{Val: 1}
	node1.Neighbors = []*Node{}
	return node1
}

func makeGraph3() *Node {
	var node1 *Node
	return node1
}

func printGraph(node *Node) {
	if node == nil {
		return
	}

	queue := []*Node{node}
	m := map[int][]*Node{}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		nodes := []*Node{}
		if m[n.Val] != nil {
			nodes = m[n.Val]
		}
		for _, neighbor := range n.Neighbors {
			if contains(neighbor, nodes) {
				continue
			}

			queue = append(queue, neighbor)
			if m[n.Val] == nil {
				m[n.Val] = []*Node{}
			}
			m[n.Val] = append(m[n.Val], neighbor)
		}
	}

	fmt.Println(m)
}

func main() {
	// result: [[2,4],[1,3],[2,4],[1,3]]
	// node := makeGraph1()

	// result: [[]]
	// node := makeGraph2()

	// result: []
	node := makeGraph3()

	// result: 
	// node := makeGraph()

	result := cloneGraph(node)
	printGraph(result)
}

