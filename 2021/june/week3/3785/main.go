package main
import (
	"fmt"
)

func swimInWater(grid [][]int) int {
    visited := make([][]int, len(grid))
    for i, row := range grid {
        visited[i] = make([]int, len(row))
        for j, _ := range visited[i] {
            visited[i][j] = -1
        }
    }

    traversal(0, 0, grid[0][0], visited, grid)
    return visited[len(grid) - 1][len(grid[0]) - 1]
}

func traversal(y int, x int, t int, visited [][]int, grid [][]int) {
    visited[y][x] = t
    if y == len(grid) - 1 && x == len(grid[0]) - 1 {
        return
    }

    maxY := len(grid)
    maxX := len(grid[0])
    // up
    if y - 1 >= 0 {
        upT := max(t, grid[y - 1][x])
        if visited[y - 1][x] == -1 || upT < visited[y - 1][x] {
            traversal(y - 1, x, upT, visited, grid)
        }
    }
    // right
    if x + 1 < maxX {
        rightT := max(t, grid[y][x + 1])
        if visited[y][x + 1] == -1 || rightT < visited[y][x + 1] {
            traversal(y, x + 1, rightT, visited, grid)
        }
    }
    // down
    if y + 1 < maxY {
        downT := max(t, grid[y + 1][x])
        if visited[y + 1][x] == -1 || downT < visited[y + 1][x] {
            traversal(y + 1, x, downT, visited, grid)
        }
    }
    // left
    if x - 1 >= 0 {
        leftT := max(t, grid[y][x - 1])
        if visited[y][x - 1] == -1 || leftT < visited[y][x - 1] {
            traversal(y, x - 1, leftT, visited, grid)
        }
    }
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func main() {
	// result: 3
	// grid := [][]int{{0, 2}, {1, 3}}

	// result: 
	grid := [][]int{
		{0,1,2,3,4},
		{24,23,22,21,5},
		{12,13,14,15,16},
		{11,17,18,19,20},
		{10,9,8,7,6},
	}

	// result: 
	// grid := [][]int{}

	result := swimInWater(grid)
	fmt.Printf("result = %v\n", result)
}

