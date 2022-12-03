package main
import (
	"fmt"
)

func halvesAreAlike(s string) bool {
	n := len(s)
	if n <= 1 {
		return false
	}

	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	left := s[:n / 2]
	right := s[n / 2:]

	leftVowelCount := int(0)
	rightVowelCount := int(0)

	for i, leftChar := range left {
		rightChar := right[i]
		if vowels[byte(leftChar)] {
			leftVowelCount++
		}
		if vowels[rightChar] {
			rightVowelCount++
		}
	}

	return leftVowelCount == rightVowelCount
}

func main() {
	// result: true
	// s := "book"

	// result: false
	s := "textbook"

	// result: 
	// s := ""

	result := halvesAreAlike(s)
	fmt.Printf("result = %v\n", result)
}

