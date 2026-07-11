package main

import (
	"fmt"
)

func sumAndMultiply(n int) int64 {
	if n == 0 {
		return 0
	}

	digits := make([]int, 0, 10)
	sum := int64(0)

	for n > 0 {
		digit := n % 10
		if digit != 0 {
			sum += int64(digit)
			digits = append(digits, digit)
		}

		n /= 10
	}

	x := int64(0)
	for i := len(digits) - 1; i >= 0; i-- {
		x = x*10 + int64(digits[i])
	}

	return x * sum
}

func main() {
	// result: 12340
	// n := int(10203004)

	// result: 1
	n := int(1000)

	// result:
	// n := int(0)

	result := sumAndMultiply(n)
	fmt.Printf("result = %v\n", result)
}
