package main
import (
	"fmt"
)

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	nextMinCol := int(-1)
	nextMin1 := int(-1)
	nextMin2 := int(-1)

	for col := 0; col < n; col++ {
		if nextMin1 == -1 || grid[n - 1][col] <= nextMin1 {
			nextMin2 = nextMin1
			nextMin1 = grid[n - 1][col]
			nextMinCol = col
		} else if  nextMin2 == -1 || grid[n - 1][col] <= nextMin2 {
			nextMin2 = grid[n - 1][col]
		}
	}

	for row := n - 2; row >= 0; row-- {
		minCol := int(-1)
		min1 := int(-1)
		min2 := int(-1)

		for col := 0; col < n; col++ {
			val := int(0)
			if col != nextMinCol {
				val = grid[row][col] + nextMin1
			} else {
				val = grid[row][col] + nextMin2
			}

			if min1 == -1 || val <= min1 {
				min2 = min1
				min1 = val
				minCol = col
			} else if min2 == -1 || val <= min2 {
				min2 = val
			}
		}

		nextMinCol = minCol
		nextMin1 = min1
		nextMin2 = min2
	}

	return nextMin1
}

func main() {
	// result: 13
	// grid := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// result: 7
	grid := [][]int{{7}}

	// result: 
	// grid := [][]int{
	// }

	result := minFallingPathSum(grid)
	fmt.Printf("result = %v\n", result)
}

