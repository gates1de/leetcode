package main
import (
	"fmt"
)

func kInversePairs(n int, k int) int {
	if k == 0 {
		return 1
	}
	modulo := int(1e9 + 7)
	dp := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k + 1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = 1
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i][j - 1] + dp[i - 1][j]
			dp[i][j] %= modulo
			if i <= j {
				dp[i][j] = (dp[i][j] - dp[i - 1][j - i] + modulo ) % modulo
			}
		}
	}
	return dp[n][k]
}

func main() {
	// result: 1
	// n := int(3)
	// k := int(0)

	// result: 2
	// n := int(3)
	// k := int(1)

	// result: 531
	n := int(7)
	k := int(9)

	// result: 
	// n := int(0)
	// k := int(0)

	result := kInversePairs(n, k)
	fmt.Printf("result = %v\n", result)
}

