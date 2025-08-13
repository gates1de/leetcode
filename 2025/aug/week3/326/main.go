package main
import (
	"fmt"
)

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	} else if n == 1 {
		return true
	} else if n % 3 != 0 {
		return false
	}

	if n / 3 <= 1 {
		return n % 3 == 0
	}
	return isPowerOfThree((n / 3) + (n % 3))
}

func main() {
	// result: true
	// n := int(27)

	// result: false
	// n := int(0)

	// result: false
	// n := int(45)

	// result: true
	// n := int(1)

	// result: false
	// n := int(19684)

	// result: false
	n := int(2147483647)

	// result: 
	// n := int(0)

	result := isPowerOfThree(n)
	fmt.Printf("result = %v\n", result)
}

