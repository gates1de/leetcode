package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	lastBit := n % 2
	n /= 2

	for n > 0 {
		if n % 2 == lastBit {
			return false
		}
		lastBit = n % 2
		n /= 2
	}

	return true
}

func main() {
	// result: true
	// n := int(5)

	// result: false
	// n := int(7)

	// result: false
	n := int(11)

	// result: 
	// n := int(0)

	result := hasAlternatingBits(n)
	fmt.Printf("result = %v\n", result)
}

