package main

import (
	"fmt"
)

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	for i0, i1 := x, x + k - 1; i0 < i1; i0, i1 = i0 + 1, i1 - 1 {
		for j := y; j < y + k; j++ {
			grid[i0][j], grid[i1][j] = grid[i1][j], grid[i0][j]
		}
	}
	return grid
}

func main() {
	// result: [[1,2,3,4],[13,14,15,8],[9,10,11,12],[5,6,7,16]]
	// grid := [][]int{}{{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16}}
	// x := int(1)
	// y := int(0)
	// k := int(3)

	// result: [[3,4,4,2],[2,3,2,3]]
	grid := [][]int{{3,4,2,3},{2,3,4,2}}
	x := int(0)
	y := int(2)
	k := int(2)

	// result: []
	// grid := [][]int{}
	// x := int(0)
	// y := int(0)
	// k := int(0)

	result := reverseSubmatrix(grid, x, y, k)
	fmt.Printf("result = %v\n", result)
}

