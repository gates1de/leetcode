package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func maxProductPath(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	maxgt := make([][]int64, m)
	minlt := make([][]int64, m)
	for i := range maxgt {
		maxgt[i] = make([]int64, n)
		minlt[i] = make([]int64, n)
	}

	maxgt[0][0] = int64(grid[0][0])
	minlt[0][0] = int64(grid[0][0])
	for i := 1; i < n; i++ {
		maxgt[0][i] = maxgt[0][i - 1] * int64(grid[0][i])
		minlt[0][i] = maxgt[0][i]
	}
	for i := 1; i < m; i++ {
		maxgt[i][0] = maxgt[i - 1][0] * int64(grid[i][0])
		minlt[i][0] = maxgt[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] >= 0 {
				maxPrev := max(maxgt[i][j - 1], maxgt[i - 1][j])
				minPrev := min(minlt[i][j - 1], minlt[i - 1][j])
				maxgt[i][j] = maxPrev * int64(grid[i][j])
				minlt[i][j] = minPrev * int64(grid[i][j])
			} else {
				maxPrev := max(maxgt[i][j - 1], maxgt[i -	1][j])
				minPrev := min(minlt[i][j - 1], minlt[i -	1][j])
				maxgt[i][j] = minPrev * int64(grid[i][j])
				minlt[i][j] = maxPrev * int64(grid[i][j])
			}
		}
	}

	if maxgt[m - 1][n - 1] < 0 {
		return -1
	}

	return int(maxgt[m - 1][n - 1] % int64(modulo))
}

func main() {
	// result: -1
	// grid := [][]int{{-1,-2,-3},{-2,-3,-3},{-3,-3,-2}}

	// result: 8
	// grid := [][]int{{1,-2,1},{1,-2,1},{3,-4,1}}

	// result: 0
	grid := [][]int{{1,3},{0,-4}}

	// result: 
	// grid := [][]int{}

	result := maxProductPath(grid)
	fmt.Printf("result = %v\n", result)
}

