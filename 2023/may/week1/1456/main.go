package main
import (
	"fmt"
)

func maxVowels(s string, k int) int {
	result := int(0)

	for i := 0; i < k; i++ {
		if i >= len(s) {
			return result
		}
		if isVowel(s[i]) {
			result++
		}
	}

	exclude := s[0]
	current := result
	for i := k; i < len(s); i++ {
		if isVowel(exclude) {
			current--
		}
		if isVowel(s[i]) {
			current++
		}
		result = max(result, current)
		exclude = s[i - k + 1]
	}

	return result
}

func isVowel(char byte) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// s := "abciiidef"
	// k := int(3)

	// result: 2
	// s := "aeiou"
	// k := int(2)

	// result: 2
	s := "leetcode"
	k := int(3)

	result := maxVowels(s, k)
	fmt.Printf("result = %v\n", result)
}

