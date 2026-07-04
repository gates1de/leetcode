package main

import (
	"fmt"
)

type Edge struct {
	To    int
	Score int
}

func minScore(n int, roads [][]int) int {
	adjacents := make([][]Edge, n + 1)
	for _, road := range roads {
		a := road[0]
		b := road[1]
		score := road[2]
		adjacents[a] = append(adjacents[a], Edge{To: b, Score: score})
		adjacents[b] = append(adjacents[b], Edge{To: a, Score: score})
	}

	return bfs(n, adjacents)
}

func bfs(n int, adjacents [][]Edge) int {
	visited := make([]bool, n + 1)
	queue := make([]int, 1, n + 1)
	queue[0] = 1
	visited[1] = true
	result := int(^uint(0) >> 1)

	for head := 0; head < len(queue); head++ {
		node := queue[head]

		for _, edge := range adjacents[node] {
			if edge.Score < result {
				result = edge.Score
			}
			if !visited[edge.To] {
				visited[edge.To] = true
				queue = append(queue, edge.To)
			}
		}
	}

	return result
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
