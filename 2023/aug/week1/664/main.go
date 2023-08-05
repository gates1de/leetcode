package main
import (
	"fmt"
)

func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	for length := 1; length <= n; length++ {
		for left := 0; left <= n - length; left++ {
			right := left + length - 1
			j := -1
			dp[left][right] = n

			for i := left; i < right; i++ {
				if s[i] != s[right] && j == -1 {
					j = i
				}

				if j != -1 {
					dp[left][right] = min(dp[left][right], 1 + dp[j][i] + dp[i + 1][right])
				}
			}

			if j == -1 {
				dp[left][right] = 0
			}
		}
	}

	return dp[0][n - 1] + 1
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// s := "aaabbb"

	// result: 2
	s := "aba"

	// result: 
	// s := ""

	result := strangePrinter(s)
	fmt.Printf("result = %v\n", result)
}

