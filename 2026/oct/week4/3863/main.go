package main

import (
	"fmt"
)

func minOperations(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	sorted := true
	minChar, maxChar := s[0], s[0]
	for i := 1; i < n; i++ {
		if s[i] < s[i-1] {
			sorted = false
		}
		if s[i] < minChar {
			minChar = s[i]
		}
		if s[i] > maxChar {
			maxChar = s[i]
		}
	}

	if sorted {
		return 0
	}
	if n == 2 {
		return -1
	}
	if s[0] == minChar || s[n-1] == maxChar {
		return 1
	}
	for i := 1; i < n-1; i++ {
		if s[i] == minChar || s[i] == maxChar {
			return 2
		}
	}
	return 3
}

func main() {
	// result: 1
	// s := "dog"

	// result: 2
	// s := "card"

	// result: -1
	s := "gf"

	// result: 0
	// s := ""

	result := minOperations(s)
	fmt.Printf("result = %v\n", result)
}
