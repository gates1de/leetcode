package main
import (
	"fmt"
)

func findMissingAndRepeatedValues(grid [][]int) []int {
	sum := int(0)
	sqrSum := int(0)
	n := len(grid)
	total := n * n

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			sum += grid[row][col]
			sqrSum += grid[row][col] * grid[row][col]
		}
	}

	sumDiff := sum - (total * (total + 1)) / 2
	sqrDiff := sqrSum - (total * (total + 1) * (2 * total + 1)) / 6
	repeat := (sqrDiff / sumDiff + sumDiff) / 2
	missing := (sqrDiff / sumDiff - sumDiff) / 2

	return []int{repeat, missing}
}

func main() {
	// result: [2,4]
	// grid := [][]int{{1,3},{2,2}}

	// result: [9,5]
	grid := [][]int{{9,1,7},{8,9,2},{3,4,6}}

	// result: []
	// grid := [][]int{}

	result := findMissingAndRepeatedValues(grid)
	fmt.Printf("result = %v\n", result)
}

