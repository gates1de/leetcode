package main

import (
	"fmt"
)

func sequentialDigits(low int, high int) []int {
	result := make([]int, 0, 36)

	for length := 2; length <= 9; length++ {
		for start := 1; start <= 10 - length; start++ {
			num := int(0)
			for digit := range length {
				num = num * 10 + start + digit
			}

			if num < low || num > high {
				continue
			}

			result = append(result, num)
		}
	}

	return result
}

func main() {
	// result: [123,234]
	// low := int(100)
	// high := int(300)

	// result: [1234,2345,3456,4567,5678,6789,12345]
	low := int(1000)
	high := int(13000)

	// result: 
	// low := int(0)
	// high := int(0)

	result := sequentialDigits(low, high)
	fmt.Printf("result = %v\n", result)
}
