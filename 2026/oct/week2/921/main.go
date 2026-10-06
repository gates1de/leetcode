package main

import (
	"fmt"
)

func minAddToMakeValid(s string) int {
	open := 0
	result := 0
	for i := range s {
		if s[i] == '(' {
			open++
		} else if open > 0 {
			open--
		} else {
			result++
		}
	}

	return result + open
}

func main() {
	// result: 1
	// s := "())"

	// result: 3
	s := "((("

	// result:
	// s := ""

	result := minAddToMakeValid(s)
	fmt.Printf("result = %v\n", result)
}
