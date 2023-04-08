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

func makeNode1() *Node {
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

func makeNode2() *Node {
	node1 := &Node{Val: 1}
	return node1
}

func makeNode() *Node {
	var node *Node
	return node
}

func main() {
	// result: [[2,4],[1,3],[2,4],[1,3]]
	// node := makeNode1()

	// result: [[]]
	node := makeNode2()

	// result: 
	// node := makeNode()

	result := cloneGraph(node)
	fmt.Printf("result = %v\n", result)
}

