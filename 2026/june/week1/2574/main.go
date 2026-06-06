package main

import (
	"fmt"
)

func leftRightDifference(nums []int) []int {
	total := int(0)
	for _, num := range nums {
		total += num
	}

	result := make([]int, len(nums))
	leftSum := int(0)
	for i, num := range nums {
		total -= num
		result[i] = leftSum - total
		if result[i] < 0 {
			result[i] = -result[i]
		}
		leftSum += num
	}

	return result
}

func main() {
	// result: [15,1,11,22]
	// nums := []int{10,4,8,3}

	// result: [0]
	nums := []int{1}

	// result: []
	// nums := []int{}

	result := leftRightDifference(nums)
	fmt.Printf("result = %v\n", result)
}
