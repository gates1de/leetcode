package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func kInversePairs(n int, k int) int {
    dp := make([][]int, n + 1)
    for i, _ := range dp {
        dp[i] = make([]int, k + 1)
    }

    for i := 1; i <= n; i++ {
        for j := 0; j <= k; j++ {
            if j == 0 {
                dp[i][j] = 1
            } else {
                for l := 0; l <= min(j, i - 1); l++ {
                    dp[i][j] = (dp[i][j] + dp[i - 1][j - l]) % modulo
                }
            }
        }
    }

    return dp[n][k]
}

func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

func main() {
	// result: 1
	// n := int(3)
	// k := int(0)

	// result: 2
	n := int(3)
	k := int(1)

	// result: 
	// n := int(0)
	// k := int(0)

	result := kInversePairs(n, k)
	fmt.Printf("result = %v\n", result)
}

