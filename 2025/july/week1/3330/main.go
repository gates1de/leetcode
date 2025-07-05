package main

import (
	"fmt"
)

func possibleStringCount(word string) int {
	result := int(1)

	for i := 1; i < len(word); i++ {
		if word[i - 1] == word[i] {
			result++
		}
	}

	return result
}

func main() {
	// result: 5
	// word := "abbcccc"

	// result: 1
	// word := "abcd"

	// result: 4
	word := "aaaa"

	// result: 
	// word := ""

	result := possibleStringCount(word)
	fmt.Printf("result = %v\n", result)
}

