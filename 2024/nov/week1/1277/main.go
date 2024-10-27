package main
import (
	"fmt"
)

func countSquares(matrix [][]int) int {
	row := len(matrix)
	col := len(matrix[0])
	dp := make([][]int, row + 1)
	for i, _ := range dp {
		dp[i] = make([]int, col + 1)
	}

	result := int(0)

	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			if matrix[i][j] == 1 {
				dp[i + 1][j + 1] = min(min(dp[i][j + 1], dp[i + 1][j]), dp[i][j]) + 1
				result += dp[i + 1][j + 1]
			}
		}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 15
	// matrix := [][]int{
	// 	{0,1,1,1},
	// 	{1,1,1,1},
	// 	{0,1,1,1},
	// }

	// result: 7
	matrix := [][]int{
		{1,0,1},
		{1,1,0},
		{1,1,0},
	}

	// result: 
	// matrix := [][]int{}

	result := countSquares(matrix)
	fmt.Printf("result = %v\n", result)
}

