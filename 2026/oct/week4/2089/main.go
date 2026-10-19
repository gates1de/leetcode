package main

import (
	"fmt"
)

func targetIndices(nums []int, target int) []int {
	less := 0
	equal := 0
	for _, value := range nums {
		if value < target {
			less++
		} else if value == target {
			equal++
		}
	}

	result := make([]int, equal)
	for i := range result {
		result[i] = less + i
	}
	return result
}

func main() {
	// result: [1,2]
	// nums := []int{1,2,5,2,3}
	// target := int(2)

	// result: [3]
	// nums := []int{1,2,5,2,3}
	// target := int(3)

	// result: [4]
	nums := []int{1, 2, 5, 2, 3}
	target := int(5)

	// result: []
	// nums := []int{}
	// target := int(0)

	result := targetIndices(nums, target)
	fmt.Printf("result = %v\n", result)
}
