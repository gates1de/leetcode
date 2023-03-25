package main
import (
	"fmt"
	"math"
)

func minScore(n int, roads [][]int) int {
	adjacents := make(map[int][][]int)
	for _, road := range roads {
		if adjacents[road[0]] == nil {
			adjacents[road[0]] = make([][]int, 0)
		}
		if adjacents[road[1]] == nil {
			adjacents[road[1]] = make([][]int, 0)
		}

		adjacents[road[0]] = append(adjacents[road[0]], []int{road[1], road[2]})
		adjacents[road[1]] = append(adjacents[road[1]], []int{road[0], road[2]})
	}
        
	return bfs(n, adjacents)
}

func bfs(n int, adjacents map[int][][]int) int {
	visited := make([]bool, n + 1)
	queue := make([]int, 0)
	result := math.MaxInt32

	queue = append(queue, 1)
	visited[1] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if adjacents[node] == nil {
			continue
		}

		for _, edge := range adjacents[node] {
			result = min(result, edge[1])
			if !visited[edge[0]] {
				visited[edge[0]] = true
				queue = append(queue, edge[0])
			}
		}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// n := int(4)
	// roads := [][]int{{1,2,9},{2,3,6},{2,4,5},{1,4,7}}

	// result: 2
	n := int(4)
	roads := [][]int{{1,2,2},{1,3,4},{3,4,7}}

	// result: 
	// n := int(0)
	// roads := [][]int{}

	result := minScore(n, roads)
	fmt.Printf("result = %v\n", result)
}

