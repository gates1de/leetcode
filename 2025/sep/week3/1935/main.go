package main

import (
	"fmt"
)

func canBeTypedWords(text string, brokenLetters string) int {
	m := make(map[rune]bool)
	for _, char := range brokenLetters {
		m[char] = true
	}

	result := int(0)
	isBroken := false
	for _, char := range text {
		if char == ' ' {
			if !isBroken {
				result++
			}

			isBroken = false
			continue
		}

		if m[char] {
			isBroken = true
		}
	}

	if !isBroken {
		result++
	}

	return result
}

func main() {
	// result: 1
	// text := "hello world"
	// brokenLetters := "ad"

	// result: 1
	// text := "leet code"
	// brokenLetters := "lt"

	// result: 0
	text := "leet code"
	brokenLetters := "e"

	// result: 
	// text := ""
	// brokenLetters := ""

	result := canBeTypedWords(text, brokenLetters)
	fmt.Printf("result = %v\n", result)
}

