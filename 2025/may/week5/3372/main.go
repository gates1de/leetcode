package main
import (
	"fmt"
)

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	n := len(edges1) + 1
	count1 := build(edges1, k)
	count2 := build(edges2, k - 1)
	maxCount2 := int(0)

	for _, c := range count2 {
		if c > maxCount2 {
			maxCount2 = c
		}
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = count1[i] + maxCount2
	}

	return result
}

func dfs(node int, parent int, k int, children [][]int) int {
	if k < 0 {
		return 0
	}

	result := int(1)
	for _, child := range children[node] {
		if child == parent {
			continue
		}

		result += dfs(child, node, k - 1, children)
	}

	return result
}

func build(edges [][]int, k int) []int {
	n := len(edges) + 1
	children := make([][]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		children[u] = append(children[u], v)
		children[v] = append(children[v], u)
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = dfs(i, -1, k, children)
	}

	return result
}

func main() {
	// result: [9,7,9,8,8]
	// edges1 := [][]int{{0,1},{0,2},{2,3},{2,4}}
	// edges2 := [][]int{{0,1},{0,2},{0,3},{2,7},{1,4},{4,5},{4,6}}
	// k := int(2)

	// result: [6,3,3,3,3]
	edges1 := [][]int{{0,1},{0,2},{0,3},{0,4}}
	edges2 := [][]int{{0,1},{1,2},{2,3}}
	k := int(1)

	// result: []
	// edges1 := [][]int{}
	// edges2 := [][]int{}
	// k := int(0)

	result := maxTargetNodes(edges1, edges2, k)
	fmt.Printf("result = %v\n", result)
}

