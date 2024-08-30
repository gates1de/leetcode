package main
import (
	"fmt"
)

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	currentShortestDistance := runDijkstra(
		edges,
		n,
		source,
		destination,
	)

	if currentShortestDistance < target {
		return [][]int{}
	}

	matchesTarget := currentShortestDistance == target

	for _, edge := range edges {
		if edge[2] > 0 {
			continue
		}

		edge[2] = 1
		if matchesTarget {
			edge[2] = int(2e9)
		}

		if !matchesTarget {
			newDistance := runDijkstra(edges, n, source, destination)

			if newDistance <= target {
				matchesTarget = true
				edge[2] += target - newDistance
			}
		}
	}

	if matchesTarget {
		return edges
	}

	return [][]int{}
}

func runDijkstra(edges [][]int, n int, source int, destination int) int {
	adjMatrix := make([][]int, n)
	minDistance := make([]int, n)
	visited := make([]bool, n)

	for i, _ := range minDistance {
		minDistance[i] = int(2e9)
	}

	for i, _ := range adjMatrix {
		adjMatrix[i] = make([]int, n)

		for j, _ := range adjMatrix[i] {
			adjMatrix[i][j] = int(2e9)
		}
	}

	minDistance[source] = 0

	for _, edge := range edges {
		if edge[2] != -1 {
			adjMatrix[edge[0]][edge[1]] = edge[2]
			adjMatrix[edge[1]][edge[0]] = edge[2]
		}
	}

	for i := 0; i < n; i++ {
		nearestUnvisitedNode := int(-1)

		for j := 0; j < n; j++ {
			if !visited[j] && (nearestUnvisitedNode == -1 || minDistance[j] < minDistance[nearestUnvisitedNode]) {
				nearestUnvisitedNode = j
			}
		}

		visited[nearestUnvisitedNode] = true

		for v := 0; v < n; v++ {
			minDistance[v] = min(minDistance[v], minDistance[nearestUnvisitedNode] + adjMatrix[nearestUnvisitedNode][v])
		}
	}

	return minDistance[destination]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// result: [[4,1,1],[2,0,1],[0,3,3],[4,3,1]]
	// n := int(5)
	// edges := [][]int{{4,1,-1},{2,0,-1},{0,3,-1},{4,3,-1}}
	// source := int(0)
	// destination := int(1)
	// target := int(5)

	// result: []
	// n := int(3)
	// edges := [][]int{{0,1,-1},{0,2,5}}
	// source := int(0)
	// destination := int(2)
	// target := int(6)

	// result: [[1,0,4],[1,2,3],[2,3,5],[0,3,1]]
	n := int(4)
	edges := [][]int{{1,0,4},{1,2,3},{2,3,5},{0,3,-1}}
	source := int(0)
	destination := int(2)
	target := int(6)

	// result: []
	// n := int(0)
	// edges := [][]int{}
	// source := int(0)
	// destination := int(0)
	// target := int(0)

	result := modifiedGraphEdges(n, edges, source, destination, target)
	fmt.Printf("result = %v\n", result)
}

