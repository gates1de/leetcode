package main

import (
	"fmt"
)

const inf = int(1 << 30)

func numberOfSpecialChars(word string) int {
	lastLower := [26]int{}
	firstUpper := [26]int{}
	for i := range 26 {
		lastLower[i] = int(-1)
		firstUpper[i] = inf
	}

	for i := 0; i < len(word); i++ {
		ch := word[i]
		if ch >= 'a' && ch <= 'z' {
			idx := ch - 'a'
			lastLower[idx] = i
		} else {
			idx := ch - 'A'
			if firstUpper[idx] == inf {
				firstUpper[idx] = i
			}
		}
	}

	result := int(0)
	for i := range 26 {
		if lastLower[i] != -1 && firstUpper[i] != inf && lastLower[i] < firstUpper[i] {
			result++
		}
	}

	return result
}

func main() {
	// result: 3
	// word := "aaAbcBC"

	// result: 0
	// word := "abc"

	// result: 0
	word := "AbBCab"

	// result: 0
	// word := ""

	result := numberOfSpecialChars(word)
	fmt.Printf("result = %v\n", result)
}
