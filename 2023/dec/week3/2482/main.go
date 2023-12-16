package main
import (
	"fmt"
)

func onesMinusZeros(grid [][]int) [][]int {
	m := len(grid)
	if m == 0 {
		return [][]int{}
	}

	n := len(grid[0])
	if n == 0 {
		return [][]int{}
	}

	rowCounts := make([]int, n)
	colCounts := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := grid[i][j]
			if val == 1 {
				rowCounts[j]++
			} else {
				rowCounts[j]--
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			val := grid[j][i]
			if val == 1 {
				colCounts[j]++
			} else {
				colCounts[j]--
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = rowCounts[j] + colCounts[i]
		}
	}
	return grid
}

func main() {
	// result: [[0,0,4],[0,0,4],[-2,-2,2]]
	// grid := [][]int{{0,1,1},{1,0,1},{0,0,1}}

	// result: [[5,5,5],[5,5,5]]
	grid := [][]int{{1,1,1},{1,1,1}}

	// result: 
	// grid := [][]int{}

	result := onesMinusZeros(grid)
	fmt.Printf("result = %v\n", result)
}

