package main
import (
	"fmt"
)

func makeConnected(n int, connections [][]int) int {
    if len(connections) < n - 1 {
        return -1
    }

	adjacents := make(map[int][]int)
	for _, connection := range connections {
		if adjacents[connection[0]] == nil {
			adjacents[connection[0]] = make([]int, 0)
		}
		if adjacents[connection[1]] == nil {
			adjacents[connection[1]] = make([]int, 0)
		}

		adjacents[connection[0]] = append(adjacents[connection[0]], connection[1])
		adjacents[connection[1]] = append(adjacents[connection[1]], connection[0])
	}

	numOfConnected := int(0)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			numOfConnected++
			dfs(i, adjacents, &visited)
		}
	}

	return numOfConnected - 1
}

func dfs(node int, adjacents map[int][]int, visited *[]bool) {
	(*visited)[node] = true
	if adjacents[node] == nil {
		return
	}

	for _, edge := range adjacents[node] {
		if !(*visited)[edge] {
			(*visited)[edge] = true
			dfs(edge, adjacents, visited)
		}
	}
}

func main() {
	// result: 1
	// n := int(4)
	// connections := [][]int{{0,1},{0,2},{1,2}}

	// result: 2
	// n := int(6)
	// connections := [][]int{{0,1},{0,2},{0,3},{1,2},{1,3}}

	// result: -1
	n := int(6)
	connections := [][]int{{0,1},{0,2},{0,3},{1,2}}

	// result: 
	// n := int(0)
	// connections := [][]int{}

	result := makeConnected(n, connections)
	fmt.Printf("result = %v\n", result)
}

