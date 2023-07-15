package main
import (
	"fmt"
)

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	indegree := make([]int, n)
	adjacents := make([][]int, n)
	for i, _ := range adjacents {
		adjacents[i] = make([]int, 0)
	}

	for i, _ := range graph {
		for _, node := range graph[i] {
			adjacents[node] = append(adjacents[node], i)
			indegree[i]++
		}
	}

	queue := make([]int, 0)
	for i, num := range indegree {
		if num == 0 {
			queue = append(queue, i)
		}
	}

	safe := make([]bool, n)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		safe[node] = true

		for _, neighbor := range adjacents[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	result := make([]int, 0)
	for i, s := range safe {
		if s {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	// result: [2,4,5,6]
	// graph := [][]int{{1,2},{2,3},{5},{0},{5},{},{}}

	// result: [4]
	graph := [][]int{{1,2,3,4},{1,2},{3,4},{0,4},{}}

	// result: 
	// graph := [][]int{}

	result := eventualSafeNodes(graph)
	fmt.Printf("result = %v\n", result)
}

