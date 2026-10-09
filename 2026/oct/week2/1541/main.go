package main

import (
	"fmt"
)

func minInsertions(s string) int {
	result := 0
	need := 0
	for i := range s {
		if s[i] == '(' {
			if need&1 == 1 {
				result++
				need--
			}
			need += 2
		} else {
			need--
			if need < 0 {
				result++
				need = 1
			}
		}
	}

	return result + need
}

func main() {
	// result: 1
	// s := "(()))"

	// result: 0
	// s := "())"

	// result: 3
	s := "))())("

	// result: 0
	// s := ""

	result := minInsertions(s)
	fmt.Printf("result = %v\n", result)
}
