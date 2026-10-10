package main

import (
	"fmt"
)

func removeInvalidParentheses(s string) []string {
	result := make([]string, 0)
	removeParentheses(s, 0, 0, '(', ')', &result)
	return result
}

func removeParentheses(s string, scanStart int, removeStart int, open byte, close byte, result *[]string) {
	balance := 0
	for i := scanStart; i < len(s); i++ {
		if s[i] == open {
			balance++
		} else if s[i] == close {
			balance--
		}

		if balance >= 0 {
			continue
		}

		for j := removeStart; j <= i; j++ {
			if s[j] == close && (j == removeStart || s[j-1] != close) {
				removeParentheses(s[:j]+s[j+1:], i, j, open, close, result)
			}
		}
		return
	}

	reversed := reverseString(s)
	if open == '(' {
		removeParentheses(reversed, 0, 0, ')', '(', result)
	} else {
		*result = append(*result, reversed)
	}
}

func reverseString(s string) string {
	bytes := []byte(s)
	for left, right := 0, len(bytes)-1; left < right; left, right = left+1, right-1 {
		bytes[left], bytes[right] = bytes[right], bytes[left]
	}
	return string(bytes)
}

func main() {
	// result: ["(())()","()()()"]
	// s := "()())()"

	// result: ["(a())()","(a)()()"]
	// s := "(a)())()"

	// result: [""]
	s := ")("

	// result: []
	// s := ""

	result := removeInvalidParentheses(s)
	fmt.Printf("result = %v\n", result)
}
