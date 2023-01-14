package main
import (
	"fmt"
)

func countSubTrees(n int, edges [][]int, labels string) []int {
	response := make([]int, n)
	node2edges := make([][]int, n)
	labelCounter := make([]int, 26)
    for _, edge := range edges {
        node2edges[edge[0]] = append(node2edges[edge[0]], edge[1])
        node2edges[edge[1]] = append(node2edges[edge[1]], edge[0])
    }
    dfs(0, -1, labels, labelCounter, node2edges, response)
    return response
}

func dfs(nodeId int, parentNodeId int, labels string, labelCounter []int, node2edges [][]int, response []int) {
    nodeLabelId := labels[nodeId] - 97
    before := labelCounter[nodeLabelId]
    labelCounter[nodeLabelId]++
    for _, nextNodeId := range node2edges[nodeId] {
        if nextNodeId == parentNodeId {
            continue
        }
        dfs(nextNodeId, nodeId, labels, labelCounter, node2edges, response)
    }
    response[nodeId] = labelCounter[nodeLabelId] - before    
}

type Node struct {
	Index int
	Label byte
	Count int
	Adjacents []*Node
	IsVisited bool
}

// Wrong Answer
func ngSolution(n int, edges [][]int, labels string) []int {
	nodes := make([]*Node, n)
	for _, edge := range edges {
		node1 := nodes[edge[0]]
		if node1 == nil {
			node1 = &Node{Index: edge[0], Label: labels[edge[0]]}
		}
		if node1.Adjacents == nil {
			node1.Adjacents = make([]*Node, 0, n)
		}

		node2 := nodes[edge[1]]
		if node2 == nil {
			node2 = &Node{Index: edge[1], Label: labels[edge[1]]}
		}
		if node2.Adjacents == nil {
			node2.Adjacents = make([]*Node, 0, n)
		}

		node1.Adjacents = append(node1.Adjacents, node2)
		node2.Adjacents = append(node2.Adjacents, node1)
		nodes[edge[0]] = node1
		nodes[edge[1]] = node2
	}
	preorder(nodes[0])
	result := make([]int, n)
	for i, node := range nodes {
		result[i] = node.Count
	}

	return result
}

func preorder(node *Node) int {
	if node == nil || node.IsVisited {
		return 0
	}

	node.IsVisited = true
	sum := int(1)
	for _, a := range node.Adjacents {
		count := preorder(a)
		if node.Label == a.Label {
			sum += count
		}
	}

	node.Count = sum
	return sum
}

func printNode(node *Node) {
	if node == nil || node.IsVisited {
		return
	}

	fmt.Println(node)
	node.IsVisited = true
	for _, a := range node.Adjacents {
		printNode(a)
	}
}

func main() {
	// result: [2,1,1,1,1,1,1]
	// n := int(7)
	// edges := [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}
	// labels := "abaedcd"

	// result: [4,2,1,1]
	// n := int(4)
	// edges := [][]int{{0,1},{1,2},{0,3}}
	// labels := "bbbb"

	// result: [3,2,1,1,1]
	n := int(5)
	edges := [][]int{{0,1},{0,2},{1,3},{0,4}}
	labels := "aabab"

	// result: 
	// n := int(0)
	// edges := [][]int{}
	// labels := ""

	result := countSubTrees(n, edges, labels)
	fmt.Printf("result = %v\n", result)
}

