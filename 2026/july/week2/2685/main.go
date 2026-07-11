package main

import (
	"fmt"
)

func countCompleteComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	size := make([]int, n)
	edgeCount := make([]int, n)
	for i := range n {
		parent[i] = i
		size[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra := find(a)
		rb := find(b)
		if ra == rb {
			return
		}
		if size[ra] < size[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		size[ra] += size[rb]
		edgeCount[ra] += edgeCount[rb]
	}

	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		ra := find(a)
		rb := find(b)
		if ra == rb {
			edgeCount[ra]++
			continue
		}
		union(ra, rb)
		edgeCount[find(a)]++
	}

	result := int(0)
	for i := range n {
		if parent[i] == i && edgeCount[i] == size[i] * (size[i] - 1) / 2 {
			result++
		}
	}

	return result
}

func main() {
	// result: 3
	// n := int(6)
	// edges := [][]int{{0,1},{0,2},{1,2},{3,4}}

	// result: 1
	n := int(6)
	edges := [][]int{{0,1},{0,2},{1,2},{3,4},{3,5}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := countCompleteComponents(n, edges)
	fmt.Printf("result = %v\n", result)
}
