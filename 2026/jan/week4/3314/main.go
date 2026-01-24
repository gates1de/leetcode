package main

import (
	"fmt"
)

func minBitwiseArray(nums []int) []int {
	result := make([]int, len(nums))

	for i, num := range nums {
		candidate := int(-1)

		for j := 1; j < num; j++ {
			if (j | (j + 1)) == num {
				candidate = j
				break
			}
		}

		result[i] = candidate
	}

	return result
}

func main() {
	// result: [-1,1,4,3]
	// nums := []int{2,3,5,7}

	// result: [9,12,15]
	nums := []int{11,13,31}

	// result: 
	// nums := []int{}

	result := minBitwiseArray(nums)
	fmt.Printf("result = %v\n", result)
}

