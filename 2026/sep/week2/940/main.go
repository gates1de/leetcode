package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func distinctSubseqII(s string) int {
	var ending [26]int
	total := int(1)

	for i := range len(s) {
		index := s[i] - 'a'
		previous := total
		total = (2 * total - ending[index] + modulo) % modulo
		ending[index] = previous
	}

	return (total + modulo - 1) % modulo
}

func main() {
	// result: 7
	// s := "abc"

	// result: 6
	// s := "aba"

	// result: 3
	s := "aaa"

	// result: 0
	// s := ""

	result := distinctSubseqII(s)
	fmt.Printf("result = %v\n", result)
}
