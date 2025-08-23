package main

import (
	"fmt"
	"math"
)

func minimumSum(grid [][]int) int {
	rgrid := rotate(grid)
	return min(helper2(grid), helper2(rgrid))
}

func rotate(vec [][]int) [][]int {
	n := len(vec)
	m := len(vec[0])
	result := make([][]int, m)

	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result[m - j - 1][i] = vec[i][j]
		}
	}

	return result
}

func helper(grid [][]int, u, d, l, r int) int {
	minI := len(grid)
	maxI := int(0)
	minJ := len(grid[0])
	maxJ := int(0)

	for i := u; i <= d; i++ {
		for j := l; j <= r; j++ {
			if grid[i][j] == 1 {
				minI = min(minI, i)
				minJ = min(minJ, j)
				maxI = max(maxI, i)
				maxJ = max(maxJ, j)
			}
		}
	}

	if minI <= maxI {
		return (maxI - minI + 1) * (maxJ - minJ + 1)
	}
	return math.MaxInt32 / 3
}

func helper2(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	result := n * m

	for i := 0; i + 1 < n; i++ {
		for j := 0; j + 1 < m; j++ {
			result = min(result, helper(grid, 0, i, 0, m - 1) +
				helper(grid, i + 1, n - 1, 0, j) +
				helper(grid, i + 1, n - 1, j + 1, m - 1))

			result = min(result, helper(grid, 0, i, 0, j) +
				helper(grid, 0, i, j + 1, m - 1) +
				helper(grid, i + 1, n - 1, 0, m - 1))
		}
	}

	for i := 0; i + 2 < n; i++ {
		for j := i + 1; j + 1 < n; j++ {
			result = min(result, helper(grid, 0, i, 0, m - 1) +
				helper(grid, i + 1, j, 0, m - 1) +
				helper(grid, j + 1, n - 1, 0, m - 1))
		}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// grid := [][]int{{1,0,1},{1,1,1}}

	// result: 5
	grid := [][]int{{1,0,1,0},{0,1,0,1}}

	// result: 
	// grid := [][]int{}

	result := minimumSum(grid)
	fmt.Printf("result = %v\n", result)
}

