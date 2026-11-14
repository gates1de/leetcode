package main

import (
	"fmt"
)

func climbStairs(n int, costs []int) int {
	if n == 0 {
		return 0
	}

	dp := [3]int{}
	for step := 1; step <= n; step++ {
		result := dp[(step-1)%3] + 1
		if step >= 2 {
			result = min(result, dp[(step-2)%3]+4)
		}
		if step >= 3 {
			result = min(result, dp[(step-3)%3]+9)
		}

		dp[step%3] = result + costs[step-1]
	}

	return dp[n%3]
}

func main() {
	// result: 13
	// n := 4
	// costs := []int{1,2,3,4}

	// result: 11
	// n := 4
	// costs := []int{5,1,6,2}

	// result: 12
	n := 3
	costs := []int{9, 8, 3}

	// result: 0
	// n := 0
	// costs := []int{}

	result := climbStairs(n, costs)
	fmt.Printf("result = %v\n", result)
}
