package main
import (
	"fmt"
)

func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	result := make([]string, 0)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	backtrack(s, wordSet, "", &result, 0)
	return result
}

func backtrack(s string, wordSet map[string]bool, current string, result *[]string, index int) {
	if index == len(s) {
		*result = append(*result, current[:len(current) - 1])
		return
	}

	for endIndex := index + 1; endIndex <= len(s); endIndex++ {
		word := s[index:endIndex]
		if wordSet[word] {
			backtrack(s, wordSet, current + word + " ", result, endIndex)
		}
	}
}

func main() {
	// result: ["cats and dog","cat sand dog"]
	s := "catsanddog"
	wordDict := []string{"cat","cats","and","sand","dog"}

	// result: ["pine apple pen apple","pineapple pen apple","pine applepen apple"]
	// s := "pineapplepenapple"
	// wordDict := []string{"apple","pen","applepen","pine","pineapple"}

	// result: []
	// s := "catsandog"
	// wordDict := []string{"cats","dog","sand","and","cat"}

	// result: []
	// s := ""
	// wordDict := []string{}

	result := wordBreak(s, wordDict)
	fmt.Printf("result = %v\n", result)
}

