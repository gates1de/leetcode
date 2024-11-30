package main
import (
	"fmt"
)

func findChampion(n int, edges [][]int) int {
	indegree := make([]int, n)

	for _, edge := range edges {
		indegree[edge[1]]++
	}

	result := int(-1)
	champCount := int(0)

	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			champCount++
			result = i
		}
	}

	if champCount > 1 {
		return -1
	}

	return result
}

func main() {
	// result: 0
	// n := int(3)
	// edges := [][]int{{0,1},{1,2}}

	// result: -1
	n := int(4)
	edges := [][]int{{0,2},{1,3},{1,2}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := findChampion(n, edges)
	fmt.Printf("result = %v\n", result)
}

