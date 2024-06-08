package main
import (
	"fmt"
)

func commonChars(words []string) []string {
	n := len(words)
	commonCounts := [26]int{}
	currentCounts := [26]int{}
	result := make([]string, 0)

	for _, char := range words[0] {
		commonCounts[char - 'a']++
	}

	for i := 1; i < n; i++ {
		currentCounts = [26]int{}

		for _, char := range words[i] {
			currentCounts[char - 'a']++
		}

		for letter := 0; letter < 26; letter++ {
			commonCounts[letter] = min(commonCounts[letter], currentCounts[letter])
		}
	}

	for letter := 0; letter < 26; letter++ {
		for count := 0; count < commonCounts[letter]; count++ {
			result = append(result, string(letter + 'a'))
		}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: ["e","l","l"]
	// words := []string{"bella","label","roller"}

	// result: ["c","o"]
	words := []string{"cool","lock","cook"}

	// result: []
	// words := []string{}

	result := commonChars(words)
	fmt.Printf("result = %v\n", result)
}

