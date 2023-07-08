package main
import (
	"fmt"
)

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		sMap := make(map[rune]int)
		for _, char := range s {
			sMap[char]++
			if sMap[char] >= 2 {
				return true
			}
		}

		return false
	}

	sSwapTargetChar := byte(0)
	goalSwapTargetChar := byte(0)
	isBuddy := false
	notEqualCount := int(0)
	for i := 0; i < len(s); i++ {
		sChar := s[i]
		goalChar := goal[i]

		if sChar != goalChar {
			notEqualCount++
			if notEqualCount >= 3 {
				return false
			}

			if sSwapTargetChar == 0 {
				sSwapTargetChar = sChar
				goalSwapTargetChar = goalChar
			} else {
				isBuddy = sSwapTargetChar == goalChar && goalSwapTargetChar == sChar
			}
		}
	}

	if notEqualCount == 1 {
		return false
	}
	return isBuddy
}

func main() {
	// result: true
	// s := "ab"
	// goal := "ba"

	// result: false
	// s := "ab"
	// goal := "ab"

	// result: true
	s := "aa"
	goal := "aa"

	// result: 
	// s := ""
	// goal := ""

	result := buddyStrings(s, goal)
	fmt.Printf("result = %v\n", result)
}

