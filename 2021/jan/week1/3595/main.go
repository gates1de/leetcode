package main
import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	result := int(0)
	characters := strings.Split(s, "")
	substring := ""
	for i := 0; i < len(characters); i++ {
		c := characters[i]
		cIndex := strings.Index(substring, c)
		if cIndex >= 0 {
			substring = substring[cIndex + 1:] + c
		} else {
			substring += c
		}

		if result < len(substring) {
			result = len(substring)
		}
	}
	return result
}

func main() {
	// s := "abcabcbb"
    // s := "pwwkew"
	// s := "dvdf"
	s := "anviaj"
	result := lengthOfLongestSubstring(s)
	fmt.Printf("result = %v\n", result)
}

