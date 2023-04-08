package main
import (
	"fmt"
)

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}

	result := int(0)
	for i, _ := range grid {
		for j, val := range grid[i] {
			if val == 0 || visited[i][j] {
				continue
			}

			result += bfs(i, j, m, n, grid, visited)
		}
	}

	return result
}

func bfs(x int, y int, m int, n int, grid [][]int, visited [][]bool) int {
	queue := make([][]int, 0)
	queue = append(queue, []int{x, y})
	visited[x][y] = true
	result := int(0)
	isClosed := true

	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for len(queue) > 0 {
		temp := queue[0]
		nextX := temp[0]
		nextY := temp[1]
		queue = queue[1:]
		result++

		for _, dir := range dirs {
			row := nextX + dir[0]
			col := nextY + dir[1]

			if row < 0 || row >= m || col < 0 || col >= n {
				isClosed = false
			} else if grid[row][col] == 1 && !visited[row][col] {
				queue = append(queue, []int{row, col})
				visited[row][col] = true
			}
		}
	}

	if !isClosed {
		return 0
	}
	return result
}

func main() {
	// result: 3
	// grid := [][]int{
	// 	{0,0,0,0},
	// 	{1,0,1,0},
	// 	{0,1,1,0},
	// 	{0,0,0,0},
	// }

	// result: 0
	// grid := [][]int{
	// 	{0,1,1,0},
	// 	{0,0,1,0},
	// 	{0,0,1,0},
	// 	{0,0,0,0},
	// }

	// result: 2
	grid := [][]int{
		{0,0,0,0},
		{0,1,0,0},
		{0,0,1,0},
		{0,0,0,0},
	}

	// result: 
	// grid := [][]int{}

	result := numEnclaves(grid)
	fmt.Printf("result = %v\n", result)
}

