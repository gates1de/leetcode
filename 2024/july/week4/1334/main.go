package main
import (
	"fmt"
	"math"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	distMatrix := make([][]int, n)
	for i, _ := range distMatrix {
		distMatrix[i] = make([]int, n)
		for j, _ := range distMatrix[i] {
			distMatrix[i][j] = math.MaxInt32
			distMatrix[i][i] = 0
		}
	}

	for _, edge := range edges {
		start := edge[0]
		end := edge[1]
		weight := edge[2]

		distMatrix[start][end] = weight
		distMatrix[end][start] = weight
	}
	
	floyd(n, distMatrix)
	return helper(n, distMatrix, distanceThreshold)
}

func floyd(n int, distMatrix [][]int) {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				distMatrix[i][j] = min(
					distMatrix[i][j],
					distMatrix[i][k] + distMatrix[k][j],
				)
			}
		}
	}
}

func helper(n int, distMatrix [][]int, threshold int) int {
	reachable := int(-1)
	fewestCount := n

	for i := 0; i < n; i++ {
		count := int(0)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			if distMatrix[i][j] <= threshold {
				count++
			}
		}

		if count <= fewestCount {
			fewestCount = count
			reachable = i
		}
	}

	return reachable
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// n := int(4)
	// edges := [][]int{{0,1,3},{1,2,1},{1,3,4},{2,3,1}}
	// distanceThreshold := int(4)

	// result: 0
	n := int(5)
	edges := [][]int{{0,1,2},{0,4,8},{1,2,3},{1,4,2},{2,3,1},{3,4,1}}
	distanceThreshold := int(2)

	// result: 
	// n := int(0)
	// edges := [][]int{}
	// distanceThreshold := int(0)

	result := findTheCity(n, edges, distanceThreshold)
	fmt.Printf("result = %v\n", result)
}

