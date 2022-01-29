package main
import (
	"fmt"
)

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	initial := word[0]
	if byte('a') <= initial && initial <= byte('z') {
		for _, char := range word[1:] {
			b := byte(char)
			if byte('A') <= b && b <= byte('Z') {
				return false
			}
		}
		return true
	}

	next := word[1]
	if byte('a') <= next && next <= byte('z') {
		for _, char := range word[1:] {
			b := byte(char)
			if byte('A') <= b && b <= byte('Z') {
				return false
			}
		}
	} else {
		for _, char := range word[1:] {
			b := byte(char)
			if byte('a') <= b && b <= byte('z') {
				return false
			}
		}
	}

	return true
}

func main() {
	// result: true
	// word := "USA"

	// result: false
	// word := "FlaG"

	// result: true
	// word := "leetcode"

	// result: true
	// word := "London"

	// result: false
	// word := "tokyO"

	// result: true
	word := "T"

	// result: 
	// word := ""

	result := detectCapitalUse(word)
	fmt.Printf("result = %v\n", result)
}

