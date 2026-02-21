package main

import (
	"fmt"
	"math/bits"
)

func countPrimeSetBits(left int, right int) int {
    result := int(0)
    for i := left; i <= right; i++ {
			if isSmallPrime(bits.OnesCount(uint(i))) {
				result++
			}
    }

    return result
}

func isSmallPrime(n int) bool {
	return (n == 2 || n == 3 || n == 5 || n == 7 ||
		n == 11 || n == 13 || n == 17 || n == 19)
}

func main() {
	// result: 4
	// left := int(6)
	// right := int(10)

	// result: 5
	left := int(10)
	right := int(15)

	// result: 
	// left := int(0)
	// right := int(0)

	result := countPrimeSetBits(left, right)
	fmt.Printf("result = %v\n", result)
}

