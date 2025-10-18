package main

import (
	"fmt"
)

func removeAnagrams(words []string) []string {
	n := len(words)
	result := make([]string, 0, len(words))
	result= append(result, words[0])

	for i := 1; i < n; i++ {
		if !compare(words[i], words[i - 1]) {
			result = append(result, words[i])
		}
	}

	return result
}

func compare(word1, word2 string) bool {
	freq := make([]int, 26)
	for _, char := range word1 {
		freq[char - 'a']++
	}
	for _, char := range word2 {
		freq[char - 'a']--
	}

	for _, count := range freq {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	// result: ["abba","cd"]
	// words := []string{"abba","baba","bbaa","cd","cd"}

	// result: ["a","b","c","d","e"]
	words := []string{"a","b","c","d","e"}

	// result: []
	// words := []string{}

	result := removeAnagrams(words)
	fmt.Printf("result = %v\n", result)
}

