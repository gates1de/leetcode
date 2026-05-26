package main

import (
	"fmt"
	"math/bits"
)

func numberOfSpecialChars(word string) int {
	var lowerMask uint32
	var upperMask uint32

	for i := range len(word) {
		ch := word[i]
		if ch >= 'a' && ch <= 'z' {
			lowerMask |= 1 << (ch - 'a')
		} else {
			upperMask |= 1 << (ch - 'A')
		}
	}

	return bits.OnesCount32(lowerMask & upperMask)
}

func main() {
	// result: 3
	// word := "aaAbcBC"

	// result: 0
	// word := "abc"

	// result: 1
	word := "abBCab"

	// result: 0
	// word := ""

	result := numberOfSpecialChars(word)
	fmt.Printf("result = %v\n", result)
}
