package main

import (
	"fmt"
)

func countSubmatrices(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	for i := 1; i < n; i++ {
		for j := range m {
			grid[i][j] += grid[i - 1][j]
		}
	}
	for i := range n {
		for j := 1; j < m; j++ {
			grid[i][j] += grid[i][j - 1];
		}
	}

	result := int(0)
	for i := range n {
		for j := range m {
			if grid[i][j] <= k {
				result++;
			}
		}
	}

	return result
}

func main() {
	// result: 4
	// grid := [][]int{{7,6,3},{6,6,1}}
	// k := int(18)

	// result: 6
	grid := [][]int{{7,2,9},{1,5,0},{2,6,6}}
	k := int(20)

	// result: 
	// grid := [][]int{}
	// k := int(0)

	result := countSubmatrices(grid, k)
	fmt.Printf("result = %v\n", result)
}

