package main
import (
	"fmt"
)

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n + 1)

	for i := n - 1; i >= 0; i-- {
		dp[i] = stoneValue[i] - dp[i + 1]
		if i + 2 <= n {
			dp[i] = max(dp[i], stoneValue[i] + stoneValue[i + 1] - dp[i + 2])
		}
		if i + 3 <= n {
			dp[i] = max(dp[i], stoneValue[i] + stoneValue[i + 1] + stoneValue[i + 2] - dp[i + 3])
		}
	}

	if dp[0] > 0 {
		return "Alice"
	}
	if dp[0] < 0 {
		return "Bob"
	}
	return "Tie"
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: "Bob"
	stoneValue := []int{1,2,3,7}

	// result: "Alice"
	// stoneValue := []int{1,2,3,-9}

	// result: "Tie"
	// stoneValue := []int{1,2,3,6}

	// result: 
	// stoneValue := []int{}

	result := stoneGameIII(stoneValue)
	fmt.Printf("result = %v\n", result)
}

