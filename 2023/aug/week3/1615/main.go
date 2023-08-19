package main
import (
	"fmt"
)

func maximalNetworkRank(n int, roads [][]int) int {
	result := int(0)
	adj := make(map[int]map[int]bool)
	for _, road := range roads {
		if adj[road[0]] == nil {
			adj[road[0]] = make(map[int]bool)
		}
		if adj[road[1]] == nil {
			adj[road[1]] = make(map[int]bool)
		}
		adj[road[0]][road[1]] = true
		adj[road[1]][road[0]] = true
	}

	for node1 := 0; node1 < n; node1++ {
		for node2 := node1 + 1; node2 < n; node2++ {
			currentRank := int(0)
			if adj[node2] != nil {
				currentRank += len(adj[node2])
			}

			if adj[node1] != nil {
				currentRank += len(adj[node1])

				for key, _ := range adj[node1] {
					if key == node2 {
						currentRank--
					}
				}
			}

			result = max(result, currentRank)
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// n := int(4)
	// roads := [][]int{{0,1},{0,3},{1,2},{1,3}}

	// result: 5
	// n := int(5)
	// roads := [][]int{{0,1},{0,3},{1,2},{1,3},{2,3},{2,4}}

	// result: 5
	n := int(8)
	roads := [][]int{{0,1},{1,2},{2,3},{2,4},{5,6},{5,7}}

	// result: 
	// n := int(0)
	// roads := [][]int{}

	result := maximalNetworkRank(n, roads)
	fmt.Printf("result = %v\n", result)
}

