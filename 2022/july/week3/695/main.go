package main
import (
	"fmt"
)

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	visited := make([][]int, len(grid))
	for i, _ := range grid {
		visited[i] = make([]int, len(grid[0]))
	}

	result := int(0)
	for y, _ := range grid {
		for x, _ := range grid[y] {
			if visited[y][x] == 1 {
				continue
			}

			count := traversal(y, x, grid, visited)

			if result < count {
				result = count
			}
		}
	}

	return result
}

func traversal(y int, x int, grid [][]int, visited [][]int) int {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	if visited[y][x] != 0 {
		return 0
	}

	visited[y][x] = 1
	if grid[y][x] == 0 {
		return 0
	}

	count := int(1)

	for _, dir := range dirs {
		nextY := y + dir[0]
		nextX := x + dir[1]
		if 0 <= nextY && nextY < len(grid) && 0 <= nextX && nextX < len(grid[0]) {
			count += traversal(nextY, nextX, grid, visited)
		}
	}

	return count
}

func main() {
	// result: 6
	// grid := [][]int{
	// 	{0,0,1,0,0,0,0,1,0,0,0,0,0},
	// 	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	// 	{0,1,1,0,1,0,0,0,0,0,0,0,0},
	// 	{0,1,0,0,1,1,0,0,1,0,1,0,0},
	// 	{0,1,0,0,1,1,0,0,1,1,1,0,0},
	// 	{0,0,0,0,0,0,0,0,0,0,1,0,0},
	// 	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	// 	{0,0,0,0,0,0,0,1,1,0,0,0,0},
	// }

	// result: 0
	// grid := [][]int{{0,0,0,0,0,0,0,0}}

	// result: 6
	grid := [][]int{
		{1,0,1,0,0,0,0,1,0,0,0,0,0},
		{1,0,0,0,0,0,0,1,1,1,0,0,0},
		{1,1,1,0,1,0,0,0,0,1,1,0,0},
		{1,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,1,1},
		{0,0,0,0,0,0,0,1,0,0,1,0,0},
		{1,0,0,0,0,0,0,1,1,1,0,0,0},
		{1,1,1,1,1,1,0,1,1,0,0,0,0},
	}

	// result: 
	// grid := [][]int{}

	result := maxAreaOfIsland(grid)
	fmt.Printf("result = %v\n", result)
}

