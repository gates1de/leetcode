package main
import (
	"fmt"
)

func matrixScore(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	result := (1 << (n - 1)) * m

	for j := 1; j < n; j++ {
		countSameBits := int(0)
		for i := 0; i < m; i++ {
			if grid[i][j] == grid[i][0] {
				countSameBits++
			}
		}

		countSameBits = max(countSameBits, m - countSameBits)
		columnScore := (1 << (n - j - 1)) * countSameBits
		result += columnScore
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 39
	// grid := [][]int{{0,0,1,1},{1,0,1,0},{1,1,0,0}}

	// result: 1
	grid := [][]int{{0}}

	// result: 
	// grid := [][]int{}

	result := matrixScore(grid)
	fmt.Printf("result = %v\n", result)
}

