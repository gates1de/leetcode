package main
import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i, _ := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	result := int(0)
	for y, _ := range grid {
		for x, _ := range grid[y] {
			if traversal(y, x, grid, visited) > 0 {
				result++
			}
		}
	}

	return result
}

func traversal(y int, x int, grid [][]byte, visited [][]bool) int {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	if visited[y][x] == true {
		return 0
	}

	visited[y][x] = true
	if grid[y][x] == byte('0') {
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
	// result: 1
	// grid := [][]byte{
	// 	{'1','1','1','1','0'},
	// 	{'1','1','0','1','0'},
	// 	{'1','1','0','0','0'},
	// 	{'0','0','0','0','0'},
	// }

	// result: 3
	// grid := [][]byte{
	// 	{'1','1','0','0','0'},
	// 	{'1','1','0','0','0'},
	// 	{'0','0','1','0','0'},
	// 	{'0','0','0','1','1'},
	// }

	// result: 0
	// grid := [][]byte{{'0'}}

	// result: 1
	grid := [][]byte{{'1'}}

	// result: 
	// grid := [][]byte{}

	result := numIslands(grid)
	fmt.Printf("result = %v\n", result)
}

