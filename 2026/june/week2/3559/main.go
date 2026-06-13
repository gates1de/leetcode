package main

import (
	"fmt"
)

const modulo = int64(1e9 + 7)

func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	if len(queries) == 0 {
		return []int{}
	}

	graph := make([][]int, n + 1)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	log := int(0)
	for (int(1) << log) <= n {
		log++
	}

	parent := make([][]int, log)
	for i := range log {
		parent[i] = make([]int, n + 1)
	}

	depth := make([]int, n + 1)
	queue := make([]int, 0, n)
	queue = append(queue, int(1))
	depth[1] = int(1)

	for head := int(0); head < len(queue); head++ {
		node := queue[head]
		for _, next := range graph[node] {
			if depth[next] != int(0) {
				continue
			}
			depth[next] = depth[node] + int(1)
			parent[0][next] = node
			queue = append(queue, next)
		}
	}

	for bit := int(1); bit < log; bit++ {
		for node := int(1); node <= n; node++ {
			parent[bit][node] = parent[bit-1][parent[bit-1][node]]
		}
	}

	pow2 := make([]int64, n+1)
	pow2[0] = int64(1)
	for i := int(1); i <= n; i++ {
		pow2[i] = (pow2[i-1] * int64(2)) % modulo
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		u := query[0]
		v := query[1]
		distance := depth[u] + depth[v] - int(2)*depth[LowestCommonAncestor(u, v, parent, depth)]
		if distance == int(0) {
			result[i] = int(0)
			continue
		}
		result[i] = int(pow2[distance-1])
	}

	return result
}

func LowestCommonAncestor(u int, v int, parent [][]int, depth []int) int {
	if depth[u] < depth[v] {
		u, v = v, u
	}

	diff := depth[u] - depth[v]
	for bit := int(0); diff > int(0); bit++ {
		if diff&int(1) == int(1) {
			u = parent[bit][u]
		}
		diff >>= 1
	}

	if u == v {
		return u
	}

	for bit := len(parent) - 1; bit >= 0; bit-- {
		if parent[bit][u] != parent[bit][v] {
			u = parent[bit][u]
			v = parent[bit][v]
		}
	}

	return parent[0][u]
}

func main() {
	// result: [0,1]
	// edges := [][]int{{1,2}}
	// queries := [][]int{{1,1},{1,2}}

	// result: [2,1,4]
	edges := [][]int{{1,2},{1,3},{3,4},{3,5}}
	queries := [][]int{{1,4},{3,4},{2,5}}

	// result: []
	// edges := [][]int{}
	// queries := [][]int{}

	result := assignEdgeWeights(edges, queries)
	fmt.Printf("result = %v\n", result)
}
