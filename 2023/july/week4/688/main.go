package main
import (
	"fmt"
)

func knightProbability(n int, k int, row int, column int) float64 {
	dirs := [][]int{
		{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
		{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
	}

	dp := make([][][]float64, k + 1)
	for i, _ := range dp {
		dp[i] = make([][]float64, n)
		for j, _ := range dp[i] {
			dp[i][j] = make([]float64, n)
		}
	}

	dp[0][row][column] = 1

	for moves := 1; moves <= k; moves++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for _, dir := range dirs {
					prevI := i - dir[0]
					prevJ := j - dir[1]
					if prevI >= 0 && prevI < n && prevJ >= 0 && prevJ < n {
						dp[moves][i][j] += dp[moves - 1][prevI][prevJ] / 8.0
					}
				}
			}
		}
	}

	result := float64(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result += dp[k][i][j]
		}
	}

	return result
}

func main() {
	// result: 0.06250
	// n := int(3)
	// k := int(2)
	// row := int(0)
	// column := int(0)

	// result: 1.00000
	n := int(1)
	k := int(0)
	row := int(0)
	column := int(0)

	// result: 
	// n := int(0)
	// k := int(0)
	// row := int(0)
	// column := int(0)

	result := knightProbability(n, k, row, column)
	fmt.Printf("result = %v\n", result)
}

