package main

import (
	"fmt"
)

func uniqueXorTriplets(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	result := int(1)
	for result <= n {
		result <<= 1
	}
	return result
}

func main() {
	// result: 2
	// nums := []int{1,2}

	// result: 4
	nums := []int{3, 1, 2}

	// result:
	// nums := []int{}

	result := uniqueXorTriplets(nums)
	fmt.Printf("result = %v\n", result)
}
