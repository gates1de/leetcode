package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func countPermutations(complexity []int) int {
	n := len(complexity)
	for i := 1; i < n; i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
	}

	result := int(1)
	for i := 2; i < n; i++ {
		result = result * i % modulo
	}

	return result
}

func main() {
	// result: 2
	// complexity := []int{1,2,3}

	// result: 0
	complexity := []int{3,3,3,4,4,4}

	// result: 
	// complexity := []int{}

	result := countPermutations(complexity)
	fmt.Printf("result = %v\n", result)
}

