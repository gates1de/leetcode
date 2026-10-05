package main

import (
	"fmt"
)

func scoreOfParentheses(s string) int {
	result := 0
	depth := 0
	for i := range s {
		if s[i] == '(' {
			depth++
		} else {
			depth--
			if i > 0 && s[i-1] == '(' {
				result += 1 << depth
			}
		}
	}
	return result
}

func main() {
	// result: 1
	// s := "()"

	// result: 2
	// s := "(())"

	// result: 2
	// s := "()()"

	// result: 3
	s := "(())()"

	// result:
	// s := ""

	result := scoreOfParentheses(s)
	fmt.Printf("result = %v\n", result)
}
