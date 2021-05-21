package main
import (
	"fmt"
)

func findAndReplacePattern(words []string, pattern string) []string {
	// fmt.Printf("m = %v\n", m)
	result := []string{}
	patternS := generatePatternString(pattern)
	// fmt.Printf("patternS = %v\n", patternS)
	for _, word := range words {
		s := generatePatternString(word)
		// fmt.Printf("s = %v\n", s)
		if s == patternS {
			result = append(result, word)
		}
	}
	return result
}

func generatePatternString(word string) string {
	result := ""
	runes := []rune{}
	for _, r := range word {
		runeIndex := findIndex(r, runes)
		// fmt.Printf("runeIndex = %v\n", runeIndex)
		if runeIndex == -1 {
			runes = append(runes, r)
			// result[len(runes) - 1]++
			result += string(int(rune('a')) + len(runes) - 1)
		} else {
			// result[runeIndex]++
			result += string(int(rune('a')) + runeIndex)
		}
	}
	return result
}

func findIndex(target rune, runes []rune) int {
	for i, r := range runes {
		if r == target {
			return i
		}
	}
	return -1
}

func main() {
	// result: ["mee", "aqq"]
	// words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	// pattern := "abb"

	// result: ["a", "b", "c"]
	// words := []string{"a", "b", "c"}
	// pattern := "a"

	// result: ["abc", "cba"]
	// words := []string{"abc", "cba", "xyx", "yxx", "yyx"}
	// pattern := "abc"

	// result: ["abab", "dede"]
	words := []string{"badc", "abab", "dddd", "dede", "yyxx"}
	pattern := "baba"

	// result: 
	// words := []string{}
	// pattern := ""

	result := findAndReplacePattern(words, pattern)
	fmt.Printf("result = %v\n", result)
}

