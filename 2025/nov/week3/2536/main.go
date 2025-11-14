package main

import (
	"fmt"
)

func rangeAddQueries(n int, queries [][]int) [][]int {
	diffs := make([][]int, n + 1)
	for i := range diffs {
		diffs[i] = make([]int, n + 1)
	}

	for _, q := range queries {
		row1, col1, row2, col2 := q[0], q[1], q[2], q[3]
		diffs[row1][col1] += 1
		diffs[row2 + 1][col1] -= 1
		diffs[row1][col2 + 1] -= 1
		diffs[row2 + 1][col2 + 1] += 1
	}

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x1 := int(0)
			if i > 0 {
				x1 = result[i - 1][j]
			}

			x2 := int(0)
			if j > 0 {
				x2 = result[i][j - 1]
			}

			x3 := int(0)
			if i > 0 && j > 0 {
				x3 = result[i - 1][j - 1]
			}
			result[i][j] = diffs[i][j] + x1 + x2 - x3
		}
	}

	return result
}

func main() {
	// result: [[1,1,0],[1,2,1],[0,1,1]]
	// n := int(3)
	// queries := [][]int{{1,1,2,2},{0,0,1,1}}

	// result: [[1,1],[1,1]]
	n := int(2)
	queries := [][]int{{0,0,1,1}}

	// result: []
	// n := int(0)
	// queries := [][]int{}

	result := rangeAddQueries(n, queries)
	fmt.Printf("result = %v\n", result)
}

