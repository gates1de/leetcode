package main
import (
	"fmt"
)

var dirs = []int{0, 1, 0, -1, 0, 0}

func getMaximumGold(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	result := int(0)

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			result = max(result, dfs(grid, m, n, row, col))
		}
	}

	return result
}

func dfs(grid [][]int, m int, n int, row int, col int) int {
	if row < 0 || col < 0 || row == m || col == n || grid[row][col] == 0 {
		return 0
	}

	maxGold := int(0)
	originalVal := grid[row][col]
	grid[row][col] = 0

	for dir := 0; dir < 4; dir++ {
		maxGold = max(maxGold, dfs(grid, m, n, dirs[dir] + row, dirs[dir + 1] + col))
	}

	grid[row][col] = originalVal
	return maxGold + originalVal
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 24
	// grid := [][]int{{0,6,0},{5,8,7},{0,9,0}}

	// result: 28
	grid := [][]int{{1,0,7},{2,0,6},{3,4,5},{0,3,0},{9,0,20}}

	// result: 
	// grid := [][]int{}

	result := getMaximumGold(grid)
	fmt.Printf("result = %v\n", result)
}

