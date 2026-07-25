package main

import (
	"fmt"
	"strconv"
)

func maxProduct(n int) int {
	maxDigit := int(-1)
	secondMaxDigit := int(-1)

	for _, character := range strconv.Itoa(n) {
		digit := int(character - '0')
		if digit >= maxDigit {
			secondMaxDigit = maxDigit
			maxDigit = digit
		} else if digit > secondMaxDigit {
			secondMaxDigit = digit
		}
	}

	result := maxDigit * secondMaxDigit
	return result
}

func main() {
	// result: 3
	// n := int(31)

	// result: 4
	// n := int(22)

	// result: 8
	n := int(124)

	// result:
	// n := int(0)

	result := maxProduct(n)
	fmt.Printf("result = %v\n", result)
}
