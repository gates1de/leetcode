package main
import (
	"fmt"
)

func minDistance(word1 string, word2 string) int {
	word1Len := len(word1)
	word2Len := len(word2)

	if word1Len == 0 {
		return word2Len
	}
	if word2Len == 0 {
		return word1Len
	}

	dp := make([][]int, word1Len + 1)
	for i, _ := range dp {
		dp[i] = make([]int, word2Len + 1)
	}

	for i := 1; i <= word1Len; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= word2Len; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= word1Len; i++ {
		for j := 1; j <= word2Len; j++ {
			if word1[i - 1] == word2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] = min(dp[i - 1][j], min(dp[i][j - 1], dp[i - 1][j - 1])) + 1
			}
		}
	}
	return dp[word1Len][word2Len]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 3
	word1 := "horse"
	word2 := "ros"

	// result: 5
	// word1 := "intention"
	// word2 := "execution"

	// result: 
	// word1 := ""
	// word2 := ""

	result := minDistance(word1, word2)
	fmt.Printf("result = %v\n", result)
}

