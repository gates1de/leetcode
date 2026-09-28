package main

import (
	"fmt"
)

func maxDepth(s string) int {
	depth := 0
	result := 0
	for i := range s {
		if s[i] == '(' {
			depth++
			if depth > result {
				result = depth
			}
		} else if s[i] == ')' {
			depth--
		}
	}

	return result
}

func main() {
	// result: 3
	// s := "(1+(2*3)+((8)/4))+1"

	// result: 3
	s := "(1)+((2))+(((3)))"

	// result:
	// s := ""

	result := maxDepth(s)
	fmt.Printf("result = %v\n", result)
}
