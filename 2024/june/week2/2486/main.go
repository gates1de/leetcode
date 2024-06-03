package main
import (
	"fmt"
)

func appendCharacters(s string, t string) int {
	first := int(0)
	longestPrefix := int(0)

	for first < len(s) && longestPrefix < len(t) {
		if s[first] == t[longestPrefix] {
			longestPrefix++
		}
		first++
	}

	return len(t) - longestPrefix
}

func main() {
	// result: 4
	// s := "coaching"
	// t := "coding"

	// result: 0
	// s := "abcde"
	// t := "a"

	// result: 5
	s := "z"
	t := "abcde"

	// result: 
	// s := ""
	// t := ""

	result := appendCharacters(s, t)
	fmt.Printf("result = %v\n", result)
}

