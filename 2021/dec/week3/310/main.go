package main
import (
	"fmt"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 || len(edges) == 0 {
		return []int{0}
	}

	m := map[int][]int{}
	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		if m[a] == nil {
			m[a] = []int{}
		}
		m[a] = append(m[a], b)

		if m[b] == nil {
			m[b] = []int{}
		}
		m[b] = append(m[b], a)
	}

	minHeight := int(2 * 10000)
	heights := map[int]int{}
	for a, nodes := range m {
		visited := map[int]bool{}
		height := helper(a, nodes, 0, m, visited)
		// fmt.Printf("height = %v\n", height)
		heights[a] = height
		if minHeight > height {
			minHeight = height
		}
	}

	result := []int{}
	for a, height := range heights {
		if minHeight == height {
			result = append(result, a)
		}
	}

	return result
}

func helper(a int, nodes []int, height int, m map[int][]int, visited map[int]bool) int {
	// fmt.Printf("a = %v, nodes = %v, height = %v\n", a, nodes, height)
	if len(nodes) == 0 {
		return height
	}

	visited[a] = true
	maxHeight := height
	for _, node := range nodes {
		if visited[node] {
			continue
		}
		h := helper(node, m[node], height + 1, m, visited)
		if maxHeight < h {
			maxHeight = h
		}
	}

	return maxHeight
}

func main() {
	// result: [1]
	// n := int(4)
	// edges := [][]int{{1,0},{1,2},{1,3}}

	// result: [3,4]
	// n := int(6)
	// edges := [][]int{{3,0},{3,1},{3,2},{3,4},{5,4}}

	// result: [0]
	// n := int(1)
	// edges := [][]int{}

	// result: [0,1]
	n := int(2)
	edges := [][]int{{0,1}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := findMinHeightTrees(n, edges)
	fmt.Printf("result = %v\n", result)
}

