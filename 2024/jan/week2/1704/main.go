package main
import (
	"fmt"
)

func halvesAreAlike(s string) bool {
	i := int(0)
	j := len(s) - 1

	leftVowelCount := int(0)
	rightVowelCount := int(0)
	for i < j {
		left := s[i]
		right := s[j]

		if left == 'a' || left == 'e' || left == 'i' || left == 'o' || left == 'u' ||
			left == 'A' || left == 'E' || left == 'I' || left == 'O' || left == 'U' {
			leftVowelCount++
		}
		if right == 'a' || right == 'e' || right == 'i' || right == 'o' || right == 'u' ||
			right == 'A' || right == 'E' || right == 'I' || right == 'O' || right == 'U' {
			rightVowelCount++
		}

		i++
		j--
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

