package main
import (
	"fmt"
)

func maximumDetonation(bombs [][]int) int {
	graph := make(map[int][]int)
	n := len(bombs)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			xi := bombs[i][0]
			yi := bombs[i][1]
			ri := bombs[i][2]
			xj := bombs[j][0]
			yj := bombs[j][1]

			if int64(ri * ri) >= (int64((xi - xj) * (xi - xj)) + int64((yi - yj) * (yi - yj))) {
				if graph[i] == nil {
					graph[i] = make([]int, 0, n)
				}
				graph[i] = append(graph[i], j)
			}
		}
	}

	result := int(0)
	for i := 0; i < n; i++ {
		result = max(result, dfs(i, graph))
	}

	return result
}

func dfs(i int, graph map[int][]int) int {
	stack := make([]int, 0)
	visited := make(map[int]bool)
	stack = append(stack, i)
	visited[i] = true
	for len(stack) > 0 {
		current := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if graph[current] == nil {
			continue
		}
		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				stack = append(stack, neighbor)
			}
		}
	}

	return len(visited)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// bombs := [][]int{{2,1,3},{6,1,4}}

	// result: 1
	// bombs := [][]int{{1,1,5},{10,10,5}}

	// result: 5
	bombs := [][]int{{1,2,3},{2,3,1},{3,4,2},{4,5,3},{5,6,4}}

	// result: 
	// bombs := [][]int{}

	result := maximumDetonation(bombs)
	fmt.Printf("result = %v\n", result)
}

