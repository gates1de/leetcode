package main

import (
	"fmt"
)

func reverseParentheses(s string) string {
	pairs := make([]int, len(s))
	top := -1
	for i := range s {
		if s[i] == '(' {
			pairs[i] = top
			top = i
		} else if s[i] == ')' {
			open := top
			top = pairs[open]
			pairs[open] = i
			pairs[i] = open
		}
	}

	result := make([]byte, 0, len(s))
	for i, direction := 0, 1; i < len(s); i += direction {
		if s[i] == '(' || s[i] == ')' {
			i = pairs[i]
			direction = -direction
			continue
		}
		result = append(result, s[i])
	}

	return string(result)
}

func main() {
	// result: "dcba"
	// s := "(abcd)"

	// result: "iloveu"
	// s := "(u(love)i)"

	// result: "leetcode"
	s := "(ed(et(oc))el)"

	// result: ""
	// s := ""

	result := reverseParentheses(s)
	fmt.Printf("result = %v\n", result)
}
