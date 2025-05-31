package main
import (
	"fmt"
)

func longestPalindrome(words []string) int {
	m := map[string]int{}
	for _, word := range words {
		m[word]++
	}

	result := int(0)
	isExistsSingleSameLetter := false
	for _, word := range words {
		if len(word) != 2 || m[word] == 0 {
			continue
		}

		if word[0] == word[1] {
			if m[word] % 2 == 0 {
				result += m[word] * 2
				m[word] = 0
			} else {
				result += (m[word] - 1) * 2
				m[word] = 1
				isExistsSingleSameLetter = true
			}
		} else {
			invertedWord := string([]byte{word[1], word[0]})
			if m[invertedWord] >= 1 {
				result += 4
				m[invertedWord] -= 1
				m[word] -= 1
			}
		}
	}

	if isExistsSingleSameLetter {
		result += 2
	}

	return result
}

func main() {
	// result: 6
	// words := []string{"lc","cl","gg"}

	// result: 8
	// words := []string{"ab","ty","yt","lc","cl","ab"}

	// result: 2
	// words := []string{"cc","ll","xx"}

	// result: 22
	// words := []string{"dd","aa","bb","dd","aa","dd","bb","dd","aa","cc","bb","cc","dd","cc"}

	// result: 14
	words := []string{"em","pe","mp","ee","pp","me","ep","em","em","me"}

	// result: 
	// words := []string{}

	result := longestPalindrome(words)
	fmt.Printf("result = %v\n", result)
}

