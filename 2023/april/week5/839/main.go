package main
import (
	"fmt"
)

func numSimilarGroups(strs []string) int {
	n := len(strs)
	adjacents := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isSimilar(strs[i], strs[j]) {
				if adjacents[i] == nil {
					adjacents[i] = make([]int, 0)
				}
				if adjacents[j] == nil {
					adjacents[j] = make([]int, 0)
				}
				adjacents[i] = append(adjacents[i], j)
				adjacents[j] = append(adjacents[j], i)
			}
		}
	}

	visited := make([]bool, n)
	result := int(0)
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, adjacents, visited)
			result++
		}
	}

	return result
}

func dfs(node int, adjacents map[int][]int, visited []bool) {
	visited[node] = true
	if adjacents[node] == nil {
		return
	}

	for _, neighbor := range adjacents[node] {
		if !visited[neighbor] {
			visited[neighbor] = true
			dfs(neighbor, adjacents, visited)
		}
	}
}

func isSimilar(a, b string) bool {
	diff := int(0)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff == 0 || diff == 2
}

func main() {
	// result: 2
	// strs := []string{"tars","rats","arts","star"}

	// result: 1
	strs := []string{"omv","ovm"}

	// result: 
	// strs := []string{}

	result := numSimilarGroups(strs)
	fmt.Printf("result = %v\n", result)
}

