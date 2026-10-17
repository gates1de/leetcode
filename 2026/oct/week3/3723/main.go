package main

import (
	"fmt"
)

func maxSumOfSquares(num int, sum int) string {
	if sum > 9*num {
		return ""
	}

	result := make([]byte, num)
	remaining := sum
	for i := range result {
		digit := remaining
		if digit > 9 {
			digit = 9
		}
		result[i] = byte('0' + digit)
		remaining -= digit
	}

	return string(result)
}

func main() {
	// result: "30"
	// num := int(2)
	// sum := int(3)

	// result: "98"
	// num := int(2)
	// sum := int(17)

	// result: ""
	num := int(1)
	sum := int(10)

	// result: ""
	// num := int(0)
	// sum := int(0)

	result := maxSumOfSquares(num, sum)
	fmt.Printf("result = %v\n", result)
}
