package main
import (
	"fmt"
)

const modulo = int64(1e9 + 7)

func numWays(words []string, target string) int {
	m := len(target)
	k := len(words[0])
	counts := make([][]int, 26)
	for i, _ := range counts {
		counts[i] = make([]int, k)
	}
	for i := 0; i < k; i++ {
		for _, word := range words {
			counts[word[i] - 97][i]++
		}
	}

	dp := make([][]int64, m + 1)
	for i, _ := range dp {
		dp[i] = make([]int64, k + 1)
	}

	dp[0][0] = 1
	for i := 0; i <= m; i++ {
		for j := 0; j < k; j++ {
			if i < m {
				dp[i + 1][j + 1] += int64(counts[target[i] - 97][j]) * dp[i][j]
				dp[i + 1][j + 1] %= modulo
			}
			dp[i][j + 1] += dp[i][j]
			dp[i][j + 1] %= modulo
		}
	}

	return int(dp[m][k])
}

func main() {
	// result: 6
	// words := []string{"acca","bbbb","caca"}
	// target := "aba"

	// result: 4
	words := []string{"abba","baab"}
	target := "bab"

	// result: 
	// words := []string{}
	// target := ""

	result := numWays(words, target)
	fmt.Printf("result = %v\n", result)
}

