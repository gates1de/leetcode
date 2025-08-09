package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
    if n == 0 {
        return false
    } else if n == 1 {
        return true
    }

    if n % 2 != 0 {
        return false
    }

    return isPowerOfTwo(n / 2)
}

func main() {
	// result: true
	// n := int(1)

	// result: true
	// n := int(16)

	// result: false
	n := int(3)

	// result: 
	// n := int(0)

	result := isPowerOfTwo(n)
	fmt.Printf("result = %v\n", result)
}

