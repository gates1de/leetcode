package main

import (
	"fmt"
)

func numOfStrings(patterns []string, word string) int {
	result := int(0)
	for _, pattern := range patterns {
		if len(pattern) > len(word) {
			continue
		}

		for i := 0; i + len(pattern) <= len(word); i++ {
			if word[i:i+len(pattern)] == pattern {
				result++
				break
			}
		}
	}

	return result
}

func main() {
	// result: 3
	// patterns := []string{"a","abc","bc","d"}
	// word := "abc"

	// result: 2
	// patterns := []string{"a","b","c"}
	// word := "aaaaabbbbb"

	// result: 3
	patterns := []string{"a","a","a"}
	word := "ab"

	// result: 
	// patterns := []string{}
	// word := ""

	result := numOfStrings(patterns, word)
	fmt.Printf("result = %v\n", result)
}
