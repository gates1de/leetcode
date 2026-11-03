package main

import (
	"fmt"
)

func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxValue := nums[0]
	for _, value := range nums[1:] {
		maxValue = max(maxValue, value)
	}

	result := 0
	for _, value := range nums {
		result += maxValue - value
	}

	return result
}

func main() {
	// result: 3
	// nums := []int{2,1,3}

	// result: 2
	nums := []int{4, 4, 5}

	// result: 0
	// nums := []int{}

	result := minMoves(nums)
	fmt.Printf("result = %v\n", result)
}
