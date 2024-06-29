package main
import (
	"fmt"
)

func getAncestors(n int, edges [][]int) [][]int {
	adjacents := make([][]int, n)
	for i, _ := range adjacents {
		adjacents[i] = make([]int, 0)
	}

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		adjacents[to] = append(adjacents[to], from)
	}

	result := make([][]int, 0, n)
	for i, _ := range adjacents {
		ancestors := make([]int, 0)
		visited := make(map[int]bool)
		findChildren(i, adjacents, visited)

		for node := 0; node < n; node++ {
			if node == i {
				continue
			}

			if visited[node] {
				ancestors = append(ancestors, node)
			}
		}

		result = append(result, ancestors)
	}

	return result
}

func findChildren(currentNode int, adajacents [][]int, visited map[int]bool) {
	visited[currentNode] = true

	for _, neighbor := range adajacents[currentNode] {
		if !visited[neighbor] {
			findChildren(neighbor, adajacents, visited)
		}
	}
}

func main() {
	// result: [[],[],[],[0,1],[0,2],[0,1,3],[0,1,2,3,4],[0,1,2,3]]
	// n := int(8)
	// edges := [][]int{{0,3},{0,4},{1,3},{2,4},{2,7},{3,5},{3,6},{3,7},{4,6}}

	// result: [[],[0],[0,1],[0,1,2],[0,1,2,3]]
	n := int(5)
	edges := [][]int{{0,1},{0,2},{0,3},{0,4},{1,2},{1,3},{1,4},{2,3},{2,4},{3,4}}

	// result: []
	// n := int(0)
	// edges := [][]int{}

	result := getAncestors(n, edges)
	fmt.Printf("result = %v\n", result)
}

