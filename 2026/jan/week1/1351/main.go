package main

import (
	"fmt"
)

func countNegatives(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	result := int(0)

	i := m - 1
	j := int(0)
	for i >= 0 && j < n {
		if grid[i][j] < 0 {
			result += n - j
			i--
		} else {
			j++
		}
	}

	return result
}

func main() {
	// result: 8
	// grid := [][]int{{4,3,2,-1},{3,2,1,-1},{1,1,-1,-2},{-1,-1,-2,-3}}

	// result: 0
	grid := [][]int{{3,2},{1,0}}

	// result: 
	// grid := [][]int{}

	result := countNegatives(grid)
	fmt.Printf("result = %v\n", result)
}

