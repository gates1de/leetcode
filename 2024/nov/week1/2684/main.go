package main
import (
	"fmt"
)

func maxMoves(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, 2)
		dp[i][0] = 1
	}

	result := int(0)

	for i := 1; i < col; i++ {
		for j := 0; j < row; j++ {
			if (grid[j][i] > grid[j][i - 1]) && dp[j][0] > 0 {
				dp[j][1] = max(dp[j][1], dp[j][0] + 1)
			}

			if j - 1 >= 0 && (grid[j][i] > grid[j - 1][i - 1]) && dp[j - 1][0] > 0 {
				dp[j][1] = max(dp[j][1], dp[j - 1][0] + 1)
			}

			if j + 1 < row && (grid[j][i] > grid[j + 1][i - 1]) && dp[j + 1][0] > 0 {
				dp[j][1] = max(dp[j][1], dp[j + 1][0] + 1)
			}

			result = max(result, dp[j][1] - 1)
		}

		for k := 0; k < row; k++ {
			dp[k][0] = dp[k][1]
			dp[k][1] = 0
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// grid := [][]int{{2,4,3,5},{5,4,9,3},{3,4,2,11},{10,9,13,15}}

	// result: 0
	grid := [][]int{{3,2,4},{2,1,9},{1,1,7}}

	// result: 
	// grid := [][]int{}

	result := maxMoves(grid)
	fmt.Printf("result = %v\n", result)
}

