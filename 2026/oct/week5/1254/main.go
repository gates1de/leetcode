package main

import (
	"fmt"
)

func closedIsland(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}

	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	stack := make([]int, 0)
	for col := range grid[0] {
		if grid[0][col] == 0 {
			stack = flood(grid, 0, col, stack[:0])
		}
		if rows > 1 && grid[rows-1][col] == 0 {
			stack = flood(grid, rows-1, col, stack[:0])
		}
	}
	for row := 1; row+1 < rows; row++ {
		if grid[row][0] == 0 {
			stack = flood(grid, row, 0, stack[:0])
		}
		if cols > 1 && grid[row][cols-1] == 0 {
			stack = flood(grid, row, cols-1, stack[:0])
		}
	}

	result := 0
	for row := 1; row+1 < rows; row++ {
		for col := 1; col+1 < cols; col++ {
			if grid[row][col] == 0 {
				result++
				stack = flood(grid, row, col, stack[:0])
			}
		}
	}

	return result
}

func flood(grid [][]int, row int, col int, stack []int) []int {
	rows := len(grid)
	cols := len(grid[0])
	grid[row][col] = 1
	stack = append(stack, row*cols+col)

	for len(stack) > 0 {
		last := len(stack) - 1
		position := stack[last]
		stack = stack[:last]

		currentRow := position / cols
		currentCol := position % cols
		if currentRow > 0 && grid[currentRow-1][currentCol] == 0 {
			grid[currentRow-1][currentCol] = 1
			stack = append(stack, position-cols)
		}
		if currentRow+1 < rows && grid[currentRow+1][currentCol] == 0 {
			grid[currentRow+1][currentCol] = 1
			stack = append(stack, position+cols)
		}
		if currentCol > 0 && grid[currentRow][currentCol-1] == 0 {
			grid[currentRow][currentCol-1] = 1
			stack = append(stack, position-1)
		}
		if currentCol+1 < cols && grid[currentRow][currentCol+1] == 0 {
			grid[currentRow][currentCol+1] = 1
			stack = append(stack, position+1)
		}
	}

	return stack
}

func main() {
	// result: 2
	// grid := [][]int{
	// 	{1,1,1,1,1,1,1,0},
	// 	{1,0,0,0,0,1,1,0},
	// 	{1,0,1,0,1,1,1,0},
	// 	{1,0,0,0,0,1,0,1},
	// 	{1,1,1,1,1,1,1,0},
	// }

	// result: 1
	// grid := [][]int{
	// 	{0,0,1,0,0},
	// 	{0,1,0,1,0},
	// 	{0,1,1,1,0},
	// }

	// result: 2
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}

	// result:
	// grid := [][]int{}

	result := closedIsland(grid)
	fmt.Printf("result = %v\n", result)
}
