package main
import (
	"fmt"
)

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n, 200000)
    for _, e := range edges {
        graph[e[0]] = append(graph[e[0]], e[1])
        graph[e[1]] = append(graph[e[1]], e[0])
    }

    visited := make(map[int]bool)
    queue := []int{source}
    visited[source] = true
    for len(queue) > 0 {
        cursor := queue[0]
        queue = queue[1:]
        if cursor == destination {
			return true
		}

        for _, node := range graph[cursor] {
            if !visited[node] {
                visited[node] = true
                queue = append(queue, node)
            }
        }
    }

    return false
}

func main() {
	// result: true
	// n := int(3)
	// edges := [][]int{{0,1},{1,2},{2,0}}
	// source := int(0)
	// destination := int(2)

	// result: false
	n := int(6)
	edges := [][]int{{0,1},{0,2},{3,5},{5,4},{4,3}}
	source := int(0)
	destination := int(5)

	// result: 
	// n := int(0)
	// edges := [][]int{}
	// source := int(0)
	// destination := int(0)

	result := validPath(n, edges, source, destination)
	fmt.Printf("result = %v\n", result)
}

