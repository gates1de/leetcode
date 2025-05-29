package main
import (
	"fmt"
)

func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
	n := len(edges1) + 1
	m := len(edges2) + 1
	color1 := make([]int, n)
	color2 := make([]int, m)
	count1 := build(edges1, color1)
	count2 := build(edges2, color2)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = count1[color1[i]] + max(count2[0], count2[1])
	}

	return result
}

func dfs(node int, parent int, depth int, color []int, children [][]int) int {
	result := 1 - depth % 2
	color[node] = depth % 2

	for _, child := range children[node] {
		if child == parent {
			continue
		}

		result += dfs(child, node, depth + 1, color, children)
	}

	return result
}

func build(edges [][]int, color []int) []int {
	n := len(edges) + 1
	children := make([][]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		children[u] = append(children[u], v)
		children[v] = append(children[v], u)
	}

	result := dfs(0, -1, 0, color, children)
	return []int{result, n - result}
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: [8,7,7,8,8]
	// edges1 := [][]int{{0,1},{0,2},{2,3},{3,4}}
	// edges2 := [][]int{{0,1},{0,2},{0,3},{2,7},{1,4},{4,5},{4,6}}

	// result: [3,6,6,6,6]
	edges1 := [][]int{{0,1},{0,2},{0,3},{0,4}}
	edges2 := [][]int{{0,1},{1,2},{2,3}}

	// result: []
	// edges1 := [][]int{}
	// edges2 := [][]int{}

	result := maxTargetNodes(edges1, edges2)
	fmt.Printf("result = %v\n", result)
}

