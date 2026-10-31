package main

import (
	"fmt"
)

func uniquePaths(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}

	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	size := rows * cols
	right := make([]int, size)
	down := make([]int, size)
	for i := range right {
		right[i] = -1
		down[i] = -1
	}

	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			position := row*cols + col
			if col+1 < cols {
				next := position + 1
				if grid[row][col+1] == 0 {
					right[position] = next
				} else {
					right[position] = down[next]
				}
			}
			if row+1 < rows {
				next := position + cols
				if grid[row+1][col] == 0 {
					down[position] = next
				} else {
					down[position] = right[next]
				}
			}
		}
	}

	const mod = 1000000007
	ways := make([]int, size)
	ways[0] = 1
	for row, cells := range grid {
		for col, value := range cells {
			if value == 1 {
				continue
			}

			position := row*cols + col
			current := ways[position]
			if current == 0 {
				continue
			}

			if next := right[position]; next != -1 {
				ways[next] += current
				if ways[next] >= mod {
					ways[next] -= mod
				}
			}
			if next := down[position]; next != -1 {
				ways[next] += current
				if ways[next] >= mod {
					ways[next] -= mod
				}
			}
		}
	}

	return ways[size-1]
}

func main() {
	// result: 5
	// grid := [][]int{{0,1,0},{0,0,1},{1,0,0}}

	// result: 2
	// grid := [][]int{{0,0},{0,0}}

	// result: 1
	grid := [][]int{{0, 1, 1}, {1, 1, 0}}

	// result: 0
	// grid := [][]int{}

	result := uniquePaths(grid)
	fmt.Printf("result = %v\n", result)
}
