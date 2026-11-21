package main

import (
	"fmt"
)

func minChanges(s string) int {
	result := 0

	for i := 1; i < len(s); i += 2 {
		if s[i-1] != s[i] {
			result++
		}
	}

	return result
}

func main() {
	// result: 2
	// s := "1001"

	// result: 1
	// s := "10"

	// result: 0
	s := "00"

	// result:
	// s := ""

	result := minChanges(s)
	fmt.Printf("result = %v\n", result)
}
