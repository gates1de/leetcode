package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numSub(s string) int {
	result := int(0)
	consecutive := int(0)
	for _, char := range s {
		if char == '0' {
			result += consecutive * (consecutive + 1) / 2
			result %= modulo
			consecutive = 0
		} else {
			consecutive++
		}
	}

	result += consecutive * (consecutive + 1) / 2
	result %= modulo
	return result
}

func main() {
	// result: 9
	// s := "0110111"

	// result: 2
	// s := "101"

	// result: 21
	s := "111111"

	// result: 
	// s := ""

	result := numSub(s)
	fmt.Printf("result = %v\n", result)
}

