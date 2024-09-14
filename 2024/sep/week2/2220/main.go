package main
import (
	"fmt"
)

func minBitFlips(start int, goal int) int {
	xor := start ^ goal
	result := int(0)

	for xor != 0 {
		xor &= xor - 1
		result++
	}

	return result
}

func main() {
	// result: 3
	// start := int(10)
	// goal := int(7)

	// result: 3
	start := int(3)
	goal := int(4)

	// result: 
	// start := int(0)
	// goal := int(0)

	result := minBitFlips(start, goal)
	fmt.Printf("result = %v\n", result)
}

