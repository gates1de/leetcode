package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numOfWays(n int) int {
	alternatingPatterns := int(6)
	allDifferentPatterns := int(6)

	for i := 2; i <= n; i++ {
		newAlternating := (3 * alternatingPatterns + 2 * allDifferentPatterns) % modulo
		newAllDifferent := (2 * alternatingPatterns + 2 * allDifferentPatterns) % modulo
		alternatingPatterns = newAlternating
		allDifferentPatterns = newAllDifferent
	}

	return (alternatingPatterns + allDifferentPatterns) % modulo
}

func main() {
	// result: 12
	// n := int(1)

	// result: 30228214
	n := int(5000)

	// result: 
	// n := int(0)

	result := numOfWays(n)
	fmt.Printf("result = %v\n", result)
}

