package main

import (
	"fmt"
)

func sortPermutation(nums []int) int {
	result := 0
	found := false

	for i, value := range nums {
		if i == value {
			continue
		}

		if !found {
			result = value
			found = true
			continue
		}

		result &= value
	}

	return result
}

func main() {
	// result: 1
	// nums := []int{0,3,2,1}

	// result: 2
	// nums := []int{0,1,3,2}

	// result: 0
	nums := []int{3, 2, 1, 0}

	// result: 0
	// nums := []int{}

	result := sortPermutation(nums)
	fmt.Printf("result = %v\n", result)
}
