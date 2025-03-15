package main
import (
	"fmt"
)

func countOfSubstrings(word string, k int) int64 {
	return atLeastK(word, k) - atLeastK(word, k + 1);
}

func atLeastK(word string, k int) int64 {
	result := int64(0)
	start := int(0)
	end := int(0)
	vowelCount := make(map[byte]int)
	consonantCount := int(0)

	for end < len(word) {
		new := word[end]
		if isVowel(new) {
			vowelCount[new]++
		} else {
			consonantCount++
		}

		for len(vowelCount) == 5 && consonantCount >= k {
			result += int64(len(word) - end)
			startLetter := word[start]
			if isVowel(startLetter) {
				vowelCount[startLetter]--

				if vowelCount[startLetter] == 0 {
					delete(vowelCount, startLetter)
				}
			} else {
				consonantCount--
			}

			start++
		}

		end++
	}

	return result
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	// result: 0
	// word := "aeioqq"
	// k := int(1)

	// result: 1
	// word := "aeiou"
	// k := int(0)

	// result: 3
	word := "ieaouqqieaouqq"
	k := int(1)

	// result: 
	// word := ""
	// k := int(0)

	result := countOfSubstrings(word, k)
	fmt.Printf("result = %v\n", result)
}

