package main
import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))

	for i, _ := range dp {
		for _, word := range wordDict {
			n := len(word)
			if i < n - 1 {
				continue
			}

			if i == n - 1 || dp[i - n] {
				if s[i - n + 1:i + 1] == word {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(s) - 1]
}

func main() {
	// result: true
	// s := "leetcode"
	// wordDict := []string{"leet","code"}

	// result: true
	// s := "applepenapple"
	// wordDict := []string{"apple","pen"}

	// result: false
	s := "catsandog"
	wordDict := []string{"cats","dog","sand","and","cat"}

	// result: 
	// s := ""
	// wordDict := []string{}

	result := wordBreak(s, wordDict)
	fmt.Printf("result = %v\n", result)
}

