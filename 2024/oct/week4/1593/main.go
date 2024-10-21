package main
import (
	"fmt"
)

func maxUniqueSplit(s string) int {
	seen := make(map[string]bool)
	return backtrack(s, 0, seen)
}

func backtrack(s string, start int, seen map[string]bool) int {
	n := len(s)
	if start == n {
		return 0
	}

	result := int(0)

	for end := start + 1; end <= n; end++ {
		substring := s[start:end]
		if !seen[substring] {
			seen[substring] = true
			result = max(result, backtrack(s, end, seen) + 1)
			delete(seen, substring)
		}
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
	// result: 5
	// s := "ababccc"

	// result: 2
	// s := "aba"

	// result: 1
	s := "aa"

	// result: 
	// s := ""

	result := maxUniqueSplit(s)
	fmt.Printf("result = %v\n", result)
}

