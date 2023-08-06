package main
import (
	"fmt"
)

const modulo = int64(1e9 + 7)

func numMusicPlaylists(n int, goal int, k int) int {
	dp := make([][]int64, goal + 1)
	for i, _ := range dp {
		dp[i] = make([]int64, n + 1)
	}
	dp[0][0] = 1

	for i := 1; i <= goal; i++ {
		for j := 1; j <= min(i, n); j++ {
			dp[i][j] = dp[i - 1][j - 1] * int64(n - j + 1) % modulo

			if j > k {
				dp[i][j] = (dp[i][j] + dp[i - 1][j] * int64(j - k)) % modulo
			}
		}
	}

	return int(dp[goal][n])
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 6
	// n := int(3)
	// goal := int(3)
	// k := int(1)

	// result: 6
	// n := int(2)
	// goal := int(3)
	// k := int(0)

	// result: 2
	n := int(2)
	goal := int(3)
	k := int(1)

	// result: 
	// n := int(0)
	// goal := int(0)
	// k := int(0)

	result := numMusicPlaylists(n, goal, k)
	fmt.Printf("result = %v\n", result)
}

