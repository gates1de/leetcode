package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	result := 0
	left, right := 0, 0
	for i := range s {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right && 2*right > result {
			result = 2 * right
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if left == right && 2*left > result {
			result = 2 * left
		} else if left > right {
			left, right = 0, 0
		}
	}

	return result
}

func main() {
	// result: 2
	// s := "(()"

	// result: 4
	// s := ")()())"

	// result: 0
	// s := ""

	// result: 2
	// s := "()(()"

	// result: 6
	// s := "()(())"

	// result: 22
	// s := ")(((((()())()()))()(()))("

	// result: 2
	// s := "))))((()(("

	// result: 4
	// s := ")()())()()("

	// result: 2
	// s := "(()(((()"

	// result: 4
	s := "(()()"

	// result:
	// s := ""

	result := longestValidParentheses(s)
	fmt.Printf("result = %v\n", result)
}
