package main
import (
	"fmt"
)

func makeGood(s string) string {
	chars := []byte(s)
	stack := make([]byte, 0, len(chars))

	for i := 0; i < len(chars); i++ {
		char := chars[i]
		if len(stack) > 0 {
			last := stack[len(stack) - 1]

			if last - byte(32) == char || last + byte(32) == char {
				stack = stack[:len(stack) - 1]
				continue
			}
		}
		stack = append(stack, char)
	}

	return string(stack)
}

func main() {
	// result: "leetcode"
	// s := "leEeetcode"

	// result: ""
	// s := "abBAcC"

	// result: "s"
	s := "s"

	// result: 
	// s := ""

	result := makeGood(s)
	fmt.Printf("result = %v\n", result)
}

