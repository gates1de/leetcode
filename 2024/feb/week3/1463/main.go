package main
import (
	"fmt"
)

func cherryPickup(grid [][]int) int {
    dp := make([][][]int, len(grid))
    for i := range dp {
        columns := make([][]int, len(grid[0]))
        for j := range columns {
            column := make([]int, len(grid[0]))
            for x := range column {
                column[x] = -1
            }
            columns[j] = column

        }
        dp[i]= columns
    }
    return dfs(grid, 0, 0, len(grid[0]) - 1, dp)
}

func dfs(grid [][]int, row int, col1 int, col2 int, dp [][][]int) int {
    if row == len(grid) {
        return 0
    }

    if dp[row][col1][col2] != -1 {
        return dp[row][col1][col2]
    }

    for i := -1; i <= 1; i++ {
        newCol1 := col1 + i
        for j := -1; j <=1; j++ {
            newCol2 := col2 + j
            if newCol1 >= 0 && newCol2 >= 0 && newCol1 < len(grid[0]) && newCol2 < len(grid[0]) {
                dp[row][col1][col2] = max(dp[row][col1][col2], dfs(grid, row + 1, newCol1, newCol2, dp))
            }
        }
    }

    if col1 == col2 {
        dp[row][col1][col2] += grid[row][col1]
    } else {
        dp[row][col1][col2] += grid[row][col1] + grid[row][col2]
    }
    return dp[row][col1][col2]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
	// result: 24
	// grid := [][]int{
	// 	{3,1,1},
	// 	{2,5,1},
	// 	{1,5,5},
	// 	{2,1,1},
	// }

	// result: 28
	grid := [][]int{
		{1,0,0,0,0,0,1},
		{2,0,0,0,0,3,0},
		{2,0,9,0,0,0,0},
		{0,3,0,5,4,0,0},
		{1,0,2,3,0,0,6},
	}

	// result: 
	// grid := [][]int{}

	result := cherryPickup(grid)
	fmt.Printf("result = %v\n", result)
}

