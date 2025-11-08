package main

import (
	"fmt"
)

func minimumOneBitOperations(n int) int {
	n ^= n >> 16
	n ^= n >> 8
	n ^= n >> 4
	n ^= n >> 2
	n ^= n >> 1
	return n
}

func main() {
	// result: 2
	// n := int(3)

	// result: 4
	n := int(6)

	// result: 
	// n := int(0)

	result := minimumOneBitOperations(n)
	fmt.Printf("result = %v\n", result)
}

