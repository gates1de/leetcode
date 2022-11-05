package main
import (
	"fmt"
)

func shortestPath(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])

	visited := make([][][]bool, m)
	for i, _ := range visited {
		visited[i] = make([][]bool, n)
		for j, _ := range visited[i] {
			visited[i][j] = make([]bool, k + 1)
		}
	}
	visited[0][0][0] = true

	result := int(0)
	queue := [][]int{{0, 0, 0}}
	for len(queue) > 0 {
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			status := queue[0]
			queue = queue[1:]

			row := status[0]
			column := status[1]
			currentK := status[2]

			if row == m - 1 && column == n - 1 {
				return result
			}

			for _, dir := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
				nextRow := dir[0] + row
				nextColumn := dir[1] + column
				nextK := currentK

				if nextRow < 0 || nextColumn < 0 || nextRow >= m || nextColumn >= n {
					continue
				}

				if grid[nextRow][nextColumn] == 1 {
					nextK++
				}

				if nextK <= k && !visited[nextRow][nextColumn][nextK] {
					visited[nextRow][nextColumn][nextK] = true
					queue = append(queue, []int{nextRow, nextColumn, nextK})
				}
			}
		}

		result++
	}

	return -1
}

// Wrong Answer
func ngSolution(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]int, m)
	for i, _ := range visited {
		visited[i] = make([]int, n)
		for j, _ := range visited[i] {
			visited[i][j] = 10000
		}
	}
	helper(0, 0, 0, grid, k, visited)

	if visited[m - 1][n - 1] == 10000 {
		return -1
	}
	return visited[m - 1][n - 1]
}

func helper(step int, x int, y int, grid [][]int, k int, visited [][]int) {
	m := len(grid)
	n := len(grid[0])
	if x == n - 1 && y == m - 1 {
		if step < visited[y][x] {
			visited[y][x] = step
		}
		return
	}

	if step > visited[y][x] {
		return
	}
	visited[y][x] = step

	for _, dir := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		nextY := y + dir[0]
		nextX := x + dir[1]
		if nextX < 0 || nextY < 0 || nextX >= n || nextY >= m {
			continue
		}

		nextPoint := grid[nextY][nextX]
		if nextPoint == 0 {
			helper(step + 1, nextX, nextY, grid, k, visited)
		} else {
			if k == 0 {
				continue
			}
			helper(step + 1, nextX, nextY, grid, k - 1, visited)
		}
	}
}

func main() {
	// result: 6
	// grid := [][]int{{0,0,0},{1,1,0},{0,0,0},{0,1,1},{0,0,0}}
	// k := int(1)

	// result: -1
	// grid := [][]int{{0,1,1},{1,1,1},{1,0,0}}
	// k := int(1)

	// result: 13
	grid := [][]int{
		{0,1,0,0,0,1,0,0},
		{0,1,0,1,0,1,0,1},
		{0,0,0,1,0,0,1,0},
	}
	k := int(1)

	// result: 
	// grid := [][]int{}
	// k := int(0)

	result := shortestPath(grid, k)
	fmt.Printf("result = %v\n", result)
}

