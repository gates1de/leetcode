package main
import (
	"fmt"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	m := map[string]bool{}
	uniqueWords := []string{}
	for _, word := range words {
		if m[word] {
			continue
		}
		m[word] = true
		uniqueWords = append(uniqueWords, word)
	}

	m = map[string]bool{}
	for i := 0; i < len(uniqueWords); i++ {
		word1 := uniqueWords[i]
		if m[word1] {
			continue
		}

		for j := i + 1; j < len(uniqueWords); j++ {
			word2 := uniqueWords[j]
			if strings.HasSuffix(word1, word2) {
				m[word2] = true
			} else if strings.HasSuffix(word2, word1) {
				m[word1] = true
			}
		}

	}

	result := int(0)
	for _, word := range uniqueWords {
		if m[word] {
			continue
		}
		result += len(word) + 1
	}

	return result
}

func main() {
	// result: 10
	// words := []string{"time", "me", "bell"}

	// result: 2
	// words := []string{"t"}

	// result: 10
	// words := []string{"ime", "time", "me", "bell"}

	// result: 20
	// words := []string{"ell","time","be","bme","me","im","bell"}

	// result: 5
	words := []string{"time", "time", "time", "time"}

	// result: 
	// words := []string{}

	result := minimumLengthEncoding(words)
	fmt.Printf("result = %v\n", result)
}

