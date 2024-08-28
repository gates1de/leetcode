package main
import (
	"fmt"
)

var directions = [][]int{
	{ 0, 1 },
	{ 1, 0 },
	{ 0, -1 },
	{ -1, 0 },
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	totalRows := len(grid2)
	totalCols := len(grid2[0])

	visited := make([][]bool, totalRows)
	for i, _ := range visited {
		visited[i] = make([]bool, totalCols)
	}

	subIslandCounts := int(0)

	for x := 0; x < totalRows; x++ {
		for y := 0; y < totalCols; y++ {
			if !visited[x][y] && isCellLand(x, y, grid2) {
				visited[x][y] = true

				if isSubIsland(x, y, grid1, grid2, visited) {
					subIslandCounts += 1
				}
			}
		}
	}

	return subIslandCounts
}

func isCellLand(x int, y int, grid [][]int) bool {
	return grid[x][y] == 1
}

func isSubIsland(
	x int,
	y int,
	grid1 [][]int,
	grid2 [][]int,
	visited [][]bool,
) bool {
	totalRows := len(grid2)
	totalCols := len(grid2[0])
	result := true

	if !isCellLand(x, y, grid1) {
		result = false
	}

	for _, direction := range directions {
		nextX := x + direction[0]
		nextY := y + direction[1]
		if (nextX >= 0 && nextY >= 0 && nextX < totalRows && nextY < totalCols && !visited[nextX][nextY] && isCellLand(nextX, nextY, grid2)) {
			visited[nextX][nextY] = true
			nextCellIsPartOfSubIsland := isSubIsland(nextX, nextY, grid1, grid2, visited)
			result = result && nextCellIsPartOfSubIsland
		}
	}

	return result
}

func main() {
	// result: 3
	// grid1 := [][]int{{1,1,1,0,0},{0,1,1,1,1},{0,0,0,0,0},{1,0,0,0,0},{1,1,0,1,1}}
	// grid2 := [][]int{{1,1,1,0,0},{0,0,1,1,1},{0,1,0,0,0},{1,0,1,1,0},{0,1,0,1,0}}

	// result: 2
	grid1 := [][]int{{1,0,1,0,1},{1,1,1,1,1},{0,0,0,0,0},{1,1,1,1,1},{1,0,1,0,1}}
	grid2 := [][]int{{0,0,0,0,0},{1,1,1,1,1},{0,1,0,1,0},{0,1,0,1,0},{1,0,0,0,1}}

	// result: 
	// grid1 := [][]int{}
	// grid2 := [][]int{}

	result := countSubIslands(grid1, grid2)
	fmt.Printf("result = %v\n", result)
}
