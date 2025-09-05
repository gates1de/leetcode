package main

import (
	"fmt"
)

func makeTheIntegerZero(num1 int, num2 int) int {
	k := int(1)

	for {
		x := int64(num1) - int64(num2) * int64(k)

		if x < int64(k) {
			return -1
		}

		if k >= bitCount(x) {
			return k
		}

		k++
	}
}

func bitCount(n int64) int {
	result := int(0)
	for n != 0 {
		result++
		n &= n - 1
	}
	return result
}

func main() {
	// result: 3
	// num1 := int(3)
	// num2 := int(-2)

	// result: -1
	num1 := int(5)
	num2 := int(7)

	// result: 
	// num1 := int(0)
	// num2 := int(0)

	result := makeTheIntegerZero(num1, num2)
	fmt.Printf("result = %v\n", result)
}

