package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)
var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    
func countPaths(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	result := int(0)
	for i, _ := range grid {
		for j, _ := range grid[i] {
			result = (result + dfs(i, j, grid, dp)) % modulo
		}
	}

	return result
}

func dfs(i int, j int, grid [][]int, dp [][]int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}

	result := int(1)
	for _, dir := range directions {
		prevI := i + dir[0]
		prevJ := j + dir[1]

		if 0 <= prevI && prevI < len(grid) &&
			0 <= prevJ && prevJ < len(grid[0]) &&
			grid[prevI][prevJ] < grid[i][j] {
			result += dfs(prevI, prevJ, grid, dp)
			result %= modulo
		}
	}

	dp[i][j] = result
	return result
}

func main() {
	// result: 8
	// grid := [][]int{{1,1},{3,4}}

	// result: 3
	grid := [][]int{{1},{2}}

	// result: 
	// grid := [][]int{}

	result := countPaths(grid)
	fmt.Printf("result = %v\n", result)
}

