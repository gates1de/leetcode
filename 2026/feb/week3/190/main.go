package main

import (
	"fmt"
)

func reverseBits(n int) int {
	if n == 0 {
		return 0
	}

	result := int(0)
	for range 32 {
		result <<= 1
		if ((n & 1) == 1) {
			result++
		}

		n >>= 1
	}

	return result
}

func main() {
	// result: 964176192
	// n := int(43261596)

	// result: 1073741822
	n := int(2147483644)

	// result: 
	// n := int(0)

	result := reverseBits(n)
	fmt.Printf("result = %v\n", result)
}

