package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := make([]byte, 0, len(s))
	for i := range s {
		current := s[i]
		if current == '(' || current == '{' || current == '[' {
			stack = append(stack, current)
			continue
		}
		if len(stack) == 0 {
			return false
		}

		last := stack[len(stack)-1]
		if last == '(' && current != ')' ||
			last == '{' && current != '}' ||
			last == '[' && current != ']' {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func main() {
	// result: true
	// s := "()"

	// result: true
	// s := "()[]{}"

	// result: false
	s := "(]"

	// result:
	// s := ""

	result := isValid(s)
	fmt.Printf("result = %v\n", result)
}
