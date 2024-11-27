package main
import (
	"fmt"
)

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	result := make([]int, 0)
	adjList := make([][]int, n)

	for i, _ := range adjList {
		adjList[i] = make([]int, 0)
	}

	for i := 0; i < n - 1; i++ {
		adjList[i] = append(adjList[i], i + 1)
	}

	for _, road := range queries {
		u := road[0]
		v := road[1]
		adjList[u] = append(adjList[u], v)
		result = append(result, findMinDistance(adjList, n))
	}

	return result
}

func findMinDistance(adjList [][]int, n int) int {
	dp := make([]int, n)
	dp[n - 1] = 0

	for node := n - 2; node >= 0; node-- {
		minDistance := n
		for _, neighbor := range adjList[node] {
			minDistance = min(minDistance, dp[neighbor] + 1)
		}
		dp[node] = minDistance
	}

	return dp[0]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: [3,2,1]
	// n := int(5)
	// queries := [][]int{{2,4},{0,2},{0,4}}

	// result: [1,1]
	n := int(4)
	queries := [][]int{{0,3},{0,2}}

	// result: []
	// n := int(0)
	// queries := [][]int{}

	result := shortestDistanceAfterQueries(n, queries)
	fmt.Printf("result = %v\n", result)
}

