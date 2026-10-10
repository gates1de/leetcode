package main

import (
	"fmt"
)

func removeOuterParentheses(s string) string {
	result := make([]byte, 0, len(s))
	depth := 0
	for i := range s {
		if s[i] == '(' {
			if depth > 0 {
				result = append(result, s[i])
			}
			depth++
		} else {
			depth--
			if depth > 0 {
				result = append(result, s[i])
			}
		}
	}
	return string(result)
}

func main() {
	// result: "()()()"
	// s := "(()())(())"

	// result: "()()()()(())"
	// s := "(()())(())(()(()))"

	// result: ""
	s := "()()"

	// result: ""
	// s := ""

	result := removeOuterParentheses(s)
	fmt.Printf("result = %v\n", result)
}
