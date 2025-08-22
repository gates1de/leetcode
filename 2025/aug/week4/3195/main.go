package main

import (
	"fmt"
)

func minimumArea(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	minI := n
	maxI := int(0)
	minJ := m
	maxJ := int(0)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				minI = min(minI, i)
				maxI = max(maxI, i)
				minJ = min(minJ, j)
				maxJ = max(maxJ, j)
			}
		}
	}

	return (maxI - minI + 1) * (maxJ - minJ + 1)
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
	// result: 6
	// grid := [][]int{{0,1,0},{1,0,1}}

	// result: 1
	grid := [][]int{{1,0},{0,0}}

	// result: 
	// grid := [][]int{}

	result := minimumArea(grid)
	fmt.Printf("result = %v\n", result)
}

