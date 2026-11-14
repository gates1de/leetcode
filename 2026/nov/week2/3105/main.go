package main

import (
	"fmt"
)

func longestMonotonicSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	increase := 1
	decrease := 1
	result := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			increase++
			decrease = 1
		} else if nums[i] < nums[i-1] {
			decrease++
			increase = 1
		} else {
			increase = 1
			decrease = 1
		}

		result = max(result, increase, decrease)
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{1,4,3,3,2}

	// result: 1
	// nums := []int{3,3,3,3}

	// result: 3
	nums := []int{3, 2, 1}

	// result:
	// nums := []int{}

	result := longestMonotonicSubarray(nums)
	fmt.Printf("result = %v\n", result)
}
