package main
import (
	"fmt"
)

func findAllConcatenatedWordsInADict(words []string) []string {
	n := len(words)
	m := make(map[string]bool, n)
	for _, word := range words {
		m[word] = true
	}
	result := make([]string, 0, n)
	for _, word := range words {
		wordLen := len(word)
		dp := make([]bool, wordLen + 1)
		dp[0] = true
		for i := 1; i <= wordLen; i++ {
			initJ := int(0)
			if wordLen == i {
				initJ = 1
			}
			for j := initJ; !dp[i] && j < i; j++ {
				dp[i] = dp[j] && m[string(word[j:i])]
			}

			if dp[wordLen] {
				result = append(result, word)
			}
		}
	}
	return result
}

func main() {
	// result: ["catsdogcats","dogcatsdog","ratcatdogcat"]
	// words := []string{"cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"}

	// result: ["catdog"]
	words := []string{"cat","dog","catdog"}

	// result: 
	// words := []string{}

	result := findAllConcatenatedWordsInADict(words)
	fmt.Printf("result = %v\n", result)
}

