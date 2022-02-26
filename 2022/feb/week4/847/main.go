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

// Time Limit Exceeded
func ngSolution(graph [][]int) int {
	result := math.MaxInt32
	for i, _ := range graph {
		// fmt.Printf("----------------- start: %v -----------------\n", i)
		remain := map[int]bool{}
		for i, _ := range graph {
			remain[i] = true
		}
		helper(i, 0, graph, remain, &result)
	}
	return result
}

func helper(index int, path int, graph [][]int, remain map[int]bool, result *int) {
	if path >= *result {
		return
	}

	// fmt.Printf("index = %v, path = %v, graph = %v, remain = %v\n", index, path, graph[index], remain)
	nodes := graph[index]
	delete(remain, index)
	if len(remain) == 0 {
		if path < *result {
			*result = path
		}
		return
	}

	for i, node := range nodes {
		copyGraph := make([][]int, len(graph))
		for j, g := range graph {
			s := make([]int, len(g))
			copy(s, g)
			copyGraph[j] = s
		}
		if len(copyGraph[index]) >= 1 {
			copyGraph[index] = append(copyGraph[index][:i], copyGraph[index][i + 1:]...)
		} else {
			copyGraph[index] = []int{}
		}
		copyRemain := map[int]bool{}
		for i, _ := range remain {
			copyRemain[i] = true
		}
		helper(node, path + 1, copyGraph, copyRemain, result)
	}
}

func main() {
	// result: 4
	// graph := [][]int{{1,2,3},{0},{0},{0}}

	// result: 4
	// graph := [][]int{{1},{0,2,4},{1,3,4},{2},{1,2}}

	// result: 0
	// graph := [][]int{{}}

	// result: 9
	// graph := [][]int{{1},{0,2,6},{1,3},{2},{5},{4,6},{1,5,7},{6}}

	// result: 8
	// graph := [][]int{{1,2,3,5},{0,4},{0,6},{0},{1},{0},{2}}

	// result: 14
	// graph := [][]int{{2,3,5,7},{2,3,7},{0,1},{0,1},{7},{0},{10},{9,10,0,1,4},{9},{7,8},{7,6}}

	// result: 11
	graph := [][]int{{1,2,3,4,5,6,7,8,9},{0,2,3,4,5,6,7,8,9},{0,1,3,4,5,6,7,8,9},{0,1,2,4,5,6,7,8,9},{0,1,2,3,5,6,7,8,9},{0,1,2,3,4,6,7,8,9},{0,1,2,3,4,5,7,8,9},{0,1,2,3,4,5,6,8,9},{0,1,2,3,4,5,6,7,9,10},{0,1,2,3,4,5,6,7,8,11},{8},{9}}

	// result: 
	// graph := [][]int{}

	result := shortestPathLength(graph)
	fmt.Printf("result = %v\n", result)
}

