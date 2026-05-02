package main

import (
	"fmt"
	"math"
)

func maxPathScore(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])

	minInt := math.MinInt32

	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)

		for j := range dp[i] {
			dp[i][j] = make([]int, k + 1)

			for c := range dp[i][j] {
				dp[i][j][c] = minInt
			}
		}
	}

	dp[0][0][0] = 0

	for i := range m {
		for j := range n {
			for c := range k + 1 {
				if dp[i][j][c] == minInt {
					continue
				}

				if i + 1 < m {
					val := grid[i + 1][j]
					cost := int(0)
					if val != 0 {
						cost = 1
					}

					if c + cost <= k {
						if dp[i + 1][j][c + cost] < dp[i][j][c] + val {
							dp[i + 1][j][c + cost] = dp[i][j][c] + val
						}
					}
				}

				if j + 1 < n {
					val := grid[i][j + 1]
					cost := int(0)
					if val != 0 {
						cost = 1
					}

					if c + cost <= k {
						if dp[i][j + 1][c + cost] < dp[i][j][c] + val {
							dp[i][j + 1][c + cost] = dp[i][j][c] + val
						}
					}
				}
			}
		}
	}

	result := minInt
	for c := range k + 1 {
		if dp[m - 1][n - 1][c] > result {
			result = dp[m - 1][n - 1][c]
		}
	}

	if result < 0 {
		return -1
	}
	return result
}

func main() {
	// result: 2
	// grid := [][]int{{0, 1},{2, 0}}
	// k := int(1)

	// result: -1
	grid := [][]int{{0, 1},{1, 2}}
	k := int(1)

	// result: 
	// grid := [][]int{}
	// k := int(0)

	result := maxPathScore(grid, k)
	fmt.Printf("result = %v\n", result)
}

