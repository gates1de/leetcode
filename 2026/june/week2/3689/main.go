package main

import (
	"fmt"
)

func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)
	if n == 0 || k == 0 {
		return int64(0)
	}

	minVal := nums[0]
	maxVal := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		}
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}

	result := int64(maxVal - minVal) * int64(k)
	return result
}

func main() {
	// result: 4
	// nums := []int{1,3,2}
	// k := int(2)

	// result: 12
	nums := []int{4,2,5,1}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maxTotalValue(nums, k)
	fmt.Printf("result = %v\n", result)
}
