package main

import (
	"fmt"
)

func countServers(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	cols := len(grid[0])
	columnCounts := make([]int, cols)
	singleServerColumns := make([]int, 0, len(grid))
	result := 0

	for _, cells := range grid {
		serverCount := 0
		singleColumn := 0

		for col, value := range cells {
			if value == 1 {
				serverCount++
				singleColumn = col
				columnCounts[col]++
			}
		}

		if serverCount > 1 {
			result += serverCount
		} else if serverCount == 1 {
			singleServerColumns = append(singleServerColumns, singleColumn)
		}
	}

	for _, col := range singleServerColumns {
		if columnCounts[col] > 1 {
			result++
		}
	}

	return result
}

func main() {
	// result: 0
	// grid := [][]int{{1,0},{0,1}}

	// result: 3
	// grid := [][]int{{1,0},{1,1}}

	// result: 4
	grid := [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}

	// result:
	// grid := [][]int{}

	result := countServers(grid)
	fmt.Printf("result = %v\n", result)
}
