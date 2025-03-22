package main
import (
	"fmt"
	"math"
)

func minimumCost(n int, edges [][]int, query [][]int) []int {
	parent := make([]int, n)
	depth := make([]int, n)

	for i, _ := range parent {
		parent[i] = -1
	}

	componentCost := make([]int, n)
	for i, _ := range componentCost {
		componentCost[i] = math.MaxInt32
	}

	for _, edge := range edges {
		union(edge[0], edge[1], parent, depth)
	}

	for _, edge := range edges {
		root := find(edge[0], parent)
		componentCost[root] &= edge[2]
	}

	result := make([]int, len(query))
	for i, q := range query {
		start := q[0]
		end := q[1]

		if find(start, parent) != find(end, parent) {
			result[i] = -1
		} else {
			root := find(start, parent)
			result[i] = componentCost[root]
		}
	}

	return result
}

func find(node int, parent []int) int {
	if parent[node] == -1 {
		return node
	}
	parent[node] = find(parent[node], parent)
	return parent[node]
}

func union(node1 int, node2 int, parent []int, depth []int) {
	root1 := find(node1, parent)
	root2 := find(node2, parent)

	if root1 == root2 {
		return
	}

	if depth[root1] > depth[root2] {
		root1, root2 = root2, root1
	}

	parent[root2] = root1

	if depth[root1] == depth[root2] {
		depth[root2]++
	}
}

func main() {
	// result: [1,-1]
	// n := int(5)
	// edges := [][]int{{0,1,7},{1,3,7},{1,2,1}}
	// query := [][]int{{0,3},{3,4}}

	// result: [0]
	n := int(3)
	edges := [][]int{{0,2,7},{0,1,15},{1,2,6},{1,2,1}}
	query := [][]int{{1,2}}

	// result: []
	// n := int(0)
	// edges := [][]int{}
	// query := [][]int{}

	result := minimumCost(n, edges, query)
	fmt.Printf("result = %v\n", result)
}

