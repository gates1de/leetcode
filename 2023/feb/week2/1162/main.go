package main
import (
	"fmt"
	"math"
)

func maxDistance(grid [][]int) int {
	n := len(grid)
	maxDistance := n * 2 + 1
	distances := make([][]int, n)
	for i, _ := range distances {
		distances[i] = make([]int, n)
		for j, _ := range distances[i] {
			distances[i][j] = maxDistance
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				distances[i][j] = 0
			} else {
				leftDist := maxDistance
				topDist := maxDistance
				if i > 0 {
					leftDist = distances[i - 1][j] + 1
				}
				if j > 0 {
					topDist = distances[i][j - 1] + 1
				}
				distances[i][j] = min(distances[i][j], min(leftDist, topDist))
			}
		}
	}

	result := math.MinInt32
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			rightDist := maxDistance
			bottomDist := maxDistance
			if i < n - 1 {
				rightDist = distances[i + 1][j] + 1
			}
			if j < n - 1 {
				bottomDist = distances[i][j + 1] + 1
			}
			distances[i][j] = min(distances[i][j], min(rightDist, bottomDist))

			result = max(result, distances[i][j])
		}
	}

	if result == 0 || result == maxDistance {
		return -1
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// grid := [][]int{{1,0,1},{0,0,0},{1,0,1}}

	// result: 4
	grid := [][]int{{1,0,0},{0,0,0},{0,0,0}}

	// result: 
	// grid := [][]int{}

	result := maxDistance(grid)
	fmt.Printf("result = %v\n", result)
}

