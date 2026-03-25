package main

import (
	"fmt"
)

func canPartitionGrid(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])
	sum := int(0)
	for i := range m {
		for j := range n {
			sum += grid[i][j]
		}
	}

	if sum % 2 != 0 {
		return false
	}

	target := sum / 2
	sum = 0

	for i := range m - 1 {
		for j := range n {
			sum += grid[i][j]
		}

		if sum == target {
			return true
		}
	}

	sum = 0
	for j := range n - 1 {
		for i := range m {
			sum += grid[i][j]
		}

		if sum == target {
			return true
		}
	}

	return false
}

func main() {
	// result: true
	// grid := [][]int{{1,4},{2,3}}

	// result: false
	grid := [][]int{{1,3},{2,4}}

	// result: false
	// grid := [][]int{}

	result := canPartitionGrid(grid)
	fmt.Printf("result = %v\n", result)
}

