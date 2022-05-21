package main
import (
	"fmt"
)

func uniquePathsWithObstacles(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])

    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                continue
            }

            if i == 0 && j == 0 {
                dp[i][j] = 1
                continue
            }

            top := int(0)
            if i > 0 {
                top = dp[i - 1][j]
            }

            left := int(0)
            if j > 0 {
                left = dp[i][j - 1]
            }

            dp[i][j] = top + left
        }
    }

    return dp[m - 1][n - 1]
}

func main() {
	// result: 2
	// obstacleGrid := [][]int{{0,0,0},{0,1,0},{0,0,0}}

	// result: 1
	// obstacleGrid := [][]int{{0,1},{0,0}}

	// result: 0
	obstacleGrid := [][]int{{0,0},{0,1}}

	// result: 
	// obstacleGrid := [][]int{}

	result := uniquePathsWithObstacles(obstacleGrid)
	fmt.Printf("result = %v\n", result)
}

