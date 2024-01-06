package main
import (
	"fmt"
)

func maxLengthBetweenEqualCharacters(s string) int {
	result := int(-1)
	if len(s) <= 1 {
		return result
	}

	m := make(map[rune]int)
	for i, char := range s {
		if _, ok := m[char]; ok {
			result = max(result, i - m[char] - 1)
			continue
		}
		m[char] = i 
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 0
	// s := "aa"

	// result: 2
	// s := "abca"

	// result: -1
	s := "cbzxy"

	// result: 
	// s := ""

	result := maxLengthBetweenEqualCharacters(s)
	fmt.Printf("result = %v\n", result)
}

