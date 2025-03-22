package main
import (
	"fmt"
)

func countCompleteComponents(n int, edges [][]int) int {
	graph := make([][]int, n)
	for i, _ := range graph {
		graph[i] = make([]int, 0)
	}

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	result := int(0)

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		component := make([]int, 0)
		queue := make([]int, 0)
		queue = append(queue, i)
		visited[i] = true

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			component = append(component, current)

			for _, neighbor := range graph[current] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
					visited[neighbor] = true
				}
			}
		}

		isComplete := true
		for _, node := range component {
			if len(graph[node]) != len(component) - 1 {
				isComplete = false
				break
			}
		}

		if isComplete {
			result++
		}
	}

	return result
}

func main() {
	// result: 3
	// n := int(6)
	// edges := [][]int{{0,1},{0,2},{1,2},{3,4}}

	// result: 1
	n := int(6)
	edges := [][]int{{0,1},{0,2},{1,2},{3,4},{3,5}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := countCompleteComponents(n, edges)
	fmt.Printf("result = %v\n", result)
}

