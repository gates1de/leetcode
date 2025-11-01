package main

import (
	"fmt"
)

func smallestNumber(n int) int {
	result := int(1)
	for result < n {
		result = result * 2 + 1
	}

	return result
}

func main() {
	// result: 7
	// n := int(5)

	// result: 15
	// n := int(10)

	// result: 3
	n := int(3)

	// result: 
	// n := int(0)

	result := smallestNumber(n)
	fmt.Printf("result = %v\n", result)
}

