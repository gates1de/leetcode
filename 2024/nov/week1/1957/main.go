package main
import (
	"fmt"
)

func makeFancyString(s string) string {
	result := make([]rune, 0, len(s))

	for _, char := range s {
		length := len(result)

		if length <= 1 {
			result = append(result, char)
			continue
		}

		if result[length - 2] == char && result[length - 1] == char {
			continue
		}

		result = append(result, char)
	}

	return string(result)
}

func main() {
	// result: "leetcode"
	// s := "leeetcode"

	// result: "aabaa"
	// s := "aaabaaaa"

	// result: "aab"
	s := "aab"

	// result: ""
	// s := ""

	result := makeFancyString(s)
	fmt.Printf("result = %v\n", result)
}

