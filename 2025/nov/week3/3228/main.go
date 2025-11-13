package main

import (
	"fmt"
)

func maxOperations(s string) int {
	countOne := int(0)
	result := int(0)
	i := int(0)

	for i < len(s) {
		if s[i] == '0' {
			for i + 1 < len(s) && s[i + 1] == '0' {
				i++
			}

			result += countOne
		} else {
			countOne++
		}

		i++
	}

	return result
}

func main() {
	// result: 4
	// s := "1001101"

	// result: 0
	s := "00111"

	// result: 
	// s := ""

	result := maxOperations(s)
	fmt.Printf("result = %v\n", result)
}

