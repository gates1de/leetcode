package main
import (
	"fmt"
)

func longestCycle(edges []int) int {
	n := len(edges)
	result := int(-1)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
        if !visited[i]  {
			distances := make(map[int]int)
			distances[i] = 1
			dfs(i, edges, distances, &visited, &result)
		}
	}

	return result
}

func dfs(node int, edges []int, distances map[int]int, visited *[]bool, result *int) {
	(*visited)[node] = true
	neighbor := edges[node]
	if neighbor == -1 {
		return
	}

	if !(*visited)[neighbor] {
		distances[neighbor] = distances[node] + 1
		dfs(neighbor, edges, distances, visited, result)
	} else if _, isExists := distances[neighbor]; isExists {
		*result = max(*result, distances[node] - distances[neighbor] + 1)
	}
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// edges := []int{3,3,4,2,3}

	// result: -1
	edges := []int{2,-1,3,1}

	// result: 
	// edges := []int{}

	result := longestCycle(edges)
	fmt.Printf("result = %v\n", result)
}

