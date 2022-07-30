package main
import (
	"fmt"
)

const a = byte(97)

func findAndReplacePattern(words []string, pattern string) []string {
	convertedPattern := convert(pattern)
	result := []string{}

	for _, word := range words {
		convertedWord := convert(word)
		if convertedPattern == convertedWord {
			result = append(result, word)
		}
	}

	return result
}

func convert(s string) string {
	m := map[rune]byte{}
	result := ""
	for _, char := range s {
		if m[char] == 0 {
			m[char] = byte(len(m) + 1)
		}
		result += string(a + m[char] - 1)
	}
	return result
}

func main() {
	// result: ["mee","aqq"]
	// words := []string{"abc","deq","mee","aqq","dkd","ccc"}
	// pattern := "abb"

	// result: ["a","b","c"]
	// words := []string{"a","b","c"}
	// pattern := "a"

	// result: ["dba","abc","zwy"]
	// words := []string{"dba", "abc", "zwy", "bww", "ghg", "take"}
	// pattern := "aop"

	// result: ["ghg"]
	words := []string{"dba", "abc", "zwy", "bww", "ghg", "take"}
	pattern := "aza"

	// result: 
	// words := []string{}
	// pattern := ""

	result := findAndReplacePattern(words, pattern)
	fmt.Printf("result = %v\n", result)
}

