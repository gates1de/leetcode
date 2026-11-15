package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numOfSubarrays(arr []int) int {
	result := 0
	prefixParity := 0
	oddCount := 0
	evenCount := 1

	for _, num := range arr {
		prefixParity ^= num & 1

		if prefixParity == 0 {
			result += oddCount
			evenCount++
		} else {
			result += evenCount
			oddCount++
		}

		if result >= modulo {
			result -= modulo
		}
	}

	return result
}

func main() {
	// result: 4
	// arr := []int{1,3,5}

	// result: 0
	// arr := []int{2,4,6}

	// result: 16
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	// result:
	// arr := []int{}

	result := numOfSubarrays(arr)
	fmt.Printf("result = %v\n", result)
}
