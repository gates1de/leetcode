package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func shortestPalindrome(s string) string {
	hashBase := int(29)
	forwardHash := int(0)
	reverseHash := int(0)
	powerVal := int(1)
	palindromeEndIndex := int(-1)

	for i, char := range s {
		forwardHash = (forwardHash * hashBase + int(char - 'a' + 1)) % modulo
		reverseHash = (reverseHash + int(char - 'a' + 1) * powerVal) % modulo
		powerVal = (powerVal * hashBase) % modulo

		if forwardHash == reverseHash {
			palindromeEndIndex = i
		}
	}

	suffix := s[palindromeEndIndex + 1:]
	reversedSuffix := reverse(suffix)

	return fmt.Sprintf("%s%s", reversedSuffix, s)
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars) - 1; i < j; i, j = i + 1, j - 1 {
			chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	// result: "aacecaaa"
	// s := "aacecaaa"

	// result: "dcbabcd"
	s := "abcd"

	// result: ""
	// s := ""

	result := shortestPalindrome(s)
	fmt.Printf("result = %v\n", result)
}

