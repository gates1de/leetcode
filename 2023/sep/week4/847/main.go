package main
import (
	"fmt"
	"math"
)

func shortestPathLength(graph [][]int) int {
    n := len(graph)
    endMask := (1 << n) - 1
    cache := make([][]int, n + 1)
    for i, _ := range cache {
        cache[i] = make([]int, endMask + 1)
    }

    result := math.MaxInt32
    for node := 0; node < n; node++ {
        result = min(result, dp(node, endMask, graph, cache))
    }

    return result
}

func dp(node int, mask int, graph [][]int, cache [][]int) int {
    if cache[node][mask] != 0 {
        return cache[node][mask]
    }

    if (mask & (mask - 1)) == 0 {
        return 0
    }

    cache[node][mask] = math.MaxInt32 - 1
    for _, n := range graph[node] {
        if (mask & (1 << n)) != 0 {
            visited := dp(n, mask, graph, cache)
            notVisited := dp(n, mask ^ (1 << node), graph, cache)
            better := min(visited, notVisited)
            cache[node][mask] = min(cache[node][mask], 1 + better)
        }
    }

    return cache[node][mask]
}

func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

func main() {
	// result: 4
	// graph := [][]int{{1,2,3},{0},{0},{0}}

	// result: 4
	graph := [][]int{{1},{0,2,4},{1,3,4},{2},{1,2}}

	// result: 
	// graph := [][]int{}

	result := shortestPathLength(graph)
	fmt.Printf("result = %v\n", result)
}

