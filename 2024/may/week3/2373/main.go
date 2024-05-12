package main
import (
	"fmt"
)

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	result := make([][]int, n - 2)

	for i := 0; i < n - 2; i++ {
		result[i] = make([]int, n - 2)
		for j := 0; j < n - 2; j++ {
			result[i][j] = findMax(grid, i, j)
		}
	}

	return result
}

func findMax(grid [][]int, x int, y int) int {
	maxElement := int(0)
	for i := x; i < x + 3; i++ {
		for j := y; j < y + 3; j++ {
			maxElement = max(maxElement, grid[i][j])
		}
	}

	return maxElement
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: [[9,9],[8,6]]
	// grid := [][]int{{9,9,8,1},{5,6,2,6},{8,2,6,4},{6,2,2,2}}

	// result: [[2,2,2],[2,2,2],[2,2,2]]
	grid := [][]int{{1,1,1,1,1},{1,1,1,1,1},{1,1,2,1,1},{1,1,1,1,1},{1,1,1,1,1}}

	// result: []
	// grid := [][]int{}

	result := largestLocal(grid)
	fmt.Printf("result = %v\n", result)
}

