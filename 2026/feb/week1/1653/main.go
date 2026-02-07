package main

import (
	"fmt"
)

func minimumDeletions(s string) int {
	n := len(s)
	bCount := int(0)
	result := int(0)

	for i := range n {
		if s[i] == 'b' {
			bCount++
		} else {
			result = min(result + 1, bCount)
		}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// s := "aababbab"

	// result: 2
	s := "bbaaaaabb"

	// result: 
	// s := ""

	result := minimumDeletions(s)
	fmt.Printf("result = %v\n", result)
}

