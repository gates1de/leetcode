package main
import (
	"fmt"
)

func longestPalindrome(s string) int {
	m := make(map[rune]bool)
	result := int(0)

	for _, char := range s {
		if m[char] {
			delete(m, char)
			result += 2
		} else {
			m[char] = true
		}
	}

	if len(m) > 0 {
		result++
	}

	return result
}

func main() {
	// result: 7
	// s := "abccccdd"

	// result: 1
	s := "a"

	// result: 
	// s := ""

	result := longestPalindrome(s)
	fmt.Printf("result = %v\n", result)
}

