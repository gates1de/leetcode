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
	// result: false
	// graph := [][]int{{1,2,3},{0,2},{0,1,3},{0,2}}

	// result: true
	graph := [][]int{{1,3},{0,2},{1,3},{0,2}}

	// result: 
	// graph := [][]int{}

	result := isBipartite(graph)
	fmt.Printf("result = %v\n", result)
}

