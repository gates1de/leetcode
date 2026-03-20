package main

import (
	"fmt"
	"math"
	"sort"
)

func minAbsDiff(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	result := make([][]int, m - k + 1)
	for i := range result {
		result[i] = make([]int, n - k + 1)
	}

	for i := 0; i + k <= m; i++ {
		for j := 0; j + k <= n; j++ {
			kGrid := make([]int, 0, k * k)
			for x := i; x < i + k; x++ {
				for y := j; y < j + k; y++ {
					kGrid = append(kGrid, grid[x][y])
				}
			}

			kMin := math.MaxInt32
			sort.Ints(kGrid)
			for t := 1; t < len(kGrid); t++ {
				if kGrid[t] == kGrid[t-1] {
					continue
				}

				kMin = min(kMin, kGrid[t] - kGrid[t - 1])
			}

			if kMin != math.MaxInt32 {
				result[i][j] = kMin
			}
		}
	}

	return result
}

func main() {
	// result: [[2]]
	// grid := [][]int{{1,8},{3,-2}}
	// k := int(2)

	// result: [[0,0]]
	// grid := [][]int{{3,-1}}
	// k := int(1)

	// result: [[1,2]]
	grid := [][]int{{1,-2,3},{2,3,5}}
	k := int(2)

	// result: []
	// grid := [][]int{}
	// k := int(0)

	result := minAbsDiff(grid, k)
	fmt.Printf("result = %v\n", result)
}

