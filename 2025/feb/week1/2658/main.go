package main
import (
	"fmt"
)

func findMaxFish(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	result := int(0)

	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}

	for row, _ := range grid {
		for col, _ := range grid[row] {
			if grid[row][col] > 0 && !visited[row][col] {
				result = max(result, calculateFishes(row, col, grid, visited))
			}
		}
	}

	return result
}

func calculateFishes(row int, col int, grid [][]int, visited [][]bool) int {
	m := len(grid)
	n := len(grid[0])
	if row < 0 ||row >= m || col < 0 || col >= n ||
			grid[row][col] == 0 || visited[row][col] {
			return 0
	}

	visited[row][col] = true

	return grid[row][col] +
			calculateFishes(row, col + 1, grid, visited) +
			calculateFishes(row, col - 1, grid, visited) +
			calculateFishes(row + 1, col, grid, visited) +
			calculateFishes(row - 1, col, grid, visited)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 7
	// grid := [][]int{{0,2,1,0},{4,0,0,3},{1,0,0,4},{0,3,2,0}}

	// result: 1
	grid := [][]int{{1,0,0,0},{0,0,0,0},{0,0,0,0},{0,0,0,1}}

	// result: 
	// grid := [][]int{}

	result := findMaxFish(grid)
	fmt.Printf("result = %v\n", result)
}

