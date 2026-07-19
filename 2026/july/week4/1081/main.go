package main

import (
	"fmt"
)

func smallestSubsequence(s string) string {
	var remaining [26]int
	for i := range s {
		remaining[s[i] - 'a']++
	}

	var inStack [26]bool
	stack := make([]byte, 0, 26)
	for i := range s {
		ch := s[i]
		index := ch - 'a'
		remaining[index]--

		if inStack[index] {
			continue
		}

		for len(stack) > 0 {
			last := stack[len(stack) - 1]
			if last <= ch || remaining[last - 'a'] == 0 {
				break
			}

			inStack[last - 'a'] = false
			stack = stack[:len(stack) - 1]
		}

		stack = append(stack, ch)
		inStack[index] = true
	}

	return string(stack)
}

func main() {
	// result: "acdb"
	// s := "cbacdcbc"

	// result: "abc"
	// s := "bcabc"

	// result: "acdb"
	s := "cbacdcbc"

	// result: ""
	// s := ""

	result := smallestSubsequence(s)
	fmt.Printf("result = %v\n", result)
}
