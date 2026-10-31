package main

import (
	"fmt"
)

func minCost(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	n := len(grid[0])
	prev := make([][1024]bool, n)
	for i, row := range grid {
		curr := make([][1024]bool, n)
		for j, value := range row {
			if i == 0 && j == 0 {
				curr[j][value] = true
				continue
			}
			for xor := range 1024 {
				if (i > 0 && prev[j][xor]) || (j > 0 && curr[j-1][xor]) {
					curr[j][xor^value] = true
				}
			}
		}
		prev = curr
	}

	for xor, reachable := range prev[n-1] {
		if reachable {
			return xor
		}
	}
	return 0
}

func main() {
	// result: 6
	// grid := [][]int{{1,2},{3,4}}

	// result: 9
	// grid := [][]int{{6,7},{5,8}}

	// result: 0
	grid := [][]int{{2,7,5}}

	// result: 0
	// grid := [][]int{}

	result := minCost(grid)
	fmt.Printf("result = %v\n", result)
}
