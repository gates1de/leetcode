package main
import (
	"fmt"
	"unicode"
)

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	hasVowel := false
	hasConsonant := false

	for _, char := range word {
		if unicode.IsLetter(char) {
			lowerChar := unicode.ToLower(char)

			if lowerChar == 'a' || lowerChar == 'e' || lowerChar == 'i' || lowerChar == 'o' || lowerChar == 'u' {
				hasVowel = true
			} else {
				hasConsonant = true
			}
		} else if !unicode.IsDigit(char) {
			return false
		}
	}

	return hasVowel && hasConsonant
}

func main() {
	// result: true
	// word := "234Adas"

	// result: false
	// word := "b3"

	// result: false
	word := "a3$e"

	// result: 
	// word := ""

	result := isValid(word)
	fmt.Printf("result = %v\n", result)
}

