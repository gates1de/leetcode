package main

import (
	"fmt"
	"sort"
)

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)

	for i := 0; i < n; i++ {
		tmp := []int{}
		for j := 0; i + j < n; j++ {
			tmp = append(tmp, grid[i + j][j])
		}

		sort.Sort(sort.Reverse(sort.IntSlice(tmp)))

		for j := 0; i + j < n; j++ {
			grid[i + j][j] = tmp[j]
		}
	}

	for j := 1; j < n; j++ {
		tmp := []int{}
		for i := 0; j + i < n; i++ {
			tmp = append(tmp, grid[i][j + i])
		}

		sort.Ints(tmp)

		for i := 0; j+i < n; i++ {
			grid[i][j + i] = tmp[i]
		}
	}

	return grid
}

func main() {
	// result: [[8,2,3],[9,6,7],[4,5,1]]
	// grid := [][]int{}

	// result: [[2,1],{1,0}]
	// grid := [][]int{{0,1},{1,2}}

	// result: [[1]]
	grid := [][]int{{1}}

	// result: []
	// grid := [][]int{}

	result := sortMatrix(grid)
	fmt.Printf("result = %v\n", result)
}

