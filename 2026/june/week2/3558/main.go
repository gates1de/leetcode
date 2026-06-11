package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func assignEdgeWeights(edges [][]int) int {
	n := len(edges) + 1
	graph := make([][]int, n + 1)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	depth := make([]int, n + 1)
	queue := make([]int, 0, n)
	queue = append(queue, 1)
	depth[1] = int(1)

	maxDepth := int(0)
	for head := int(0); head < len(queue); head++ {
		node := queue[head]
		for _, next := range graph[node] {
			if depth[next] != 0 {
				continue
			}
			depth[next] = depth[node] + 1
			if depth[next] - 1 > maxDepth {
				maxDepth = depth[next] - 1
			}
			queue = append(queue, next)
		}
	}

	if maxDepth == 0 {
		return int(0)
	}

	result := int(1)
	for range maxDepth - 1 {
		result = (result * 2) % modulo
	}

	return result
}

func main() {
	// result: 1
	// edges := [][]int{{1, 2}}

	// result: 2
	edges := [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}}

	// result: 0
	// edges := [][]int{}

	result := assignEdgeWeights(edges)
	fmt.Printf("result = %v\n", result)
}
