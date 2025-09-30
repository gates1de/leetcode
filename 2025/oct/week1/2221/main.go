package main

import (
	"fmt"
)

func triangularSum(nums []int) int {
	for len(nums) >= 2 {
		for i := 0; i < len(nums) - 1; i++ {
			nums[i] = (nums[i] + nums[i + 1]) % 10
		}
		nums = nums[:len(nums) - 1]
	}

	return nums[0]
}

func main() {
	// result: 8
	// nums := []int{1,2,3,4,5}

	// result: 5
	nums := []int{5}

	// result: 
	// nums := []int{}

	result := triangularSum(nums)
	fmt.Printf("result = %v\n", result)
}

