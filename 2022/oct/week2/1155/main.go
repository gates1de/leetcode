package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numRollsToTarget(n int, k int, target int) int {
	memo := make([][]int, n + 1)
	for i, _ := range memo {
		memo[i] = make([]int, target + 1)
		for j, _ := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dp(n, k, target, memo) % modulo
}

func dp(n int, k int, target int, memo [][]int) int {
	if n == 0 {
		if target > 0 {
			return 0
		}
		return 1
	}

	if memo[n][target] >= 0 {
		return memo[n][target]
	}

	result := int(0)
	for i := 1; i <= k; i++ {
		remaining := target - i
		if remaining < 0 {
			break
		}
		result += dp(n - 1, k, remaining, memo)
		result %= modulo
	}
	memo[n][target] = result
	return result
}

func main() {
	// result: 1
	// n := int(1)
	// k := int(6)
	// target := int(3)

	// result: 6
	// n := int(2)
	// k := int(6)
	// target := int(7)

	// result: 222616187
	n := int(30)
	k := int(30)
	target := int(500)

	// result: 
	// n := int(0)
	// k := int(0)
	// target := int(0)

	result := numRollsToTarget(n, k, target)
	fmt.Printf("result = %v\n", result)
}

