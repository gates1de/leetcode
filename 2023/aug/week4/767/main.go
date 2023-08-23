package main
import (
	"fmt"
)

func reorganizeString(s string) string {
	charCounts := make([]int, 26)
	for _, c := range s {
		charCounts[c - 97]++
	}

	maxCount := int(0)
	letter := int(0)
	for i, count := range charCounts {
		if count > maxCount {
			maxCount = count
			letter = i
		}
	}

	resultChars := make([]byte, len(s))
	if maxCount > (len(s) + 1) / 2 {
		return ""
	}

	index := int(0)
	for charCounts[letter] != 0 {
		resultChars[index] = byte(letter + 97)
		index += 2
		charCounts[letter]--
	}

	for i, _ := range charCounts {
		for charCounts[i] > 0 {
			if index >= len(s) {
				index = 1
			}
			resultChars[index] = byte(i + 97)
			index += 2
			charCounts[i]--
		}
	}

	return string(resultChars)
}

func main() {
	// result: "aba"
	// s := "aab"

	// result: ""
	// s := "aaab"

	// result: "acabc"
	s := "aabcc"

	// result: 
	// s := ""

	result := reorganizeString(s)
	fmt.Printf("result = %v\n", result)
}

