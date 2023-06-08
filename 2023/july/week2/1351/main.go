package main
import (
	"fmt"
)

func countNegatives(grid [][]int) int {
	result := int(0)

	for _, row := range grid {
		for _, val := range row {
			if val < 0 {
				result++
			}
		}
	}

	return result
}

func main() {
	// result: 8
	// grid := [][]int{{4,3,2,-1},{3,2,1,-1},{1,1,-1,-2},{-1,-1,-2,-3}}

	// result: 0
	grid := [][]int{{3,2},{1,0}}

	// result: 
	// grid := [][]int{}

	result := countNegatives(grid)
	fmt.Printf("result = %v\n", result)
}

