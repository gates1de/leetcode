package main
import (
	"fmt"
)

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([]int, n)
	dpPrev := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[j] = dpPrev[j - 1] + 2
			} else {
				dp[j] = max(dpPrev[j], dp[j - 1])
			}
		}
		copy(dpPrev, dp)
	}
	return dp[n - 1]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// s := "bbbab"

	// result: 2
	s := "cbbd"

	// result: 
	// s := ""

	result := longestPalindromeSubseq(s)
	fmt.Printf("result = %v\n", result)
}

