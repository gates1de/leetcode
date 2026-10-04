package main

import (
	"fmt"
)

func checkValidString(s string) bool {
	minimum := 0
	maximum := 0
	for i := range s {
		switch s[i] {
		case '(':
			minimum++
			maximum++
		case ')':
			if minimum > 0 {
				minimum--
			}
			maximum--
		default:
			if minimum > 0 {
				minimum--
			}
			maximum++
		}

		if maximum < 0 {
			return false
		}
	}

	return minimum == 0
}

func main() {
	// result: true
	// s := "()"

	// result: true
	// s := "(*)"

	// result: true
	s := "(*))"

	// result:
	// s := ""

	result := checkValidString(s)
	fmt.Printf("result = %v\n", result)
}
