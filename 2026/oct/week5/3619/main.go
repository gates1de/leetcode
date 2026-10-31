package main

import (
	"fmt"
)

func countIslands(grid [][]int, k int) int {
	if len(grid) == 0 || len(grid[0]) == 0 || k <= 0 {
		return 0
	}

	cols := len(grid[0])
	stack := make([]int, 0)
	result := 0

	for row, cells := range grid {
		for col, value := range cells {
			if value <= 0 {
				continue
			}

			grid[row][col] = -value
			stack = append(stack, row*cols+col)
			sum := int64(0)

			for len(stack) > 0 {
				last := len(stack) - 1
				position := stack[last]
				stack = stack[:last]

				currentRow := position / cols
				currentCol := position % cols
				sum -= int64(grid[currentRow][currentCol])

				if currentRow > 0 && grid[currentRow-1][currentCol] > 0 {
					grid[currentRow-1][currentCol] = -grid[currentRow-1][currentCol]
					stack = append(stack, position-cols)
				}
				if currentRow+1 < len(grid) && grid[currentRow+1][currentCol] > 0 {
					grid[currentRow+1][currentCol] = -grid[currentRow+1][currentCol]
					stack = append(stack, position+cols)
				}
				if currentCol > 0 && grid[currentRow][currentCol-1] > 0 {
					grid[currentRow][currentCol-1] = -grid[currentRow][currentCol-1]
					stack = append(stack, position-1)
				}
				if currentCol+1 < cols && grid[currentRow][currentCol+1] > 0 {
					grid[currentRow][currentCol+1] = -grid[currentRow][currentCol+1]
					stack = append(stack, position+1)
				}
			}

			if sum%int64(k) == 0 {
				result++
			}
		}
	}

	return result
}

func main() {
	// result: 2
	// grid := [][]int{{0,2,1,0,0},{0,5,0,0,5},{0,0,1,0,0},{0,1,4,7,0},{0,2,0,0,8}}
	// k := int(5)

	// result: 6
	grid := [][]int{{3, 0, 3, 0}, {0, 3, 0, 3}, {3, 0, 3, 0}}
	k := int(3)

	// result: 0
	// grid := [][]int{}
	// k := int(0)

	result := countIslands(grid, k)
	fmt.Printf("result = %v\n", result)
}
