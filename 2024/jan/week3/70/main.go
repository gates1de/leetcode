package main
import (
	"fmt"
)

func climbStairs(n int) int {
	dp := make([]int, n + 1)
	for i, _ := range dp {
		dp[i] = -1
	}
	return helper(n, dp)
}

func helper(n int, dp []int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	}

	if dp[n] != -1 {
		return dp[n]
	}

	dp[n] = helper(n - 1, dp) + helper(n - 2, dp)
	return dp[n]
}

func main() {
	// result: 2
	// n := int(2)

	// result: 3
	n := int(3)

	// result: 
	// n := int(0)

	result := climbStairs(n)
	fmt.Printf("result = %v\n", result)
}

