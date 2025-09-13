package main

import (
	"fmt"
)

func maxFreqSum(s string) int {
	m := make(map[byte]int)
	for _, char := range s {
		m[byte(char)]++
	}

	vowelCount := int(0)
	consonantCount := int(0)
	for char, count := range m {
		if isVowel(char) {
			vowelCount = max(vowelCount, count)
		} else {
			consonantCount = max(consonantCount, count)
		}
	}

	return vowelCount + consonantCount
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 6
	// s := "successes"

	// result: 3
	s := "aeiaeia"

	// result: 
	// s := ""

	result := maxFreqSum(s)
	fmt.Printf("result = %v\n", result)
}

