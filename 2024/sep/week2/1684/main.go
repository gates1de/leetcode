package main
import (
	"fmt"
)

func countConsistentStrings(allowed string, words []string) int {
	m := make(map[rune]bool)
	for _, char := range allowed {
		m[char] = true
	}

	result := int(0)
	TOP: for _, word := range words {
		for _, char := range word {
			if !m[char] {
				continue TOP
			}
		}
		result++
	}

	return result
}

func main() {
	// result: 2
	// allowed := "ab"
	// words := []string{"ad","bd","aaab","baa","badab"}

	// result: 7
	// allowed := "abc"
	// words := []string{"a","b","c","ab","ac","bc","abc"}

	// result: 4
	allowed := "cad"
	words := []string{"cc","acd","b","ba","bac","bad","ac","d"}

	// result: 
	// allowed := ""
	// words := []string{}

	result := countConsistentStrings(allowed, words)
	fmt.Printf("result = %v\n", result)
}

