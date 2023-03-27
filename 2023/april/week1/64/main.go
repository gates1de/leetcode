package main
import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

    for i := 1; i < m; i++ {
    	grid[i][0] += grid[i - 1][0]
    }

    for j := 1; j < n; j++{
        grid[0][j] += grid[0][j - 1]
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            grid[i][j] += min(grid[i - 1][j], grid[i][j - 1])
        }
    }

    return grid[m - 1][n - 1]
}

func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

func main() {
	// result: 7
	// grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}

	// result: 12
	// grid := [][]int{{1,2,3},{4,5,6}}

	// result: 
	grid := [][]int{{0,0},{0,0}}

	// result: 
	// grid := [][]int{}

	result := minPathSum(grid)
	fmt.Printf("result = %v\n", result)
}

