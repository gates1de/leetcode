package main

import (
	"fmt"
)

func checkString(s string) bool {
	seenB := false
	for _, character := range s {
		if character == 'b' {
			seenB = true
		} else if seenB {
			return false
		}
	}

	return true
}

func main() {
	// result: true
	// s := "aaabbb"

	// result: false
	// s := "abab"

	// result: true
	s := "bbb"

	// result: false
	// s := ""

	result := checkString(s)
	fmt.Printf("result = %v\n", result)
}
