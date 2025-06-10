package main
import (
	"fmt"
)

func maxDifference(s string) int {
	a1 := int(0)
	a2 := len(s)
	m := make(map[rune]int)

	for _, char := range s {
		m[char]++
	}

	for _, count := range m {
		if count % 2 == 1 {
			a1 = max(a1, count)
		} else {
			a2 = min(a2, count)
		}
	}

	return a1 - a2
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// s := "aaaaabbc"

	// result: 1
	s := "abcabcab"

	// result: 
	// s := ""

	result := maxDifference(s)
	fmt.Printf("result = %v\n", result)
}

