package main
import (
	"fmt"
)

func detectCapitalUse(word string) bool {
	if len(word) == 0 {
		return false
	}

	isUpperFirst := 65 <= word[0] && word[0] <= 90
	upperCount := int(0)
	lowerCount := int(0)
	for _, char := range word {
		if 65 <= char && char <= 90  {
			upperCount++
		} else if 97 <= char && char <= 122  {
			lowerCount++
		}
	}

	if upperCount == len(word) {
		return true
	} else if lowerCount == len(word) {
		return true
	} else if isUpperFirst && upperCount == 1 && lowerCount + 1 == len(word) {
		return true
	}
	return false
}

func main() {
	// result: true
	// word := "USA"

	// result: false
	// word := "FlaG"

	// result: true
	word := "Google"

	// result: 
	// word := ""

	result := detectCapitalUse(word)
	fmt.Printf("result = %v\n", result)
}

