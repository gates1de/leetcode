package main
import (
	"fmt"
)

func minimumDeleteSum(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)
	dp := make([][]int, m + 1)
	for i, _ := range dp {
		dp[i] = make([]int, n + 1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i - 1][0] + int(s1[i - 1])
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j - 1] + int(s2[j - 1])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i - 1] == s2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] = min(int(s1[i - 1]) + dp[i - 1][j], int(s2[j - 1]) + dp[i][j - 1])
			}
		}
	}

	return dp[m][n]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 231
	// s1 := "sea"
	// s2 := "eat"

	// result: 403
	s1 := "delete"
	s2 := "leet"

	// result: 
	// s1 := ""
	// s2 := ""

	result := minimumDeleteSum(s1, s2)
	fmt.Printf("result = %v\n", result)
}

