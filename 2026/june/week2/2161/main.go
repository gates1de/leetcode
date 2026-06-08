package main

import (
	"fmt"
)

func pivotArray(nums []int, pivot int) []int {
	result := make([]int, len(nums))
	index := int(0)

	for _, num := range nums {
		if num < pivot {
			result[index] = num
			index++
		}
	}

	for _, num := range nums {
		if num == pivot {
			result[index] = num
			index++
		}
	}

	for _, num := range nums {
		if num > pivot {
			result[index] = num
			index++
		}
	}

	return result
}

func main() {
	// result: [9,5,3,10,10,12,14]
	// nums := []int{9,12,5,10,14,3,10}
	// pivot := int(10)

	// result: [-3,2,4,3]
	nums := []int{-3,4,3,2}
	pivot := int(2)

	// result: []
	// nums := []int{}
	// pivot := int(0)

	result := pivotArray(nums, pivot)
	fmt.Printf("result = %v\n", result)
}
