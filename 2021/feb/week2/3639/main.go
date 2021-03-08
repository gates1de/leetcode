package main
import (
	"fmt"
)

func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))
	for i, _ := range graph {
		if colors[i] == 0 && !dfs(i, graph, colors, 1) {
			return false
		}
	}
	return true
}

func dfs(i int, graph [][]int, colors []int, color int) bool {
	// fmt.Printf("node %v: color = %v, saved color = %v\n", i, color, colors)
	if colors[i] != 0 {
		return colors[i] == color
	}

	colors[i] = color
	for _, to := range graph[i] {
		if !dfs(to, graph, colors, -color) {
			return false
		}
	}
	return true
}

func main() {
	// graph := [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
	// graph := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	// graph := [][]int{{1}, {0, 3}, {3}, {1, 2}}
	graph := [][]int{{1, 2, 3}, {0, 3, 4}, {0, 3}, {0, 1, 2}, {1}}
	result := isBipartite(graph)
	fmt.Printf("result = %v\n", result)
}

