package main
import (
	"fmt"
)

func maxAreaOfIsland(grid [][]int) int {
	visited := make([][]int, len(grid))
	for i, row := range grid {
		visited[i] = make([]int, len(row))
	}
	result := int(0)
	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 || visited[i][j] == 1 {
				continue
			}
			areaOfIsland := traversal(i, j, grid, visited)
			fmt.Printf("[%v, %v]: area of island = %v\n", i, j, areaOfIsland)
			if result < areaOfIsland {
				result = areaOfIsland
			}
		}
	}
	return result
}

func traversal(i int, j int, grid [][]int, visited [][]int) int {
	result := int(0)

	queue := [][]int{{i, j}}

	for len(queue) > 0 {
		fmt.Printf("queue = %v\n", queue)
		coord := queue[len(queue) - 1]
		if len(queue) == 1 {
			queue = [][]int{}
		} else {
			queue = queue[:len(queue) - 1]
		}

		y := coord[0]
		x := coord[1]
		if visited[y][x] == 0 && grid[y][x] == 1 {
			result += 1
		}
		if visited[y][x] == 1 {
			continue
		}
		visited[y][x] = 1

		maxY := len(grid)
		maxX := len(grid[0])
		// up
		if y - 1 >= 0 && grid[y - 1][x] == 1 {
			queue = append(queue, []int{y - 1, x})
		}
		// right
		if x + 1 < maxX && grid[y][x + 1] == 1 {
			queue = append(queue, []int{y, x + 1})
		}
		// down
		if y + 1 < maxY && grid[y + 1][x] == 1 {
			queue = append(queue, []int{y + 1, x})
		}
		// left
		if x - 1 >= 0 && grid[y][x - 1] == 1 {
			queue = append(queue, []int{y, x - 1})
		}
	}
	return result
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

	// result: 11
	// grid := [][]int{
	// 	{0,1,0,0,0,1,0},
	// 	{0,1,1,0,0,1,0},
	// 	{0,0,0,0,0,1,0},
	// 	{0,1,0,0,0,1,0},
	// 	{0,1,1,1,0,1,0},
	// 	{0,1,0,0,0,1,0},
	// 	{0,1,0,0,0,1,0},
	// 	{0,1,0,0,0,1,0},
	// 	{0,1,0,0,0,1,0},
	// 	{0,1,0,0,0,1,0},
	// 	{0,1,0,0,0,1,0},
	// }

	// result: 13
	// grid := [][]int{
	// 	{0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0},
	// 	{0,0,0,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0},
	// 	{0,1,1,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0},
	// 	{0,1,0,0,1,1,0,0,1,0,1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0},
	// 	{0,1,0,0,1,1,0,0,1,1,1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0},
	// 	{0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0},
	// 	{0,0,0,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0},
	// 	{0,1,1,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0},
	// 	{0,1,0,0,1,1,0,0,1,0,1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0},
	// 	{0,1,0,0,1,1,0,0,1,1,1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0},
	// 	{0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0},
	// 	{0,0,0,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0},
	// 	{0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0},
	// }

	// result: 3
	grid := [][]int{{0,1},{1,1}}

	// result: 
	// grid := [][]int{
	// }

	result := maxAreaOfIsland(grid)
	fmt.Printf("result = %v\n", result)
}
