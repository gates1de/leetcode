package main

import (
	"fmt"
)

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	count := int(1)
	preCount := int(0)
	result := int(0)

	for i := 1; i < n; i++ {
		if nums[i] > nums[i - 1] {
			count++
		} else {
			preCount = count
			count = 1
		}

		result = max(result, min(preCount, count))
		result = max(result, count / 2)
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// nums := []int{2,5,7,8,9,2,3,4,3,1}

	// result: 2
	nums := []int{1,2,3,4,4,4,4,5,6,7}

	// result: 
	// nums := []int{}

	result := maxIncreasingSubarrays(nums)
	fmt.Printf("result = %v\n", result)
}

