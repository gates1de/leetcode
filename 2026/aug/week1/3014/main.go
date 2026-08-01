package main

import (
	"fmt"
)

const keys = int(8)

func minimumPushes(word string) int {
	result := int(0)
	for index := range word {
		result += index / keys + 1
	}

	return result
}

func main() {
	// result: 5
	// word := "abcde"

	// result: 12
	word := "xycdefghij"

	// result:
	// word := ""

	result := minimumPushes(word)
	fmt.Printf("result = %v\n", result)
}
