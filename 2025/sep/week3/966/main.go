package main

import (
	"fmt"
	"strings"
)

func spellchecker(wordlist []string, queries []string) []string {
	if len(wordlist) == 0 || len(queries) == 0 {
		return []string{}
	}

	wordSet := make(map[string]bool)
	wordLowers := make(map[string]string)
	wordDevowels := make(map[string]string)

	for _, word := range wordlist {
		wordSet[word] = true

		lowerWord := strings.ToLower(word)
		if _, exists := wordLowers[lowerWord]; !exists {
			wordLowers[lowerWord] = word
		}

		devowelLowerWord:= devowel(lowerWord)
		if _, exists := wordDevowels[devowelLowerWord]; !exists {
			wordDevowels[devowelLowerWord] = word
		}
	}

	result := make([]string, len(queries))
	i := int(0)
	for _, query := range queries {
		result[i] = solve(query, wordSet, wordLowers, wordDevowels)
		i++
	}

	return result
}

func solve(
	query string,
	wordSet map[string]bool,
	wordLowers map[string]string,
	wordDevowels map[string]string,
) string {
	if wordSet[query] {
		return query
	}

	lowerQuery := strings.ToLower(query)
	if q, exists := wordLowers[lowerQuery]; exists {
		return q
	}

	devowelLowerQuery := devowel(lowerQuery)
	if q, exists := wordDevowels[devowelLowerQuery]; exists {
		return q
	}

	return ""
}

func devowel(word string) string {
	result := make([]byte, len(word))

	for i, char := range word {
		b := byte(char)
		if isVowel(b) {
			result[i] = '*'
		} else {
			result[i] = b
		}
	}

	return string(result)
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	// result: ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]
	// wordlist := []string{"KiTe", "kite", "hare", "Hare"}
	// queries := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}

	// result: ["yellow"]
	wordlist := []string{"yellow"}
	queries := []string{"YellOw"}

	// result: []
	// wordlist := []string{}
	// queries := []string{}

	result := spellchecker(wordlist, queries)
	fmt.Printf("result = %v\n", result)
}

