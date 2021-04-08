package main
import (
	"fmt"
	"strings"
)

func halvesAreAlike(s string) bool {
	front := strings.ToLower(s[:len(s) / 2])
	behind := strings.ToLower(s[len(s) / 2:])
	// fmt.Printf("front = %v, behind = %v\n", front, behind)

	frontVowelCount := int(0)
	behindVowelCount := int(0)
	for _, vowel := range []string{"a", "e", "i", "o", "u"} {
		frontVowelCount += strings.Count(front, vowel)
		behindVowelCount += strings.Count(behind, vowel)
	}
	return frontVowelCount == behindVowelCount
}

func main() {
	// result: true
	// s := "book"

	// result: false
	// s := "textbook"

	// result: false
	// s := "MerryChristmas"

	// result: true
	s := "AbCdEfGh"

	result := halvesAreAlike(s)
	fmt.Printf("result = %v\n", result)
}

