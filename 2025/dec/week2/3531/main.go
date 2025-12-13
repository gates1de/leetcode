package main

import (
	"fmt"
)

func countCoveredBuildings(n int, buildings [][]int) int {
	maxRow := make([]int, n + 1)
	minRow := make([]int, n + 1)
	maxCol := make([]int, n + 1)
	minCol := make([]int, n + 1)

	for i := range minRow {
		minRow[i] = n + 1
		minCol[i] = n + 1
	}

	for _, building := range buildings {
		x:= building[0]
		y := building[1]
		maxRow[y] = max(maxRow[y], x)
		minRow[y] = min(minRow[y], x)
		maxCol[x] = max(maxCol[x], y)
		minCol[x] = min(minCol[x], y)
	}

	result := int(0)
	for _, building := range buildings {
		x:= building[0]
		y := building[1]

		if x > minRow[y] && x < maxRow[y] &&
			y > minCol[x] && y < maxCol[x] {
			result++
		}
	}

	return result
}

func main() {
	// result: 1
	// n := int(3)
	// buildings := [][]int{{1,2},{2,2},{3,2},{2,1},{2,3}}

	// result: 0
	// n := int(3)
	// buildings := [][]int{{1,1},{1,2},{2,1},{2,2}}

	// result: 1
	n := int(5)
	buildings := [][]int{{1,3},{3,2},{3,3},{3,5},{5,3}}

	// result: 
	// n := int(0)
	// buildings := [][]int{}

	result := countCoveredBuildings(n, buildings)
	fmt.Printf("result = %v\n", result)
}

