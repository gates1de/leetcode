package main

import (
	"fmt"
)

func minimumOperations(nums []int) int {
	result := int(0)
	for _, x := range nums {
		remainder := x % 3
		result += min(remainder, 3 - remainder)
	}

	return result
}

func main() {
	// result: 3
	// nums := []int{1,2,3,4}

	// result: 0
	nums := []int{3,6,9}

	// result: 
	// nums := []int{}

	result := minimumOperations(nums)
	fmt.Printf("result = %v\n", result)
}

