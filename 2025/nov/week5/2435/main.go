package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numberOfPaths(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][][]int64, m + 1)
	for i := range dp {
		dp[i] = make([][]int64, n + 1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, k)
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				dp[i][j][grid[0][0] % k] = 1
				continue
			}

			value := grid[i - 1][j - 1] % k
			for r := range k {
				prevMod := (r - value + k) % k
				dp[i][j][r] = (dp[i - 1][j][prevMod] + dp[i][j - 1][prevMod]) % int64(modulo)
			}
		}
	}

	return int(dp[m][n][0])
}

func main() {
	// result: 2
	// grid := [][]int{{5,2,4},{3,0,5},{0,7,2}}
	// k := int(3)

	// result: 1
	// grid := [][]int{{0,0}}
	// k := int(5)

	// result: 10
	grid := [][]int{{7,3,4,9},{2,3,6,2},{2,3,7,0}}
	k := int(1)

	// result: 
	// grid := [][]int{}
	// k := int(0)

	result := numberOfPaths(grid, k)
	fmt.Printf("result = %v\n", result)
}

