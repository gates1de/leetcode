package main
import (
	"fmt"
)

func countPairs(n int, edges [][]int) int64 {
	adjacents := make(map[int][]int)
	for _, edge := range edges {
		if adjacents[edge[0]] == nil {
			adjacents[edge[0]] = make([]int, 0)
		}
		if adjacents[edge[1]] == nil {
			adjacents[edge[1]] = make([]int, 0)
		}

		adjacents[edge[0]] = append(adjacents[edge[0]], edge[1])
		adjacents[edge[1]] = append(adjacents[edge[1]], edge[0])
	}

	result := int64(0)
	sizeOfComponent := int64(0)
	remainingNodes := int64(n)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			sizeOfComponent = dfs(i, adjacents, &visited)
			result += sizeOfComponent * (remainingNodes - sizeOfComponent)
			remainingNodes -= sizeOfComponent
		}
	}
	return result
}

func dfs(node int, adjacents map[int][]int, visited *[]bool) int64 {
	count := int64(1)
	(*visited)[node] = true

	if adjacents[node] == nil {
		return count
	}

	for _, neighbor := range adjacents[node] {
		if !(*visited)[neighbor] {
			count += dfs(neighbor, adjacents, visited)
		}
	}

	return count
}

func main() {
	// result: 0
	// n := int(3)
	// edges := [][]int{{0,1},{0,2},{1,2}}

	// result: 14
	n := int(7)
	edges := [][]int{{0,2},{0,5},{2,4},{1,6},{5,4}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := countPairs(n, edges)
	fmt.Printf("result = %v\n", result)
}

