package main
import (
	"fmt"
)

func getLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	dp := make([]int, n)
	prev := make([]int, n)
	maxLen := int(1)
	endIndex := int(0)

	for i := 0; i < n; i++ {
		dp[i] = 1
		prev[i] = -1
	}

	for i := 1; i < n; i++ {
		bestLen := int(1)
		bestPrev := int(-1)

		for j := i - 1; j >= 0; j-- {
			if groups[i] != groups[j] && dp[j] + 1 > bestLen {
				bestLen, bestPrev = dp[j] + 1, j
			}
		}

		dp[i] = bestLen
		prev[i] = bestPrev

		if dp[i] > maxLen {
			maxLen = dp[i]
			endIndex = i
		}
	}

	result := make([]string, 0)
	for i := endIndex; i != -1; i = prev[i] {
		result = append(result, words[i])
	}

	reverse(result)
	return result
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// result: ["e","b"]
	// words := []string{"e","a","b"}
	// groups := []int{0,0,1}

	// result: ["a","b","c"]
	words := []string{"a","b","c","d"}
	groups := []int{1,0,1,1}

	// result: []
	// words := []string{}
	// groups := []int{}

	result := getLongestSubsequence(words, groups)
	fmt.Printf("result = %v\n", result)
}

