package main

import (
	"fmt"
)

func minMoves(nums []int) int {
	minValue := nums[0]
	sum := int64(0)
	for _, value := range nums {
		minValue = min(minValue, value)
		sum += int64(value)
	}

	result := sum - int64(minValue)*int64(len(nums))
	return int(result)
}

func main() {
	// result: 3
	// nums := []int{1,2,3}

	// result: 0
	nums := []int{1, 1, 1}

	// result: 0
	// nums := []int{}

	result := minMoves(nums)
	fmt.Printf("result = %v\n", result)
}
