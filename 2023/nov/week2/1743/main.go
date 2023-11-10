package main
import (
	"fmt"
	"math"
)

func restoreArray(adjacentPairs [][]int) []int {
	graph := make(map[int][]int)

	for _, edge := range adjacentPairs {
		x := edge[0]
		y := edge[1]

		if graph[x] == nil {
			graph[x] = make([]int, 0)
		}
		if graph[y] == nil {
			graph[y] = make([]int, 0)
		}

		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	root := int(0)
	for key, val := range graph {
		if len(val) == 1 {
			root = key
			break
		}
	}

	result := make([]int, len(graph))
	dfs(root, math.MaxInt32, result, 0, graph)
	return result
}

func dfs(node int, prev int, result []int, i int, graph map[int][]int) {
	result[i] = node
	for _, neighbor := range graph[node] {
		if neighbor != prev {
			dfs(neighbor, node, result, i + 1, graph)
		}
	}
}

func main() {
	// result: [1,2,3,4]
	// adjacentPairs := [][]int{{2,1},{3,4},{3,2}}

	// result: [-2,4,1,-3]
	adjacentPairs := [][]int{{4,-2},{1,4},{-3,1}}

	// result: [100000,-100000]
	// adjacentPairs := [][]int{{100000,-100000}}

	// result: []
	// adjacentPairs := [][]int{}

	result := restoreArray(adjacentPairs)
	fmt.Printf("result = %v\n", result)
}

