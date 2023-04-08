package main
import (
	"fmt"
)

func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}

	result := int(0)
	for i, _ := range grid {
		for j, val := range grid[i] {
			if val == 0 && !visited[i][j] && bfs(i, j, m, n, grid, visited) {
				result++
			}
		}
	}

	return result
}

func bfs(x int, y int, m int, n int, grid [][]int, visited [][]bool) bool {
	queue := make([][]int, 0)
	queue = append(queue, []int{x, y})
	visited[x][y] = true
	isClosed := true

	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for len(queue) > 0 {
		temp := queue[0]
		nextX := temp[0]
		nextY := temp[1]
		queue = queue[1:]

		for _, dir := range dirs {
			row := nextX + dir[0]
			col := nextY + dir[1]

			if row < 0 || row >= m || col < 0 || col >= n {
				isClosed = false
			} else if grid[row][col] == 0 && !visited[row][col] {
				queue = append(queue, []int{row, col})
				visited[row][col] = true
			}
		}
	}

	return isClosed
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
		{1,1,1,1,1,1,1},
		{1,0,0,0,0,0,1},
		{1,0,1,1,1,0,1},
		{1,0,1,0,1,0,1},
		{1,0,1,1,1,0,1},
		{1,0,0,0,0,0,1},
		{1,1,1,1,1,1,1},
	}

	// result: 
	// grid := [][]int{}

	result := closedIsland(grid)
	fmt.Printf("result = %v\n", result)
}

