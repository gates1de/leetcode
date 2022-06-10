package main
import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	result := int(0)
	sub := ""
	for _, r := range s {
		char := string(r)
		index := strings.Index(sub, char)
		if index == -1 {
			sub += char
			continue
		}

		if result < len(sub) {
			result = len(sub)
		}
		sub = string(sub[index + 1:]) + char
	}

	if result < len(sub) {
		result = len(sub)
	}

	return result
}

func main() {
	// result: 3
	// s := "abcabcbb"

	// result: 1
	// s := "bbbbb"

	// result: 3
	// s := "pwwkew"

	// result: 3
	// s := "dvdf"

	// result: 0
	// s := ""

	// result: 5
	s := "anviaj"

	// result: 
	// s := ""

	result := lengthOfLongestSubstring(s)
	fmt.Printf("result = %v\n", result)
}

