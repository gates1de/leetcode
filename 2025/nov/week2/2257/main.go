package main

import (
	"fmt"
)

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i, _ := range grid {
		grid[i] = make([]int, n)
	}

	for _, guard := range guards {
		grid[guard[0]][guard[1]] = 2
	}

	for _, wall := range walls {
		grid[wall[0]][wall[1]] = 3
	}

	for _, guard := range guards {
		dfs(guard[0] - 1, guard[1], grid, 'U')
		dfs(guard[0] + 1, guard[1], grid, 'D')
		dfs(guard[0], guard[1] - 1, grid, 'L')
		dfs(guard[0], guard[1] + 1, grid, 'R')
	}

	result := int(0)
	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				result++
			}
		}
	}

	return result
}

func dfs(row int, col int, grid [][]int, dir byte) {
	if row < 0 ||
		row >= len(grid) ||
		col < 0 ||
		col >= len(grid[0]) ||
		grid[row][col] == 2 ||
		grid[row][col] == 3 {
		return
	}

	grid[row][col] = 1
	if dir == 'U' {
		dfs(row - 1, col, grid, 'U')
	}
	if dir == 'D' {
		dfs(row + 1, col, grid, 'D')
	}
	if dir == 'L' {
		dfs(row, col - 1, grid, 'L')
	}
	if dir == 'R' {
		dfs(row, col + 1, grid, 'R')
	}
}

func main() {
	// result: 7
	// m := int(4)
	// n := int(6)
	// guards := [][]int{{0,0},{1,1},{2,3}}
	// walls := [][]int{{0,1},{2,2},{1,4}}

	// result: 4
	m := int(3)
	n := int(3)
	guards := [][]int{{1,1}}
	walls := [][]int{{0,1},{1,0},{2,1},{1,2}}

	// result: 
	// m := int(0)
	// n := int(0)
	// guards := [][]int{}
	// walls := [][]int{}

	result := countUnguarded(m, n, guards, walls)
	fmt.Printf("result = %v\n", result)
}

