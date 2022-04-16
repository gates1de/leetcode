package main
import (
	"fmt"
)

func shiftGrid(grid [][]int, k int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return grid
	}

	for i := 0; i < k; i++ {
		shift(grid)
	}

	return grid
}

func shift(grid [][]int) {
	m := len(grid)
	n := len(grid[0])
	target := grid[0][0]
	for i, row := range grid {
		for j, _ := range row {
			if j == n - 1 {
				if i == m - 1 {
					grid[0][0] = target
					return
				}
				tmp := grid[i + 1][0]
				grid[i + 1][0] = target
				target = tmp
			} else {
				tmp := grid[i][j + 1]
				grid[i][j + 1] = target
				target = tmp
			}
		}
	}
}

func main() {
	// result: [[9,1,2],[3,4,5],[6,7,8]]
	// grid := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	// k := int(1)

	// result: [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
	// grid := [][]int{{3,8,1,9},{19,7,2,5},{4,6,11,10},{12,0,21,13}}
	// k := int(4)

	// result: [[1,2,3],[4,5,6],[7,8,9]]
	grid := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	k := int(9)

	// result: 
	// grid := [][]int{}
	// k := int(0)

	result := shiftGrid(grid, k)
	fmt.Printf("result = %v\n", result)
}

