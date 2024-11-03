package main
import (
	"fmt"
)

func rotateString(s string, goal string) bool {
	if len(s) == 0 || len(goal) == 0 {
		return false
	}

	firstChar := rune(s[0])
	for i, char := range goal {
		if firstChar != char {
			continue
		}

		goalChars := []byte(goal[i:])
		goalChars = append(goalChars, []byte(goal[:i])...)
		if s == string(goalChars) {
			return true
		}
	}

	return false
}

func main() {
	// result: true
	// s := "abcde"
	// goal := "cdeab"

	// result: false
	s := "abcde"
	goal := "abced"

	// result: false
	// s := ""
	// goal := ""

	result := rotateString(s, goal)
	fmt.Printf("result = %v\n", result)
}

