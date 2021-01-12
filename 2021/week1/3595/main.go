package main
import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	splitS := split(s)
	longestIndex := int(0)
	for i := 0; i < len(splitS); i++ {
		substring := splitS[i]
		if len(substring) >= len(splitS[longestIndex]) {
			longestIndex = i
		}
	}

	longestSubstring := splitS[longestIndex]
    // fmt.Printf("longestSubstring = %v\n", longestSubstring)
	result := len(longestSubstring)
	if longestIndex == 0 {
		return result
	} else {
		preSubstring := splitS[longestIndex - 1]
		psCharacters := strings.Split(preSubstring, "")
		// fmt.Printf("preSubstring = %v\n", preSubstring)
		for i := len(psCharacters) - 1; i >= 0; i-- {
			c := psCharacters[i]
			// fmt.Printf("c = %v\n", c)
			if strings.Contains(longestSubstring, c) {
				break
			}
			result++
		}
	}
	return result
}

func split(s string) []string {
	result := []string{}

	resultIndex := int(0)
	characters := strings.Split(s, "")
	for i := 0; i < len(characters); i++ {
		c := characters[i]
		if len(result) == 0 {
			result = append(result, c)
			continue
		}
		substring := result[resultIndex]
		if strings.Contains(substring, c) {
			result = append(result, c)
			resultIndex++
		} else {
			result[resultIndex] = substring + c
		}
		// fmt.Printf("result = %v\n", result)
	}
	return result
}

func main() {
	// s := "abcabcbb"
	// s := "pwwkew"
	s := "dvdf"
	result := lengthOfLongestSubstring(s)
	fmt.Printf("result = %v\n", result)
}

