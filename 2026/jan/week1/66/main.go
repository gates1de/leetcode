package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)
	if n == 0 {
		return []int{1}
	}

	carry := int(1)
	for i := n - 1; i >= 0; i-- {
		num := digits[i]

		if i == 0 {
			if num + carry >= 10 {
				result := make([]int, n + 1)
				copy(result[1:], digits)
				result[0] = (num + carry) / 10 
				result[1] = (num + carry) % 10
				return result
			}

			digits[i] += carry
			return digits
		}

		digits[i] = (num + carry) % 10
		carry = (num + carry) / 10
	}

	return digits
}

func main() {
	// result: [1,2,4]
	// digits := []int{1,2,3}

	// result: [4,3,2,2]
	// digits := []int{4,3,2,1}

	// result: [1,0]
	digits := []int{9}

	// result: []
	// digits := []int{}

	result := plusOne(digits)
	fmt.Printf("result = %v\n", result)
}

