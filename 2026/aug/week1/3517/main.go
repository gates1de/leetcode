package main

import (
	"fmt"
	"strings"
)

const alphabetSize = 26

func smallestPalindrome(s string) string {
	counts := [alphabetSize]int{}
	for _, character := range s {
		counts[character-'a']++
	}

	half := strings.Builder{}
	for character, count := range counts {
		half.WriteString(strings.Repeat(string(rune('a'+character)), count/2))
	}
	halfString := half.String()

	resultBuilder := strings.Builder{}
	resultBuilder.WriteString(halfString)
	for character, count := range counts {
		if count%2 == 1 {
			resultBuilder.WriteByte(byte('a' + character))
			break
		}
	}

	for index := range halfString {
		resultBuilder.WriteByte(halfString[len(halfString)-1-index])
	}

	return resultBuilder.String()
}

func main() {
	// result: "z"
	// s := "z"

	// result: "abbba"
	// s := "babab"

	// result: "acddca"
	s := "daccad"

	// result: ""
	// s := ""

	result := smallestPalindrome(s)
	fmt.Printf("result = %v\n", result)
}
