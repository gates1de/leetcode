package main

import (
	"fmt"
)

func mirrorDistance(n int) int {
	diff := n - reverse(n)
	if diff < 0 {
		return -diff
	}
	return diff
}

func reverse(n int) int {
	result := int(0)
	for n > 0 {
		result = result * 10 + n % 10
		n /= 10
	}
	return result
}

func main() {
	// result: 27
	// n := int(25)

	// result: 9
	// n := int(10)

	// result: 0
	n := int(7)

	// result: 
	// n := int(0)

	result := mirrorDistance(n)
	fmt.Printf("result = %v\n", result)
}

